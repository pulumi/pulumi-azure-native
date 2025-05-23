// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor
{
    /// <summary>
    /// The scheduled query rule resource.
    /// 
    /// Uses Azure REST API version 2024-01-01-preview.
    /// 
    /// Other available API versions: 2023-12-01, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:monitor:ScheduledQueryRule")]
    public partial class ScheduledQueryRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Actions to invoke when the alert fires.
        /// </summary>
        [Output("actions")]
        public Output<Outputs.ActionsResponse?> Actions { get; private set; } = null!;

        /// <summary>
        /// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("autoMitigate")]
        public Output<bool?> AutoMitigate { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("checkWorkspaceAlertsStorageConfigured")]
        public Output<bool?> CheckWorkspaceAlertsStorageConfigured { get; private set; } = null!;

        /// <summary>
        /// The api-version used when creating this alert rule
        /// </summary>
        [Output("createdWithApiVersion")]
        public Output<string> CreatedWithApiVersion { get; private set; } = null!;

        /// <summary>
        /// The rule criteria that defines the conditions of the scheduled query rule.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.ScheduledQueryRuleCriteriaResponse> Criteria { get; private set; } = null!;

        /// <summary>
        /// The description of the scheduled query rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the alert rule
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. 
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Output("evaluationFrequency")]
        public Output<string?> EvaluationFrequency { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// True if alert rule is legacy Log Analytic rule
        /// </summary>
        [Output("isLegacyLogAnalyticsRule")]
        public Output<bool> IsLegacyLogAnalyticsRule { get; private set; } = null!;

        /// <summary>
        /// The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.
        /// </summary>
        [Output("isWorkspaceAlertsStorageConfigured")]
        public Output<bool> IsWorkspaceAlertsStorageConfigured { get; private set; } = null!;

        /// <summary>
        /// Indicates the type of scheduled query rule. The default is LogAlert.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("muteActionsDuration")]
        public Output<string?> MuteActionsDuration { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("overrideQueryTimeRange")]
        public Output<string?> OverrideQueryTimeRange { get; private set; } = null!;

        /// <summary>
        /// Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("resolveConfiguration")]
        public Output<Outputs.RuleResolveConfigurationResponse?> ResolveConfiguration { get; private set; } = null!;

        /// <summary>
        /// The list of resource id's that this scheduled query rule is scoped to.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Output("severity")]
        public Output<double?> Severity { get; private set; } = null!;

        /// <summary>
        /// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Output("skipQueryValidation")]
        public Output<bool?> SkipQueryValidation { get; private set; } = null!;

        /// <summary>
        /// SystemData of ScheduledQueryRule.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
        /// </summary>
        [Output("targetResourceTypes")]
        public Output<ImmutableArray<string>> TargetResourceTypes { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Output("windowSize")]
        public Output<string?> WindowSize { get; private set; } = null!;


        /// <summary>
        /// Create a ScheduledQueryRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScheduledQueryRule(string name, ScheduledQueryRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:monitor:ScheduledQueryRule", name, args ?? new ScheduledQueryRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScheduledQueryRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:monitor:ScheduledQueryRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20180416:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20200501preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20220801preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20230315preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20231201:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights/v20240101preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:insights:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20180416:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20200501preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20210201preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20210801:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20220615:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20220801preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20230315preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20231201:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20240101preview:ScheduledQueryRule" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20250101preview:ScheduledQueryRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ScheduledQueryRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScheduledQueryRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ScheduledQueryRule(name, id, options);
        }
    }

    public sealed class ScheduledQueryRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actions to invoke when the alert fires.
        /// </summary>
        [Input("actions")]
        public Input<Inputs.ActionsArgs>? Actions { get; set; }

        /// <summary>
        /// The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("autoMitigate")]
        public Input<bool>? AutoMitigate { get; set; }

        /// <summary>
        /// The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("checkWorkspaceAlertsStorageConfigured")]
        public Input<bool>? CheckWorkspaceAlertsStorageConfigured { get; set; }

        /// <summary>
        /// The rule criteria that defines the conditions of the scheduled query rule.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.ScheduledQueryRuleCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// The description of the scheduled query rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the alert rule
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Input("evaluationFrequency")]
        public Input<string>? EvaluationFrequency { get; set; }

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Indicates the type of scheduled query rule. The default is LogAlert.
        /// </summary>
        [Input("kind")]
        public InputUnion<string, Pulumi.AzureNative.Monitor.Kind>? Kind { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("muteActionsDuration")]
        public Input<string>? MuteActionsDuration { get; set; }

        /// <summary>
        /// If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("overrideQueryTimeRange")]
        public Input<string>? OverrideQueryTimeRange { get; set; }

        /// <summary>
        /// Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("resolveConfiguration")]
        public Input<Inputs.RuleResolveConfigurationArgs>? ResolveConfiguration { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// The list of resource id's that this scheduled query rule is scoped to.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Input("severity")]
        public Input<double>? Severity { get; set; }

        /// <summary>
        /// The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.
        /// </summary>
        [Input("skipQueryValidation")]
        public Input<bool>? SkipQueryValidation { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetResourceTypes")]
        private InputList<string>? _targetResourceTypes;

        /// <summary>
        /// List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert
        /// </summary>
        public InputList<string> TargetResourceTypes
        {
            get => _targetResourceTypes ?? (_targetResourceTypes = new InputList<string>());
            set => _targetResourceTypes = value;
        }

        /// <summary>
        /// The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.
        /// </summary>
        [Input("windowSize")]
        public Input<string>? WindowSize { get; set; }

        public ScheduledQueryRuleArgs()
        {
        }
        public static new ScheduledQueryRuleArgs Empty => new ScheduledQueryRuleArgs();
    }
}
