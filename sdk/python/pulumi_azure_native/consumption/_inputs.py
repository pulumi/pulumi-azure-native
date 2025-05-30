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
    'BudgetComparisonExpressionArgs',
    'BudgetComparisonExpressionArgsDict',
    'BudgetFilterPropertiesArgs',
    'BudgetFilterPropertiesArgsDict',
    'BudgetFilterArgs',
    'BudgetFilterArgsDict',
    'BudgetTimePeriodArgs',
    'BudgetTimePeriodArgsDict',
    'NotificationArgs',
    'NotificationArgsDict',
]

MYPY = False

if not MYPY:
    class BudgetComparisonExpressionArgsDict(TypedDict):
        """
        The comparison expression to be used in the budgets.
        """
        name: pulumi.Input[builtins.str]
        """
        The name of the column to use in comparison.
        """
        operator: pulumi.Input[Union[builtins.str, 'BudgetOperatorType']]
        """
        The operator to use for comparison.
        """
        values: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]
        """
        Array of values to use for comparison
        """
elif False:
    BudgetComparisonExpressionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BudgetComparisonExpressionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 operator: pulumi.Input[Union[builtins.str, 'BudgetOperatorType']],
                 values: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        """
        The comparison expression to be used in the budgets.
        :param pulumi.Input[builtins.str] name: The name of the column to use in comparison.
        :param pulumi.Input[Union[builtins.str, 'BudgetOperatorType']] operator: The operator to use for comparison.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] values: Array of values to use for comparison
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the column to use in comparison.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[Union[builtins.str, 'BudgetOperatorType']]:
        """
        The operator to use for comparison.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[Union[builtins.str, 'BudgetOperatorType']]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        Array of values to use for comparison
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "values", value)


if not MYPY:
    class BudgetFilterPropertiesArgsDict(TypedDict):
        """
        The Dimensions or Tags to filter a budget by.
        """
        dimensions: NotRequired[pulumi.Input['BudgetComparisonExpressionArgsDict']]
        """
        Has comparison expression for a dimension
        """
        tags: NotRequired[pulumi.Input['BudgetComparisonExpressionArgsDict']]
        """
        Has comparison expression for a tag
        """
elif False:
    BudgetFilterPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BudgetFilterPropertiesArgs:
    def __init__(__self__, *,
                 dimensions: Optional[pulumi.Input['BudgetComparisonExpressionArgs']] = None,
                 tags: Optional[pulumi.Input['BudgetComparisonExpressionArgs']] = None):
        """
        The Dimensions or Tags to filter a budget by.
        :param pulumi.Input['BudgetComparisonExpressionArgs'] dimensions: Has comparison expression for a dimension
        :param pulumi.Input['BudgetComparisonExpressionArgs'] tags: Has comparison expression for a tag
        """
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input['BudgetComparisonExpressionArgs']]:
        """
        Has comparison expression for a dimension
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input['BudgetComparisonExpressionArgs']]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['BudgetComparisonExpressionArgs']]:
        """
        Has comparison expression for a tag
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['BudgetComparisonExpressionArgs']]):
        pulumi.set(self, "tags", value)


if not MYPY:
    class BudgetFilterArgsDict(TypedDict):
        """
        May be used to filter budgets by resource group, resource, or meter.
        """
        and_: NotRequired[pulumi.Input[Sequence[pulumi.Input['BudgetFilterPropertiesArgsDict']]]]
        """
        The logical "AND" expression. Must have at least 2 items.
        """
        dimensions: NotRequired[pulumi.Input['BudgetComparisonExpressionArgsDict']]
        """
        Has comparison expression for a dimension
        """
        tags: NotRequired[pulumi.Input['BudgetComparisonExpressionArgsDict']]
        """
        Has comparison expression for a tag
        """
elif False:
    BudgetFilterArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BudgetFilterArgs:
    def __init__(__self__, *,
                 and_: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetFilterPropertiesArgs']]]] = None,
                 dimensions: Optional[pulumi.Input['BudgetComparisonExpressionArgs']] = None,
                 tags: Optional[pulumi.Input['BudgetComparisonExpressionArgs']] = None):
        """
        May be used to filter budgets by resource group, resource, or meter.
        :param pulumi.Input[Sequence[pulumi.Input['BudgetFilterPropertiesArgs']]] and_: The logical "AND" expression. Must have at least 2 items.
        :param pulumi.Input['BudgetComparisonExpressionArgs'] dimensions: Has comparison expression for a dimension
        :param pulumi.Input['BudgetComparisonExpressionArgs'] tags: Has comparison expression for a tag
        """
        if and_ is not None:
            pulumi.set(__self__, "and_", and_)
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="and")
    def and_(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BudgetFilterPropertiesArgs']]]]:
        """
        The logical "AND" expression. Must have at least 2 items.
        """
        return pulumi.get(self, "and_")

    @and_.setter
    def and_(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetFilterPropertiesArgs']]]]):
        pulumi.set(self, "and_", value)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input['BudgetComparisonExpressionArgs']]:
        """
        Has comparison expression for a dimension
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input['BudgetComparisonExpressionArgs']]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['BudgetComparisonExpressionArgs']]:
        """
        Has comparison expression for a tag
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['BudgetComparisonExpressionArgs']]):
        pulumi.set(self, "tags", value)


if not MYPY:
    class BudgetTimePeriodArgsDict(TypedDict):
        """
        The start and end date for a budget.
        """
        start_date: pulumi.Input[builtins.str]
        """
        The start date for the budget.
        """
        end_date: NotRequired[pulumi.Input[builtins.str]]
        """
        The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
elif False:
    BudgetTimePeriodArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BudgetTimePeriodArgs:
    def __init__(__self__, *,
                 start_date: pulumi.Input[builtins.str],
                 end_date: Optional[pulumi.Input[builtins.str]] = None):
        """
        The start and end date for a budget.
        :param pulumi.Input[builtins.str] start_date: The start date for the budget.
        :param pulumi.Input[builtins.str] end_date: The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        pulumi.set(__self__, "start_date", start_date)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Input[builtins.str]:
        """
        The start date for the budget.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "end_date", value)


if not MYPY:
    class NotificationArgsDict(TypedDict):
        """
        The notification associated with a budget.
        """
        contact_emails: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]
        """
        Email addresses to send the budget notification to when the threshold is exceeded. Must have at least one contact email or contact group specified at the Subscription or Resource Group scopes. All other scopes must have at least one contact email specified.
        """
        enabled: pulumi.Input[builtins.bool]
        """
        The notification is enabled or not.
        """
        operator: pulumi.Input[Union[builtins.str, 'OperatorType']]
        """
        The comparison operator.
        """
        threshold: pulumi.Input[builtins.float]
        """
        Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        """
        contact_groups: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        Action groups to send the budget notification to when the threshold is exceeded. Must be provided as a fully qualified Azure resource id. Only supported at Subscription or Resource Group scopes.
        """
        contact_roles: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        Contact roles to send the budget notification to when the threshold is exceeded.
        """
        locale: NotRequired[pulumi.Input[Union[builtins.str, 'CultureCode']]]
        """
        Language in which the recipient will receive the notification
        """
        threshold_type: NotRequired[pulumi.Input[Union[builtins.str, 'ThresholdType']]]
        """
        The type of threshold
        """
elif False:
    NotificationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NotificationArgs:
    def __init__(__self__, *,
                 contact_emails: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 enabled: pulumi.Input[builtins.bool],
                 operator: pulumi.Input[Union[builtins.str, 'OperatorType']],
                 threshold: pulumi.Input[builtins.float],
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 contact_roles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 locale: Optional[pulumi.Input[Union[builtins.str, 'CultureCode']]] = None,
                 threshold_type: Optional[pulumi.Input[Union[builtins.str, 'ThresholdType']]] = None):
        """
        The notification associated with a budget.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] contact_emails: Email addresses to send the budget notification to when the threshold is exceeded. Must have at least one contact email or contact group specified at the Subscription or Resource Group scopes. All other scopes must have at least one contact email specified.
        :param pulumi.Input[builtins.bool] enabled: The notification is enabled or not.
        :param pulumi.Input[Union[builtins.str, 'OperatorType']] operator: The comparison operator.
        :param pulumi.Input[builtins.float] threshold: Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] contact_groups: Action groups to send the budget notification to when the threshold is exceeded. Must be provided as a fully qualified Azure resource id. Only supported at Subscription or Resource Group scopes.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] contact_roles: Contact roles to send the budget notification to when the threshold is exceeded.
        :param pulumi.Input[Union[builtins.str, 'CultureCode']] locale: Language in which the recipient will receive the notification
        :param pulumi.Input[Union[builtins.str, 'ThresholdType']] threshold_type: The type of threshold
        """
        pulumi.set(__self__, "contact_emails", contact_emails)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "threshold", threshold)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if contact_roles is not None:
            pulumi.set(__self__, "contact_roles", contact_roles)
        if locale is not None:
            pulumi.set(__self__, "locale", locale)
        if threshold_type is None:
            threshold_type = 'Actual'
        if threshold_type is not None:
            pulumi.set(__self__, "threshold_type", threshold_type)

    @property
    @pulumi.getter(name="contactEmails")
    def contact_emails(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        Email addresses to send the budget notification to when the threshold is exceeded. Must have at least one contact email or contact group specified at the Subscription or Resource Group scopes. All other scopes must have at least one contact email specified.
        """
        return pulumi.get(self, "contact_emails")

    @contact_emails.setter
    def contact_emails(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "contact_emails", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[builtins.bool]:
        """
        The notification is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[Union[builtins.str, 'OperatorType']]:
        """
        The comparison operator.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[Union[builtins.str, 'OperatorType']]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Input[builtins.float]:
        """
        Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        """
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: pulumi.Input[builtins.float]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Action groups to send the budget notification to when the threshold is exceeded. Must be provided as a fully qualified Azure resource id. Only supported at Subscription or Resource Group scopes.
        """
        return pulumi.get(self, "contact_groups")

    @contact_groups.setter
    def contact_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "contact_groups", value)

    @property
    @pulumi.getter(name="contactRoles")
    def contact_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Contact roles to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_roles")

    @contact_roles.setter
    def contact_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "contact_roles", value)

    @property
    @pulumi.getter
    def locale(self) -> Optional[pulumi.Input[Union[builtins.str, 'CultureCode']]]:
        """
        Language in which the recipient will receive the notification
        """
        return pulumi.get(self, "locale")

    @locale.setter
    def locale(self, value: Optional[pulumi.Input[Union[builtins.str, 'CultureCode']]]):
        pulumi.set(self, "locale", value)

    @property
    @pulumi.getter(name="thresholdType")
    def threshold_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'ThresholdType']]]:
        """
        The type of threshold
        """
        return pulumi.get(self, "threshold_type")

    @threshold_type.setter
    def threshold_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'ThresholdType']]]):
        pulumi.set(self, "threshold_type", value)


