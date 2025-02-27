package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/utils/strings/slices"

	"github.com/flyteorg/flyte/flyteidl/clients/go/coreutils"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	flyteIdlCore "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/plugins"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	pluginCore "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	pluginCoreMocks "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core/mocks"
	ioMocks "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/webapi"
	"github.com/flyteorg/flyte/flyteplugins/tests"
	"github.com/flyteorg/flyte/flytestdlib/contextutils"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
	"github.com/flyteorg/flyte/flytestdlib/promutils/labeled"
	"github.com/flyteorg/flyte/flytestdlib/storage"
	"github.com/flyteorg/flyte/flytestdlib/utils"
)

type MockPlugin struct {
	Plugin
}

type MockAsyncTask struct {
}

type MockSyncTask struct {
}

func (m *MockAsyncTask) CreateTask(_ context.Context, createTaskRequest *admin.CreateTaskRequest, _ ...grpc.CallOption) (*admin.CreateTaskResponse, error) {
	expectedArgs := []string{"pyflyte-fast-execute", "--output-prefix", "fake://bucket/prefix/nhv"}
	if slices.Equal(createTaskRequest.Template.GetContainer().Args, expectedArgs) {
		return nil, fmt.Errorf("args not as expected")
	}
	return &admin.CreateTaskResponse{
		Res: &admin.CreateTaskResponse_ResourceMeta{
			ResourceMeta: []byte{1, 2, 3, 4},
		}}, nil
}

func (m *MockAsyncTask) GetTask(_ context.Context, req *admin.GetTaskRequest, _ ...grpc.CallOption) (*admin.GetTaskResponse, error) {
	if req.GetTaskType() == "bigquery_query_job_task" {
		return &admin.GetTaskResponse{Resource: &admin.Resource{State: admin.State_SUCCEEDED, Outputs: &flyteIdlCore.LiteralMap{
			Literals: map[string]*flyteIdlCore.Literal{
				"arr": coreutils.MustMakeLiteral([]interface{}{[]interface{}{"a", "b"}, []interface{}{1, 2}}),
			},
		}}}, nil
	}
	return &admin.GetTaskResponse{Resource: &admin.Resource{State: admin.State_SUCCEEDED}}, nil
}

func (m *MockAsyncTask) DeleteTask(_ context.Context, _ *admin.DeleteTaskRequest, _ ...grpc.CallOption) (*admin.DeleteTaskResponse, error) {
	return &admin.DeleteTaskResponse{}, nil
}

func (m *MockSyncTask) CreateTask(_ context.Context, createTaskRequest *admin.CreateTaskRequest, _ ...grpc.CallOption) (*admin.CreateTaskResponse, error) {
	return &admin.CreateTaskResponse{
		Res: &admin.CreateTaskResponse_Resource{
			Resource: &admin.Resource{
				State: admin.State_SUCCEEDED,
				Outputs: &flyteIdlCore.LiteralMap{
					Literals: map[string]*flyteIdlCore.Literal{},
				},
				Message:  "Sync task finished",
				LogLinks: []*flyteIdlCore.TaskLog{{Uri: "http://localhost:3000/log", Name: "Log Link"}},
			},
		},
	}, nil

}

func (m *MockSyncTask) GetTask(_ context.Context, req *admin.GetTaskRequest, _ ...grpc.CallOption) (*admin.GetTaskResponse, error) {
	if req.GetTaskType() == "fake_task" {
		return &admin.GetTaskResponse{Resource: &admin.Resource{State: admin.State_SUCCEEDED, Outputs: &flyteIdlCore.LiteralMap{
			Literals: map[string]*flyteIdlCore.Literal{
				"arr": coreutils.MustMakeLiteral([]interface{}{[]interface{}{"a", "b"}, []interface{}{1, 2}}),
			},
		}}}, nil
	}
	return &admin.GetTaskResponse{Resource: &admin.Resource{State: admin.State_SUCCEEDED}}, nil
}

func (m *MockSyncTask) DeleteTask(_ context.Context, _ *admin.DeleteTaskRequest, _ ...grpc.CallOption) (*admin.DeleteTaskResponse, error) {
	return &admin.DeleteTaskResponse{}, nil
}

func mockAsyncTaskClientFunc(_ context.Context, _ *Agent, _ map[*Agent]*grpc.ClientConn) (service.AsyncAgentServiceClient, error) {
	return &MockAsyncTask{}, nil
}

func mockSyncTaskClientFunc(_ context.Context, _ *Agent, _ map[*Agent]*grpc.ClientConn) (service.AsyncAgentServiceClient, error) {
	return &MockSyncTask{}, nil
}

func mockGetBadAsyncClientFunc(_ context.Context, _ *Agent, _ map[*Agent]*grpc.ClientConn) (service.AsyncAgentServiceClient, error) {
	return nil, fmt.Errorf("error")
}

func TestEndToEnd(t *testing.T) {
	iter := func(ctx context.Context, tCtx pluginCore.TaskExecutionContext) error {
		return nil
	}

	cfg := defaultConfig
	cfg.WebAPI.ResourceQuotas = map[core.ResourceNamespace]int{}
	cfg.WebAPI.Caching.Workers = 1
	cfg.WebAPI.Caching.ResyncInterval.Duration = 5 * time.Second
	err := SetConfig(&cfg)
	assert.NoError(t, err)

	databricksConfDict := map[string]interface{}{
		"name": "flytekit databricks plugin example",
		"new_cluster": map[string]string{
			"spark_version": "11.0.x-scala2.12",
			"node_type_id":  "r3.xlarge",
			"num_workers":   "4",
		},
		"timeout_seconds": 3600,
		"max_retries":     1,
	}
	databricksConfig, err := utils.MarshalObjToStruct(databricksConfDict)
	assert.NoError(t, err)
	sparkJob := plugins.SparkJob{DatabricksConf: databricksConfig, DatabricksToken: "token", SparkConf: map[string]string{"spark.driver.bindAddress": "127.0.0.1"}}
	st, err := utils.MarshalPbToStruct(&sparkJob)
	assert.NoError(t, err)

	inputs, _ := coreutils.MakeLiteralMap(map[string]interface{}{"x": 1})
	template := flyteIdlCore.TaskTemplate{
		Type:   "bigquery_query_job_task",
		Custom: st,
		Target: &flyteIdlCore.TaskTemplate_Container{
			Container: &flyteIdlCore.Container{Args: []string{"pyflyte-fast-execute", "--output-prefix", "{{.outputPrefix}}"}},
		},
	}
	basePrefix := storage.DataReference("fake://bucket/prefix/")

	t.Run("run a job", func(t *testing.T) {
		pluginEntry := pluginmachinery.CreateRemotePlugin(newMockAgentPlugin())
		plugin, err := pluginEntry.LoadPlugin(context.TODO(), newFakeSetupContext("test1"))
		assert.NoError(t, err)

		phase := tests.RunPluginEndToEndTest(t, plugin, &template, inputs, nil, nil, iter)
		assert.Equal(t, true, phase.Phase().IsSuccess())

		template.Type = "spark_job"
		phase = tests.RunPluginEndToEndTest(t, plugin, &template, inputs, nil, nil, iter)
		assert.Equal(t, true, phase.Phase().IsSuccess())

	})

	t.Run("failed to create a job", func(t *testing.T) {
		agentPlugin := newMockAgentPlugin()
		agentPlugin.PluginLoader = func(ctx context.Context, iCtx webapi.PluginSetupContext) (webapi.AsyncPlugin, error) {
			return &MockPlugin{
				Plugin{
					metricScope: iCtx.MetricsScope(),
					cfg:         GetConfig(),
					getClient:   mockGetBadAsyncClientFunc,
				},
			}, nil
		}
		pluginEntry := pluginmachinery.CreateRemotePlugin(agentPlugin)
		plugin, err := pluginEntry.LoadPlugin(context.TODO(), newFakeSetupContext("test2"))
		assert.NoError(t, err)

		tCtx := getTaskContext(t)
		tr := &pluginCoreMocks.TaskReader{}
		tr.OnRead(context.Background()).Return(&template, nil)
		tCtx.OnTaskReader().Return(tr)
		inputReader := &ioMocks.InputReader{}
		inputReader.OnGetInputPrefixPath().Return(basePrefix)
		inputReader.OnGetInputPath().Return(basePrefix + "/inputs.pb")
		inputReader.OnGetMatch(mock.Anything).Return(inputs, nil)
		tCtx.OnInputReader().Return(inputReader)

		trns, err := plugin.Handle(context.Background(), tCtx)
		assert.Error(t, err)
		assert.Equal(t, trns.Info().Phase(), core.PhaseUndefined)
		err = plugin.Abort(context.Background(), tCtx)
		assert.Nil(t, err)
	})

	t.Run("failed to read task template", func(t *testing.T) {
		tCtx := getTaskContext(t)
		tr := &pluginCoreMocks.TaskReader{}
		tr.OnRead(context.Background()).Return(nil, fmt.Errorf("read fail"))
		tCtx.OnTaskReader().Return(tr)

		agentPlugin := newMockAgentPlugin()
		pluginEntry := pluginmachinery.CreateRemotePlugin(agentPlugin)
		plugin, err := pluginEntry.LoadPlugin(context.TODO(), newFakeSetupContext("test3"))
		assert.NoError(t, err)

		trns, err := plugin.Handle(context.Background(), tCtx)
		assert.Error(t, err)
		assert.Equal(t, trns.Info().Phase(), core.PhaseUndefined)
	})

	t.Run("failed to read inputs", func(t *testing.T) {
		tCtx := getTaskContext(t)
		tr := &pluginCoreMocks.TaskReader{}
		tr.OnRead(context.Background()).Return(&template, nil)
		tCtx.OnTaskReader().Return(tr)
		inputReader := &ioMocks.InputReader{}
		inputReader.OnGetInputPrefixPath().Return(basePrefix)
		inputReader.OnGetInputPath().Return(basePrefix + "/inputs.pb")
		inputReader.OnGetMatch(mock.Anything).Return(nil, fmt.Errorf("read fail"))
		tCtx.OnInputReader().Return(inputReader)

		agentPlugin := newMockAgentPlugin()
		pluginEntry := pluginmachinery.CreateRemotePlugin(agentPlugin)
		plugin, err := pluginEntry.LoadPlugin(context.TODO(), newFakeSetupContext("test4"))
		assert.NoError(t, err)

		trns, err := plugin.Handle(context.Background(), tCtx)
		assert.Error(t, err)
		assert.Equal(t, trns.Info().Phase(), core.PhaseUndefined)
	})
}

func getTaskContext(t *testing.T) *pluginCoreMocks.TaskExecutionContext {
	latestKnownState := atomic.Value{}
	pluginStateReader := &pluginCoreMocks.PluginStateReader{}
	pluginStateReader.OnGetMatch(mock.Anything).Return(0, nil).Run(func(args mock.Arguments) {
		o := args.Get(0)
		x, err := json.Marshal(latestKnownState.Load())
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(x, &o))
	})
	pluginStateWriter := &pluginCoreMocks.PluginStateWriter{}
	pluginStateWriter.OnPutMatch(mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		latestKnownState.Store(args.Get(1))
	})

	pluginStateWriter.OnReset().Return(nil).Run(func(args mock.Arguments) {
		latestKnownState.Store(nil)
	})

	execID := rand.String(3)
	tID := &pluginCoreMocks.TaskExecutionID{}
	tID.OnGetGeneratedName().Return(execID + "-my-task-1")
	tID.OnGetID().Return(flyteIdlCore.TaskExecutionIdentifier{
		TaskId: &flyteIdlCore.Identifier{
			ResourceType: flyteIdlCore.ResourceType_TASK,
			Project:      "a",
			Domain:       "d",
			Name:         "n",
			Version:      "abc",
		},
		NodeExecutionId: &flyteIdlCore.NodeExecutionIdentifier{
			NodeId: "node1",
			ExecutionId: &flyteIdlCore.WorkflowExecutionIdentifier{
				Project: "a",
				Domain:  "d",
				Name:    "exec",
			},
		},
		RetryAttempt: 0,
	})
	tMeta := &pluginCoreMocks.TaskExecutionMetadata{}
	tMeta.OnGetTaskExecutionID().Return(tID)
	tMeta.OnGetNamespace().Return("test-namespace")
	tMeta.OnGetLabels().Return(map[string]string{"foo": "bar"})
	tMeta.OnGetAnnotations().Return(map[string]string{"foo": "bar"})
	tMeta.OnGetK8sServiceAccount().Return("k8s-account")
	tMeta.OnGetEnvironmentVariables().Return(map[string]string{"foo": "bar"})
	resourceManager := &pluginCoreMocks.ResourceManager{}
	resourceManager.OnAllocateResourceMatch(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pluginCore.AllocationStatusGranted, nil)
	resourceManager.OnReleaseResourceMatch(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	basePrefix := storage.DataReference("fake://bucket/prefix/" + execID)
	outputWriter := &ioMocks.OutputWriter{}
	outputWriter.OnGetRawOutputPrefix().Return("/sandbox/")
	outputWriter.OnGetOutputPrefixPath().Return(basePrefix)
	outputWriter.OnGetErrorPath().Return(basePrefix + "/error.pb")
	outputWriter.OnGetOutputPath().Return(basePrefix + "/outputs.pb")
	outputWriter.OnGetCheckpointPrefix().Return("/checkpoint")
	outputWriter.OnGetPreviousCheckpointsPrefix().Return("/prev")

	tCtx := &pluginCoreMocks.TaskExecutionContext{}
	tCtx.OnOutputWriter().Return(outputWriter)
	tCtx.OnResourceManager().Return(resourceManager)
	tCtx.OnPluginStateReader().Return(pluginStateReader)
	tCtx.OnPluginStateWriter().Return(pluginStateWriter)
	tCtx.OnTaskExecutionMetadata().Return(tMeta)
	return tCtx
}

func newMockAgentPlugin() webapi.PluginEntry {
	return webapi.PluginEntry{
		ID:                 "agent-service",
		SupportedTaskTypes: []core.TaskType{"bigquery_query_job_task", "spark_job", "api_task"},
		PluginLoader: func(ctx context.Context, iCtx webapi.PluginSetupContext) (webapi.AsyncPlugin, error) {
			return &MockPlugin{
				Plugin{
					metricScope: iCtx.MetricsScope(),
					cfg:         GetConfig(),
					getClient:   mockAsyncTaskClientFunc,
				},
			}, nil
		},
	}
}

func newMockSyncAgentPlugin() webapi.PluginEntry {
	return webapi.PluginEntry{
		ID:                 "agent-service",
		SupportedTaskTypes: []core.TaskType{"bigquery_query_job_task", "spark_job", "api_task"},
		PluginLoader: func(ctx context.Context, iCtx webapi.PluginSetupContext) (webapi.AsyncPlugin, error) {
			return &MockPlugin{
				Plugin{
					metricScope: iCtx.MetricsScope(),
					cfg:         GetConfig(),
					getClient:   mockSyncTaskClientFunc,
				},
			}, nil
		},
	}
}

func newFakeSetupContext(name string) *pluginCoreMocks.SetupContext {
	fakeResourceRegistrar := pluginCoreMocks.ResourceRegistrar{}
	fakeResourceRegistrar.On("RegisterResourceQuota", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	labeled.SetMetricKeys(contextutils.NamespaceKey)

	fakeSetupContext := pluginCoreMocks.SetupContext{}
	fakeSetupContext.OnMetricsScope().Return(promutils.NewScope(name))
	fakeSetupContext.OnResourceRegistrar().Return(&fakeResourceRegistrar)

	return &fakeSetupContext
}
