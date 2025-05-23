// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement
{
    /// <summary>
    /// A budget resource.
    /// 
    /// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.
    /// 
    /// Other available API versions: 2019-04-01-preview, 2023-04-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:costmanagement:Budget")]
    public partial class Budget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The total amount of cost to track with the budget.
        /// 
        ///  Supported for CategoryType(s): Cost.
        /// 
        ///  Required for CategoryType(s): Cost.
        /// </summary>
        [Output("amount")]
        public Output<double?> Amount { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The category of the budget.
        /// - 'Cost' defines a Budget.
        /// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// The current amount of cost which is being tracked for a budget.
        /// 
        ///  Supported for CategoryType(s): Cost.
        /// </summary>
        [Output("currentSpend")]
        public Output<Outputs.CurrentSpendResponse> CurrentSpend { get; private set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// May be used to filter budgets by user-specified dimensions and/or tags.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Output("filter")]
        public Output<Outputs.BudgetFilterResponse?> Filter { get; private set; } = null!;

        /// <summary>
        /// The forecasted cost which is being tracked for a budget.
        /// 
        ///  Supported for CategoryType(s): Cost.
        /// </summary>
        [Output("forecastSpend")]
        public Output<Outputs.ForecastSpendResponse> ForecastSpend { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Dictionary of notifications associated with the budget.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        /// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
        /// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableDictionary<string, Outputs.NotificationResponse>?> Notifications { get; private set; } = null!;

        /// <summary>
        /// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
        /// 
        /// Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        ///  Supported timeGrainTypes for **CategoryType: Cost**
        /// 
        /// - Monthly
        /// - Quarterly
        /// - Annually
        /// - BillingMonth*
        /// - BillingQuarter*
        /// - BillingAnnual*
        /// 
        ///   *only supported for Web Direct customers.
        /// 
        ///  Supported timeGrainTypes for **CategoryType: ReservationUtilization**
        /// - Last7Days
        /// - Last30Days
        /// 
        ///  Required for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Output("timeGrain")]
        public Output<string> TimeGrain { get; private set; } = null!;

        /// <summary>
        /// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        ///  Required for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Output("timePeriod")]
        public Output<Outputs.BudgetTimePeriodResponse> TimePeriod { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Budget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Budget(string name, BudgetArgs args, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:Budget", name, args ?? new BudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Budget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:Budget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20190401preview:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230401preview:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230801:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230901:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20231101:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20240801:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20241001preview:Budget" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20250301:Budget" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Budget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Budget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Budget(name, id, options);
        }
    }

    public sealed class BudgetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The total amount of cost to track with the budget.
        /// 
        ///  Supported for CategoryType(s): Cost.
        /// 
        ///  Required for CategoryType(s): Cost.
        /// </summary>
        [Input("amount")]
        public Input<double>? Amount { get; set; }

        /// <summary>
        /// Budget Name.
        /// </summary>
        [Input("budgetName")]
        public Input<string>? BudgetName { get; set; }

        /// <summary>
        /// The category of the budget.
        /// - 'Cost' defines a Budget.
        /// - 'ReservationUtilization' defines a Reservation Utilization Alert Rule.
        /// </summary>
        [Input("category", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.CategoryType> Category { get; set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// May be used to filter budgets by user-specified dimensions and/or tags.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BudgetFilterArgs>? Filter { get; set; }

        [Input("notifications")]
        private InputMap<Inputs.NotificationArgs>? _notifications;

        /// <summary>
        /// Dictionary of notifications associated with the budget.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        /// - Constraints for **CategoryType: Cost** - Budget can have up to 5 notifications with thresholdType: Actual and 5 notifications with thresholdType: Forecasted.
        /// - Constraints for **CategoryType: ReservationUtilization** - Only one notification allowed. thresholdType is not applicable.
        /// </summary>
        public InputMap<Inputs.NotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputMap<Inputs.NotificationArgs>());
            set => _notifications = value;
        }

        /// <summary>
        /// The scope associated with budget operations.
        /// 
        ///  Supported scopes for **CategoryType: Cost**
        /// 
        ///  Azure RBAC Scopes:
        /// - '/subscriptions/{subscriptionId}/' for subscription scope
        /// - '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope
        /// - '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope
        /// 
        ///  EA (Enterprise Agreement) Scopes:
        /// 
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope
        /// 
        ///  MCA (Modern Customer Agreement) Scopes:
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
        /// 
        ///  Supported scopes for **CategoryType: ReservationUtilization**
        /// 
        ///  EA (Enterprise Agreement) Scopes:
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account Scope
        /// 
        /// MCA (Modern Customer Agreement) Scopes:
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope (non-CSP only)
        /// - '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' for customer scope (CSP only)
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
        /// 
        /// Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        ///  Supported timeGrainTypes for **CategoryType: Cost**
        /// 
        /// - Monthly
        /// - Quarterly
        /// - Annually
        /// - BillingMonth*
        /// - BillingQuarter*
        /// - BillingAnnual*
        /// 
        ///   *only supported for Web Direct customers.
        /// 
        ///  Supported timeGrainTypes for **CategoryType: ReservationUtilization**
        /// - Last7Days
        /// - Last30Days
        /// 
        ///  Required for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Input("timeGrain", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.TimeGrainType> TimeGrain { get; set; } = null!;

        /// <summary>
        /// The time period that defines the active period of the budget. The budget will evaluate data on or after the startDate and will expire on the endDate.
        /// 
        ///  Supported for CategoryType(s): Cost, ReservationUtilization.
        /// 
        ///  Required for CategoryType(s): Cost, ReservationUtilization.
        /// </summary>
        [Input("timePeriod", required: true)]
        public Input<Inputs.BudgetTimePeriodArgs> TimePeriod { get; set; } = null!;

        public BudgetArgs()
        {
        }
        public static new BudgetArgs Empty => new BudgetArgs();
    }
}
