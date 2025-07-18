// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class GetLocalRuleCounters
    {
        /// <summary>
        /// Get counters
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetLocalRuleCountersResult> InvokeAsync(GetLocalRuleCountersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalRuleCountersResult>("azure-native:cloudngfw:getLocalRuleCounters", args ?? new GetLocalRuleCountersArgs(), options.WithDefaults());

        /// <summary>
        /// Get counters
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLocalRuleCountersResult> Invoke(GetLocalRuleCountersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalRuleCountersResult>("azure-native:cloudngfw:getLocalRuleCounters", args ?? new GetLocalRuleCountersInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get counters
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLocalRuleCountersResult> Invoke(GetLocalRuleCountersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalRuleCountersResult>("azure-native:cloudngfw:getLocalRuleCounters", args ?? new GetLocalRuleCountersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalRuleCountersArgs : global::Pulumi.InvokeArgs
    {
        [Input("firewallName")]
        public string? FirewallName { get; set; }

        /// <summary>
        /// LocalRulestack resource name
        /// </summary>
        [Input("localRulestackName", required: true)]
        public string LocalRulestackName { get; set; } = null!;

        /// <summary>
        /// Local Rule priority
        /// </summary>
        [Input("priority", required: true)]
        public string Priority { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLocalRuleCountersArgs()
        {
        }
        public static new GetLocalRuleCountersArgs Empty => new GetLocalRuleCountersArgs();
    }

    public sealed class GetLocalRuleCountersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("firewallName")]
        public Input<string>? FirewallName { get; set; }

        /// <summary>
        /// LocalRulestack resource name
        /// </summary>
        [Input("localRulestackName", required: true)]
        public Input<string> LocalRulestackName { get; set; } = null!;

        /// <summary>
        /// Local Rule priority
        /// </summary>
        [Input("priority", required: true)]
        public Input<string> Priority { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetLocalRuleCountersInvokeArgs()
        {
        }
        public static new GetLocalRuleCountersInvokeArgs Empty => new GetLocalRuleCountersInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalRuleCountersResult
    {
        /// <summary>
        /// apps seen
        /// </summary>
        public readonly Outputs.AppSeenDataResponse? AppSeen;
        /// <summary>
        /// firewall name
        /// </summary>
        public readonly string? FirewallName;
        /// <summary>
        /// hit count
        /// </summary>
        public readonly int? HitCount;
        /// <summary>
        /// last updated timestamp
        /// </summary>
        public readonly string? LastUpdatedTimestamp;
        /// <summary>
        /// priority number
        /// </summary>
        public readonly string Priority;
        /// <summary>
        /// timestamp of request
        /// </summary>
        public readonly string? RequestTimestamp;
        /// <summary>
        /// rule list name
        /// </summary>
        public readonly string? RuleListName;
        /// <summary>
        /// rule name
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// rule Stack Name
        /// </summary>
        public readonly string? RuleStackName;
        /// <summary>
        /// timestamp of response
        /// </summary>
        public readonly string? Timestamp;

        [OutputConstructor]
        private GetLocalRuleCountersResult(
            Outputs.AppSeenDataResponse? appSeen,

            string? firewallName,

            int? hitCount,

            string? lastUpdatedTimestamp,

            string priority,

            string? requestTimestamp,

            string? ruleListName,

            string ruleName,

            string? ruleStackName,

            string? timestamp)
        {
            AppSeen = appSeen;
            FirewallName = firewallName;
            HitCount = hitCount;
            LastUpdatedTimestamp = lastUpdatedTimestamp;
            Priority = priority;
            RequestTimestamp = requestTimestamp;
            RuleListName = ruleListName;
            RuleName = ruleName;
            RuleStackName = ruleStackName;
            Timestamp = timestamp;
        }
    }
}
