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
    'GetTrustedIdProviderResult',
    'AwaitableGetTrustedIdProviderResult',
    'get_trusted_id_provider',
    'get_trusted_id_provider_output',
]

@pulumi.output_type
class GetTrustedIdProviderResult:
    """
    Data Lake Store trusted identity provider information.
    """
    def __init__(__self__, azure_api_version=None, id=None, id_provider=None, name=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if id_provider and not isinstance(id_provider, str):
            raise TypeError("Expected argument 'id_provider' to be a str")
        pulumi.set(__self__, "id_provider", id_provider)
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
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idProvider")
    def id_provider(self) -> builtins.str:
        """
        The URL of this trusted identity provider.
        """
        return pulumi.get(self, "id_provider")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetTrustedIdProviderResult(GetTrustedIdProviderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrustedIdProviderResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            id_provider=self.id_provider,
            name=self.name,
            type=self.type)


def get_trusted_id_provider(account_name: Optional[builtins.str] = None,
                            resource_group_name: Optional[builtins.str] = None,
                            trusted_id_provider_name: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrustedIdProviderResult:
    """
    Gets the specified Data Lake Store trusted identity provider.

    Uses Azure REST API version 2016-11-01.


    :param builtins.str account_name: The name of the Data Lake Store account.
    :param builtins.str resource_group_name: The name of the Azure resource group.
    :param builtins.str trusted_id_provider_name: The name of the trusted identity provider to retrieve.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['trustedIdProviderName'] = trusted_id_provider_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:datalakestore:getTrustedIdProvider', __args__, opts=opts, typ=GetTrustedIdProviderResult).value

    return AwaitableGetTrustedIdProviderResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        id_provider=pulumi.get(__ret__, 'id_provider'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'))
def get_trusted_id_provider_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                   trusted_id_provider_name: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTrustedIdProviderResult]:
    """
    Gets the specified Data Lake Store trusted identity provider.

    Uses Azure REST API version 2016-11-01.


    :param builtins.str account_name: The name of the Data Lake Store account.
    :param builtins.str resource_group_name: The name of the Azure resource group.
    :param builtins.str trusted_id_provider_name: The name of the trusted identity provider to retrieve.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['trustedIdProviderName'] = trusted_id_provider_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:datalakestore:getTrustedIdProvider', __args__, opts=opts, typ=GetTrustedIdProviderResult)
    return __ret__.apply(lambda __response__: GetTrustedIdProviderResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        id_provider=pulumi.get(__response__, 'id_provider'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type')))
