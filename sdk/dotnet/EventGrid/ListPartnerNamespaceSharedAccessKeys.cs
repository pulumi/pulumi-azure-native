// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class ListPartnerNamespaceSharedAccessKeys
    {
        /// <summary>
        /// List the two keys used to publish to a partner namespace.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListPartnerNamespaceSharedAccessKeysResult> InvokeAsync(ListPartnerNamespaceSharedAccessKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListPartnerNamespaceSharedAccessKeysResult>("azure-native:eventgrid:listPartnerNamespaceSharedAccessKeys", args ?? new ListPartnerNamespaceSharedAccessKeysArgs(), options.WithDefaults());

        /// <summary>
        /// List the two keys used to publish to a partner namespace.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPartnerNamespaceSharedAccessKeysResult> Invoke(ListPartnerNamespaceSharedAccessKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListPartnerNamespaceSharedAccessKeysResult>("azure-native:eventgrid:listPartnerNamespaceSharedAccessKeys", args ?? new ListPartnerNamespaceSharedAccessKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List the two keys used to publish to a partner namespace.
        /// 
        /// Uses Azure REST API version 2025-02-15.
        /// 
        /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListPartnerNamespaceSharedAccessKeysResult> Invoke(ListPartnerNamespaceSharedAccessKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListPartnerNamespaceSharedAccessKeysResult>("azure-native:eventgrid:listPartnerNamespaceSharedAccessKeys", args ?? new ListPartnerNamespaceSharedAccessKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListPartnerNamespaceSharedAccessKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the partner namespace.
        /// </summary>
        [Input("partnerNamespaceName", required: true)]
        public string PartnerNamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListPartnerNamespaceSharedAccessKeysArgs()
        {
        }
        public static new ListPartnerNamespaceSharedAccessKeysArgs Empty => new ListPartnerNamespaceSharedAccessKeysArgs();
    }

    public sealed class ListPartnerNamespaceSharedAccessKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the partner namespace.
        /// </summary>
        [Input("partnerNamespaceName", required: true)]
        public Input<string> PartnerNamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListPartnerNamespaceSharedAccessKeysInvokeArgs()
        {
        }
        public static new ListPartnerNamespaceSharedAccessKeysInvokeArgs Empty => new ListPartnerNamespaceSharedAccessKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListPartnerNamespaceSharedAccessKeysResult
    {
        /// <summary>
        /// Shared access key1 for the partner namespace.
        /// </summary>
        public readonly string? Key1;
        /// <summary>
        /// Shared access key2 for the partner namespace.
        /// </summary>
        public readonly string? Key2;

        [OutputConstructor]
        private ListPartnerNamespaceSharedAccessKeysResult(
            string? key1,

            string? key2)
        {
            Key1 = key1;
            Key2 = key2;
        }
    }
}
