// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20230901Preview
{
    public static class GetTagOperationLink
    {
        /// <summary>
        /// Gets the operation link for the tag.
        /// </summary>
        public static Task<GetTagOperationLinkResult> InvokeAsync(GetTagOperationLinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagOperationLinkResult>("azure-native:apimanagement/v20230901preview:getTagOperationLink", args ?? new GetTagOperationLinkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the operation link for the tag.
        /// </summary>
        public static Output<GetTagOperationLinkResult> Invoke(GetTagOperationLinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagOperationLinkResult>("azure-native:apimanagement/v20230901preview:getTagOperationLink", args ?? new GetTagOperationLinkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the operation link for the tag.
        /// </summary>
        public static Output<GetTagOperationLinkResult> Invoke(GetTagOperationLinkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagOperationLinkResult>("azure-native:apimanagement/v20230901preview:getTagOperationLink", args ?? new GetTagOperationLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagOperationLinkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Tag-operation link identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationLinkId", required: true)]
        public string OperationLinkId { get; set; } = null!;

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

        public GetTagOperationLinkArgs()
        {
        }
        public static new GetTagOperationLinkArgs Empty => new GetTagOperationLinkArgs();
    }

    public sealed class GetTagOperationLinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Tag-operation link identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationLinkId", required: true)]
        public Input<string> OperationLinkId { get; set; } = null!;

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

        public GetTagOperationLinkInvokeArgs()
        {
        }
        public static new GetTagOperationLinkInvokeArgs Empty => new GetTagOperationLinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagOperationLinkResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Full resource Id of an API operation.
        /// </summary>
        public readonly string OperationId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagOperationLinkResult(
            string id,

            string name,

            string operationId,

            string type)
        {
            Id = id;
            Name = name;
            OperationId = operationId;
            Type = type;
        }
    }
}
