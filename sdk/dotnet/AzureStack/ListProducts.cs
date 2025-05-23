// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStack
{
    public static class ListProducts
    {
        /// <summary>
        /// Returns a list of products.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Task<ListProductsResult> InvokeAsync(ListProductsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListProductsResult>("azure-native:azurestack:listProducts", args ?? new ListProductsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of products.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Output<ListProductsResult> Invoke(ListProductsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsResult>("azure-native:azurestack:listProducts", args ?? new ListProductsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of products.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Output<ListProductsResult> Invoke(ListProductsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListProductsResult>("azure-native:azurestack:listProducts", args ?? new ListProductsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListProductsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("productName", required: true)]
        public string ProductName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public string RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public string ResourceGroup { get; set; } = null!;

        public ListProductsArgs()
        {
        }
        public static new ListProductsArgs Empty => new ListProductsArgs();
    }

    public sealed class ListProductsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the product.
        /// </summary>
        [Input("productName", required: true)]
        public Input<string> ProductName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public Input<string> RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        public ListProductsInvokeArgs()
        {
        }
        public static new ListProductsInvokeArgs Empty => new ListProductsInvokeArgs();
    }


    [OutputType]
    public sealed class ListProductsResult
    {
        /// <summary>
        /// URI to the next page.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of products.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProductResponse> Value;

        [OutputConstructor]
        private ListProductsResult(
            string? nextLink,

            ImmutableArray<Outputs.ProductResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
