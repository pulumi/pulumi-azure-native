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
    'PerimeterBasedAccessRuleResponse',
    'SubResourceResponse',
    'SubscriptionIdResponse',
]

@pulumi.output_type
class PerimeterBasedAccessRuleResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "perimeterGuid":
            suggest = "perimeter_guid"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PerimeterBasedAccessRuleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PerimeterBasedAccessRuleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PerimeterBasedAccessRuleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 location: str,
                 perimeter_guid: str):
        """
        :param str id: NSP id in the ARM id format.
        :param str location: Location of the NSP supplied.
        :param str perimeter_guid: Resource guid of the NSP supplied.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "perimeter_guid", perimeter_guid)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        NSP id in the ARM id format.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Location of the NSP supplied.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="perimeterGuid")
    def perimeter_guid(self) -> str:
        """
        Resource guid of the NSP supplied.
        """
        return pulumi.get(self, "perimeter_guid")


@pulumi.output_type
class SubResourceResponse(dict):
    """
    Reference to another subresource.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        Reference to another subresource.
        :param str id: Resource ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class SubscriptionIdResponse(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        :param str id: Subscription id in the ARM id format.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Subscription id in the ARM id format.
        """
        return pulumi.get(self, "id")


