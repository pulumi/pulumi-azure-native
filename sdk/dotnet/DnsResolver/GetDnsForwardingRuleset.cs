// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DnsResolver
{
    public static class GetDnsForwardingRuleset
    {
        /// <summary>
        /// Gets a DNS forwarding ruleset properties.
        /// 
        /// Uses Azure REST API version 2023-07-01-preview.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDnsForwardingRulesetResult> InvokeAsync(GetDnsForwardingRulesetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsForwardingRulesetResult>("azure-native:dnsresolver:getDnsForwardingRuleset", args ?? new GetDnsForwardingRulesetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a DNS forwarding ruleset properties.
        /// 
        /// Uses Azure REST API version 2023-07-01-preview.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDnsForwardingRulesetResult> Invoke(GetDnsForwardingRulesetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsForwardingRulesetResult>("azure-native:dnsresolver:getDnsForwardingRuleset", args ?? new GetDnsForwardingRulesetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a DNS forwarding ruleset properties.
        /// 
        /// Uses Azure REST API version 2023-07-01-preview.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDnsForwardingRulesetResult> Invoke(GetDnsForwardingRulesetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsForwardingRulesetResult>("azure-native:dnsresolver:getDnsForwardingRuleset", args ?? new GetDnsForwardingRulesetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsForwardingRulesetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNS forwarding ruleset.
        /// </summary>
        [Input("dnsForwardingRulesetName", required: true)]
        public string DnsForwardingRulesetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDnsForwardingRulesetArgs()
        {
        }
        public static new GetDnsForwardingRulesetArgs Empty => new GetDnsForwardingRulesetArgs();
    }

    public sealed class GetDnsForwardingRulesetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNS forwarding ruleset.
        /// </summary>
        [Input("dnsForwardingRulesetName", required: true)]
        public Input<string> DnsForwardingRulesetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDnsForwardingRulesetInvokeArgs()
        {
        }
        public static new GetDnsForwardingRulesetInvokeArgs Empty => new GetDnsForwardingRulesetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsForwardingRulesetResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The reference to the DNS resolver outbound endpoints that are used to route DNS queries matching the forwarding rules in the ruleset to the target DNS servers.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> DnsResolverOutboundEndpoints;
        /// <summary>
        /// ETag of the DNS forwarding ruleset.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resourceGuid for the DNS forwarding ruleset.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDnsForwardingRulesetResult(
            string azureApiVersion,

            ImmutableArray<Outputs.SubResourceResponse> dnsResolverOutboundEndpoints,

            string etag,

            string id,

            string location,

            string name,

            string provisioningState,

            string resourceGuid,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DnsResolverOutboundEndpoints = dnsResolverOutboundEndpoints;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
