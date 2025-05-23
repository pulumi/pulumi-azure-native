// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Syntex
{
    public static class GetDocumentProcessor
    {
        /// <summary>
        /// Returns a document processor for a given name.
        /// 
        /// Uses Azure REST API version 2022-09-15-preview.
        /// </summary>
        public static Task<GetDocumentProcessorResult> InvokeAsync(GetDocumentProcessorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDocumentProcessorResult>("azure-native:syntex:getDocumentProcessor", args ?? new GetDocumentProcessorArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a document processor for a given name.
        /// 
        /// Uses Azure REST API version 2022-09-15-preview.
        /// </summary>
        public static Output<GetDocumentProcessorResult> Invoke(GetDocumentProcessorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentProcessorResult>("azure-native:syntex:getDocumentProcessor", args ?? new GetDocumentProcessorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a document processor for a given name.
        /// 
        /// Uses Azure REST API version 2022-09-15-preview.
        /// </summary>
        public static Output<GetDocumentProcessorResult> Invoke(GetDocumentProcessorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDocumentProcessorResult>("azure-native:syntex:getDocumentProcessor", args ?? new GetDocumentProcessorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentProcessorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of document processor resource.
        /// </summary>
        [Input("processorName", required: true)]
        public string ProcessorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDocumentProcessorArgs()
        {
        }
        public static new GetDocumentProcessorArgs Empty => new GetDocumentProcessorArgs();
    }

    public sealed class GetDocumentProcessorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of document processor resource.
        /// </summary>
        [Input("processorName", required: true)]
        public Input<string> ProcessorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDocumentProcessorInvokeArgs()
        {
        }
        public static new GetDocumentProcessorInvokeArgs Empty => new GetDocumentProcessorInvokeArgs();
    }


    [OutputType]
    public sealed class GetDocumentProcessorResult
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
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Document processor properties.
        /// </summary>
        public readonly Outputs.DocumentProcessorPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDocumentProcessorResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.DocumentProcessorPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
