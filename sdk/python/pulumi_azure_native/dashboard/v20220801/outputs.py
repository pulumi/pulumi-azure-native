# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'AzureMonitorWorkspaceIntegrationResponse',
    'GrafanaIntegrationsResponse',
    'ManagedGrafanaPropertiesResponse',
    'ManagedServiceIdentityResponse',
    'PrivateEndpointConnectionResponse',
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'ResourceSkuResponse',
    'SystemDataResponse',
    'UserAssignedIdentityResponse',
]

@pulumi.output_type
class AzureMonitorWorkspaceIntegrationResponse(dict):
    """
    Integrations for Azure Monitor Workspace.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "azureMonitorWorkspaceResourceId":
            suggest = "azure_monitor_workspace_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureMonitorWorkspaceIntegrationResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureMonitorWorkspaceIntegrationResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureMonitorWorkspaceIntegrationResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 azure_monitor_workspace_resource_id: Optional[str] = None):
        """
        Integrations for Azure Monitor Workspace.
        :param str azure_monitor_workspace_resource_id: The resource Id of the connected Azure Monitor Workspace.
        """
        if azure_monitor_workspace_resource_id is not None:
            pulumi.set(__self__, "azure_monitor_workspace_resource_id", azure_monitor_workspace_resource_id)

    @property
    @pulumi.getter(name="azureMonitorWorkspaceResourceId")
    def azure_monitor_workspace_resource_id(self) -> Optional[str]:
        """
        The resource Id of the connected Azure Monitor Workspace.
        """
        return pulumi.get(self, "azure_monitor_workspace_resource_id")


@pulumi.output_type
class GrafanaIntegrationsResponse(dict):
    """
    GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards, alerting defaults) for common monitoring scenarios.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "azureMonitorWorkspaceIntegrations":
            suggest = "azure_monitor_workspace_integrations"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GrafanaIntegrationsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GrafanaIntegrationsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GrafanaIntegrationsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 azure_monitor_workspace_integrations: Optional[Sequence['outputs.AzureMonitorWorkspaceIntegrationResponse']] = None):
        """
        GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards, alerting defaults) for common monitoring scenarios.
        """
        if azure_monitor_workspace_integrations is not None:
            pulumi.set(__self__, "azure_monitor_workspace_integrations", azure_monitor_workspace_integrations)

    @property
    @pulumi.getter(name="azureMonitorWorkspaceIntegrations")
    def azure_monitor_workspace_integrations(self) -> Optional[Sequence['outputs.AzureMonitorWorkspaceIntegrationResponse']]:
        return pulumi.get(self, "azure_monitor_workspace_integrations")


@pulumi.output_type
class ManagedGrafanaPropertiesResponse(dict):
    """
    Properties specific to the grafana resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "grafanaVersion":
            suggest = "grafana_version"
        elif key == "outboundIPs":
            suggest = "outbound_ips"
        elif key == "privateEndpointConnections":
            suggest = "private_endpoint_connections"
        elif key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "apiKey":
            suggest = "api_key"
        elif key == "autoGeneratedDomainNameLabelScope":
            suggest = "auto_generated_domain_name_label_scope"
        elif key == "deterministicOutboundIP":
            suggest = "deterministic_outbound_ip"
        elif key == "grafanaIntegrations":
            suggest = "grafana_integrations"
        elif key == "publicNetworkAccess":
            suggest = "public_network_access"
        elif key == "zoneRedundancy":
            suggest = "zone_redundancy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedGrafanaPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedGrafanaPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedGrafanaPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 endpoint: str,
                 grafana_version: str,
                 outbound_ips: Sequence[str],
                 private_endpoint_connections: Sequence['outputs.PrivateEndpointConnectionResponse'],
                 provisioning_state: str,
                 api_key: Optional[str] = None,
                 auto_generated_domain_name_label_scope: Optional[str] = None,
                 deterministic_outbound_ip: Optional[str] = None,
                 grafana_integrations: Optional['outputs.GrafanaIntegrationsResponse'] = None,
                 public_network_access: Optional[str] = None,
                 zone_redundancy: Optional[str] = None):
        """
        Properties specific to the grafana resource.
        :param str endpoint: The endpoint of the Grafana instance.
        :param str grafana_version: The Grafana software version.
        :param Sequence[str] outbound_ips: List of outbound IPs if deterministicOutboundIP is enabled.
        :param Sequence['PrivateEndpointConnectionResponse'] private_endpoint_connections: The private endpoint connections of the Grafana instance.
        :param str provisioning_state: Provisioning state of the resource.
        :param str api_key: The api key setting of the Grafana instance.
        :param str auto_generated_domain_name_label_scope: Scope for dns deterministic name hash calculation.
        :param str deterministic_outbound_ip: Whether a Grafana instance uses deterministic outbound IPs.
        :param 'GrafanaIntegrationsResponse' grafana_integrations: GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards, alerting defaults) for common monitoring scenarios.
        :param str public_network_access: Indicate the state for enable or disable traffic over the public interface.
        :param str zone_redundancy: The zone redundancy setting of the Grafana instance.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "grafana_version", grafana_version)
        pulumi.set(__self__, "outbound_ips", outbound_ips)
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if auto_generated_domain_name_label_scope is not None:
            pulumi.set(__self__, "auto_generated_domain_name_label_scope", auto_generated_domain_name_label_scope)
        if deterministic_outbound_ip is not None:
            pulumi.set(__self__, "deterministic_outbound_ip", deterministic_outbound_ip)
        if grafana_integrations is not None:
            pulumi.set(__self__, "grafana_integrations", grafana_integrations)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if zone_redundancy is not None:
            pulumi.set(__self__, "zone_redundancy", zone_redundancy)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The endpoint of the Grafana instance.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="grafanaVersion")
    def grafana_version(self) -> str:
        """
        The Grafana software version.
        """
        return pulumi.get(self, "grafana_version")

    @property
    @pulumi.getter(name="outboundIPs")
    def outbound_ips(self) -> Sequence[str]:
        """
        List of outbound IPs if deterministicOutboundIP is enabled.
        """
        return pulumi.get(self, "outbound_ips")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Sequence['outputs.PrivateEndpointConnectionResponse']:
        """
        The private endpoint connections of the Grafana instance.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[str]:
        """
        The api key setting of the Grafana instance.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="autoGeneratedDomainNameLabelScope")
    def auto_generated_domain_name_label_scope(self) -> Optional[str]:
        """
        Scope for dns deterministic name hash calculation.
        """
        return pulumi.get(self, "auto_generated_domain_name_label_scope")

    @property
    @pulumi.getter(name="deterministicOutboundIP")
    def deterministic_outbound_ip(self) -> Optional[str]:
        """
        Whether a Grafana instance uses deterministic outbound IPs.
        """
        return pulumi.get(self, "deterministic_outbound_ip")

    @property
    @pulumi.getter(name="grafanaIntegrations")
    def grafana_integrations(self) -> Optional['outputs.GrafanaIntegrationsResponse']:
        """
        GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards, alerting defaults) for common monitoring scenarios.
        """
        return pulumi.get(self, "grafana_integrations")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        Indicate the state for enable or disable traffic over the public interface.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="zoneRedundancy")
    def zone_redundancy(self) -> Optional[str]:
        """
        The zone redundancy setting of the Grafana instance.
        """
        return pulumi.get(self, "zone_redundancy")


@pulumi.output_type
class ManagedServiceIdentityResponse(dict):
    """
    Managed service identity (system assigned and/or user assigned identities)
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"
        elif key == "userAssignedIdentities":
            suggest = "user_assigned_identities"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedServiceIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedServiceIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedServiceIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str,
                 user_assigned_identities: Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']] = None):
        """
        Managed service identity (system assigned and/or user assigned identities)
        :param str principal_id: The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
        :param str tenant_id: The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
        :param str type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
        :param Mapping[str, 'UserAssignedIdentityResponse'] user_assigned_identities: The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']]:
        """
        The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        """
        return pulumi.get(self, "user_assigned_identities")


@pulumi.output_type
class PrivateEndpointConnectionResponse(dict):
    """
    The Private Endpoint Connection resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "privateLinkServiceConnectionState":
            suggest = "private_link_service_connection_state"
        elif key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "systemData":
            suggest = "system_data"
        elif key == "groupIds":
            suggest = "group_ids"
        elif key == "privateEndpoint":
            suggest = "private_endpoint"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateEndpointConnectionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateEndpointConnectionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateEndpointConnectionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 name: str,
                 private_link_service_connection_state: 'outputs.PrivateLinkServiceConnectionStateResponse',
                 provisioning_state: str,
                 system_data: 'outputs.SystemDataResponse',
                 type: str,
                 group_ids: Optional[Sequence[str]] = None,
                 private_endpoint: Optional['outputs.PrivateEndpointResponse'] = None):
        """
        The Private Endpoint Connection resource.
        :param str id: Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        :param str name: The name of the resource
        :param 'PrivateLinkServiceConnectionStateResponse' private_link_service_connection_state: A collection of information about the state of the connection between service consumer and provider.
        :param str provisioning_state: The provisioning state of the private endpoint connection resource.
        :param 'SystemDataResponse' system_data: Azure Resource Manager metadata containing createdBy and modifiedBy information.
        :param str type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        :param Sequence[str] group_ids: The private endpoint connection group ids.
        :param 'PrivateEndpointResponse' private_endpoint: The resource of private end point.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        pulumi.set(__self__, "system_data", system_data)
        pulumi.set(__self__, "type", type)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> 'outputs.PrivateLinkServiceConnectionStateResponse':
        """
        A collection of information about the state of the connection between service consumer and provider.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the private endpoint connection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[Sequence[str]]:
        """
        The private endpoint connection group ids.
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.PrivateEndpointResponse']:
        """
        The resource of private end point.
        """
        return pulumi.get(self, "private_endpoint")


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    The Private Endpoint resource.
    """
    def __init__(__self__, *,
                 id: str):
        """
        The Private Endpoint resource.
        :param str id: The ARM identifier for Private Endpoint
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ARM identifier for Private Endpoint
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    A collection of information about the state of the connection between service consumer and provider.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionsRequired":
            suggest = "actions_required"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateLinkServiceConnectionStateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 actions_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        A collection of information about the state of the connection between service consumer and provider.
        :param str actions_required: A message indicating if changes on the service provider require any updates on the consumer.
        :param str description: The reason for approval/rejection of the connection.
        :param str status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        if actions_required is not None:
            pulumi.set(__self__, "actions_required", actions_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> Optional[str]:
        """
        A message indicating if changes on the service provider require any updates on the consumer.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The reason for approval/rejection of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ResourceSkuResponse(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


@pulumi.output_type
class UserAssignedIdentityResponse(dict):
    """
    User assigned identity properties
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "principalId":
            suggest = "principal_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in UserAssignedIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 principal_id: str):
        """
        User assigned identity properties
        :param str client_id: The client ID of the assigned identity.
        :param str principal_id: The principal ID of the assigned identity.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The client ID of the assigned identity.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal ID of the assigned identity.
        """
        return pulumi.get(self, "principal_id")


