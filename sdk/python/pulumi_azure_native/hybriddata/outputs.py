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
    'CustomerSecretResponse',
    'ScheduleResponse',
    'SkuResponse',
]

@pulumi.output_type
class CustomerSecretResponse(dict):
    """
    The pair of customer secret.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyIdentifier":
            suggest = "key_identifier"
        elif key == "keyValue":
            suggest = "key_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomerSecretResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomerSecretResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomerSecretResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 algorithm: builtins.str,
                 key_identifier: builtins.str,
                 key_value: builtins.str):
        """
        The pair of customer secret.
        :param builtins.str algorithm: The encryption algorithm used to encrypt data.
        :param builtins.str key_identifier: The identifier to the data service input object which this secret corresponds to.
        :param builtins.str key_value: It contains the encrypted customer secret.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "key_identifier", key_identifier)
        pulumi.set(__self__, "key_value", key_value)

    @property
    @pulumi.getter
    def algorithm(self) -> builtins.str:
        """
        The encryption algorithm used to encrypt data.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="keyIdentifier")
    def key_identifier(self) -> builtins.str:
        """
        The identifier to the data service input object which this secret corresponds to.
        """
        return pulumi.get(self, "key_identifier")

    @property
    @pulumi.getter(name="keyValue")
    def key_value(self) -> builtins.str:
        """
        It contains the encrypted customer secret.
        """
        return pulumi.get(self, "key_value")


@pulumi.output_type
class ScheduleResponse(dict):
    """
    Schedule for the job run.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyList":
            suggest = "policy_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ScheduleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ScheduleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ScheduleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: Optional[builtins.str] = None,
                 policy_list: Optional[Sequence[builtins.str]] = None):
        """
        Schedule for the job run.
        :param builtins.str name: Name of the schedule.
        :param Sequence[builtins.str] policy_list: A list of repetition intervals in ISO 8601 format.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_list is not None:
            pulumi.set(__self__, "policy_list", policy_list)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        Name of the schedule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyList")
    def policy_list(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of repetition intervals in ISO 8601 format.
        """
        return pulumi.get(self, "policy_list")


@pulumi.output_type
class SkuResponse(dict):
    """
    The sku type.
    """
    def __init__(__self__, *,
                 name: Optional[builtins.str] = None,
                 tier: Optional[builtins.str] = None):
        """
        The sku type.
        :param builtins.str name: The sku name. Required for data manager creation, optional for update.
        :param builtins.str tier: The sku tier. This is based on the SKU name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The sku name. Required for data manager creation, optional for update.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> Optional[builtins.str]:
        """
        The sku tier. This is based on the SKU name.
        """
        return pulumi.get(self, "tier")


