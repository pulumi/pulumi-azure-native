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
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    Represents a Database.
    """
    def __init__(__self__, azure_api_version=None, charset=None, collation=None, id=None, name=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if charset and not isinstance(charset, str):
            raise TypeError("Expected argument 'charset' to be a str")
        pulumi.set(__self__, "charset", charset)
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter
    def charset(self) -> Optional[builtins.str]:
        """
        Character set of the database.
        """
        return pulumi.get(self, "charset")

    @property
    @pulumi.getter
    def collation(self) -> Optional[builtins.str]:
        """
        Collation of the database.
        """
        return pulumi.get(self, "collation")

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


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            azure_api_version=self.azure_api_version,
            charset=self.charset,
            collation=self.collation,
            id=self.id,
            name=self.name,
            system_data=self.system_data,
            type=self.type)


def get_database(database_name: Optional[builtins.str] = None,
                 resource_group_name: Optional[builtins.str] = None,
                 server_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Gets information about an existing database.

    Uses Azure REST API version 2024-08-01.

    Other available API versions: 2022-12-01, 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: Name of the database.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:dbforpostgresql:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        charset=pulumi.get(__ret__, 'charset'),
        collation=pulumi.get(__ret__, 'collation'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_database_output(database_name: Optional[pulumi.Input[builtins.str]] = None,
                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        server_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Gets information about an existing database.

    Uses Azure REST API version 2024-08-01.

    Other available API versions: 2022-12-01, 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: Name of the database.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:dbforpostgresql:getDatabase', __args__, opts=opts, typ=GetDatabaseResult)
    return __ret__.apply(lambda __response__: GetDatabaseResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        charset=pulumi.get(__response__, 'charset'),
        collation=pulumi.get(__response__, 'collation'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
