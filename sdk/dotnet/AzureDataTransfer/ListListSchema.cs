// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer
{
    public static class ListListSchema
    {
        /// <summary>
        /// Lists the schemas for the specified connection in a pipeline.
        /// 
        /// Uses Azure REST API version 2024-09-27.
        /// 
        /// Other available API versions: 2023-10-11-preview, 2024-01-25, 2024-05-07, 2024-09-11, 2025-03-01-preview, 2025-04-11-preview, 2025-05-21. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListListSchemaResult> InvokeAsync(ListListSchemaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListListSchemaResult>("azure-native:azuredatatransfer:listListSchema", args ?? new ListListSchemaArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the schemas for the specified connection in a pipeline.
        /// 
        /// Uses Azure REST API version 2024-09-27.
        /// 
        /// Other available API versions: 2023-10-11-preview, 2024-01-25, 2024-05-07, 2024-09-11, 2025-03-01-preview, 2025-04-11-preview, 2025-05-21. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListListSchemaResult> Invoke(ListListSchemaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListListSchemaResult>("azure-native:azuredatatransfer:listListSchema", args ?? new ListListSchemaInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the schemas for the specified connection in a pipeline.
        /// 
        /// Uses Azure REST API version 2024-09-27.
        /// 
        /// Other available API versions: 2023-10-11-preview, 2024-01-25, 2024-05-07, 2024-09-11, 2025-03-01-preview, 2025-04-11-preview, 2025-05-21. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListListSchemaResult> Invoke(ListListSchemaInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListListSchemaResult>("azure-native:azuredatatransfer:listListSchema", args ?? new ListListSchemaInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListListSchemaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connection ID associated with this schema
        /// </summary>
        [Input("connectionId")]
        public string? ConnectionId { get; set; }

        /// <summary>
        /// Content of the schema
        /// </summary>
        [Input("content")]
        public string? Content { get; set; }

        /// <summary>
        /// The direction of the schema.
        /// </summary>
        [Input("direction")]
        public Union<string, Pulumi.AzureNative.AzureDataTransfer.SchemaDirection>? Direction { get; set; }

        /// <summary>
        /// ID associated with this schema
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Name of the schema
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The name for the pipeline that is to be requested.
        /// </summary>
        [Input("pipelineName", required: true)]
        public string PipelineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Schema Type
        /// </summary>
        [Input("schemaType")]
        public Union<string, Pulumi.AzureNative.AzureDataTransfer.SchemaType>? SchemaType { get; set; }

        /// <summary>
        /// Uri containing SAS token for the zipped schema
        /// </summary>
        [Input("schemaUri")]
        public string? SchemaUri { get; set; }

        /// <summary>
        /// Status of the schema
        /// </summary>
        [Input("status")]
        public Union<string, Pulumi.AzureNative.AzureDataTransfer.SchemaStatus>? Status { get; set; }

        public ListListSchemaArgs()
        {
        }
        public static new ListListSchemaArgs Empty => new ListListSchemaArgs();
    }

    public sealed class ListListSchemaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connection ID associated with this schema
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// Content of the schema
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The direction of the schema.
        /// </summary>
        [Input("direction")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaDirection>? Direction { get; set; }

        /// <summary>
        /// ID associated with this schema
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the schema
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name for the pipeline that is to be requested.
        /// </summary>
        [Input("pipelineName", required: true)]
        public Input<string> PipelineName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Schema Type
        /// </summary>
        [Input("schemaType")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaType>? SchemaType { get; set; }

        /// <summary>
        /// Uri containing SAS token for the zipped schema
        /// </summary>
        [Input("schemaUri")]
        public Input<string>? SchemaUri { get; set; }

        /// <summary>
        /// Status of the schema
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaStatus>? Status { get; set; }

        public ListListSchemaInvokeArgs()
        {
        }
        public static new ListListSchemaInvokeArgs Empty => new ListListSchemaInvokeArgs();
    }


    [OutputType]
    public sealed class ListListSchemaResult
    {
        /// <summary>
        /// Schemas array.
        /// </summary>
        public readonly ImmutableArray<Outputs.SchemaResponse> Value;

        [OutputConstructor]
        private ListListSchemaResult(ImmutableArray<Outputs.SchemaResponse> value)
        {
            Value = value;
        }
    }
}
