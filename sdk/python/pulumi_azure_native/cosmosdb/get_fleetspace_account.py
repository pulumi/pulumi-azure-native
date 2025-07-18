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
    'GetFleetspaceAccountResult',
    'AwaitableGetFleetspaceAccountResult',
    'get_fleetspace_account',
    'get_fleetspace_account_output',
]

@pulumi.output_type
class GetFleetspaceAccountResult:
    """
    An Azure Cosmos DB Fleetspace Account
    """
    def __init__(__self__, azure_api_version=None, global_database_account_properties=None, id=None, name=None, provisioning_state=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if global_database_account_properties and not isinstance(global_database_account_properties, dict):
            raise TypeError("Expected argument 'global_database_account_properties' to be a dict")
        pulumi.set(__self__, "global_database_account_properties", global_database_account_properties)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
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
    @pulumi.getter(name="globalDatabaseAccountProperties")
    def global_database_account_properties(self) -> Optional['outputs.FleetspaceAccountPropertiesResponseGlobalDatabaseAccountProperties']:
        """
        Configuration for fleetspace Account in the fleetspace.
        """
        return pulumi.get(self, "global_database_account_properties")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

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
        A provisioning state of the Fleetspace Account.
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
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetFleetspaceAccountResult(GetFleetspaceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFleetspaceAccountResult(
            azure_api_version=self.azure_api_version,
            global_database_account_properties=self.global_database_account_properties,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            type=self.type)


def get_fleetspace_account(fleet_name: Optional[builtins.str] = None,
                           fleetspace_account_name: Optional[builtins.str] = None,
                           fleetspace_name: Optional[builtins.str] = None,
                           resource_group_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFleetspaceAccountResult:
    """
    Retrieves the properties of an existing Azure Cosmos DB fleetspace account under a fleetspace

    Uses Azure REST API version 2025-05-01-preview.


    :param builtins.str fleet_name: Cosmos DB fleet name. Needs to be unique under a subscription.
    :param builtins.str fleetspace_account_name: Cosmos DB fleetspace account name.
    :param builtins.str fleetspace_name: Cosmos DB fleetspace name. Needs to be unique under a fleet.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['fleetName'] = fleet_name
    __args__['fleetspaceAccountName'] = fleetspace_account_name
    __args__['fleetspaceName'] = fleetspace_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cosmosdb:getFleetspaceAccount', __args__, opts=opts, typ=GetFleetspaceAccountResult).value

    return AwaitableGetFleetspaceAccountResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        global_database_account_properties=pulumi.get(__ret__, 'global_database_account_properties'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_fleetspace_account_output(fleet_name: Optional[pulumi.Input[builtins.str]] = None,
                                  fleetspace_account_name: Optional[pulumi.Input[builtins.str]] = None,
                                  fleetspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                  resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFleetspaceAccountResult]:
    """
    Retrieves the properties of an existing Azure Cosmos DB fleetspace account under a fleetspace

    Uses Azure REST API version 2025-05-01-preview.


    :param builtins.str fleet_name: Cosmos DB fleet name. Needs to be unique under a subscription.
    :param builtins.str fleetspace_account_name: Cosmos DB fleetspace account name.
    :param builtins.str fleetspace_name: Cosmos DB fleetspace name. Needs to be unique under a fleet.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['fleetName'] = fleet_name
    __args__['fleetspaceAccountName'] = fleetspace_account_name
    __args__['fleetspaceName'] = fleetspace_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cosmosdb:getFleetspaceAccount', __args__, opts=opts, typ=GetFleetspaceAccountResult)
    return __ret__.apply(lambda __response__: GetFleetspaceAccountResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        global_database_account_properties=pulumi.get(__response__, 'global_database_account_properties'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
