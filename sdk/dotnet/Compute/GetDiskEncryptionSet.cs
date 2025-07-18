// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute
{
    public static class GetDiskEncryptionSet
    {
        /// <summary>
        /// Gets information about a disk encryption set.
        /// 
        /// Uses Azure REST API version 2024-03-02.
        /// 
        /// Other available API versions: 2022-07-02, 2023-01-02, 2023-04-02, 2023-10-02, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDiskEncryptionSetResult> InvokeAsync(GetDiskEncryptionSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiskEncryptionSetResult>("azure-native:compute:getDiskEncryptionSet", args ?? new GetDiskEncryptionSetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a disk encryption set.
        /// 
        /// Uses Azure REST API version 2024-03-02.
        /// 
        /// Other available API versions: 2022-07-02, 2023-01-02, 2023-04-02, 2023-10-02, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDiskEncryptionSetResult> Invoke(GetDiskEncryptionSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskEncryptionSetResult>("azure-native:compute:getDiskEncryptionSet", args ?? new GetDiskEncryptionSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a disk encryption set.
        /// 
        /// Uses Azure REST API version 2024-03-02.
        /// 
        /// Other available API versions: 2022-07-02, 2023-01-02, 2023-04-02, 2023-10-02, 2025-01-02. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDiskEncryptionSetResult> Invoke(GetDiskEncryptionSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskEncryptionSetResult>("azure-native:compute:getDiskEncryptionSet", args ?? new GetDiskEncryptionSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiskEncryptionSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
        /// </summary>
        [Input("diskEncryptionSetName", required: true)]
        public string DiskEncryptionSetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDiskEncryptionSetArgs()
        {
        }
        public static new GetDiskEncryptionSetArgs Empty => new GetDiskEncryptionSetArgs();
    }

    public sealed class GetDiskEncryptionSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
        /// </summary>
        [Input("diskEncryptionSetName", required: true)]
        public Input<string> DiskEncryptionSetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDiskEncryptionSetInvokeArgs()
        {
        }
        public static new GetDiskEncryptionSetInvokeArgs Empty => new GetDiskEncryptionSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiskEncryptionSetResult
    {
        /// <summary>
        /// The key vault key which is currently used by this disk encryption set.
        /// </summary>
        public readonly Outputs.KeyForDiskEncryptionSetResponse? ActiveKey;
        /// <summary>
        /// The error that was encountered during auto-key rotation. If an error is present, then auto-key rotation will not be attempted until the error on this disk encryption set is fixed.
        /// </summary>
        public readonly Outputs.ApiErrorResponse AutoKeyRotationError;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The type of key used to encrypt the data of the disk.
        /// </summary>
        public readonly string? EncryptionType;
        /// <summary>
        /// Multi-tenant application client id to access key vault in a different tenant. Setting the value to 'None' will clear the property.
        /// </summary>
        public readonly string? FederatedClientId;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
        /// </summary>
        public readonly Outputs.EncryptionSetIdentityResponse? Identity;
        /// <summary>
        /// The time when the active key of this disk encryption set was updated.
        /// </summary>
        public readonly string LastKeyRotationTimestamp;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyForDiskEncryptionSetResponse> PreviousKeys;
        /// <summary>
        /// The disk encryption set provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Set this flag to true to enable auto-updating of this disk encryption set to the latest key version.
        /// </summary>
        public readonly bool? RotationToLatestKeyVersionEnabled;
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
        private GetDiskEncryptionSetResult(
            Outputs.KeyForDiskEncryptionSetResponse? activeKey,

            Outputs.ApiErrorResponse autoKeyRotationError,

            string azureApiVersion,

            string? encryptionType,

            string? federatedClientId,

            string id,

            Outputs.EncryptionSetIdentityResponse? identity,

            string lastKeyRotationTimestamp,

            string location,

            string name,

            ImmutableArray<Outputs.KeyForDiskEncryptionSetResponse> previousKeys,

            string provisioningState,

            bool? rotationToLatestKeyVersionEnabled,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            ActiveKey = activeKey;
            AutoKeyRotationError = autoKeyRotationError;
            AzureApiVersion = azureApiVersion;
            EncryptionType = encryptionType;
            FederatedClientId = federatedClientId;
            Id = id;
            Identity = identity;
            LastKeyRotationTimestamp = lastKeyRotationTimestamp;
            Location = location;
            Name = name;
            PreviousKeys = previousKeys;
            ProvisioningState = provisioningState;
            RotationToLatestKeyVersionEnabled = rotationToLatestKeyVersionEnabled;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
