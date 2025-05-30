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
    'AccountResourcePropertiesArgs',
    'AccountResourcePropertiesArgsDict',
]

MYPY = False

if not MYPY:
    class AccountResourcePropertiesArgsDict(TypedDict):
        """
        Property bag from billing account
        """
        app_id: pulumi.Input[builtins.str]
        """
        Customer owned application ID
        """
elif False:
    AccountResourcePropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AccountResourcePropertiesArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[builtins.str]):
        """
        Property bag from billing account
        :param pulumi.Input[builtins.str] app_id: Customer owned application ID
        """
        pulumi.set(__self__, "app_id", app_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[builtins.str]:
        """
        Customer owned application ID
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "app_id", value)


