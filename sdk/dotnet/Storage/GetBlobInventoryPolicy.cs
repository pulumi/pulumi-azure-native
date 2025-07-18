// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage
{
    public static class GetBlobInventoryPolicy
    {
        /// <summary>
        /// Gets the blob inventory policy associated with the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetBlobInventoryPolicyResult> InvokeAsync(GetBlobInventoryPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBlobInventoryPolicyResult>("azure-native:storage:getBlobInventoryPolicy", args ?? new GetBlobInventoryPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the blob inventory policy associated with the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBlobInventoryPolicyResult> Invoke(GetBlobInventoryPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlobInventoryPolicyResult>("azure-native:storage:getBlobInventoryPolicy", args ?? new GetBlobInventoryPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the blob inventory policy associated with the specified storage account.
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetBlobInventoryPolicyResult> Invoke(GetBlobInventoryPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlobInventoryPolicyResult>("azure-native:storage:getBlobInventoryPolicy", args ?? new GetBlobInventoryPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBlobInventoryPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the storage account blob inventory policy. It should always be 'default'
        /// </summary>
        [Input("blobInventoryPolicyName", required: true)]
        public string BlobInventoryPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBlobInventoryPolicyArgs()
        {
        }
        public static new GetBlobInventoryPolicyArgs Empty => new GetBlobInventoryPolicyArgs();
    }

    public sealed class GetBlobInventoryPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the storage account blob inventory policy. It should always be 'default'
        /// </summary>
        [Input("blobInventoryPolicyName", required: true)]
        public Input<string> BlobInventoryPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetBlobInventoryPolicyInvokeArgs()
        {
        }
        public static new GetBlobInventoryPolicyInvokeArgs Empty => new GetBlobInventoryPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetBlobInventoryPolicyResult
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
        /// Returns the last modified date and time of the blob inventory policy.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The storage account blob inventory policy object. It is composed of policy rules.
        /// </summary>
        public readonly Outputs.BlobInventoryPolicySchemaResponse Policy;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBlobInventoryPolicyResult(
            string azureApiVersion,

            string id,

            string lastModifiedTime,

            string name,

            Outputs.BlobInventoryPolicySchemaResponse policy,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Policy = policy;
            SystemData = systemData;
            Type = type;
        }
    }
}
