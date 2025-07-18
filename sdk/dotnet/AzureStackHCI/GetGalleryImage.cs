// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI
{
    public static class GetGalleryImage
    {
        /// <summary>
        /// Gets a gallery image
        /// 
        /// Uses Azure REST API version 2025-02-01-preview.
        /// 
        /// Other available API versions: 2022-12-15-preview, 2023-07-01-preview, 2023-09-01-preview, 2024-01-01, 2024-02-01-preview, 2024-05-01-preview, 2024-07-15-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetGalleryImageResult> InvokeAsync(GetGalleryImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGalleryImageResult>("azure-native:azurestackhci:getGalleryImage", args ?? new GetGalleryImageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a gallery image
        /// 
        /// Uses Azure REST API version 2025-02-01-preview.
        /// 
        /// Other available API versions: 2022-12-15-preview, 2023-07-01-preview, 2023-09-01-preview, 2024-01-01, 2024-02-01-preview, 2024-05-01-preview, 2024-07-15-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetGalleryImageResult> Invoke(GetGalleryImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGalleryImageResult>("azure-native:azurestackhci:getGalleryImage", args ?? new GetGalleryImageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a gallery image
        /// 
        /// Uses Azure REST API version 2025-02-01-preview.
        /// 
        /// Other available API versions: 2022-12-15-preview, 2023-07-01-preview, 2023-09-01-preview, 2024-01-01, 2024-02-01-preview, 2024-05-01-preview, 2024-07-15-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetGalleryImageResult> Invoke(GetGalleryImageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGalleryImageResult>("azure-native:azurestackhci:getGalleryImage", args ?? new GetGalleryImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGalleryImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the gallery image
        /// </summary>
        [Input("galleryImageName", required: true)]
        public string GalleryImageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGalleryImageArgs()
        {
        }
        public static new GetGalleryImageArgs Empty => new GetGalleryImageArgs();
    }

    public sealed class GetGalleryImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the gallery image
        /// </summary>
        [Input("galleryImageName", required: true)]
        public Input<string> GalleryImageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetGalleryImageInvokeArgs()
        {
        }
        public static new GetGalleryImageInvokeArgs Empty => new GetGalleryImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetGalleryImageResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
        /// </summary>
        public readonly string? CloudInitDataSource;
        /// <summary>
        /// Storage ContainerID of the storage container to be used for gallery image
        /// </summary>
        public readonly string? ContainerId;
        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// The hypervisor generation of the Virtual Machine [V1, V2]
        /// </summary>
        public readonly string? HyperVGeneration;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This is the gallery image definition identifier.
        /// </summary>
        public readonly Outputs.GalleryImageIdentifierResponse? Identifier;
        /// <summary>
        /// location of the image the gallery image should be created from
        /// </summary>
        public readonly string? ImagePath;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Operating system type that the gallery image uses [Windows, Linux]
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Provisioning state of the gallery image.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource ID of the source virtual machine from whose OS disk the gallery image is created.
        /// </summary>
        public readonly string? SourceVirtualMachineId;
        /// <summary>
        /// The observed state of gallery images
        /// </summary>
        public readonly Outputs.GalleryImageStatusResponse Status;
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
        /// <summary>
        /// Specifies information about the gallery image version that you want to create or update.
        /// </summary>
        public readonly Outputs.GalleryImageVersionResponse? Version;
        /// <summary>
        /// The credentials used to login to the image repository that has access to the specified image
        /// </summary>
        public readonly Outputs.VmImageRepositoryCredentialsResponse? VmImageRepositoryCredentials;

        [OutputConstructor]
        private GetGalleryImageResult(
            string azureApiVersion,

            string? cloudInitDataSource,

            string? containerId,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string? hyperVGeneration,

            string id,

            Outputs.GalleryImageIdentifierResponse? identifier,

            string? imagePath,

            string location,

            string name,

            string osType,

            string provisioningState,

            string? sourceVirtualMachineId,

            Outputs.GalleryImageStatusResponse status,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.GalleryImageVersionResponse? version,

            Outputs.VmImageRepositoryCredentialsResponse? vmImageRepositoryCredentials)
        {
            AzureApiVersion = azureApiVersion;
            CloudInitDataSource = cloudInitDataSource;
            ContainerId = containerId;
            ExtendedLocation = extendedLocation;
            HyperVGeneration = hyperVGeneration;
            Id = id;
            Identifier = identifier;
            ImagePath = imagePath;
            Location = location;
            Name = name;
            OsType = osType;
            ProvisioningState = provisioningState;
            SourceVirtualMachineId = sourceVirtualMachineId;
            Status = status;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            Version = version;
            VmImageRepositoryCredentials = vmImageRepositoryCredentials;
        }
    }
}
