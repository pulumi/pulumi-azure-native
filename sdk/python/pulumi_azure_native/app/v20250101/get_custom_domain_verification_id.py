# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ... import _utilities

__all__ = [
    'GetCustomDomainVerificationIdResult',
    'AwaitableGetCustomDomainVerificationIdResult',
    'get_custom_domain_verification_id',
    'get_custom_domain_verification_id_output',
]

@pulumi.output_type
class GetCustomDomainVerificationIdResult:
    """
    Custom domain verification Id of a subscription
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


class AwaitableGetCustomDomainVerificationIdResult(GetCustomDomainVerificationIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomDomainVerificationIdResult(
            value=self.value)


def get_custom_domain_verification_id(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomDomainVerificationIdResult:
    """
    Get the verification id of a subscription used for verifying custom domains
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:app/v20250101:getCustomDomainVerificationId', __args__, opts=opts, typ=GetCustomDomainVerificationIdResult).value

    return AwaitableGetCustomDomainVerificationIdResult(
        value=pulumi.get(__ret__, 'value'))
def get_custom_domain_verification_id_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCustomDomainVerificationIdResult]:
    """
    Get the verification id of a subscription used for verifying custom domains
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:app/v20250101:getCustomDomainVerificationId', __args__, opts=opts, typ=GetCustomDomainVerificationIdResult)
    return __ret__.apply(lambda __response__: GetCustomDomainVerificationIdResult(
        value=pulumi.get(__response__, 'value')))
