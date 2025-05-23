// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement
{
    public static class GetSmartDetectorAlertRule
    {
        /// <summary>
        /// Get a specific Smart Detector alert rule.
        /// 
        /// Uses Azure REST API version 2021-04-01.
        /// 
        /// Other available API versions: 2019-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native alertsmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSmartDetectorAlertRuleResult> InvokeAsync(GetSmartDetectorAlertRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSmartDetectorAlertRuleResult>("azure-native:alertsmanagement:getSmartDetectorAlertRule", args ?? new GetSmartDetectorAlertRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Get a specific Smart Detector alert rule.
        /// 
        /// Uses Azure REST API version 2021-04-01.
        /// 
        /// Other available API versions: 2019-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native alertsmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSmartDetectorAlertRuleResult> Invoke(GetSmartDetectorAlertRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSmartDetectorAlertRuleResult>("azure-native:alertsmanagement:getSmartDetectorAlertRule", args ?? new GetSmartDetectorAlertRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a specific Smart Detector alert rule.
        /// 
        /// Uses Azure REST API version 2021-04-01.
        /// 
        /// Other available API versions: 2019-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native alertsmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSmartDetectorAlertRuleResult> Invoke(GetSmartDetectorAlertRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSmartDetectorAlertRuleResult>("azure-native:alertsmanagement:getSmartDetectorAlertRule", args ?? new GetSmartDetectorAlertRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSmartDetectorAlertRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("alertRuleName", required: true)]
        public string AlertRuleName { get; set; } = null!;

        /// <summary>
        /// Indicates if Smart Detector should be expanded.
        /// </summary>
        [Input("expandDetector")]
        public bool? ExpandDetector { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSmartDetectorAlertRuleArgs()
        {
        }
        public static new GetSmartDetectorAlertRuleArgs Empty => new GetSmartDetectorAlertRuleArgs();
    }

    public sealed class GetSmartDetectorAlertRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alert rule.
        /// </summary>
        [Input("alertRuleName", required: true)]
        public Input<string> AlertRuleName { get; set; } = null!;

        /// <summary>
        /// Indicates if Smart Detector should be expanded.
        /// </summary>
        [Input("expandDetector")]
        public Input<bool>? ExpandDetector { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSmartDetectorAlertRuleInvokeArgs()
        {
        }
        public static new GetSmartDetectorAlertRuleInvokeArgs Empty => new GetSmartDetectorAlertRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetSmartDetectorAlertRuleResult
    {
        /// <summary>
        /// The alert rule actions.
        /// </summary>
        public readonly Outputs.ActionGroupsInformationResponse ActionGroups;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The alert rule description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The alert rule's detector.
        /// </summary>
        public readonly Outputs.DetectorResponse Detector;
        /// <summary>
        /// The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 1 minute, depending on the detector.
        /// </summary>
        public readonly string Frequency;
        /// <summary>
        /// The resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The alert rule resources scope.
        /// </summary>
        public readonly ImmutableArray<string> Scope;
        /// <summary>
        /// The alert rule severity.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The alert rule state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The alert rule throttling information.
        /// </summary>
        public readonly Outputs.ThrottlingInformationResponse? Throttling;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSmartDetectorAlertRuleResult(
            Outputs.ActionGroupsInformationResponse actionGroups,

            string azureApiVersion,

            string? description,

            Outputs.DetectorResponse detector,

            string frequency,

            string id,

            string? location,

            string name,

            ImmutableArray<string> scope,

            string severity,

            string state,

            ImmutableDictionary<string, string>? tags,

            Outputs.ThrottlingInformationResponse? throttling,

            string type)
        {
            ActionGroups = actionGroups;
            AzureApiVersion = azureApiVersion;
            Description = description;
            Detector = detector;
            Frequency = frequency;
            Id = id;
            Location = location;
            Name = name;
            Scope = scope;
            Severity = severity;
            State = state;
            Tags = tags;
            Throttling = throttling;
            Type = type;
        }
    }
}
