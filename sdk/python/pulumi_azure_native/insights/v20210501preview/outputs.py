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
from . import outputs

__all__ = [
    'DiagnosticSettingsCategoryResourceResponse',
    'LogSettingsResponse',
    'ManagementGroupLogSettingsResponse',
    'MetricSettingsResponse',
    'RetentionPolicyResponse',
    'SubscriptionLogSettingsResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class DiagnosticSettingsCategoryResourceResponse(dict):
    """
    The diagnostic settings category resource.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 system_data: 'outputs.SystemDataResponse',
                 type: str,
                 category_groups: Optional[Sequence[str]] = None,
                 category_type: Optional[str] = None):
        """
        The diagnostic settings category resource.
        :param str id: Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        :param str name: The name of the resource
        :param 'SystemDataResponse' system_data: The system metadata related to this resource.
        :param str type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        :param Sequence[str] category_groups: the collection of what category groups are supported.
        :param str category_type: The type of the diagnostic settings category.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "system_data", system_data)
        pulumi.set(__self__, "type", type)
        if category_groups is not None:
            pulumi.set(__self__, "category_groups", category_groups)
        if category_type is not None:
            pulumi.set(__self__, "category_type", category_type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata related to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="categoryGroups")
    def category_groups(self) -> Optional[Sequence[str]]:
        """
        the collection of what category groups are supported.
        """
        return pulumi.get(self, "category_groups")

    @property
    @pulumi.getter(name="categoryType")
    def category_type(self) -> Optional[str]:
        """
        The type of the diagnostic settings category.
        """
        return pulumi.get(self, "category_type")


@pulumi.output_type
class LogSettingsResponse(dict):
    """
    Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "categoryGroup":
            suggest = "category_group"
        elif key == "retentionPolicy":
            suggest = "retention_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LogSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LogSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LogSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 category: Optional[str] = None,
                 category_group: Optional[str] = None,
                 retention_policy: Optional['outputs.RetentionPolicyResponse'] = None):
        """
        Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
        :param bool enabled: a value indicating whether this log is enabled.
        :param str category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        :param str category_group: Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        :param 'RetentionPolicyResponse' retention_policy: the retention policy for this log.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if category_group is not None:
            pulumi.set(__self__, "category_group", category_group)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        a value indicating whether this log is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="categoryGroup")
    def category_group(self) -> Optional[str]:
        """
        Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        """
        return pulumi.get(self, "category_group")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional['outputs.RetentionPolicyResponse']:
        """
        the retention policy for this log.
        """
        return pulumi.get(self, "retention_policy")


@pulumi.output_type
class ManagementGroupLogSettingsResponse(dict):
    """
    Part of Management Group diagnostic setting. Specifies the settings for a particular log.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "categoryGroup":
            suggest = "category_group"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagementGroupLogSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagementGroupLogSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagementGroupLogSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 category: Optional[str] = None,
                 category_group: Optional[str] = None):
        """
        Part of Management Group diagnostic setting. Specifies the settings for a particular log.
        :param bool enabled: a value indicating whether this log is enabled.
        :param str category: Name of a Management Group Diagnostic Log category for a resource type this setting is applied to.
        :param str category_group: Name of a Management Group Diagnostic Log category group for a resource type this setting is applied to.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if category_group is not None:
            pulumi.set(__self__, "category_group", category_group)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        a value indicating whether this log is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        Name of a Management Group Diagnostic Log category for a resource type this setting is applied to.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="categoryGroup")
    def category_group(self) -> Optional[str]:
        """
        Name of a Management Group Diagnostic Log category group for a resource type this setting is applied to.
        """
        return pulumi.get(self, "category_group")


@pulumi.output_type
class MetricSettingsResponse(dict):
    """
    Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "retentionPolicy":
            suggest = "retention_policy"
        elif key == "timeGrain":
            suggest = "time_grain"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MetricSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MetricSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MetricSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 category: Optional[str] = None,
                 retention_policy: Optional['outputs.RetentionPolicyResponse'] = None,
                 time_grain: Optional[str] = None):
        """
        Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
        :param bool enabled: a value indicating whether this category is enabled.
        :param str category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        :param 'RetentionPolicyResponse' retention_policy: the retention policy for this category.
        :param str time_grain: the timegrain of the metric in ISO8601 format.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)
        if time_grain is not None:
            pulumi.set(__self__, "time_grain", time_grain)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        a value indicating whether this category is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional['outputs.RetentionPolicyResponse']:
        """
        the retention policy for this category.
        """
        return pulumi.get(self, "retention_policy")

    @property
    @pulumi.getter(name="timeGrain")
    def time_grain(self) -> Optional[str]:
        """
        the timegrain of the metric in ISO8601 format.
        """
        return pulumi.get(self, "time_grain")


@pulumi.output_type
class RetentionPolicyResponse(dict):
    """
    Specifies the retention policy for the log.
    """
    def __init__(__self__, *,
                 days: int,
                 enabled: bool):
        """
        Specifies the retention policy for the log.
        :param int days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.
        :param bool enabled: a value indicating whether the retention policy is enabled.
        """
        pulumi.set(__self__, "days", days)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def days(self) -> int:
        """
        the number of days for the retention in days. A value of 0 will retain the events indefinitely.
        """
        return pulumi.get(self, "days")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        a value indicating whether the retention policy is enabled.
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class SubscriptionLogSettingsResponse(dict):
    """
    Part of Subscription diagnostic setting. Specifies the settings for a particular log.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "categoryGroup":
            suggest = "category_group"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionLogSettingsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionLogSettingsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionLogSettingsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled: bool,
                 category: Optional[str] = None,
                 category_group: Optional[str] = None):
        """
        Part of Subscription diagnostic setting. Specifies the settings for a particular log.
        :param bool enabled: a value indicating whether this log is enabled.
        :param str category: Name of a Subscription Diagnostic Log category for a resource type this setting is applied to.
        :param str category_group: Name of a Subscription Diagnostic Log category group for a resource type this setting is applied to.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if category_group is not None:
            pulumi.set(__self__, "category_group", category_group)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        a value indicating whether this log is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        Name of a Subscription Diagnostic Log category for a resource type this setting is applied to.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="categoryGroup")
    def category_group(self) -> Optional[str]:
        """
        Name of a Subscription Diagnostic Log category group for a resource type this setting is applied to.
        """
        return pulumi.get(self, "category_group")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


