// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20230301Preview
{
    public static class GetContentType
    {
        /// <summary>
        /// Gets the details of the developer portal's content type. Content types describe content items' properties, validation rules, and constraints.
        /// </summary>
        public static Task<GetContentTypeResult> InvokeAsync(GetContentTypeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContentTypeResult>("azure-native:apimanagement/v20230301preview:getContentType", args ?? new GetContentTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the developer portal's content type. Content types describe content items' properties, validation rules, and constraints.
        /// </summary>
        public static Output<GetContentTypeResult> Invoke(GetContentTypeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContentTypeResult>("azure-native:apimanagement/v20230301preview:getContentType", args ?? new GetContentTypeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the developer portal's content type. Content types describe content items' properties, validation rules, and constraints.
        /// </summary>
        public static Output<GetContentTypeResult> Invoke(GetContentTypeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContentTypeResult>("azure-native:apimanagement/v20230301preview:getContentType", args ?? new GetContentTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContentTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Content type identifier.
        /// </summary>
        [Input("contentTypeId", required: true)]
        public string ContentTypeId { get; set; } = null!;

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

        public GetContentTypeArgs()
        {
        }
        public static new GetContentTypeArgs Empty => new GetContentTypeArgs();
    }

    public sealed class GetContentTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Content type identifier.
        /// </summary>
        [Input("contentTypeId", required: true)]
        public Input<string> ContentTypeId { get; set; } = null!;

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

        public GetContentTypeInvokeArgs()
        {
        }
        public static new GetContentTypeInvokeArgs Empty => new GetContentTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetContentTypeResult
    {
        /// <summary>
        /// Content type description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Content type schema.
        /// </summary>
        public readonly object? Schema;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Content type version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetContentTypeResult(
            string? description,

            string id,

            string name,

            object? schema,

            string type,

            string? version)
        {
            Description = description;
            Id = id;
            Name = name;
            Schema = schema;
            Type = type;
            Version = version;
        }
    }
}
