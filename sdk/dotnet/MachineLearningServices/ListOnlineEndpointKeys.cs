// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    public static class ListOnlineEndpointKeys
    {
        /// <summary>
        /// Keys for endpoint authentication.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListOnlineEndpointKeysResult> InvokeAsync(ListOnlineEndpointKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListOnlineEndpointKeysResult>("azure-native:machinelearningservices:listOnlineEndpointKeys", args ?? new ListOnlineEndpointKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Keys for endpoint authentication.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListOnlineEndpointKeysResult> Invoke(ListOnlineEndpointKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListOnlineEndpointKeysResult>("azure-native:machinelearningservices:listOnlineEndpointKeys", args ?? new ListOnlineEndpointKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Keys for endpoint authentication.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListOnlineEndpointKeysResult> Invoke(ListOnlineEndpointKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListOnlineEndpointKeysResult>("azure-native:machinelearningservices:listOnlineEndpointKeys", args ?? new ListOnlineEndpointKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListOnlineEndpointKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Online Endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public ListOnlineEndpointKeysArgs()
        {
        }
        public static new ListOnlineEndpointKeysArgs Empty => new ListOnlineEndpointKeysArgs();
    }

    public sealed class ListOnlineEndpointKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Online Endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ListOnlineEndpointKeysInvokeArgs()
        {
        }
        public static new ListOnlineEndpointKeysInvokeArgs Empty => new ListOnlineEndpointKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListOnlineEndpointKeysResult
    {
        /// <summary>
        /// The primary key.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// The secondary key.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListOnlineEndpointKeysResult(
            string? primaryKey,

            string? secondaryKey)
        {
            PrimaryKey = primaryKey;
            SecondaryKey = secondaryKey;
        }
    }
}
