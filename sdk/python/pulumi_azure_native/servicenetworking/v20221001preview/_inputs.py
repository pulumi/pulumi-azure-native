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
    'AssociationSubnetArgs',
    'AssociationSubnetArgsDict',
    'FrontendPropertiesIPAddressArgs',
    'FrontendPropertiesIPAddressArgsDict',
]

MYPY = False

if not MYPY:
    class AssociationSubnetArgsDict(TypedDict):
        """
        Association Subnet.
        """
        id: pulumi.Input[str]
        """
        Association ID.
        """
elif False:
    AssociationSubnetArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AssociationSubnetArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str]):
        """
        Association Subnet.
        :param pulumi.Input[str] id: Association ID.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Association ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


if not MYPY:
    class FrontendPropertiesIPAddressArgsDict(TypedDict):
        """
        Frontend IP Address.
        """
        id: pulumi.Input[str]
        """
        IP Address.
        """
elif False:
    FrontendPropertiesIPAddressArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FrontendPropertiesIPAddressArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str]):
        """
        Frontend IP Address.
        :param pulumi.Input[str] id: IP Address.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        IP Address.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


