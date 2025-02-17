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
from ._enums import *

__all__ = [
    'IdentityArgs',
    'IdentityArgsDict',
    'PolicyDefinitionReferenceArgs',
    'PolicyDefinitionReferenceArgsDict',
    'PolicySkuArgs',
    'PolicySkuArgsDict',
]

MYPY = False

if not MYPY:
    class IdentityArgsDict(TypedDict):
        """
        Identity for the resource.
        """
        type: NotRequired[pulumi.Input['ResourceIdentityType']]
        """
        The identity type.
        """
elif False:
    IdentityArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input['ResourceIdentityType']] = None):
        """
        Identity for the resource.
        :param pulumi.Input['ResourceIdentityType'] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['ResourceIdentityType']]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['ResourceIdentityType']]):
        pulumi.set(self, "type", value)


if not MYPY:
    class PolicyDefinitionReferenceArgsDict(TypedDict):
        """
        The policy definition reference.
        """
        parameters: NotRequired[Any]
        """
        Required if a parameter is used in policy rule.
        """
        policy_definition_id: NotRequired[pulumi.Input[str]]
        """
        The ID of the policy definition or policy set definition.
        """
elif False:
    PolicyDefinitionReferenceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PolicyDefinitionReferenceArgs:
    def __init__(__self__, *,
                 parameters: Optional[Any] = None,
                 policy_definition_id: Optional[pulumi.Input[str]] = None):
        """
        The policy definition reference.
        :param Any parameters: Required if a parameter is used in policy rule.
        :param pulumi.Input[str] policy_definition_id: The ID of the policy definition or policy set definition.
        """
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_definition_id is not None:
            pulumi.set(__self__, "policy_definition_id", policy_definition_id)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        """
        Required if a parameter is used in policy rule.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[Any]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the policy definition or policy set definition.
        """
        return pulumi.get(self, "policy_definition_id")

    @policy_definition_id.setter
    def policy_definition_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_definition_id", value)


if not MYPY:
    class PolicySkuArgsDict(TypedDict):
        """
        The policy sku. This property is optional, obsolete, and will be ignored.
        """
        name: pulumi.Input[str]
        """
        The name of the policy sku. Possible values are A0 and A1.
        """
        tier: NotRequired[pulumi.Input[str]]
        """
        The policy sku tier. Possible values are Free and Standard.
        """
elif False:
    PolicySkuArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PolicySkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 tier: Optional[pulumi.Input[str]] = None):
        """
        The policy sku. This property is optional, obsolete, and will be ignored.
        :param pulumi.Input[str] name: The name of the policy sku. Possible values are A0 and A1.
        :param pulumi.Input[str] tier: The policy sku tier. Possible values are Free and Standard.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the policy sku. Possible values are A0 and A1.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The policy sku tier. Possible values are Free and Standard.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


