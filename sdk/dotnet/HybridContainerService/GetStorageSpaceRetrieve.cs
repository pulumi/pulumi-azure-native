// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService
{
    public static class GetStorageSpaceRetrieve
    {
        /// <summary>
        /// Gets the Hybrid AKS storage space object
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// </summary>
        public static Task<GetStorageSpaceRetrieveResult> InvokeAsync(GetStorageSpaceRetrieveArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageSpaceRetrieveResult>("azure-native:hybridcontainerservice:getStorageSpaceRetrieve", args ?? new GetStorageSpaceRetrieveArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Hybrid AKS storage space object
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// </summary>
        public static Output<GetStorageSpaceRetrieveResult> Invoke(GetStorageSpaceRetrieveInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageSpaceRetrieveResult>("azure-native:hybridcontainerservice:getStorageSpaceRetrieve", args ?? new GetStorageSpaceRetrieveInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Hybrid AKS storage space object
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// </summary>
        public static Output<GetStorageSpaceRetrieveResult> Invoke(GetStorageSpaceRetrieveInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageSpaceRetrieveResult>("azure-native:hybridcontainerservice:getStorageSpaceRetrieve", args ?? new GetStorageSpaceRetrieveInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageSpaceRetrieveArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Parameter for the name of the storage object
        /// </summary>
        [Input("storageSpacesName", required: true)]
        public string StorageSpacesName { get; set; } = null!;

        public GetStorageSpaceRetrieveArgs()
        {
        }
        public static new GetStorageSpaceRetrieveArgs Empty => new GetStorageSpaceRetrieveArgs();
    }

    public sealed class GetStorageSpaceRetrieveInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Parameter for the name of the storage object
        /// </summary>
        [Input("storageSpacesName", required: true)]
        public Input<string> StorageSpacesName { get; set; } = null!;

        public GetStorageSpaceRetrieveInvokeArgs()
        {
        }
        public static new GetStorageSpaceRetrieveInvokeArgs Empty => new GetStorageSpaceRetrieveInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageSpaceRetrieveResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        public readonly Outputs.StorageSpacesResponseExtendedLocation? ExtendedLocation;
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
        /// HybridAKSStorageSpec defines the desired state of HybridAKSStorage
        /// </summary>
        public readonly Outputs.StorageSpacesPropertiesResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
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
        private GetStorageSpaceRetrieveResult(
            string azureApiVersion,

            Outputs.StorageSpacesResponseExtendedLocation? extendedLocation,

            string id,

            string location,

            string name,

            Outputs.StorageSpacesPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ExtendedLocation = extendedLocation;
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
