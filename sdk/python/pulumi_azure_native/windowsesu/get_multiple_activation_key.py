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
    'GetMultipleActivationKeyResult',
    'AwaitableGetMultipleActivationKeyResult',
    'get_multiple_activation_key',
    'get_multiple_activation_key_output',
]

@pulumi.output_type
class GetMultipleActivationKeyResult:
    """
    MAK key details.
    """
    def __init__(__self__, agreement_number=None, azure_api_version=None, expiration_date=None, id=None, installed_server_number=None, is_eligible=None, location=None, multiple_activation_key=None, name=None, os_type=None, provisioning_state=None, support_type=None, tags=None, type=None):
        if agreement_number and not isinstance(agreement_number, str):
            raise TypeError("Expected argument 'agreement_number' to be a str")
        pulumi.set(__self__, "agreement_number", agreement_number)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if expiration_date and not isinstance(expiration_date, str):
            raise TypeError("Expected argument 'expiration_date' to be a str")
        pulumi.set(__self__, "expiration_date", expiration_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if installed_server_number and not isinstance(installed_server_number, int):
            raise TypeError("Expected argument 'installed_server_number' to be a int")
        pulumi.set(__self__, "installed_server_number", installed_server_number)
        if is_eligible and not isinstance(is_eligible, bool):
            raise TypeError("Expected argument 'is_eligible' to be a bool")
        pulumi.set(__self__, "is_eligible", is_eligible)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if multiple_activation_key and not isinstance(multiple_activation_key, str):
            raise TypeError("Expected argument 'multiple_activation_key' to be a str")
        pulumi.set(__self__, "multiple_activation_key", multiple_activation_key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if support_type and not isinstance(support_type, str):
            raise TypeError("Expected argument 'support_type' to be a str")
        pulumi.set(__self__, "support_type", support_type)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="agreementNumber")
    def agreement_number(self) -> Optional[builtins.str]:
        """
        Agreement number under which the key is requested.
        """
        return pulumi.get(self, "agreement_number")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> builtins.str:
        """
        End of support of security updates activated by the MAK key.
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="installedServerNumber")
    def installed_server_number(self) -> Optional[builtins.int]:
        """
        Number of activations/servers using the MAK key.
        """
        return pulumi.get(self, "installed_server_number")

    @property
    @pulumi.getter(name="isEligible")
    def is_eligible(self) -> Optional[builtins.bool]:
        """
        <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
        """
        return pulumi.get(self, "is_eligible")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="multipleActivationKey")
    def multiple_activation_key(self) -> builtins.str:
        """
        MAK 5x5 key.
        """
        return pulumi.get(self, "multiple_activation_key")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[builtins.str]:
        """
        Type of OS for which the key is requested.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="supportType")
    def support_type(self) -> Optional[builtins.str]:
        """
        Type of support
        """
        return pulumi.get(self, "support_type")

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


class AwaitableGetMultipleActivationKeyResult(GetMultipleActivationKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMultipleActivationKeyResult(
            agreement_number=self.agreement_number,
            azure_api_version=self.azure_api_version,
            expiration_date=self.expiration_date,
            id=self.id,
            installed_server_number=self.installed_server_number,
            is_eligible=self.is_eligible,
            location=self.location,
            multiple_activation_key=self.multiple_activation_key,
            name=self.name,
            os_type=self.os_type,
            provisioning_state=self.provisioning_state,
            support_type=self.support_type,
            tags=self.tags,
            type=self.type)


def get_multiple_activation_key(multiple_activation_key_name: Optional[builtins.str] = None,
                                resource_group_name: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMultipleActivationKeyResult:
    """
    Get a MAK key.

    Uses Azure REST API version 2019-09-16-preview.


    :param builtins.str multiple_activation_key_name: The name of the MAK key.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['multipleActivationKeyName'] = multiple_activation_key_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:windowsesu:getMultipleActivationKey', __args__, opts=opts, typ=GetMultipleActivationKeyResult).value

    return AwaitableGetMultipleActivationKeyResult(
        agreement_number=pulumi.get(__ret__, 'agreement_number'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        expiration_date=pulumi.get(__ret__, 'expiration_date'),
        id=pulumi.get(__ret__, 'id'),
        installed_server_number=pulumi.get(__ret__, 'installed_server_number'),
        is_eligible=pulumi.get(__ret__, 'is_eligible'),
        location=pulumi.get(__ret__, 'location'),
        multiple_activation_key=pulumi.get(__ret__, 'multiple_activation_key'),
        name=pulumi.get(__ret__, 'name'),
        os_type=pulumi.get(__ret__, 'os_type'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        support_type=pulumi.get(__ret__, 'support_type'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_multiple_activation_key_output(multiple_activation_key_name: Optional[pulumi.Input[builtins.str]] = None,
                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMultipleActivationKeyResult]:
    """
    Get a MAK key.

    Uses Azure REST API version 2019-09-16-preview.


    :param builtins.str multiple_activation_key_name: The name of the MAK key.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['multipleActivationKeyName'] = multiple_activation_key_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:windowsesu:getMultipleActivationKey', __args__, opts=opts, typ=GetMultipleActivationKeyResult)
    return __ret__.apply(lambda __response__: GetMultipleActivationKeyResult(
        agreement_number=pulumi.get(__response__, 'agreement_number'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        expiration_date=pulumi.get(__response__, 'expiration_date'),
        id=pulumi.get(__response__, 'id'),
        installed_server_number=pulumi.get(__response__, 'installed_server_number'),
        is_eligible=pulumi.get(__response__, 'is_eligible'),
        location=pulumi.get(__response__, 'location'),
        multiple_activation_key=pulumi.get(__response__, 'multiple_activation_key'),
        name=pulumi.get(__response__, 'name'),
        os_type=pulumi.get(__response__, 'os_type'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        support_type=pulumi.get(__response__, 'support_type'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
