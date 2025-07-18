// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder
{
    public static class ListProductsAndConfigurationProductFamilies
    {
        /// <summary>
        /// List product families for the given subscription.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native edgeorder [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListProductsAndConfigurationProductFamiliesResult> InvokeAsync(ListProductsAndConfigurationProductFamiliesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListProductsAndConfigurationProductFamiliesResult>("azure-native:edgeorder:listProductsAndConfigurationProductFamilies", args ?? new ListProductsAndConfigurationProductFamiliesArgs(), options.WithDefaults());

        /// <summary>
        /// List product families for the given subscription.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native edgeorder [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListProductsAndConfigurationProductFamiliesResult> Invoke(ListProductsAndConfigurationProductFamiliesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsAndConfigurationProductFamiliesResult>("azure-native:edgeorder:listProductsAndConfigurationProductFamilies", args ?? new ListProductsAndConfigurationProductFamiliesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List product families for the given subscription.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native edgeorder [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListProductsAndConfigurationProductFamiliesResult> Invoke(ListProductsAndConfigurationProductFamiliesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsAndConfigurationProductFamiliesResult>("azure-native:edgeorder:listProductsAndConfigurationProductFamilies", args ?? new ListProductsAndConfigurationProductFamiliesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListProductsAndConfigurationProductFamiliesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
        /// </summary>
        [Input("customerSubscriptionDetails")]
        public Inputs.CustomerSubscriptionDetails? CustomerSubscriptionDetails { get; set; }

        /// <summary>
        /// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        [Input("filterableProperties", required: true)]
        private Dictionary<string, ImmutableArray<Inputs.FilterableProperty>>? _filterableProperties;

        /// <summary>
        /// Dictionary of filterable properties on product family.
        /// </summary>
        public Dictionary<string, ImmutableArray<Inputs.FilterableProperty>> FilterableProperties
        {
            get => _filterableProperties ?? (_filterableProperties = new Dictionary<string, ImmutableArray<Inputs.FilterableProperty>>());
            set => _filterableProperties = value;
        }

        /// <summary>
        /// $skipToken is supported on list of product families, which provides the next page in the list of product families.
        /// </summary>
        [Input("skipToken")]
        public string? SkipToken { get; set; }

        public ListProductsAndConfigurationProductFamiliesArgs()
        {
        }
        public static new ListProductsAndConfigurationProductFamiliesArgs Empty => new ListProductsAndConfigurationProductFamiliesArgs();
    }

    public sealed class ListProductsAndConfigurationProductFamiliesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details.
        /// </summary>
        [Input("customerSubscriptionDetails")]
        public Input<Inputs.CustomerSubscriptionDetailsArgs>? CustomerSubscriptionDetails { get; set; }

        /// <summary>
        /// $expand is supported on configurations parameter for product, which provides details on the configurations for the product.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        [Input("filterableProperties", required: true)]
        private InputMap<ImmutableArray<Inputs.FilterablePropertyArgs>>? _filterableProperties;

        /// <summary>
        /// Dictionary of filterable properties on product family.
        /// </summary>
        public InputMap<ImmutableArray<Inputs.FilterablePropertyArgs>> FilterableProperties
        {
            get => _filterableProperties ?? (_filterableProperties = new InputMap<ImmutableArray<Inputs.FilterablePropertyArgs>>());
            set => _filterableProperties = value;
        }

        /// <summary>
        /// $skipToken is supported on list of product families, which provides the next page in the list of product families.
        /// </summary>
        [Input("skipToken")]
        public Input<string>? SkipToken { get; set; }

        public ListProductsAndConfigurationProductFamiliesInvokeArgs()
        {
        }
        public static new ListProductsAndConfigurationProductFamiliesInvokeArgs Empty => new ListProductsAndConfigurationProductFamiliesInvokeArgs();
    }


    [OutputType]
    public sealed class ListProductsAndConfigurationProductFamiliesResult
    {
        /// <summary>
        /// The link to the next page of items
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// The ProductFamily items on this page
        /// </summary>
        public readonly ImmutableArray<Outputs.ProductFamilyResponse> Value;

        [OutputConstructor]
        private ListProductsAndConfigurationProductFamiliesResult(
            string? nextLink,

            ImmutableArray<Outputs.ProductFamilyResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
