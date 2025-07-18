// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class ListInferenceGroupDeltaModelsAsync
    {
        /// <summary>
        /// A paginated list of String entities.
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListInferenceGroupDeltaModelsAsyncResult> InvokeAsync(ListInferenceGroupDeltaModelsAsyncArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListInferenceGroupDeltaModelsAsyncResult>("azure-native:machinelearningservices:listInferenceGroupDeltaModelsAsync", args ?? new ListInferenceGroupDeltaModelsAsyncArgs(), options.WithDefaults());

        /// <summary>
        /// A paginated list of String entities.
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListInferenceGroupDeltaModelsAsyncResult> Invoke(ListInferenceGroupDeltaModelsAsyncInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListInferenceGroupDeltaModelsAsyncResult>("azure-native:machinelearningservices:listInferenceGroupDeltaModelsAsync", args ?? new ListInferenceGroupDeltaModelsAsyncInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A paginated list of String entities.
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListInferenceGroupDeltaModelsAsyncResult> Invoke(ListInferenceGroupDeltaModelsAsyncInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListInferenceGroupDeltaModelsAsyncResult>("azure-native:machinelearningservices:listInferenceGroupDeltaModelsAsync", args ?? new ListInferenceGroupDeltaModelsAsyncInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListInferenceGroupDeltaModelsAsyncArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets or sets number of delta models to return. Default: -1, means that all will be returned.
        /// </summary>
        [Input("count")]
        public int? Count { get; set; }

        /// <summary>
        /// InferenceGroup name.
        /// </summary>
        [Input("groupName", required: true)]
        public string GroupName { get; set; } = null!;

        /// <summary>
        /// InferencePool name.
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets skip token for paginated response.
        /// </summary>
        [Input("skipToken")]
        public string? SkipToken { get; set; }

        /// <summary>
        /// Gets or sets target base model.
        /// </summary>
        [Input("targetBaseModel")]
        public string? TargetBaseModel { get; set; }

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListInferenceGroupDeltaModelsAsyncArgs()
        {
            Count = -1;
        }
        public static new ListInferenceGroupDeltaModelsAsyncArgs Empty => new ListInferenceGroupDeltaModelsAsyncArgs();
    }

    public sealed class ListInferenceGroupDeltaModelsAsyncInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets or sets number of delta models to return. Default: -1, means that all will be returned.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// InferenceGroup name.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// InferencePool name.
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets skip token for paginated response.
        /// </summary>
        [Input("skipToken")]
        public Input<string>? SkipToken { get; set; }

        /// <summary>
        /// Gets or sets target base model.
        /// </summary>
        [Input("targetBaseModel")]
        public Input<string>? TargetBaseModel { get; set; }

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ListInferenceGroupDeltaModelsAsyncInvokeArgs()
        {
            Count = -1;
        }
        public static new ListInferenceGroupDeltaModelsAsyncInvokeArgs Empty => new ListInferenceGroupDeltaModelsAsyncInvokeArgs();
    }


    [OutputType]
    public sealed class ListInferenceGroupDeltaModelsAsyncResult
    {
        /// <summary>
        /// The link to the next page of String objects. If null, there are no additional pages.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// An array of objects of type String.
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private ListInferenceGroupDeltaModelsAsyncResult(
            string? nextLink,

            ImmutableArray<string> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
