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
    'GetWorkloadGroupResult',
    'AwaitableGetWorkloadGroupResult',
    'get_workload_group',
    'get_workload_group_output',
]

@pulumi.output_type
class GetWorkloadGroupResult:
    """
    Workload group operations for a data warehouse
    """
    def __init__(__self__, azure_api_version=None, id=None, importance=None, max_resource_percent=None, max_resource_percent_per_request=None, min_resource_percent=None, min_resource_percent_per_request=None, name=None, query_execution_timeout=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if importance and not isinstance(importance, str):
            raise TypeError("Expected argument 'importance' to be a str")
        pulumi.set(__self__, "importance", importance)
        if max_resource_percent and not isinstance(max_resource_percent, int):
            raise TypeError("Expected argument 'max_resource_percent' to be a int")
        pulumi.set(__self__, "max_resource_percent", max_resource_percent)
        if max_resource_percent_per_request and not isinstance(max_resource_percent_per_request, float):
            raise TypeError("Expected argument 'max_resource_percent_per_request' to be a float")
        pulumi.set(__self__, "max_resource_percent_per_request", max_resource_percent_per_request)
        if min_resource_percent and not isinstance(min_resource_percent, int):
            raise TypeError("Expected argument 'min_resource_percent' to be a int")
        pulumi.set(__self__, "min_resource_percent", min_resource_percent)
        if min_resource_percent_per_request and not isinstance(min_resource_percent_per_request, float):
            raise TypeError("Expected argument 'min_resource_percent_per_request' to be a float")
        pulumi.set(__self__, "min_resource_percent_per_request", min_resource_percent_per_request)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query_execution_timeout and not isinstance(query_execution_timeout, int):
            raise TypeError("Expected argument 'query_execution_timeout' to be a int")
        pulumi.set(__self__, "query_execution_timeout", query_execution_timeout)
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
    def importance(self) -> Optional[builtins.str]:
        """
        The workload group importance level.
        """
        return pulumi.get(self, "importance")

    @property
    @pulumi.getter(name="maxResourcePercent")
    def max_resource_percent(self) -> builtins.int:
        """
        The workload group cap percentage resource.
        """
        return pulumi.get(self, "max_resource_percent")

    @property
    @pulumi.getter(name="maxResourcePercentPerRequest")
    def max_resource_percent_per_request(self) -> Optional[builtins.float]:
        """
        The workload group request maximum grant percentage.
        """
        return pulumi.get(self, "max_resource_percent_per_request")

    @property
    @pulumi.getter(name="minResourcePercent")
    def min_resource_percent(self) -> builtins.int:
        """
        The workload group minimum percentage resource.
        """
        return pulumi.get(self, "min_resource_percent")

    @property
    @pulumi.getter(name="minResourcePercentPerRequest")
    def min_resource_percent_per_request(self) -> builtins.float:
        """
        The workload group request minimum grant percentage.
        """
        return pulumi.get(self, "min_resource_percent_per_request")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryExecutionTimeout")
    def query_execution_timeout(self) -> Optional[builtins.int]:
        """
        The workload group query execution timeout.
        """
        return pulumi.get(self, "query_execution_timeout")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetWorkloadGroupResult(GetWorkloadGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkloadGroupResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            importance=self.importance,
            max_resource_percent=self.max_resource_percent,
            max_resource_percent_per_request=self.max_resource_percent_per_request,
            min_resource_percent=self.min_resource_percent,
            min_resource_percent_per_request=self.min_resource_percent_per_request,
            name=self.name,
            query_execution_timeout=self.query_execution_timeout,
            type=self.type)


def get_workload_group(database_name: Optional[builtins.str] = None,
                       resource_group_name: Optional[builtins.str] = None,
                       server_name: Optional[builtins.str] = None,
                       workload_group_name: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkloadGroupResult:
    """
    Gets a workload group

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    :param builtins.str workload_group_name: The name of the workload group.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['workloadGroupName'] = workload_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getWorkloadGroup', __args__, opts=opts, typ=GetWorkloadGroupResult).value

    return AwaitableGetWorkloadGroupResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        importance=pulumi.get(__ret__, 'importance'),
        max_resource_percent=pulumi.get(__ret__, 'max_resource_percent'),
        max_resource_percent_per_request=pulumi.get(__ret__, 'max_resource_percent_per_request'),
        min_resource_percent=pulumi.get(__ret__, 'min_resource_percent'),
        min_resource_percent_per_request=pulumi.get(__ret__, 'min_resource_percent_per_request'),
        name=pulumi.get(__ret__, 'name'),
        query_execution_timeout=pulumi.get(__ret__, 'query_execution_timeout'),
        type=pulumi.get(__ret__, 'type'))
def get_workload_group_output(database_name: Optional[pulumi.Input[builtins.str]] = None,
                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              server_name: Optional[pulumi.Input[builtins.str]] = None,
                              workload_group_name: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkloadGroupResult]:
    """
    Gets a workload group

    Uses Azure REST API version 2023-08-01.

    Other available API versions: 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str database_name: The name of the database.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str server_name: The name of the server.
    :param builtins.str workload_group_name: The name of the workload group.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['workloadGroupName'] = workload_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getWorkloadGroup', __args__, opts=opts, typ=GetWorkloadGroupResult)
    return __ret__.apply(lambda __response__: GetWorkloadGroupResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        importance=pulumi.get(__response__, 'importance'),
        max_resource_percent=pulumi.get(__response__, 'max_resource_percent'),
        max_resource_percent_per_request=pulumi.get(__response__, 'max_resource_percent_per_request'),
        min_resource_percent=pulumi.get(__response__, 'min_resource_percent'),
        min_resource_percent_per_request=pulumi.get(__response__, 'min_resource_percent_per_request'),
        name=pulumi.get(__response__, 'name'),
        query_execution_timeout=pulumi.get(__response__, 'query_execution_timeout'),
        type=pulumi.get(__response__, 'type')))
