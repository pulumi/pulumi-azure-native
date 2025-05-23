// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataLakeStore
{
    public static class GetVirtualNetworkRule
    {
        /// <summary>
        /// Gets the specified Data Lake Store virtual network rule.
        /// 
        /// Uses Azure REST API version 2016-11-01.
        /// </summary>
        public static Task<GetVirtualNetworkRuleResult> InvokeAsync(GetVirtualNetworkRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkRuleResult>("azure-native:datalakestore:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Data Lake Store virtual network rule.
        /// 
        /// Uses Azure REST API version 2016-11-01.
        /// </summary>
        public static Output<GetVirtualNetworkRuleResult> Invoke(GetVirtualNetworkRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkRuleResult>("azure-native:datalakestore:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Data Lake Store virtual network rule.
        /// 
        /// Uses Azure REST API version 2016-11-01.
        /// </summary>
        public static Output<GetVirtualNetworkRuleResult> Invoke(GetVirtualNetworkRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkRuleResult>("azure-native:datalakestore:getVirtualNetworkRule", args ?? new GetVirtualNetworkRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Data Lake Store account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule to retrieve.
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
        /// The name of the Data Lake Store account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule to retrieve.
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
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource identifier for the subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVirtualNetworkRuleResult(
            string azureApiVersion,

            string id,

            string name,

            string subnetId,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            SubnetId = subnetId;
            Type = type;
        }
    }
}
