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
    'StreamingLocatorUserDefinedContentKeyArgs',
    'StreamingLocatorUserDefinedContentKeyArgsDict',
]

MYPY = False

if not MYPY:
    class StreamingLocatorUserDefinedContentKeyArgsDict(TypedDict):
        """
        Describes the properties of a user-defined content key in the Streaming Locator
        """
        id: pulumi.Input[str]
        """
        ID of Content Key
        """
        label: NotRequired[pulumi.Input[str]]
        """
        The Content Key description
        """
        value: NotRequired[pulumi.Input[str]]
        """
        The Content Key secret
        """
elif False:
    StreamingLocatorUserDefinedContentKeyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StreamingLocatorUserDefinedContentKeyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 label: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Describes the properties of a user-defined content key in the Streaming Locator
        :param pulumi.Input[str] id: ID of Content Key
        :param pulumi.Input[str] label: The Content Key description
        :param pulumi.Input[str] value: The Content Key secret
        """
        pulumi.set(__self__, "id", id)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        ID of Content Key
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        """
        The Content Key description
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The Content Key secret
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


