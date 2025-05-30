// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    public static class GetScheduledAlertRule
    {
        /// <summary>
        /// Gets the alert rule.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Task<GetScheduledAlertRuleResult> InvokeAsync(GetScheduledAlertRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScheduledAlertRuleResult>("azure-native:securityinsights:getScheduledAlertRule", args ?? new GetScheduledAlertRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the alert rule.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetScheduledAlertRuleResult> Invoke(GetScheduledAlertRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScheduledAlertRuleResult>("azure-native:securityinsights:getScheduledAlertRule", args ?? new GetScheduledAlertRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the alert rule.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetScheduledAlertRuleResult> Invoke(GetScheduledAlertRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetScheduledAlertRuleResult>("azure-native:securityinsights:getScheduledAlertRule", args ?? new GetScheduledAlertRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScheduledAlertRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Alert rule ID
        /// </summary>
        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetScheduledAlertRuleArgs()
        {
        }
        public static new GetScheduledAlertRuleArgs Empty => new GetScheduledAlertRuleArgs();
    }

    public sealed class GetScheduledAlertRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Alert rule ID
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetScheduledAlertRuleInvokeArgs()
        {
        }
        public static new GetScheduledAlertRuleInvokeArgs Empty => new GetScheduledAlertRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetScheduledAlertRuleResult
    {
        /// <summary>
        /// The alert details override settings
        /// </summary>
        public readonly Outputs.AlertDetailsOverrideResponse? AlertDetailsOverride;
        /// <summary>
        /// The Name of the alert rule template used to create this rule.
        /// </summary>
        public readonly string? AlertRuleTemplateName;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Dictionary of string key-value pairs of columns to be attached to the alert
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomDetails;
        /// <summary>
        /// The description of the alert rule.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The display name for alerts created by this alert rule.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Determines whether this alert rule is enabled or disabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Array of the entity mappings of the alert rule
        /// </summary>
        public readonly ImmutableArray<Outputs.EntityMappingResponse> EntityMappings;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The event grouping settings.
        /// </summary>
        public readonly Outputs.EventGroupingSettingsResponse? EventGroupingSettings;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The settings of the incidents that created from alerts triggered by this analytics rule
        /// </summary>
        public readonly Outputs.IncidentConfigurationResponse? IncidentConfiguration;
        /// <summary>
        /// The kind of the alert rule
        /// Expected value is 'Scheduled'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The last time that this alert rule has been modified.
        /// </summary>
        public readonly string LastModifiedUtc;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The query that creates alerts for this rule.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// The frequency (in ISO 8601 duration format) for this alert rule to run.
        /// </summary>
        public readonly string QueryFrequency;
        /// <summary>
        /// The period (in ISO 8601 duration format) that this alert rule looks at.
        /// </summary>
        public readonly string QueryPeriod;
        /// <summary>
        /// The severity for alerts created by this alert rule.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The suppression (in ISO 8601 duration format) to wait since last time this alert rule been triggered.
        /// </summary>
        public readonly string SuppressionDuration;
        /// <summary>
        /// Determines whether the suppression for this alert rule is enabled or disabled.
        /// </summary>
        public readonly bool SuppressionEnabled;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The tactics of the alert rule
        /// </summary>
        public readonly ImmutableArray<string> Tactics;
        /// <summary>
        /// The techniques of the alert rule
        /// </summary>
        public readonly ImmutableArray<string> Techniques;
        /// <summary>
        /// The version of the alert rule template used to create this rule - in format &lt;a.b.c&gt;, where all are numbers, for example 0 &lt;1.0.2&gt;
        /// </summary>
        public readonly string? TemplateVersion;
        /// <summary>
        /// The operation against the threshold that triggers alert rule.
        /// </summary>
        public readonly string TriggerOperator;
        /// <summary>
        /// The threshold triggers this alert rule.
        /// </summary>
        public readonly int TriggerThreshold;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetScheduledAlertRuleResult(
            Outputs.AlertDetailsOverrideResponse? alertDetailsOverride,

            string? alertRuleTemplateName,

            string azureApiVersion,

            ImmutableDictionary<string, string>? customDetails,

            string? description,

            string displayName,

            bool enabled,

            ImmutableArray<Outputs.EntityMappingResponse> entityMappings,

            string? etag,

            Outputs.EventGroupingSettingsResponse? eventGroupingSettings,

            string id,

            Outputs.IncidentConfigurationResponse? incidentConfiguration,

            string kind,

            string lastModifiedUtc,

            string name,

            string query,

            string queryFrequency,

            string queryPeriod,

            string severity,

            string suppressionDuration,

            bool suppressionEnabled,

            Outputs.SystemDataResponse systemData,

            ImmutableArray<string> tactics,

            ImmutableArray<string> techniques,

            string? templateVersion,

            string triggerOperator,

            int triggerThreshold,

            string type)
        {
            AlertDetailsOverride = alertDetailsOverride;
            AlertRuleTemplateName = alertRuleTemplateName;
            AzureApiVersion = azureApiVersion;
            CustomDetails = customDetails;
            Description = description;
            DisplayName = displayName;
            Enabled = enabled;
            EntityMappings = entityMappings;
            Etag = etag;
            EventGroupingSettings = eventGroupingSettings;
            Id = id;
            IncidentConfiguration = incidentConfiguration;
            Kind = kind;
            LastModifiedUtc = lastModifiedUtc;
            Name = name;
            Query = query;
            QueryFrequency = queryFrequency;
            QueryPeriod = queryPeriod;
            Severity = severity;
            SuppressionDuration = suppressionDuration;
            SuppressionEnabled = suppressionEnabled;
            SystemData = systemData;
            Tactics = tactics;
            Techniques = techniques;
            TemplateVersion = templateVersion;
            TriggerOperator = triggerOperator;
            TriggerThreshold = triggerThreshold;
            Type = type;
        }
    }
}
