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
    /// Scheduled action definition.
    /// 
    /// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
    /// 
    /// Other available API versions: 2022-04-01-preview, 2022-06-01-preview, 2022-10-01, 2023-03-01, 2023-04-01-preview, 2023-07-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:costmanagement:ScheduledAction")]
    public partial class ScheduledAction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Scheduled action name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
        /// </summary>
        [Output("eTag")]
        public Output<string> ETag { get; private set; } = null!;

        /// <summary>
        /// Destination format of the view data. This is optional.
        /// </summary>
        [Output("fileDestination")]
        public Output<Outputs.FileDestinationResponse?> FileDestination { get; private set; } = null!;

        /// <summary>
        /// Kind of the scheduled action.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notification properties based on scheduled action kind.
        /// </summary>
        [Output("notification")]
        public Output<Outputs.NotificationPropertiesResponse> Notification { get; private set; } = null!;

        /// <summary>
        /// Email address of the point of contact that should get the unsubscribe requests and notification emails.
        /// </summary>
        [Output("notificationEmail")]
        public Output<string?> NotificationEmail { get; private set; } = null!;

        /// <summary>
        /// Schedule of the scheduled action.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.SchedulePropertiesResponse> Schedule { get; private set; } = null!;

        /// <summary>
        /// For private scheduled action(Create or Update), scope will be empty.&lt;br /&gt; For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// Status of the scheduled action.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Kind of the scheduled action.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
        /// </summary>
        [Output("viewId")]
        public Output<string> ViewId { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledAction(string name, ScheduledActionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:ScheduledAction", name, args ?? new ScheduledActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledAction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:ScheduledAction", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20220401preview:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20220601preview:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20221001:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230301:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230401preview:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230701preview:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230801:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230901:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20231101:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20240801:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20241001preview:ScheduledAction" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20250301:ScheduledAction" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledAction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ScheduledAction(name, id, options);
        }
    }

    public sealed class ScheduledActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Scheduled action name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Destination format of the view data. This is optional.
        /// </summary>
        [Input("fileDestination")]
        public Input<Inputs.FileDestinationArgs>? FileDestination { get; set; }

        /// <summary>
        /// Kind of the scheduled action.
        /// </summary>
        [Input("kind")]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.ScheduledActionKind>? Kind { get; set; }

        /// <summary>
        /// Scheduled action name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Notification properties based on scheduled action kind.
        /// </summary>
        [Input("notification", required: true)]
        public Input<Inputs.NotificationPropertiesArgs> Notification { get; set; } = null!;

        /// <summary>
        /// Email address of the point of contact that should get the unsubscribe requests and notification emails.
        /// </summary>
        [Input("notificationEmail")]
        public Input<string>? NotificationEmail { get; set; }

        /// <summary>
        /// Schedule of the scheduled action.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.SchedulePropertiesArgs> Schedule { get; set; } = null!;

        /// <summary>
        /// For private scheduled action(Create or Update), scope will be empty.&lt;br /&gt; For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Status of the scheduled action.
        /// </summary>
        [Input("status", required: true)]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.ScheduledActionStatus> Status { get; set; } = null!;

        /// <summary>
        /// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
        /// </summary>
        [Input("viewId", required: true)]
        public Input<string> ViewId { get; set; } = null!;

        public ScheduledActionArgs()
        {
        }
        public static new ScheduledActionArgs Empty => new ScheduledActionArgs();
    }
}
