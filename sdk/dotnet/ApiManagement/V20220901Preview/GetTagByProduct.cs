// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20220901Preview
{
    public static class GetTagByProduct
    {
        /// <summary>
        /// Get tag associated with the Product.
        /// </summary>
        public static Task<GetTagByProductResult> InvokeAsync(GetTagByProductArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagByProductResult>("azure-native:apimanagement/v20220901preview:getTagByProduct", args ?? new GetTagByProductArgs(), options.WithDefaults());

        /// <summary>
        /// Get tag associated with the Product.
        /// </summary>
        public static Output<GetTagByProductResult> Invoke(GetTagByProductInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagByProductResult>("azure-native:apimanagement/v20220901preview:getTagByProduct", args ?? new GetTagByProductInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get tag associated with the Product.
        /// </summary>
        public static Output<GetTagByProductResult> Invoke(GetTagByProductInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagByProductResult>("azure-native:apimanagement/v20220901preview:getTagByProduct", args ?? new GetTagByProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagByProductArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId", required: true)]
        public string TagId { get; set; } = null!;

        public GetTagByProductArgs()
        {
        }
        public static new GetTagByProductArgs Empty => new GetTagByProductArgs();
    }

    public sealed class GetTagByProductInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId", required: true)]
        public Input<string> TagId { get; set; } = null!;

        public GetTagByProductInvokeArgs()
        {
        }
        public static new GetTagByProductInvokeArgs Empty => new GetTagByProductInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagByProductResult
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagByProductResult(
            string displayName,

            string id,

            string name,

            string type)
        {
            DisplayName = displayName;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
