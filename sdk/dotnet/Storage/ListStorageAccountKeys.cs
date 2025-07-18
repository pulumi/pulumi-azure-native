// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage
{
    public static class ListStorageAccountKeys
    {
        /// <summary>
        /// Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListStorageAccountKeysResult> InvokeAsync(ListStorageAccountKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListStorageAccountKeysResult>("azure-native:storage:listStorageAccountKeys", args ?? new ListStorageAccountKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListStorageAccountKeysResult> Invoke(ListStorageAccountKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListStorageAccountKeysResult>("azure-native:storage:listStorageAccountKeys", args ?? new ListStorageAccountKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the access keys or Kerberos keys (if active directory enabled) for the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListStorageAccountKeysResult> Invoke(ListStorageAccountKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListStorageAccountKeysResult>("azure-native:storage:listStorageAccountKeys", args ?? new ListStorageAccountKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListStorageAccountKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies type of the key to be listed. Possible value is kerb.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListStorageAccountKeysArgs()
        {
        }
        public static new ListStorageAccountKeysArgs Empty => new ListStorageAccountKeysArgs();
    }

    public sealed class ListStorageAccountKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Specifies type of the key to be listed. Possible value is kerb.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListStorageAccountKeysInvokeArgs()
        {
        }
        public static new ListStorageAccountKeysInvokeArgs Empty => new ListStorageAccountKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListStorageAccountKeysResult
    {
        /// <summary>
        /// Gets the list of storage account keys and their properties for the specified storage account.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountKeyResponse> Keys;

        [OutputConstructor]
        private ListStorageAccountKeysResult(ImmutableArray<Outputs.StorageAccountKeyResponse> keys)
        {
            Keys = keys;
        }
    }
}
