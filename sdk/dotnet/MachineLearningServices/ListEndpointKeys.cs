// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class ListEndpointKeys
    {
        /// <summary>
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListEndpointKeysResult> InvokeAsync(ListEndpointKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEndpointKeysResult>("azure-native:machinelearningservices:listEndpointKeys", args ?? new ListEndpointKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListEndpointKeysResult> Invoke(ListEndpointKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEndpointKeysResult>("azure-native:machinelearningservices:listEndpointKeys", args ?? new ListEndpointKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2024-01-01-preview, 2024-07-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListEndpointKeysResult> Invoke(ListEndpointKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListEndpointKeysResult>("azure-native:machinelearningservices:listEndpointKeys", args ?? new ListEndpointKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListEndpointKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint resource.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Azure Machine Learning Workspace Name
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListEndpointKeysArgs()
        {
        }
        public static new ListEndpointKeysArgs Empty => new ListEndpointKeysArgs();
    }

    public sealed class ListEndpointKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint resource.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Azure Machine Learning Workspace Name
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ListEndpointKeysInvokeArgs()
        {
        }
        public static new ListEndpointKeysInvokeArgs Empty => new ListEndpointKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListEndpointKeysResult
    {
        /// <summary>
        /// Dictionary of Keys for the endpoint.
        /// </summary>
        public readonly Outputs.AccountApiKeysResponse? Keys;

        [OutputConstructor]
        private ListEndpointKeysResult(Outputs.AccountApiKeysResponse? keys)
        {
            Keys = keys;
        }
    }
}
