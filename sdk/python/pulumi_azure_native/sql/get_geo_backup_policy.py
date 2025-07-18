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

__all__ = [
    'GetGeoBackupPolicyResult',
    'AwaitableGetGeoBackupPolicyResult',
    'get_geo_backup_policy',
    'get_geo_backup_policy_output',
]

@pulumi.output_type
class GetGeoBackupPolicyResult:
    """
    A Geo backup policy.
    """
    def __init__(__self__, azure_api_version=None, id=None, kind=None, location=None, name=None, state=None, storage_type=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if storage_type and not isinstance(storage_type, str):
            raise TypeError("Expected argument 'storage_type' to be a str")
        pulumi.set(__self__, "storage_type", storage_type)
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
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> builtins.str:
        """
        Kind of geo backup policy.  This is metadata used for the Azure portal experience.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Backup policy location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        """
        The state of the geo backup policy.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> builtins.str:
        """
        The storage type of the geo backup policy.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetGeoBackupPolicyResult(GetGeoBackupPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGeoBackupPolicyResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            kind=self.kind,
            location=self.location,
            name=self.name,
            state=self.state,
            storage_type=self.storage_type,
            type=self.type)


def get_geo_backup_policy(database_name: Optional[builtins.str] = None,
                          geo_backup_policy_name: Optional[builtins.str] = None,
                          resource_group_name: Optional[builtins.str] = None,
                          server_name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGeoBackupPolicyResult:
    """
    Gets a Geo backup policy for the given database resource.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str geo_backup_policy_name: The name of the Geo backup policy. This should always be 'Default'.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['geoBackupPolicyName'] = geo_backup_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getGeoBackupPolicy', __args__, opts=opts, typ=GetGeoBackupPolicyResult).value

    return AwaitableGetGeoBackupPolicyResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        storage_type=pulumi.get(__ret__, 'storage_type'),
        type=pulumi.get(__ret__, 'type'))
def get_geo_backup_policy_output(database_name: Optional[pulumi.Input[builtins.str]] = None,
                                 geo_backup_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGeoBackupPolicyResult]:
    """
    Gets a Geo backup policy for the given database resource.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str geo_backup_policy_name: The name of the Geo backup policy. This should always be 'Default'.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['geoBackupPolicyName'] = geo_backup_policy_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getGeoBackupPolicy', __args__, opts=opts, typ=GetGeoBackupPolicyResult)
    return __ret__.apply(lambda __response__: GetGeoBackupPolicyResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        state=pulumi.get(__response__, 'state'),
        storage_type=pulumi.get(__response__, 'storage_type'),
        type=pulumi.get(__response__, 'type')))
