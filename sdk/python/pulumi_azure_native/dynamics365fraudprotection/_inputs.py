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
    'DFPInstanceAdministratorsArgs',
    'DFPInstanceAdministratorsArgsDict',
]

MYPY = False

if not MYPY:
    class DFPInstanceAdministratorsArgsDict(TypedDict):
        """
        An array of administrator user identities
        """
        members: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        An array of administrator user identities.
        """
elif False:
    DFPInstanceAdministratorsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DFPInstanceAdministratorsArgs:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        An array of administrator user identities
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] members: An array of administrator user identities.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        An array of administrator user identities.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "members", value)


