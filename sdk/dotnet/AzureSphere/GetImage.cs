// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureSphere
{
    public static class GetImage
    {
        /// <summary>
        /// Get a Image
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuresphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("azure-native:azuresphere:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Image
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuresphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("azure-native:azuresphere:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Image
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2022-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuresphere [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("azure-native:azuresphere:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of catalog
        /// </summary>
        [Input("catalogName", required: true)]
        public string CatalogName { get; set; } = null!;

        /// <summary>
        /// Image name. Use an image GUID for GA versions of the API.
        /// </summary>
        [Input("imageName", required: true)]
        public string ImageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of catalog
        /// </summary>
        [Input("catalogName", required: true)]
        public Input<string> CatalogName { get; set; } = null!;

        /// <summary>
        /// Image name. Use an image GUID for GA versions of the API.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The image component id.
        /// </summary>
        public readonly string ComponentId;
        /// <summary>
        /// The image description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
        /// </summary>
        public readonly string? Image;
        /// <summary>
        /// Image ID
        /// </summary>
        public readonly string? ImageId;
        /// <summary>
        /// Image name
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// The image type.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Regional data boundary for an image
        /// </summary>
        public readonly string? RegionalDataBoundary;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Location the image
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetImageResult(
            string azureApiVersion,

            string componentId,

            string description,

            string id,

            string? image,

            string? imageId,

            string imageName,

            string imageType,

            string name,

            string provisioningState,

            string? regionalDataBoundary,

            Outputs.SystemDataResponse systemData,

            string type,

            string uri)
        {
            AzureApiVersion = azureApiVersion;
            ComponentId = componentId;
            Description = description;
            Id = id;
            Image = image;
            ImageId = imageId;
            ImageName = imageName;
            ImageType = imageType;
            Name = name;
            ProvisioningState = provisioningState;
            RegionalDataBoundary = regionalDataBoundary;
            SystemData = systemData;
            Type = type;
            Uri = uri;
        }
    }
}
