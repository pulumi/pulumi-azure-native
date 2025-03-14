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
    'PrivateEndpointPropertyResponse',
    'PrivateLinkServiceConnectionStatePropertyResponse',
]

@pulumi.output_type
class PrivateEndpointPropertyResponse(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        :param str id: Resource id of the private endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource id of the private endpoint.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PrivateLinkServiceConnectionStatePropertyResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionsRequired":
            suggest = "actions_required"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateLinkServiceConnectionStatePropertyResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateLinkServiceConnectionStatePropertyResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateLinkServiceConnectionStatePropertyResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 actions_required: str,
                 description: str,
                 status: str):
        """
        :param str actions_required: The actions required for private link service connection.
        :param str description: The private link service connection description.
        :param str status: The private link service connection status.
        """
        pulumi.set(__self__, "actions_required", actions_required)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> str:
        """
        The actions required for private link service connection.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The private link service connection description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The private link service connection status.
        """
        return pulumi.get(self, "status")


