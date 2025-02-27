# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from flyteidl.admin import agent_pb2 as flyteidl_dot_admin_dot_agent__pb2


class AsyncAgentServiceStub(object):
    """AsyncAgentService defines an RPC Service that allows propeller to send the request to the agent server.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateTask = channel.unary_unary(
                '/flyteidl.service.AsyncAgentService/CreateTask',
                request_serializer=flyteidl_dot_admin_dot_agent__pb2.CreateTaskRequest.SerializeToString,
                response_deserializer=flyteidl_dot_admin_dot_agent__pb2.CreateTaskResponse.FromString,
                )
        self.GetTask = channel.unary_unary(
                '/flyteidl.service.AsyncAgentService/GetTask',
                request_serializer=flyteidl_dot_admin_dot_agent__pb2.GetTaskRequest.SerializeToString,
                response_deserializer=flyteidl_dot_admin_dot_agent__pb2.GetTaskResponse.FromString,
                )
        self.DeleteTask = channel.unary_unary(
                '/flyteidl.service.AsyncAgentService/DeleteTask',
                request_serializer=flyteidl_dot_admin_dot_agent__pb2.DeleteTaskRequest.SerializeToString,
                response_deserializer=flyteidl_dot_admin_dot_agent__pb2.DeleteTaskResponse.FromString,
                )


class AsyncAgentServiceServicer(object):
    """AsyncAgentService defines an RPC Service that allows propeller to send the request to the agent server.
    """

    def CreateTask(self, request, context):
        """Send a task create request to the agent server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTask(self, request, context):
        """Get job status.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteTask(self, request, context):
        """Delete the task resource.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AsyncAgentServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateTask': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateTask,
                    request_deserializer=flyteidl_dot_admin_dot_agent__pb2.CreateTaskRequest.FromString,
                    response_serializer=flyteidl_dot_admin_dot_agent__pb2.CreateTaskResponse.SerializeToString,
            ),
            'GetTask': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTask,
                    request_deserializer=flyteidl_dot_admin_dot_agent__pb2.GetTaskRequest.FromString,
                    response_serializer=flyteidl_dot_admin_dot_agent__pb2.GetTaskResponse.SerializeToString,
            ),
            'DeleteTask': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteTask,
                    request_deserializer=flyteidl_dot_admin_dot_agent__pb2.DeleteTaskRequest.FromString,
                    response_serializer=flyteidl_dot_admin_dot_agent__pb2.DeleteTaskResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'flyteidl.service.AsyncAgentService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AsyncAgentService(object):
    """AsyncAgentService defines an RPC Service that allows propeller to send the request to the agent server.
    """

    @staticmethod
    def CreateTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AsyncAgentService/CreateTask',
            flyteidl_dot_admin_dot_agent__pb2.CreateTaskRequest.SerializeToString,
            flyteidl_dot_admin_dot_agent__pb2.CreateTaskResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AsyncAgentService/GetTask',
            flyteidl_dot_admin_dot_agent__pb2.GetTaskRequest.SerializeToString,
            flyteidl_dot_admin_dot_agent__pb2.GetTaskResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AsyncAgentService/DeleteTask',
            flyteidl_dot_admin_dot_agent__pb2.DeleteTaskRequest.SerializeToString,
            flyteidl_dot_admin_dot_agent__pb2.DeleteTaskResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class AgentMetadataServiceStub(object):
    """AgentMetadataService defines an RPC service that is also served over HTTP via grpc-gateway.
    This service allows propeller or users to get the metadata of agents.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetAgent = channel.unary_unary(
                '/flyteidl.service.AgentMetadataService/GetAgent',
                request_serializer=flyteidl_dot_admin_dot_agent__pb2.GetAgentRequest.SerializeToString,
                response_deserializer=flyteidl_dot_admin_dot_agent__pb2.GetAgentResponse.FromString,
                )
        self.ListAgents = channel.unary_unary(
                '/flyteidl.service.AgentMetadataService/ListAgents',
                request_serializer=flyteidl_dot_admin_dot_agent__pb2.ListAgentsRequest.SerializeToString,
                response_deserializer=flyteidl_dot_admin_dot_agent__pb2.ListAgentsResponse.FromString,
                )


class AgentMetadataServiceServicer(object):
    """AgentMetadataService defines an RPC service that is also served over HTTP via grpc-gateway.
    This service allows propeller or users to get the metadata of agents.
    """

    def GetAgent(self, request, context):
        """Fetch a :ref:`ref_flyteidl.admin.Agent` definition.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListAgents(self, request, context):
        """Fetch a list of :ref:`ref_flyteidl.admin.Agent` definitions.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AgentMetadataServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetAgent': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAgent,
                    request_deserializer=flyteidl_dot_admin_dot_agent__pb2.GetAgentRequest.FromString,
                    response_serializer=flyteidl_dot_admin_dot_agent__pb2.GetAgentResponse.SerializeToString,
            ),
            'ListAgents': grpc.unary_unary_rpc_method_handler(
                    servicer.ListAgents,
                    request_deserializer=flyteidl_dot_admin_dot_agent__pb2.ListAgentsRequest.FromString,
                    response_serializer=flyteidl_dot_admin_dot_agent__pb2.ListAgentsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'flyteidl.service.AgentMetadataService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AgentMetadataService(object):
    """AgentMetadataService defines an RPC service that is also served over HTTP via grpc-gateway.
    This service allows propeller or users to get the metadata of agents.
    """

    @staticmethod
    def GetAgent(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AgentMetadataService/GetAgent',
            flyteidl_dot_admin_dot_agent__pb2.GetAgentRequest.SerializeToString,
            flyteidl_dot_admin_dot_agent__pb2.GetAgentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListAgents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AgentMetadataService/ListAgents',
            flyteidl_dot_admin_dot_agent__pb2.ListAgentsRequest.SerializeToString,
            flyteidl_dot_admin_dot_agent__pb2.ListAgentsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
