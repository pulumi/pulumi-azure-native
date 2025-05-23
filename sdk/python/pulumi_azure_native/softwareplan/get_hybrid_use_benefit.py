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
    'GetHybridUseBenefitResult',
    'AwaitableGetHybridUseBenefitResult',
    'get_hybrid_use_benefit',
    'get_hybrid_use_benefit_output',
]

@pulumi.output_type
class GetHybridUseBenefitResult:
    """
    Response on GET of a hybrid use benefit
    """
    def __init__(__self__, azure_api_version=None, created_date=None, etag=None, id=None, last_updated_date=None, name=None, provisioning_state=None, sku=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if etag and not isinstance(etag, int):
            raise TypeError("Expected argument 'etag' to be a int")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_updated_date and not isinstance(last_updated_date, str):
            raise TypeError("Expected argument 'last_updated_date' to be a str")
        pulumi.set(__self__, "last_updated_date", last_updated_date)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
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
    @pulumi.getter(name="createdDate")
    def created_date(self) -> builtins.str:
        """
        Created date
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def etag(self) -> builtins.int:
        """
        Indicates the revision of the hybrid use benefit
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> builtins.str:
        """
        Last updated date
        """
        return pulumi.get(self, "last_updated_date")

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
        Provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        Hybrid use benefit SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetHybridUseBenefitResult(GetHybridUseBenefitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHybridUseBenefitResult(
            azure_api_version=self.azure_api_version,
            created_date=self.created_date,
            etag=self.etag,
            id=self.id,
            last_updated_date=self.last_updated_date,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            type=self.type)


def get_hybrid_use_benefit(plan_id: Optional[builtins.str] = None,
                           scope: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHybridUseBenefitResult:
    """
    Gets a given plan ID

    Uses Azure REST API version 2019-12-01.


    :param builtins.str plan_id: This is a unique identifier for a plan. Should be a guid.
    :param builtins.str scope: The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
    """
    __args__ = dict()
    __args__['planId'] = plan_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:softwareplan:getHybridUseBenefit', __args__, opts=opts, typ=GetHybridUseBenefitResult).value

    return AwaitableGetHybridUseBenefitResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created_date=pulumi.get(__ret__, 'created_date'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        last_updated_date=pulumi.get(__ret__, 'last_updated_date'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sku=pulumi.get(__ret__, 'sku'),
        type=pulumi.get(__ret__, 'type'))
def get_hybrid_use_benefit_output(plan_id: Optional[pulumi.Input[builtins.str]] = None,
                                  scope: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetHybridUseBenefitResult]:
    """
    Gets a given plan ID

    Uses Azure REST API version 2019-12-01.


    :param builtins.str plan_id: This is a unique identifier for a plan. Should be a guid.
    :param builtins.str scope: The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
    """
    __args__ = dict()
    __args__['planId'] = plan_id
    __args__['scope'] = scope
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:softwareplan:getHybridUseBenefit', __args__, opts=opts, typ=GetHybridUseBenefitResult)
    return __ret__.apply(lambda __response__: GetHybridUseBenefitResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created_date=pulumi.get(__response__, 'created_date'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        last_updated_date=pulumi.get(__response__, 'last_updated_date'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sku=pulumi.get(__response__, 'sku'),
        type=pulumi.get(__response__, 'type')))
