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
from ._enums import *

__all__ = [
    'SkuArgs',
    'SkuArgsDict',
]

MYPY = False

if not MYPY:
    class SkuArgsDict(TypedDict):
        """
        Model representing SKU for Azure Dev Spaces Controller.
        """
        name: pulumi.Input[Union[builtins.str, 'SkuName']]
        """
        The name of the SKU for Azure Dev Spaces Controller.
        """
        tier: NotRequired[pulumi.Input[Union[builtins.str, 'SkuTier']]]
        """
        The tier of the SKU for Azure Dev Spaces Controller.
        """
elif False:
    SkuArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[Union[builtins.str, 'SkuName']],
                 tier: Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]] = None):
        """
        Model representing SKU for Azure Dev Spaces Controller.
        :param pulumi.Input[Union[builtins.str, 'SkuName']] name: The name of the SKU for Azure Dev Spaces Controller.
        :param pulumi.Input[Union[builtins.str, 'SkuTier']] tier: The tier of the SKU for Azure Dev Spaces Controller.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[Union[builtins.str, 'SkuName']]:
        """
        The name of the SKU for Azure Dev Spaces Controller.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[Union[builtins.str, 'SkuName']]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]]:
        """
        The tier of the SKU for Azure Dev Spaces Controller.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]]):
        pulumi.set(self, "tier", value)


