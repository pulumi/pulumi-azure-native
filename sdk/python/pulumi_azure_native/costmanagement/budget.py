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
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['BudgetArgs', 'Budget']

@pulumi.input_type
class BudgetArgs:
    def __init__(__self__, *,
                 category: pulumi.Input[Union[builtins.str, 'CategoryType']],
                 scope: pulumi.Input[builtins.str],
                 time_grain: pulumi.Input[Union[builtins.str, 'TimeGrainType']],
                 time_period: pulumi.Input['BudgetTimePeriodArgs'],
                 amount: Optional[pulumi.Input[builtins.float]] = None,
                 budget_name: Optional[pulumi.Input[builtins.str]] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input['BudgetFilterArgs']] = None,
                 notifications: Optional[pulumi.Input[Mapping[str, pulumi.Input['NotificationArgs']]]] = None):
        """
        The set of arguments for constructing a Budget resource.
        :param pulumi.Input[Union[builtins.str, 'CategoryType']] category: The category of the budget.
               - 'Cost' defines a Budget.
               - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        :param pulumi.Input[builtins.str] scope: The scope associated with budget operations.
               
                Supported scopes for **CategoryType: Cost**
               
                Azure RBAC Scopes:
               - '/subscriptions/{subscriptionId}/' for subscription scope
               - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
               - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope
               
                EA (Enterprise Agreement) Scopes:
               
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope
               
                MCA (Modern Customer Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
               
                Supported scopes for **CategoryType: ReservationUtilization**
               
                EA (Enterprise Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope
               
               MCA (Modern Customer Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
        :param pulumi.Input[Union[builtins.str, 'TimeGrainType']] time_grain: The time covered by a budget. Tracking of the amount will be reset based on the time grain.
               
               Supported for CategoryType(s): Cost, ReservationUtilization.
               
                Supported timeGrainTypes for **CategoryType: Cost**
               
               - Monthly
               - Quarterly
               - Annually
               - BillingMonth*
               - BillingQuarter*
               - BillingAnnual*
               
                 *only supported for Web Direct customers.
               
                Supported timeGrainTypes for **CategoryType: ReservationUtilization**
               - Last7Days
               - Last30Days
               
                Required for CategoryType(s): Cost, ReservationUtilization.
        :param pulumi.Input['BudgetTimePeriodArgs'] time_period: The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
               
                Required for CategoryType(s): Cost, ReservationUtilization.
        :param pulumi.Input[builtins.float] amount: The total amount of cost to track with the budget.
               
                Supported for CategoryType(s): Cost.
               
                Required for CategoryType(s): Cost.
        :param pulumi.Input[builtins.str] budget_name: Budget Name.
        :param pulumi.Input[builtins.str] e_tag: eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        :param pulumi.Input['BudgetFilterArgs'] filter: May be used to filter budgets by user-specified dimensions and/or tags.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
        :param pulumi.Input[Mapping[str, pulumi.Input['NotificationArgs']]] notifications: Dictionary of notifications associated with the budget.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
               
               - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
               - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "scope", scope)
        pulumi.set(__self__, "time_grain", time_grain)
        pulumi.set(__self__, "time_period", time_period)
        if amount is not None:
            pulumi.set(__self__, "amount", amount)
        if budget_name is not None:
            pulumi.set(__self__, "budget_name", budget_name)
        if e_tag is not None:
            pulumi.set(__self__, "e_tag", e_tag)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Input[Union[builtins.str, 'CategoryType']]:
        """
        The category of the budget.
        - 'Cost' defines a Budget.
        - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: pulumi.Input[Union[builtins.str, 'CategoryType']]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[builtins.str]:
        """
        The scope associated with budget operations.

         Supported scopes for **CategoryType: Cost**

         Azure RBAC Scopes:
        - '/subscriptions/{subscriptionId}/' for subscription scope
        - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
        - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope

         EA (Enterprise Agreement) Scopes:

        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope

         MCA (Modern Customer Agreement) Scopes:
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)

         Supported scopes for **CategoryType: ReservationUtilization**

         EA (Enterprise Agreement) Scopes:
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope

        MCA (Modern Customer Agreement) Scopes:
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
        - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="timeGrain")
    def time_grain(self) -> pulumi.Input[Union[builtins.str, 'TimeGrainType']]:
        """
        The time covered by a budget. Tracking of the amount will be reset based on the time grain.

        Supported for CategoryType(s): Cost, ReservationUtilization.

         Supported timeGrainTypes for **CategoryType: Cost**

        - Monthly
        - Quarterly
        - Annually
        - BillingMonth*
        - BillingQuarter*
        - BillingAnnual*

          *only supported for Web Direct customers.

         Supported timeGrainTypes for **CategoryType: ReservationUtilization**
        - Last7Days
        - Last30Days

         Required for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "time_grain")

    @time_grain.setter
    def time_grain(self, value: pulumi.Input[Union[builtins.str, 'TimeGrainType']]):
        pulumi.set(self, "time_grain", value)

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> pulumi.Input['BudgetTimePeriodArgs']:
        """
        The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.

         Supported for CategoryType(s): Cost, ReservationUtilization.

         Required for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "time_period")

    @time_period.setter
    def time_period(self, value: pulumi.Input['BudgetTimePeriodArgs']):
        pulumi.set(self, "time_period", value)

    @property
    @pulumi.getter
    def amount(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        The total amount of cost to track with the budget.

         Supported for CategoryType(s): Cost.

         Required for CategoryType(s): Cost.
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter(name="budgetName")
    def budget_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Budget Name.
        """
        return pulumi.get(self, "budget_name")

    @budget_name.setter
    def budget_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "budget_name", value)

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        return pulumi.get(self, "e_tag")

    @e_tag.setter
    def e_tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "e_tag", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input['BudgetFilterArgs']]:
        """
        May be used to filter budgets by user-specified dimensions and/or tags.

         Supported for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input['BudgetFilterArgs']]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['NotificationArgs']]]]:
        """
        Dictionary of notifications associated with the budget.

         Supported for CategoryType(s): Cost, ReservationUtilization.

        - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
        - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['NotificationArgs']]]]):
        pulumi.set(self, "notifications", value)


@pulumi.type_token("azure-native:costmanagement:Budget")
class Budget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[builtins.float]] = None,
                 budget_name: Optional[pulumi.Input[builtins.str]] = None,
                 category: Optional[pulumi.Input[Union[builtins.str, 'CategoryType']]] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input[Union['BudgetFilterArgs', 'BudgetFilterArgsDict']]] = None,
                 notifications: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']]]]] = None,
                 scope: Optional[pulumi.Input[builtins.str]] = None,
                 time_grain: Optional[pulumi.Input[Union[builtins.str, 'TimeGrainType']]] = None,
                 time_period: Optional[pulumi.Input[Union['BudgetTimePeriodArgs', 'BudgetTimePeriodArgsDict']]] = None,
                 __props__=None):
        """
        A budget resource.

        Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.

        Other available API versions: 2019-04-01-preview, 2023-04-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.float] amount: The total amount of cost to track with the budget.
               
                Supported for CategoryType(s): Cost.
               
                Required for CategoryType(s): Cost.
        :param pulumi.Input[builtins.str] budget_name: Budget Name.
        :param pulumi.Input[Union[builtins.str, 'CategoryType']] category: The category of the budget.
               - 'Cost' defines a Budget.
               - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        :param pulumi.Input[builtins.str] e_tag: eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        :param pulumi.Input[Union['BudgetFilterArgs', 'BudgetFilterArgsDict']] filter: May be used to filter budgets by user-specified dimensions and/or tags.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
        :param pulumi.Input[Mapping[str, pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']]]] notifications: Dictionary of notifications associated with the budget.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
               
               - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
               - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        :param pulumi.Input[builtins.str] scope: The scope associated with budget operations.
               
                Supported scopes for **CategoryType: Cost**
               
                Azure RBAC Scopes:
               - '/subscriptions/{subscriptionId}/' for subscription scope
               - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
               - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope
               
                EA (Enterprise Agreement) Scopes:
               
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope
               
                MCA (Modern Customer Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
               
                Supported scopes for **CategoryType: ReservationUtilization**
               
                EA (Enterprise Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope
               
               MCA (Modern Customer Agreement) Scopes:
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
               - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
        :param pulumi.Input[Union[builtins.str, 'TimeGrainType']] time_grain: The time covered by a budget. Tracking of the amount will be reset based on the time grain.
               
               Supported for CategoryType(s): Cost, ReservationUtilization.
               
                Supported timeGrainTypes for **CategoryType: Cost**
               
               - Monthly
               - Quarterly
               - Annually
               - BillingMonth*
               - BillingQuarter*
               - BillingAnnual*
               
                 *only supported for Web Direct customers.
               
                Supported timeGrainTypes for **CategoryType: ReservationUtilization**
               - Last7Days
               - Last30Days
               
                Required for CategoryType(s): Cost, ReservationUtilization.
        :param pulumi.Input[Union['BudgetTimePeriodArgs', 'BudgetTimePeriodArgsDict']] time_period: The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
               
                Supported for CategoryType(s): Cost, ReservationUtilization.
               
                Required for CategoryType(s): Cost, ReservationUtilization.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BudgetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A budget resource.

        Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.

        Other available API versions: 2019-04-01-preview, 2023-04-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param BudgetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BudgetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amount: Optional[pulumi.Input[builtins.float]] = None,
                 budget_name: Optional[pulumi.Input[builtins.str]] = None,
                 category: Optional[pulumi.Input[Union[builtins.str, 'CategoryType']]] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input[Union['BudgetFilterArgs', 'BudgetFilterArgsDict']]] = None,
                 notifications: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['NotificationArgs', 'NotificationArgsDict']]]]] = None,
                 scope: Optional[pulumi.Input[builtins.str]] = None,
                 time_grain: Optional[pulumi.Input[Union[builtins.str, 'TimeGrainType']]] = None,
                 time_period: Optional[pulumi.Input[Union['BudgetTimePeriodArgs', 'BudgetTimePeriodArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BudgetArgs.__new__(BudgetArgs)

            __props__.__dict__["amount"] = amount
            __props__.__dict__["budget_name"] = budget_name
            if category is None and not opts.urn:
                raise TypeError("Missing required property 'category'")
            __props__.__dict__["category"] = category
            __props__.__dict__["e_tag"] = e_tag
            __props__.__dict__["filter"] = filter
            __props__.__dict__["notifications"] = notifications
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__.__dict__["scope"] = scope
            if time_grain is None and not opts.urn:
                raise TypeError("Missing required property 'time_grain'")
            __props__.__dict__["time_grain"] = time_grain
            if time_period is None and not opts.urn:
                raise TypeError("Missing required property 'time_period'")
            __props__.__dict__["time_period"] = time_period
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["current_spend"] = None
            __props__.__dict__["forecast_spend"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:costmanagement/v20190401preview:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20230401preview:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20230801:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20230901:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20231101:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20240801:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20241001preview:Budget"), pulumi.Alias(type_="azure-native:costmanagement/v20250301:Budget")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Budget, __self__).__init__(
            'azure-native:costmanagement:Budget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Budget':
        """
        Get an existing Budget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BudgetArgs.__new__(BudgetArgs)

        __props__.__dict__["amount"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["current_spend"] = None
        __props__.__dict__["e_tag"] = None
        __props__.__dict__["filter"] = None
        __props__.__dict__["forecast_spend"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notifications"] = None
        __props__.__dict__["time_grain"] = None
        __props__.__dict__["time_period"] = None
        __props__.__dict__["type"] = None
        return Budget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        The total amount of cost to track with the budget.

         Supported for CategoryType(s): Cost.

         Required for CategoryType(s): Cost.
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[builtins.str]:
        """
        The category of the budget.
        - 'Cost' defines a Budget.
        - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="currentSpend")
    def current_spend(self) -> pulumi.Output['outputs.CurrentSpendResponse']:
        """
        The current amount of cost which is being tracked for a budget.

         Supported for CategoryType(s): Cost.
        """
        return pulumi.get(self, "current_spend")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional['outputs.BudgetFilterResponse']]:
        """
        May be used to filter budgets by user-specified dimensions and/or tags.

         Supported for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="forecastSpend")
    def forecast_spend(self) -> pulumi.Output['outputs.ForecastSpendResponse']:
        """
        The forecasted cost which is being tracked for a budget.

         Supported for CategoryType(s): Cost.
        """
        return pulumi.get(self, "forecast_spend")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.NotificationResponse']]]:
        """
        Dictionary of notifications associated with the budget.

         Supported for CategoryType(s): Cost, ReservationUtilization.

        - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
        - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="timeGrain")
    def time_grain(self) -> pulumi.Output[builtins.str]:
        """
        The time covered by a budget. Tracking of the amount will be reset based on the time grain.

        Supported for CategoryType(s): Cost, ReservationUtilization.

         Supported timeGrainTypes for **CategoryType: Cost**

        - Monthly
        - Quarterly
        - Annually
        - BillingMonth*
        - BillingQuarter*
        - BillingAnnual*

          *only supported for Web Direct customers.

         Supported timeGrainTypes for **CategoryType: ReservationUtilization**
        - Last7Days
        - Last30Days

         Required for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "time_grain")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> pulumi.Output['outputs.BudgetTimePeriodResponse']:
        """
        The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.

         Supported for CategoryType(s): Cost, ReservationUtilization.

         Required for CategoryType(s): Cost, ReservationUtilization.
        """
        return pulumi.get(self, "time_period")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

