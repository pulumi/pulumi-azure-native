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
    'GetTagByOperationResult',
    'AwaitableGetTagByOperationResult',
    'get_tag_by_operation',
    'get_tag_by_operation_output',
]

@pulumi.output_type
class GetTagByOperationResult:
    """
    Tag Contract details.
    """
    def __init__(__self__, azure_api_version=None, display_name=None, id=None, name=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="displayName")
    def display_name(self) -> builtins.str:
        """
        Tag name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetTagByOperationResult(GetTagByOperationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagByOperationResult(
            azure_api_version=self.azure_api_version,
            display_name=self.display_name,
            id=self.id,
            name=self.name,
            type=self.type)


def get_tag_by_operation(api_id: Optional[builtins.str] = None,
                         operation_id: Optional[builtins.str] = None,
                         resource_group_name: Optional[builtins.str] = None,
                         service_name: Optional[builtins.str] = None,
                         tag_id: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagByOperationResult:
    """
    Get tag associated with the Operation.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param builtins.str operation_id: Operation identifier within an API. Must be unique in the current API Management service instance.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    :param builtins.str tag_id: Tag identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['operationId'] = operation_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement:getTagByOperation', __args__, opts=opts, typ=GetTagByOperationResult).value

    return AwaitableGetTagByOperationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        display_name=pulumi.get(__ret__, 'display_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'))
def get_tag_by_operation_output(api_id: Optional[pulumi.Input[builtins.str]] = None,
                                operation_id: Optional[pulumi.Input[builtins.str]] = None,
                                resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                service_name: Optional[pulumi.Input[builtins.str]] = None,
                                tag_id: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTagByOperationResult]:
    """
    Get tag associated with the Operation.

    Uses Azure REST API version 2022-09-01-preview.

    Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param builtins.str operation_id: Operation identifier within an API. Must be unique in the current API Management service instance.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str service_name: The name of the API Management service.
    :param builtins.str tag_id: Tag identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['operationId'] = operation_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement:getTagByOperation', __args__, opts=opts, typ=GetTagByOperationResult)
    return __ret__.apply(lambda __response__: GetTagByOperationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        display_name=pulumi.get(__response__, 'display_name'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type')))
