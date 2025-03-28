// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder
{
    public static class ListProductsAndConfigurations
    {
        /// <summary>
        /// List configurations for the given product family, product line and product for the given subscription.
        /// 
        /// Uses Azure REST API version 2022-05-01-preview.
        /// 
        /// Other available API versions: 2024-02-01.
        /// </summary>
        public static Task<ListProductsAndConfigurationsResult> InvokeAsync(ListProductsAndConfigurationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListProductsAndConfigurationsResult>("azure-native:edgeorder:listProductsAndConfigurations", args ?? new ListProductsAndConfigurationsArgs(), options.WithDefaults());

        /// <summary>
        /// List configurations for the given product family, product line and product for the given subscription.
        /// 
        /// Uses Azure REST API version 2022-05-01-preview.
        /// 
        /// Other available API versions: 2024-02-01.
        /// </summary>
        public static Output<ListProductsAndConfigurationsResult> Invoke(ListProductsAndConfigurationsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsAndConfigurationsResult>("azure-native:edgeorder:listProductsAndConfigurations", args ?? new ListProductsAndConfigurationsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List configurations for the given product family, product line and product for the given subscription.
        /// 
        /// Uses Azure REST API version 2022-05-01-preview.
        /// 
        /// Other available API versions: 2024-02-01.
        /// </summary>
        public static Output<ListProductsAndConfigurationsResult> Invoke(ListProductsAndConfigurationsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsAndConfigurationsResult>("azure-native:edgeorder:listProductsAndConfigurations", args ?? new ListProductsAndConfigurationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListProductsAndConfigurationsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Holds details about product hierarchy information and filterable property.
        /// </summary>
        [Input("configurationFilter")]
        public Inputs.ConfigurationFilter? ConfigurationFilter { get; set; }

        /// <summary>
        /// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
        /// </summary>
        [Input("customerSubscriptionDetails")]
        public Inputs.CustomerSubscriptionDetails? CustomerSubscriptionDetails { get; set; }

        /// <summary>
        /// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
        /// </summary>
        [Input("skipToken")]
        public string? SkipToken { get; set; }

        public ListProductsAndConfigurationsArgs()
        {
        }
        public static new ListProductsAndConfigurationsArgs Empty => new ListProductsAndConfigurationsArgs();
    }

    public sealed class ListProductsAndConfigurationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Holds details about product hierarchy information and filterable property.
        /// </summary>
        [Input("configurationFilter")]
        public Input<Inputs.ConfigurationFilterArgs>? ConfigurationFilter { get; set; }

        /// <summary>
        /// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
        /// </summary>
        [Input("customerSubscriptionDetails")]
        public Input<Inputs.CustomerSubscriptionDetailsArgs>? CustomerSubscriptionDetails { get; set; }

        /// <summary>
        /// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
        /// </summary>
        [Input("skipToken")]
        public Input<string>? SkipToken { get; set; }

        public ListProductsAndConfigurationsInvokeArgs()
        {
        }
        public static new ListProductsAndConfigurationsInvokeArgs Empty => new ListProductsAndConfigurationsInvokeArgs();
    }


    [OutputType]
    public sealed class ListProductsAndConfigurationsResult
    {
        /// <summary>
        /// Link for the next set of configurations.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigurationResponse> Value;

        [OutputConstructor]
        private ListProductsAndConfigurationsResult(
            string? nextLink,

            ImmutableArray<Outputs.ConfigurationResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
