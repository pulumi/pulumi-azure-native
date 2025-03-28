// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20230601
{
    public static class ListFirewallPolicyIdpsSignaturesFilterValue
    {
        /// <summary>
        /// Retrieves the current filter values for the signatures overrides
        /// </summary>
        public static Task<ListFirewallPolicyIdpsSignaturesFilterValueResult> InvokeAsync(ListFirewallPolicyIdpsSignaturesFilterValueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListFirewallPolicyIdpsSignaturesFilterValueResult>("azure-native:network/v20230601:listFirewallPolicyIdpsSignaturesFilterValue", args ?? new ListFirewallPolicyIdpsSignaturesFilterValueArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current filter values for the signatures overrides
        /// </summary>
        public static Output<ListFirewallPolicyIdpsSignaturesFilterValueResult> Invoke(ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListFirewallPolicyIdpsSignaturesFilterValueResult>("azure-native:network/v20230601:listFirewallPolicyIdpsSignaturesFilterValue", args ?? new ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current filter values for the signatures overrides
        /// </summary>
        public static Output<ListFirewallPolicyIdpsSignaturesFilterValueResult> Invoke(ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListFirewallPolicyIdpsSignaturesFilterValueResult>("azure-native:network/v20230601:listFirewallPolicyIdpsSignaturesFilterValue", args ?? new ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListFirewallPolicyIdpsSignaturesFilterValueArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Describes the name of the column which values will be returned
        /// </summary>
        [Input("filterName")]
        public string? FilterName { get; set; }

        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public string FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListFirewallPolicyIdpsSignaturesFilterValueArgs()
        {
        }
        public static new ListFirewallPolicyIdpsSignaturesFilterValueArgs Empty => new ListFirewallPolicyIdpsSignaturesFilterValueArgs();
    }

    public sealed class ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Describes the name of the column which values will be returned
        /// </summary>
        [Input("filterName")]
        public Input<string>? FilterName { get; set; }

        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public Input<string> FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs()
        {
        }
        public static new ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs Empty => new ListFirewallPolicyIdpsSignaturesFilterValueInvokeArgs();
    }


    [OutputType]
    public sealed class ListFirewallPolicyIdpsSignaturesFilterValueResult
    {
        /// <summary>
        /// Describes the possible values
        /// </summary>
        public readonly ImmutableArray<string> FilterValues;

        [OutputConstructor]
        private ListFirewallPolicyIdpsSignaturesFilterValueResult(ImmutableArray<string> filterValues)
        {
            FilterValues = filterValues;
        }
    }
}
