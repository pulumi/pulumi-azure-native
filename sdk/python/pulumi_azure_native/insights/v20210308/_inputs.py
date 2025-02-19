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
    'MyWorkbookManagedIdentityArgs',
    'MyWorkbookManagedIdentityArgsDict',
    'WorkbookManagedIdentityArgs',
    'WorkbookManagedIdentityArgsDict',
]

MYPY = False

if not MYPY:
    class MyWorkbookManagedIdentityArgsDict(TypedDict):
        """
        Customer Managed Identity
        """
        type: NotRequired[pulumi.Input[str]]
        """
        The identity type.
        """
elif False:
    MyWorkbookManagedIdentityArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MyWorkbookManagedIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Customer Managed Identity
        :param pulumi.Input[str] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


if not MYPY:
    class WorkbookManagedIdentityArgsDict(TypedDict):
        """
        Customer Managed Identity
        """
        type: NotRequired[pulumi.Input[str]]
        """
        The identity type.
        """
elif False:
    WorkbookManagedIdentityArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class WorkbookManagedIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Customer Managed Identity
        :param pulumi.Input[str] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


