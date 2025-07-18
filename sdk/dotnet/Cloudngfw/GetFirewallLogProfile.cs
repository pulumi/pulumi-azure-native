// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw
{
    public static class GetFirewallLogProfile
    {
        /// <summary>
        /// Log Profile for Firewall
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetFirewallLogProfileResult> InvokeAsync(GetFirewallLogProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallLogProfileResult>("azure-native:cloudngfw:getFirewallLogProfile", args ?? new GetFirewallLogProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Log Profile for Firewall
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallLogProfileResult> Invoke(GetFirewallLogProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallLogProfileResult>("azure-native:cloudngfw:getFirewallLogProfile", args ?? new GetFirewallLogProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Log Profile for Firewall
        /// 
        /// Uses Azure REST API version 2025-02-06-preview.
        /// 
        /// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetFirewallLogProfileResult> Invoke(GetFirewallLogProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallLogProfileResult>("azure-native:cloudngfw:getFirewallLogProfile", args ?? new GetFirewallLogProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallLogProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Firewall resource name
        /// </summary>
        [Input("firewallName", required: true)]
        public string FirewallName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFirewallLogProfileArgs()
        {
        }
        public static new GetFirewallLogProfileArgs Empty => new GetFirewallLogProfileArgs();
    }

    public sealed class GetFirewallLogProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Firewall resource name
        /// </summary>
        [Input("firewallName", required: true)]
        public Input<string> FirewallName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetFirewallLogProfileInvokeArgs()
        {
        }
        public static new GetFirewallLogProfileInvokeArgs Empty => new GetFirewallLogProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallLogProfileResult
    {
        /// <summary>
        /// Application Insight details
        /// </summary>
        public readonly Outputs.ApplicationInsightsResponse? ApplicationInsights;
        /// <summary>
        /// Common destination configurations
        /// </summary>
        public readonly Outputs.LogDestinationResponse? CommonDestination;
        /// <summary>
        /// Decrypt destination configurations
        /// </summary>
        public readonly Outputs.LogDestinationResponse? DecryptLogDestination;
        /// <summary>
        /// Log option SAME/INDIVIDUAL
        /// </summary>
        public readonly string? LogOption;
        /// <summary>
        /// One of possible log type
        /// </summary>
        public readonly string? LogType;
        /// <summary>
        /// Threat destination configurations
        /// </summary>
        public readonly Outputs.LogDestinationResponse? ThreatLogDestination;
        /// <summary>
        /// Traffic destination configurations
        /// </summary>
        public readonly Outputs.LogDestinationResponse? TrafficLogDestination;

        [OutputConstructor]
        private GetFirewallLogProfileResult(
            Outputs.ApplicationInsightsResponse? applicationInsights,

            Outputs.LogDestinationResponse? commonDestination,

            Outputs.LogDestinationResponse? decryptLogDestination,

            string? logOption,

            string? logType,

            Outputs.LogDestinationResponse? threatLogDestination,

            Outputs.LogDestinationResponse? trafficLogDestination)
        {
            ApplicationInsights = applicationInsights;
            CommonDestination = commonDestination;
            DecryptLogDestination = decryptLogDestination;
            LogOption = logOption;
            LogType = logType;
            ThreatLogDestination = threatLogDestination;
            TrafficLogDestination = trafficLogDestination;
        }
    }
}
