from flyteidl.event import event_pb2 as _event_pb2
from flyteidl.core import literals_pb2 as _literals_pb2
from flyteidl.core import interface_pb2 as _interface_pb2
from flyteidl.core import artifact_id_pb2 as _artifact_id_pb2
from flyteidl.core import identifier_pb2 as _identifier_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CloudEventWorkflowExecution(_message.Message):
    __slots__ = ["raw_event", "output_data", "output_interface", "input_data", "artifact_ids", "reference_execution", "principal", "launch_plan_id"]
    RAW_EVENT_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_DATA_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_INTERFACE_FIELD_NUMBER: _ClassVar[int]
    INPUT_DATA_FIELD_NUMBER: _ClassVar[int]
    ARTIFACT_IDS_FIELD_NUMBER: _ClassVar[int]
    REFERENCE_EXECUTION_FIELD_NUMBER: _ClassVar[int]
    PRINCIPAL_FIELD_NUMBER: _ClassVar[int]
    LAUNCH_PLAN_ID_FIELD_NUMBER: _ClassVar[int]
    raw_event: _event_pb2.WorkflowExecutionEvent
    output_data: _literals_pb2.LiteralMap
    output_interface: _interface_pb2.TypedInterface
    input_data: _literals_pb2.LiteralMap
    artifact_ids: _containers.RepeatedCompositeFieldContainer[_artifact_id_pb2.ArtifactID]
    reference_execution: _identifier_pb2.WorkflowExecutionIdentifier
    principal: str
    launch_plan_id: _identifier_pb2.Identifier
    def __init__(self, raw_event: _Optional[_Union[_event_pb2.WorkflowExecutionEvent, _Mapping]] = ..., output_data: _Optional[_Union[_literals_pb2.LiteralMap, _Mapping]] = ..., output_interface: _Optional[_Union[_interface_pb2.TypedInterface, _Mapping]] = ..., input_data: _Optional[_Union[_literals_pb2.LiteralMap, _Mapping]] = ..., artifact_ids: _Optional[_Iterable[_Union[_artifact_id_pb2.ArtifactID, _Mapping]]] = ..., reference_execution: _Optional[_Union[_identifier_pb2.WorkflowExecutionIdentifier, _Mapping]] = ..., principal: _Optional[str] = ..., launch_plan_id: _Optional[_Union[_identifier_pb2.Identifier, _Mapping]] = ...) -> None: ...

class CloudEventNodeExecution(_message.Message):
    __slots__ = ["raw_event", "task_exec_id", "output_data", "output_interface", "input_data", "artifact_ids", "principal", "launch_plan_id"]
    RAW_EVENT_FIELD_NUMBER: _ClassVar[int]
    TASK_EXEC_ID_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_DATA_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_INTERFACE_FIELD_NUMBER: _ClassVar[int]
    INPUT_DATA_FIELD_NUMBER: _ClassVar[int]
    ARTIFACT_IDS_FIELD_NUMBER: _ClassVar[int]
    PRINCIPAL_FIELD_NUMBER: _ClassVar[int]
    LAUNCH_PLAN_ID_FIELD_NUMBER: _ClassVar[int]
    raw_event: _event_pb2.NodeExecutionEvent
    task_exec_id: _identifier_pb2.TaskExecutionIdentifier
    output_data: _literals_pb2.LiteralMap
    output_interface: _interface_pb2.TypedInterface
    input_data: _literals_pb2.LiteralMap
    artifact_ids: _containers.RepeatedCompositeFieldContainer[_artifact_id_pb2.ArtifactID]
    principal: str
    launch_plan_id: _identifier_pb2.Identifier
    def __init__(self, raw_event: _Optional[_Union[_event_pb2.NodeExecutionEvent, _Mapping]] = ..., task_exec_id: _Optional[_Union[_identifier_pb2.TaskExecutionIdentifier, _Mapping]] = ..., output_data: _Optional[_Union[_literals_pb2.LiteralMap, _Mapping]] = ..., output_interface: _Optional[_Union[_interface_pb2.TypedInterface, _Mapping]] = ..., input_data: _Optional[_Union[_literals_pb2.LiteralMap, _Mapping]] = ..., artifact_ids: _Optional[_Iterable[_Union[_artifact_id_pb2.ArtifactID, _Mapping]]] = ..., principal: _Optional[str] = ..., launch_plan_id: _Optional[_Union[_identifier_pb2.Identifier, _Mapping]] = ...) -> None: ...

class CloudEventTaskExecution(_message.Message):
    __slots__ = ["raw_event"]
    RAW_EVENT_FIELD_NUMBER: _ClassVar[int]
    raw_event: _event_pb2.TaskExecutionEvent
    def __init__(self, raw_event: _Optional[_Union[_event_pb2.TaskExecutionEvent, _Mapping]] = ...) -> None: ...

class CloudEventExecutionStart(_message.Message):
    __slots__ = ["execution_id", "launch_plan_id", "workflow_id", "artifact_ids", "artifact_keys", "principal"]
    EXECUTION_ID_FIELD_NUMBER: _ClassVar[int]
    LAUNCH_PLAN_ID_FIELD_NUMBER: _ClassVar[int]
    WORKFLOW_ID_FIELD_NUMBER: _ClassVar[int]
    ARTIFACT_IDS_FIELD_NUMBER: _ClassVar[int]
    ARTIFACT_KEYS_FIELD_NUMBER: _ClassVar[int]
    PRINCIPAL_FIELD_NUMBER: _ClassVar[int]
    execution_id: _identifier_pb2.WorkflowExecutionIdentifier
    launch_plan_id: _identifier_pb2.Identifier
    workflow_id: _identifier_pb2.Identifier
    artifact_ids: _containers.RepeatedCompositeFieldContainer[_artifact_id_pb2.ArtifactID]
    artifact_keys: _containers.RepeatedScalarFieldContainer[str]
    principal: str
    def __init__(self, execution_id: _Optional[_Union[_identifier_pb2.WorkflowExecutionIdentifier, _Mapping]] = ..., launch_plan_id: _Optional[_Union[_identifier_pb2.Identifier, _Mapping]] = ..., workflow_id: _Optional[_Union[_identifier_pb2.Identifier, _Mapping]] = ..., artifact_ids: _Optional[_Iterable[_Union[_artifact_id_pb2.ArtifactID, _Mapping]]] = ..., artifact_keys: _Optional[_Iterable[str]] = ..., principal: _Optional[str] = ...) -> None: ...
