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
    'GetFailoverGroupResult',
    'AwaitableGetFailoverGroupResult',
    'get_failover_group',
    'get_failover_group_output',
]

@pulumi.output_type
class GetFailoverGroupResult:
    """
    A failover group.
    """
    def __init__(__self__, azure_api_version=None, databases=None, id=None, location=None, name=None, partner_servers=None, read_only_endpoint=None, read_write_endpoint=None, replication_role=None, replication_state=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partner_servers and not isinstance(partner_servers, list):
            raise TypeError("Expected argument 'partner_servers' to be a list")
        pulumi.set(__self__, "partner_servers", partner_servers)
        if read_only_endpoint and not isinstance(read_only_endpoint, dict):
            raise TypeError("Expected argument 'read_only_endpoint' to be a dict")
        pulumi.set(__self__, "read_only_endpoint", read_only_endpoint)
        if read_write_endpoint and not isinstance(read_write_endpoint, dict):
            raise TypeError("Expected argument 'read_write_endpoint' to be a dict")
        pulumi.set(__self__, "read_write_endpoint", read_write_endpoint)
        if replication_role and not isinstance(replication_role, str):
            raise TypeError("Expected argument 'replication_role' to be a str")
        pulumi.set(__self__, "replication_role", replication_role)
        if replication_state and not isinstance(replication_state, str):
            raise TypeError("Expected argument 'replication_state' to be a str")
        pulumi.set(__self__, "replication_state", replication_state)
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
    def databases(self) -> Optional[Sequence[builtins.str]]:
        """
        List of databases in the failover group.
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource location.
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
    @pulumi.getter(name="partnerServers")
    def partner_servers(self) -> Sequence['outputs.PartnerInfoResponse']:
        """
        List of partner server information for the failover group.
        """
        return pulumi.get(self, "partner_servers")

    @property
    @pulumi.getter(name="readOnlyEndpoint")
    def read_only_endpoint(self) -> Optional['outputs.FailoverGroupReadOnlyEndpointResponse']:
        """
        Read-only endpoint of the failover group instance.
        """
        return pulumi.get(self, "read_only_endpoint")

    @property
    @pulumi.getter(name="readWriteEndpoint")
    def read_write_endpoint(self) -> 'outputs.FailoverGroupReadWriteEndpointResponse':
        """
        Read-write endpoint of the failover group instance.
        """
        return pulumi.get(self, "read_write_endpoint")

    @property
    @pulumi.getter(name="replicationRole")
    def replication_role(self) -> builtins.str:
        """
        Local replication role of the failover group instance.
        """
        return pulumi.get(self, "replication_role")

    @property
    @pulumi.getter(name="replicationState")
    def replication_state(self) -> builtins.str:
        """
        Replication state of the failover group instance.
        """
        return pulumi.get(self, "replication_state")

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
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetFailoverGroupResult(GetFailoverGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFailoverGroupResult(
            azure_api_version=self.azure_api_version,
            databases=self.databases,
            id=self.id,
            location=self.location,
            name=self.name,
            partner_servers=self.partner_servers,
            read_only_endpoint=self.read_only_endpoint,
            read_write_endpoint=self.read_write_endpoint,
            replication_role=self.replication_role,
            replication_state=self.replication_state,
            tags=self.tags,
            type=self.type)


def get_failover_group(failover_group_name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       server_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFailoverGroupResult:
    """
    Gets a failover group.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2015-05-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str failover_group_name: The name of the failover group.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server containing the failover group.
    """
    __args__ = dict()
    __args__['failoverGroupName'] = failover_group_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getFailoverGroup', __args__, opts=opts, typ=GetFailoverGroupResult).value

    return AwaitableGetFailoverGroupResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        databases=pulumi.get(__ret__, 'databases'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        partner_servers=pulumi.get(__ret__, 'partner_servers'),
        read_only_endpoint=pulumi.get(__ret__, 'read_only_endpoint'),
        read_write_endpoint=pulumi.get(__ret__, 'read_write_endpoint'),
        replication_role=pulumi.get(__ret__, 'replication_role'),
        replication_state=pulumi.get(__ret__, 'replication_state'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_failover_group_output(failover_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              server_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFailoverGroupResult]:
    """
    Gets a failover group.

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2015-05-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str failover_group_name: The name of the failover group.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server containing the failover group.
    """
    __args__ = dict()
    __args__['failoverGroupName'] = failover_group_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getFailoverGroup', __args__, opts=opts, typ=GetFailoverGroupResult)
    return __ret__.apply(lambda __response__: GetFailoverGroupResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        databases=pulumi.get(__response__, 'databases'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        partner_servers=pulumi.get(__response__, 'partner_servers'),
        read_only_endpoint=pulumi.get(__response__, 'read_only_endpoint'),
        read_write_endpoint=pulumi.get(__response__, 'read_write_endpoint'),
        replication_role=pulumi.get(__response__, 'replication_role'),
        replication_state=pulumi.get(__response__, 'replication_state'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
