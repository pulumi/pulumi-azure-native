// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService
{
    public static class ListSignalRKeys
    {
        /// <summary>
        /// Get the access keys of the resource.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native signalrservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListSignalRKeysResult> InvokeAsync(ListSignalRKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListSignalRKeysResult>("azure-native:signalrservice:listSignalRKeys", args ?? new ListSignalRKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Get the access keys of the resource.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native signalrservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListSignalRKeysResult> Invoke(ListSignalRKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListSignalRKeysResult>("azure-native:signalrservice:listSignalRKeys", args ?? new ListSignalRKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the access keys of the resource.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native signalrservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListSignalRKeysResult> Invoke(ListSignalRKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListSignalRKeysResult>("azure-native:signalrservice:listSignalRKeys", args ?? new ListSignalRKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListSignalRKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListSignalRKeysArgs()
        {
        }
        public static new ListSignalRKeysArgs Empty => new ListSignalRKeysArgs();
    }

    public sealed class ListSignalRKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public ListSignalRKeysInvokeArgs()
        {
        }
        public static new ListSignalRKeysInvokeArgs Empty => new ListSignalRKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListSignalRKeysResult
    {
        /// <summary>
        /// Connection string constructed via the primaryKey
        /// </summary>
        public readonly string? PrimaryConnectionString;
        /// <summary>
        /// The primary access key.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// Connection string constructed via the secondaryKey
        /// </summary>
        public readonly string? SecondaryConnectionString;
        /// <summary>
        /// The secondary access key.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListSignalRKeysResult(
            string? primaryConnectionString,

            string? primaryKey,

            string? secondaryConnectionString,

            string? secondaryKey)
        {
            PrimaryConnectionString = primaryConnectionString;
            PrimaryKey = primaryKey;
            SecondaryConnectionString = secondaryConnectionString;
            SecondaryKey = secondaryKey;
        }
    }
}
