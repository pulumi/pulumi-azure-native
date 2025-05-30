// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB
{
    public static class GetMongoClusterFirewallRule
    {
        /// <summary>
        /// Gets information about a mongo cluster firewall rule.
        /// 
        /// Uses Azure REST API version 2024-02-15-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-03-15-preview, 2023-09-15-preview, 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMongoClusterFirewallRuleResult> InvokeAsync(GetMongoClusterFirewallRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMongoClusterFirewallRuleResult>("azure-native:cosmosdb:getMongoClusterFirewallRule", args ?? new GetMongoClusterFirewallRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a mongo cluster firewall rule.
        /// 
        /// Uses Azure REST API version 2024-02-15-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-03-15-preview, 2023-09-15-preview, 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMongoClusterFirewallRuleResult> Invoke(GetMongoClusterFirewallRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoClusterFirewallRuleResult>("azure-native:cosmosdb:getMongoClusterFirewallRule", args ?? new GetMongoClusterFirewallRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a mongo cluster firewall rule.
        /// 
        /// Uses Azure REST API version 2024-02-15-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-03-15-preview, 2023-09-15-preview, 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMongoClusterFirewallRuleResult> Invoke(GetMongoClusterFirewallRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongoClusterFirewallRuleResult>("azure-native:cosmosdb:getMongoClusterFirewallRule", args ?? new GetMongoClusterFirewallRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMongoClusterFirewallRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the mongo cluster firewall rule.
        /// </summary>
        [Input("firewallRuleName", required: true)]
        public string FirewallRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the mongo cluster.
        /// </summary>
        [Input("mongoClusterName", required: true)]
        public string MongoClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMongoClusterFirewallRuleArgs()
        {
        }
        public static new GetMongoClusterFirewallRuleArgs Empty => new GetMongoClusterFirewallRuleArgs();
    }

    public sealed class GetMongoClusterFirewallRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the mongo cluster firewall rule.
        /// </summary>
        [Input("firewallRuleName", required: true)]
        public Input<string> FirewallRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the mongo cluster.
        /// </summary>
        [Input("mongoClusterName", required: true)]
        public Input<string> MongoClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetMongoClusterFirewallRuleInvokeArgs()
        {
        }
        public static new GetMongoClusterFirewallRuleInvokeArgs Empty => new GetMongoClusterFirewallRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetMongoClusterFirewallRuleResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        public readonly string EndIpAddress;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the firewall rule.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
        /// </summary>
        public readonly string StartIpAddress;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMongoClusterFirewallRuleResult(
            string azureApiVersion,

            string endIpAddress,

            string id,

            string name,

            string provisioningState,

            string startIpAddress,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            EndIpAddress = endIpAddress;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            StartIpAddress = startIpAddress;
            SystemData = systemData;
            Type = type;
        }
    }
}
