# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetRegisteredServerResult',
    'AwaitableGetRegisteredServerResult',
    'get_registered_server',
    'get_registered_server_output',
]

@pulumi.output_type
class GetRegisteredServerResult:
    """
    Registered Server resource.
    """
    def __init__(__self__, active_auth_type=None, agent_version=None, agent_version_expiration_date=None, agent_version_status=None, application_id=None, azure_api_version=None, cluster_id=None, cluster_name=None, discovery_endpoint_uri=None, friendly_name=None, id=None, identity=None, last_heart_beat=None, last_operation_name=None, last_workflow_id=None, latest_application_id=None, management_endpoint_uri=None, monitoring_configuration=None, monitoring_endpoint_uri=None, name=None, provisioning_state=None, resource_location=None, server_certificate=None, server_id=None, server_management_error_code=None, server_name=None, server_os_version=None, server_role=None, service_location=None, storage_sync_service_uid=None, system_data=None, type=None):
        if active_auth_type and not isinstance(active_auth_type, str):
            raise TypeError("Expected argument 'active_auth_type' to be a str")
        pulumi.set(__self__, "active_auth_type", active_auth_type)
        if agent_version and not isinstance(agent_version, str):
            raise TypeError("Expected argument 'agent_version' to be a str")
        pulumi.set(__self__, "agent_version", agent_version)
        if agent_version_expiration_date and not isinstance(agent_version_expiration_date, str):
            raise TypeError("Expected argument 'agent_version_expiration_date' to be a str")
        pulumi.set(__self__, "agent_version_expiration_date", agent_version_expiration_date)
        if agent_version_status and not isinstance(agent_version_status, str):
            raise TypeError("Expected argument 'agent_version_status' to be a str")
        pulumi.set(__self__, "agent_version_status", agent_version_status)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if discovery_endpoint_uri and not isinstance(discovery_endpoint_uri, str):
            raise TypeError("Expected argument 'discovery_endpoint_uri' to be a str")
        pulumi.set(__self__, "discovery_endpoint_uri", discovery_endpoint_uri)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, bool):
            raise TypeError("Expected argument 'identity' to be a bool")
        pulumi.set(__self__, "identity", identity)
        if last_heart_beat and not isinstance(last_heart_beat, str):
            raise TypeError("Expected argument 'last_heart_beat' to be a str")
        pulumi.set(__self__, "last_heart_beat", last_heart_beat)
        if last_operation_name and not isinstance(last_operation_name, str):
            raise TypeError("Expected argument 'last_operation_name' to be a str")
        pulumi.set(__self__, "last_operation_name", last_operation_name)
        if last_workflow_id and not isinstance(last_workflow_id, str):
            raise TypeError("Expected argument 'last_workflow_id' to be a str")
        pulumi.set(__self__, "last_workflow_id", last_workflow_id)
        if latest_application_id and not isinstance(latest_application_id, str):
            raise TypeError("Expected argument 'latest_application_id' to be a str")
        pulumi.set(__self__, "latest_application_id", latest_application_id)
        if management_endpoint_uri and not isinstance(management_endpoint_uri, str):
            raise TypeError("Expected argument 'management_endpoint_uri' to be a str")
        pulumi.set(__self__, "management_endpoint_uri", management_endpoint_uri)
        if monitoring_configuration and not isinstance(monitoring_configuration, str):
            raise TypeError("Expected argument 'monitoring_configuration' to be a str")
        pulumi.set(__self__, "monitoring_configuration", monitoring_configuration)
        if monitoring_endpoint_uri and not isinstance(monitoring_endpoint_uri, str):
            raise TypeError("Expected argument 'monitoring_endpoint_uri' to be a str")
        pulumi.set(__self__, "monitoring_endpoint_uri", monitoring_endpoint_uri)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_location and not isinstance(resource_location, str):
            raise TypeError("Expected argument 'resource_location' to be a str")
        pulumi.set(__self__, "resource_location", resource_location)
        if server_certificate and not isinstance(server_certificate, str):
            raise TypeError("Expected argument 'server_certificate' to be a str")
        pulumi.set(__self__, "server_certificate", server_certificate)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if server_management_error_code and not isinstance(server_management_error_code, int):
            raise TypeError("Expected argument 'server_management_error_code' to be a int")
        pulumi.set(__self__, "server_management_error_code", server_management_error_code)
        if server_name and not isinstance(server_name, str):
            raise TypeError("Expected argument 'server_name' to be a str")
        pulumi.set(__self__, "server_name", server_name)
        if server_os_version and not isinstance(server_os_version, str):
            raise TypeError("Expected argument 'server_os_version' to be a str")
        pulumi.set(__self__, "server_os_version", server_os_version)
        if server_role and not isinstance(server_role, str):
            raise TypeError("Expected argument 'server_role' to be a str")
        pulumi.set(__self__, "server_role", server_role)
        if service_location and not isinstance(service_location, str):
            raise TypeError("Expected argument 'service_location' to be a str")
        pulumi.set(__self__, "service_location", service_location)
        if storage_sync_service_uid and not isinstance(storage_sync_service_uid, str):
            raise TypeError("Expected argument 'storage_sync_service_uid' to be a str")
        pulumi.set(__self__, "storage_sync_service_uid", storage_sync_service_uid)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="activeAuthType")
    def active_auth_type(self) -> builtins.str:
        """
        Server auth type.
        """
        return pulumi.get(self, "active_auth_type")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> Optional[builtins.str]:
        """
        Registered Server Agent Version
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="agentVersionExpirationDate")
    def agent_version_expiration_date(self) -> builtins.str:
        """
        Registered Server Agent Version Expiration Date
        """
        return pulumi.get(self, "agent_version_expiration_date")

    @property
    @pulumi.getter(name="agentVersionStatus")
    def agent_version_status(self) -> builtins.str:
        """
        Registered Server Agent Version Status
        """
        return pulumi.get(self, "agent_version_status")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[builtins.str]:
        """
        Server Application Id
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[builtins.str]:
        """
        Registered Server clusterId
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[builtins.str]:
        """
        Registered Server clusterName
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="discoveryEndpointUri")
    def discovery_endpoint_uri(self) -> Optional[builtins.str]:
        """
        Resource discoveryEndpointUri
        """
        return pulumi.get(self, "discovery_endpoint_uri")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[builtins.str]:
        """
        Friendly Name
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> builtins.bool:
        """
        Apply server with newly discovered ApplicationId if available.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="lastHeartBeat")
    def last_heart_beat(self) -> Optional[builtins.str]:
        """
        Registered Server last heart beat
        """
        return pulumi.get(self, "last_heart_beat")

    @property
    @pulumi.getter(name="lastOperationName")
    def last_operation_name(self) -> Optional[builtins.str]:
        """
        Resource Last Operation Name
        """
        return pulumi.get(self, "last_operation_name")

    @property
    @pulumi.getter(name="lastWorkflowId")
    def last_workflow_id(self) -> Optional[builtins.str]:
        """
        Registered Server lastWorkflowId
        """
        return pulumi.get(self, "last_workflow_id")

    @property
    @pulumi.getter(name="latestApplicationId")
    def latest_application_id(self) -> Optional[builtins.str]:
        """
        Latest Server Application Id discovered from the server. It is not yet applied.
        """
        return pulumi.get(self, "latest_application_id")

    @property
    @pulumi.getter(name="managementEndpointUri")
    def management_endpoint_uri(self) -> Optional[builtins.str]:
        """
        Management Endpoint Uri
        """
        return pulumi.get(self, "management_endpoint_uri")

    @property
    @pulumi.getter(name="monitoringConfiguration")
    def monitoring_configuration(self) -> Optional[builtins.str]:
        """
        Monitoring Configuration
        """
        return pulumi.get(self, "monitoring_configuration")

    @property
    @pulumi.getter(name="monitoringEndpointUri")
    def monitoring_endpoint_uri(self) -> Optional[builtins.str]:
        """
        Telemetry Endpoint Uri
        """
        return pulumi.get(self, "monitoring_endpoint_uri")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[builtins.str]:
        """
        Registered Server Provisioning State
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceLocation")
    def resource_location(self) -> Optional[builtins.str]:
        """
        Resource Location
        """
        return pulumi.get(self, "resource_location")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> Optional[builtins.str]:
        """
        Registered Server Certificate
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[builtins.str]:
        """
        Registered Server serverId
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="serverManagementErrorCode")
    def server_management_error_code(self) -> Optional[builtins.int]:
        """
        Registered Server Management Error Code
        """
        return pulumi.get(self, "server_management_error_code")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> builtins.str:
        """
        Server name
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="serverOSVersion")
    def server_os_version(self) -> Optional[builtins.str]:
        """
        Registered Server OS Version
        """
        return pulumi.get(self, "server_os_version")

    @property
    @pulumi.getter(name="serverRole")
    def server_role(self) -> Optional[builtins.str]:
        """
        Registered Server serverRole
        """
        return pulumi.get(self, "server_role")

    @property
    @pulumi.getter(name="serviceLocation")
    def service_location(self) -> Optional[builtins.str]:
        """
        Service Location
        """
        return pulumi.get(self, "service_location")

    @property
    @pulumi.getter(name="storageSyncServiceUid")
    def storage_sync_service_uid(self) -> Optional[builtins.str]:
        """
        Registered Server storageSyncServiceUid
        """
        return pulumi.get(self, "storage_sync_service_uid")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetRegisteredServerResult(GetRegisteredServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegisteredServerResult(
            active_auth_type=self.active_auth_type,
            agent_version=self.agent_version,
            agent_version_expiration_date=self.agent_version_expiration_date,
            agent_version_status=self.agent_version_status,
            application_id=self.application_id,
            azure_api_version=self.azure_api_version,
            cluster_id=self.cluster_id,
            cluster_name=self.cluster_name,
            discovery_endpoint_uri=self.discovery_endpoint_uri,
            friendly_name=self.friendly_name,
            id=self.id,
            identity=self.identity,
            last_heart_beat=self.last_heart_beat,
            last_operation_name=self.last_operation_name,
            last_workflow_id=self.last_workflow_id,
            latest_application_id=self.latest_application_id,
            management_endpoint_uri=self.management_endpoint_uri,
            monitoring_configuration=self.monitoring_configuration,
            monitoring_endpoint_uri=self.monitoring_endpoint_uri,
            name=self.name,
            provisioning_state=self.provisioning_state,
            resource_location=self.resource_location,
            server_certificate=self.server_certificate,
            server_id=self.server_id,
            server_management_error_code=self.server_management_error_code,
            server_name=self.server_name,
            server_os_version=self.server_os_version,
            server_role=self.server_role,
            service_location=self.service_location,
            storage_sync_service_uid=self.storage_sync_service_uid,
            system_data=self.system_data,
            type=self.type)


def get_registered_server(resource_group_name: Optional[builtins.str] = None,
                          server_id: Optional[builtins.str] = None,
                          storage_sync_service_name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegisteredServerResult:
    """
    Get a given registered server.

    Uses Azure REST API version 2022-09-01.

    Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_id: GUID identifying the on-premises server.
    :param builtins.str storage_sync_service_name: Name of Storage Sync Service resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverId'] = server_id
    __args__['storageSyncServiceName'] = storage_sync_service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:storagesync:getRegisteredServer', __args__, opts=opts, typ=GetRegisteredServerResult).value

    return AwaitableGetRegisteredServerResult(
        active_auth_type=pulumi.get(__ret__, 'active_auth_type'),
        agent_version=pulumi.get(__ret__, 'agent_version'),
        agent_version_expiration_date=pulumi.get(__ret__, 'agent_version_expiration_date'),
        agent_version_status=pulumi.get(__ret__, 'agent_version_status'),
        application_id=pulumi.get(__ret__, 'application_id'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        cluster_name=pulumi.get(__ret__, 'cluster_name'),
        discovery_endpoint_uri=pulumi.get(__ret__, 'discovery_endpoint_uri'),
        friendly_name=pulumi.get(__ret__, 'friendly_name'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        last_heart_beat=pulumi.get(__ret__, 'last_heart_beat'),
        last_operation_name=pulumi.get(__ret__, 'last_operation_name'),
        last_workflow_id=pulumi.get(__ret__, 'last_workflow_id'),
        latest_application_id=pulumi.get(__ret__, 'latest_application_id'),
        management_endpoint_uri=pulumi.get(__ret__, 'management_endpoint_uri'),
        monitoring_configuration=pulumi.get(__ret__, 'monitoring_configuration'),
        monitoring_endpoint_uri=pulumi.get(__ret__, 'monitoring_endpoint_uri'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        resource_location=pulumi.get(__ret__, 'resource_location'),
        server_certificate=pulumi.get(__ret__, 'server_certificate'),
        server_id=pulumi.get(__ret__, 'server_id'),
        server_management_error_code=pulumi.get(__ret__, 'server_management_error_code'),
        server_name=pulumi.get(__ret__, 'server_name'),
        server_os_version=pulumi.get(__ret__, 'server_os_version'),
        server_role=pulumi.get(__ret__, 'server_role'),
        service_location=pulumi.get(__ret__, 'service_location'),
        storage_sync_service_uid=pulumi.get(__ret__, 'storage_sync_service_uid'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_registered_server_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 server_id: Optional[pulumi.Input[builtins.str]] = None,
                                 storage_sync_service_name: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegisteredServerResult]:
    """
    Get a given registered server.

    Uses Azure REST API version 2022-09-01.

    Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_id: GUID identifying the on-premises server.
    :param builtins.str storage_sync_service_name: Name of Storage Sync Service resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverId'] = server_id
    __args__['storageSyncServiceName'] = storage_sync_service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:storagesync:getRegisteredServer', __args__, opts=opts, typ=GetRegisteredServerResult)
    return __ret__.apply(lambda __response__: GetRegisteredServerResult(
        active_auth_type=pulumi.get(__response__, 'active_auth_type'),
        agent_version=pulumi.get(__response__, 'agent_version'),
        agent_version_expiration_date=pulumi.get(__response__, 'agent_version_expiration_date'),
        agent_version_status=pulumi.get(__response__, 'agent_version_status'),
        application_id=pulumi.get(__response__, 'application_id'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        cluster_name=pulumi.get(__response__, 'cluster_name'),
        discovery_endpoint_uri=pulumi.get(__response__, 'discovery_endpoint_uri'),
        friendly_name=pulumi.get(__response__, 'friendly_name'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        last_heart_beat=pulumi.get(__response__, 'last_heart_beat'),
        last_operation_name=pulumi.get(__response__, 'last_operation_name'),
        last_workflow_id=pulumi.get(__response__, 'last_workflow_id'),
        latest_application_id=pulumi.get(__response__, 'latest_application_id'),
        management_endpoint_uri=pulumi.get(__response__, 'management_endpoint_uri'),
        monitoring_configuration=pulumi.get(__response__, 'monitoring_configuration'),
        monitoring_endpoint_uri=pulumi.get(__response__, 'monitoring_endpoint_uri'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        resource_location=pulumi.get(__response__, 'resource_location'),
        server_certificate=pulumi.get(__response__, 'server_certificate'),
        server_id=pulumi.get(__response__, 'server_id'),
        server_management_error_code=pulumi.get(__response__, 'server_management_error_code'),
        server_name=pulumi.get(__response__, 'server_name'),
        server_os_version=pulumi.get(__response__, 'server_os_version'),
        server_role=pulumi.get(__response__, 'server_role'),
        service_location=pulumi.get(__response__, 'service_location'),
        storage_sync_service_uid=pulumi.get(__response__, 'storage_sync_service_uid'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
