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
    'GetNetworkFabricControllerResult',
    'AwaitableGetNetworkFabricControllerResult',
    'get_network_fabric_controller',
    'get_network_fabric_controller_output',
]

@pulumi.output_type
class GetNetworkFabricControllerResult:
    """
    The Network Fabric Controller resource definition.
    """
    def __init__(__self__, annotation=None, azure_api_version=None, id=None, infrastructure_express_route_connections=None, infrastructure_services=None, ipv4_address_space=None, ipv6_address_space=None, is_workload_management_network_enabled=None, location=None, managed_resource_group_configuration=None, name=None, network_fabric_ids=None, nfc_sku=None, provisioning_state=None, system_data=None, tags=None, tenant_internet_gateway_ids=None, type=None, workload_express_route_connections=None, workload_management_network=None, workload_services=None):
        if annotation and not isinstance(annotation, str):
            raise TypeError("Expected argument 'annotation' to be a str")
        pulumi.set(__self__, "annotation", annotation)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if infrastructure_express_route_connections and not isinstance(infrastructure_express_route_connections, list):
            raise TypeError("Expected argument 'infrastructure_express_route_connections' to be a list")
        pulumi.set(__self__, "infrastructure_express_route_connections", infrastructure_express_route_connections)
        if infrastructure_services and not isinstance(infrastructure_services, dict):
            raise TypeError("Expected argument 'infrastructure_services' to be a dict")
        pulumi.set(__self__, "infrastructure_services", infrastructure_services)
        if ipv4_address_space and not isinstance(ipv4_address_space, str):
            raise TypeError("Expected argument 'ipv4_address_space' to be a str")
        pulumi.set(__self__, "ipv4_address_space", ipv4_address_space)
        if ipv6_address_space and not isinstance(ipv6_address_space, str):
            raise TypeError("Expected argument 'ipv6_address_space' to be a str")
        pulumi.set(__self__, "ipv6_address_space", ipv6_address_space)
        if is_workload_management_network_enabled and not isinstance(is_workload_management_network_enabled, str):
            raise TypeError("Expected argument 'is_workload_management_network_enabled' to be a str")
        pulumi.set(__self__, "is_workload_management_network_enabled", is_workload_management_network_enabled)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if managed_resource_group_configuration and not isinstance(managed_resource_group_configuration, dict):
            raise TypeError("Expected argument 'managed_resource_group_configuration' to be a dict")
        pulumi.set(__self__, "managed_resource_group_configuration", managed_resource_group_configuration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_fabric_ids and not isinstance(network_fabric_ids, list):
            raise TypeError("Expected argument 'network_fabric_ids' to be a list")
        pulumi.set(__self__, "network_fabric_ids", network_fabric_ids)
        if nfc_sku and not isinstance(nfc_sku, str):
            raise TypeError("Expected argument 'nfc_sku' to be a str")
        pulumi.set(__self__, "nfc_sku", nfc_sku)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tenant_internet_gateway_ids and not isinstance(tenant_internet_gateway_ids, list):
            raise TypeError("Expected argument 'tenant_internet_gateway_ids' to be a list")
        pulumi.set(__self__, "tenant_internet_gateway_ids", tenant_internet_gateway_ids)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if workload_express_route_connections and not isinstance(workload_express_route_connections, list):
            raise TypeError("Expected argument 'workload_express_route_connections' to be a list")
        pulumi.set(__self__, "workload_express_route_connections", workload_express_route_connections)
        if workload_management_network and not isinstance(workload_management_network, bool):
            raise TypeError("Expected argument 'workload_management_network' to be a bool")
        pulumi.set(__self__, "workload_management_network", workload_management_network)
        if workload_services and not isinstance(workload_services, dict):
            raise TypeError("Expected argument 'workload_services' to be a dict")
        pulumi.set(__self__, "workload_services", workload_services)

    @property
    @pulumi.getter
    def annotation(self) -> Optional[builtins.str]:
        """
        Switch configuration description.
        """
        return pulumi.get(self, "annotation")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="infrastructureExpressRouteConnections")
    def infrastructure_express_route_connections(self) -> Optional[Sequence['outputs.ExpressRouteConnectionInformationResponse']]:
        """
        As part of an update, the Infrastructure ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Infrastructure services. (This is a Mandatory attribute)
        """
        return pulumi.get(self, "infrastructure_express_route_connections")

    @property
    @pulumi.getter(name="infrastructureServices")
    def infrastructure_services(self) -> 'outputs.ControllerServicesResponse':
        """
        InfrastructureServices IP ranges.
        """
        return pulumi.get(self, "infrastructure_services")

    @property
    @pulumi.getter(name="ipv4AddressSpace")
    def ipv4_address_space(self) -> Optional[builtins.str]:
        """
        IPv4 Network Fabric Controller Address Space.
        """
        return pulumi.get(self, "ipv4_address_space")

    @property
    @pulumi.getter(name="ipv6AddressSpace")
    def ipv6_address_space(self) -> Optional[builtins.str]:
        """
        IPv6 Network Fabric Controller Address Space.
        """
        return pulumi.get(self, "ipv6_address_space")

    @property
    @pulumi.getter(name="isWorkloadManagementNetworkEnabled")
    def is_workload_management_network_enabled(self) -> Optional[builtins.str]:
        """
        A workload management network is required for all the tenant (workload) traffic. This traffic is only dedicated for Tenant workloads which are required to access internet or any other MSFT/Public endpoints.
        """
        return pulumi.get(self, "is_workload_management_network_enabled")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> Optional['outputs.ManagedResourceGroupConfigurationResponse']:
        """
        Managed Resource Group configuration properties.
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkFabricIds")
    def network_fabric_ids(self) -> Sequence[builtins.str]:
        """
        The NF-ID will be an input parameter used by the NF to link and get associated with the parent NFC Service.
        """
        return pulumi.get(self, "network_fabric_ids")

    @property
    @pulumi.getter(name="nfcSku")
    def nfc_sku(self) -> Optional[builtins.str]:
        """
        Network Fabric Controller SKU.
        """
        return pulumi.get(self, "nfc_sku")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provides you the latest status of the NFC service, whether it is Accepted, updating, Succeeded or Failed. During this process, the states keep changing based on the status of NFC provisioning.
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
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantInternetGatewayIds")
    def tenant_internet_gateway_ids(self) -> Sequence[builtins.str]:
        """
        List of tenant InternetGateway resource IDs
        """
        return pulumi.get(self, "tenant_internet_gateway_ids")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workloadExpressRouteConnections")
    def workload_express_route_connections(self) -> Optional[Sequence['outputs.ExpressRouteConnectionInformationResponse']]:
        """
        As part of an update, the workload ExpressRoute CircuitID should be provided to create and Provision a NFC. This Express route is dedicated for Workload services. (This is a Mandatory attribute).
        """
        return pulumi.get(self, "workload_express_route_connections")

    @property
    @pulumi.getter(name="workloadManagementNetwork")
    def workload_management_network(self) -> builtins.bool:
        """
        A workload management network is required for all the tenant (workload) traffic. This traffic is only dedicated for Tenant workloads which are required to access internet or any other MSFT/Public endpoints. This is used for the backward compatibility.
        """
        return pulumi.get(self, "workload_management_network")

    @property
    @pulumi.getter(name="workloadServices")
    def workload_services(self) -> 'outputs.ControllerServicesResponse':
        """
        WorkloadServices IP ranges.
        """
        return pulumi.get(self, "workload_services")


class AwaitableGetNetworkFabricControllerResult(GetNetworkFabricControllerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkFabricControllerResult(
            annotation=self.annotation,
            azure_api_version=self.azure_api_version,
            id=self.id,
            infrastructure_express_route_connections=self.infrastructure_express_route_connections,
            infrastructure_services=self.infrastructure_services,
            ipv4_address_space=self.ipv4_address_space,
            ipv6_address_space=self.ipv6_address_space,
            is_workload_management_network_enabled=self.is_workload_management_network_enabled,
            location=self.location,
            managed_resource_group_configuration=self.managed_resource_group_configuration,
            name=self.name,
            network_fabric_ids=self.network_fabric_ids,
            nfc_sku=self.nfc_sku,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            tenant_internet_gateway_ids=self.tenant_internet_gateway_ids,
            type=self.type,
            workload_express_route_connections=self.workload_express_route_connections,
            workload_management_network=self.workload_management_network,
            workload_services=self.workload_services)


def get_network_fabric_controller(network_fabric_controller_name: Optional[builtins.str] = None,
                                  resource_group_name: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkFabricControllerResult:
    """
    Shows the provisioning status of Network Fabric Controller.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_fabric_controller_name: Name of the Network Fabric Controller.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkFabricControllerName'] = network_fabric_controller_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:managednetworkfabric:getNetworkFabricController', __args__, opts=opts, typ=GetNetworkFabricControllerResult).value

    return AwaitableGetNetworkFabricControllerResult(
        annotation=pulumi.get(__ret__, 'annotation'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        infrastructure_express_route_connections=pulumi.get(__ret__, 'infrastructure_express_route_connections'),
        infrastructure_services=pulumi.get(__ret__, 'infrastructure_services'),
        ipv4_address_space=pulumi.get(__ret__, 'ipv4_address_space'),
        ipv6_address_space=pulumi.get(__ret__, 'ipv6_address_space'),
        is_workload_management_network_enabled=pulumi.get(__ret__, 'is_workload_management_network_enabled'),
        location=pulumi.get(__ret__, 'location'),
        managed_resource_group_configuration=pulumi.get(__ret__, 'managed_resource_group_configuration'),
        name=pulumi.get(__ret__, 'name'),
        network_fabric_ids=pulumi.get(__ret__, 'network_fabric_ids'),
        nfc_sku=pulumi.get(__ret__, 'nfc_sku'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_internet_gateway_ids=pulumi.get(__ret__, 'tenant_internet_gateway_ids'),
        type=pulumi.get(__ret__, 'type'),
        workload_express_route_connections=pulumi.get(__ret__, 'workload_express_route_connections'),
        workload_management_network=pulumi.get(__ret__, 'workload_management_network'),
        workload_services=pulumi.get(__ret__, 'workload_services'))
def get_network_fabric_controller_output(network_fabric_controller_name: Optional[pulumi.Input[builtins.str]] = None,
                                         resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkFabricControllerResult]:
    """
    Shows the provisioning status of Network Fabric Controller.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_fabric_controller_name: Name of the Network Fabric Controller.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkFabricControllerName'] = network_fabric_controller_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:managednetworkfabric:getNetworkFabricController', __args__, opts=opts, typ=GetNetworkFabricControllerResult)
    return __ret__.apply(lambda __response__: GetNetworkFabricControllerResult(
        annotation=pulumi.get(__response__, 'annotation'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        infrastructure_express_route_connections=pulumi.get(__response__, 'infrastructure_express_route_connections'),
        infrastructure_services=pulumi.get(__response__, 'infrastructure_services'),
        ipv4_address_space=pulumi.get(__response__, 'ipv4_address_space'),
        ipv6_address_space=pulumi.get(__response__, 'ipv6_address_space'),
        is_workload_management_network_enabled=pulumi.get(__response__, 'is_workload_management_network_enabled'),
        location=pulumi.get(__response__, 'location'),
        managed_resource_group_configuration=pulumi.get(__response__, 'managed_resource_group_configuration'),
        name=pulumi.get(__response__, 'name'),
        network_fabric_ids=pulumi.get(__response__, 'network_fabric_ids'),
        nfc_sku=pulumi.get(__response__, 'nfc_sku'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        tenant_internet_gateway_ids=pulumi.get(__response__, 'tenant_internet_gateway_ids'),
        type=pulumi.get(__response__, 'type'),
        workload_express_route_connections=pulumi.get(__response__, 'workload_express_route_connections'),
        workload_management_network=pulumi.get(__response__, 'workload_management_network'),
        workload_services=pulumi.get(__response__, 'workload_services')))
