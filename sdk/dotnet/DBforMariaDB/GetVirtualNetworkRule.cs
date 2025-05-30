// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforMariaDB
{
    public static class GetVirtualNetworkRule
    {
        /// <summary>
        /// Gets a virtual network rule.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Task<GetVirtualNetworkRuleResult> InvokeAsync(GetVirtualNetworkRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkRuleResult>("azure-native:dbformariadb:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a virtual network rule.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Output<GetVirtualNetworkRuleResult> Invoke(GetVirtualNetworkRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkRuleResult>("azure-native:dbformariadb:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a virtual network rule.
        /// 
        /// Uses Azure REST API version 2018-06-01.
        /// </summary>
        public static Output<GetVirtualNetworkRuleResult> Invoke(GetVirtualNetworkRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkRuleResult>("azure-native:dbformariadb:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule.
        /// </summary>
        [Input("virtualNetworkRuleName", required: true)]
        public string VirtualNetworkRuleName { get; set; } = null!;

        public GetVirtualNetworkRuleArgs()
        {
        }
        public static new GetVirtualNetworkRuleArgs Empty => new GetVirtualNetworkRuleArgs();
    }

    public sealed class GetVirtualNetworkRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule.
        /// </summary>
        [Input("virtualNetworkRuleName", required: true)]
        public Input<string> VirtualNetworkRuleName { get; set; } = null!;

        public GetVirtualNetworkRuleInvokeArgs()
        {
        }
        public static new GetVirtualNetworkRuleInvokeArgs Empty => new GetVirtualNetworkRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualNetworkRuleResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Create firewall rule before the virtual network has vnet service endpoint enabled.
        /// </summary>
        public readonly bool? IgnoreMissingVnetServiceEndpoint;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Virtual Network Rule State
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The ARM resource id of the virtual network subnet.
        /// </summary>
        public readonly string VirtualNetworkSubnetId;

        [OutputConstructor]
        private GetVirtualNetworkRuleResult(
            string azureApiVersion,

            string id,

            bool? ignoreMissingVnetServiceEndpoint,

            string name,

            string state,

            string type,

            string virtualNetworkSubnetId)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            IgnoreMissingVnetServiceEndpoint = ignoreMissingVnetServiceEndpoint;
            Name = name;
            State = state;
            Type = type;
            VirtualNetworkSubnetId = virtualNetworkSubnetId;
        }
    }
}
