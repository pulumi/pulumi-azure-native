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
    'GetSapVirtualInstanceResult',
    'AwaitableGetSapVirtualInstanceResult',
    'get_sap_virtual_instance',
    'get_sap_virtual_instance_output',
]

@pulumi.output_type
class GetSapVirtualInstanceResult:
    """
    Define the Virtual Instance for SAP solutions resource.
    """
    def __init__(__self__, azure_api_version=None, configuration=None, environment=None, errors=None, health=None, id=None, identity=None, location=None, managed_resource_group_configuration=None, managed_resources_network_access_type=None, name=None, provisioning_state=None, sap_product=None, state=None, status=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if errors and not isinstance(errors, dict):
            raise TypeError("Expected argument 'errors' to be a dict")
        pulumi.set(__self__, "errors", errors)
        if health and not isinstance(health, str):
            raise TypeError("Expected argument 'health' to be a str")
        pulumi.set(__self__, "health", health)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if managed_resource_group_configuration and not isinstance(managed_resource_group_configuration, dict):
            raise TypeError("Expected argument 'managed_resource_group_configuration' to be a dict")
        pulumi.set(__self__, "managed_resource_group_configuration", managed_resource_group_configuration)
        if managed_resources_network_access_type and not isinstance(managed_resources_network_access_type, str):
            raise TypeError("Expected argument 'managed_resources_network_access_type' to be a str")
        pulumi.set(__self__, "managed_resources_network_access_type", managed_resources_network_access_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sap_product and not isinstance(sap_product, str):
            raise TypeError("Expected argument 'sap_product' to be a str")
        pulumi.set(__self__, "sap_product", sap_product)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def configuration(self) -> Any:
        """
        Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def environment(self) -> builtins.str:
        """
        Defines the environment type - Production/Non Production.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def errors(self) -> 'outputs.SAPVirtualInstanceErrorResponse':
        """
        Indicates any errors on the Virtual Instance for SAP solutions resource.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def health(self) -> builtins.str:
        """
        Defines the health of SAP Instances.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.SAPVirtualInstanceIdentityResponse']:
        """
        The managed service identities assigned to this resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedResourceGroupConfiguration")
    def managed_resource_group_configuration(self) -> Optional['outputs.ManagedRGConfigurationResponse']:
        """
        Managed resource group configuration
        """
        return pulumi.get(self, "managed_resource_group_configuration")

    @property
    @pulumi.getter(name="managedResourcesNetworkAccessType")
    def managed_resources_network_access_type(self) -> Optional[builtins.str]:
        """
        Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
        """
        return pulumi.get(self, "managed_resources_network_access_type")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Defines the provisioning states.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sapProduct")
    def sap_product(self) -> builtins.str:
        """
        Defines the SAP Product type.
        """
        return pulumi.get(self, "sap_product")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        """
        Defines the Virtual Instance for SAP state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Defines the SAP Instance status.
        """
        return pulumi.get(self, "status")

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
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetSapVirtualInstanceResult(GetSapVirtualInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSapVirtualInstanceResult(
            azure_api_version=self.azure_api_version,
            configuration=self.configuration,
            environment=self.environment,
            errors=self.errors,
            health=self.health,
            id=self.id,
            identity=self.identity,
            location=self.location,
            managed_resource_group_configuration=self.managed_resource_group_configuration,
            managed_resources_network_access_type=self.managed_resources_network_access_type,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sap_product=self.sap_product,
            state=self.state,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_sap_virtual_instance(resource_group_name: Optional[builtins.str] = None,
                             sap_virtual_instance_name: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSapVirtualInstanceResult:
    """
    Gets a Virtual Instance for SAP solutions resource

    Uses Azure REST API version 2024-09-01.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str sap_virtual_instance_name: The name of the Virtual Instances for SAP solutions resource
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['sapVirtualInstanceName'] = sap_virtual_instance_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:workloads:getSapVirtualInstance', __args__, opts=opts, typ=GetSapVirtualInstanceResult).value

    return AwaitableGetSapVirtualInstanceResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        configuration=pulumi.get(__ret__, 'configuration'),
        environment=pulumi.get(__ret__, 'environment'),
        errors=pulumi.get(__ret__, 'errors'),
        health=pulumi.get(__ret__, 'health'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        managed_resource_group_configuration=pulumi.get(__ret__, 'managed_resource_group_configuration'),
        managed_resources_network_access_type=pulumi.get(__ret__, 'managed_resources_network_access_type'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sap_product=pulumi.get(__ret__, 'sap_product'),
        state=pulumi.get(__ret__, 'state'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_sap_virtual_instance_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                    sap_virtual_instance_name: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSapVirtualInstanceResult]:
    """
    Gets a Virtual Instance for SAP solutions resource

    Uses Azure REST API version 2024-09-01.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str sap_virtual_instance_name: The name of the Virtual Instances for SAP solutions resource
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['sapVirtualInstanceName'] = sap_virtual_instance_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:workloads:getSapVirtualInstance', __args__, opts=opts, typ=GetSapVirtualInstanceResult)
    return __ret__.apply(lambda __response__: GetSapVirtualInstanceResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        configuration=pulumi.get(__response__, 'configuration'),
        environment=pulumi.get(__response__, 'environment'),
        errors=pulumi.get(__response__, 'errors'),
        health=pulumi.get(__response__, 'health'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        managed_resource_group_configuration=pulumi.get(__response__, 'managed_resource_group_configuration'),
        managed_resources_network_access_type=pulumi.get(__response__, 'managed_resources_network_access_type'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sap_product=pulumi.get(__response__, 'sap_product'),
        state=pulumi.get(__response__, 'state'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
