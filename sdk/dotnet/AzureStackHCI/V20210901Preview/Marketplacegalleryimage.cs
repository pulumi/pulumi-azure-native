// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20210901Preview
{
    /// <summary>
    /// The marketplace gallery image resource definition.
    /// </summary>
    [AzureNativeResourceType("azure-native:azurestackhci/v20210901preview:Marketplacegalleryimage")]
    public partial class Marketplacegalleryimage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
        /// </summary>
        [Output("cloudInitDataSource")]
        public Output<string?> CloudInitDataSource { get; private set; } = null!;

        /// <summary>
        /// Container Name for storage container
        /// </summary>
        [Output("containerName")]
        public Output<string?> ContainerName { get; private set; } = null!;

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The hypervisor generation of the Virtual Machine [V1, V2]
        /// </summary>
        [Output("hyperVGeneration")]
        public Output<string?> HyperVGeneration { get; private set; } = null!;

        /// <summary>
        /// This is the gallery image definition identifier.
        /// </summary>
        [Output("identifier")]
        public Output<Outputs.GalleryImageIdentifierResponse?> Identifier { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// operating system type that the gallery image uses. Expected to be linux or windows
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the gallery image.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// name of the object to be used in moc
        /// </summary>
        [Output("resourceName")]
        public Output<string?> ResourceName { get; private set; } = null!;

        /// <summary>
        /// MarketplaceGalleryImageStatus defines the observed state of marketplacegalleryimages
        /// </summary>
        [Output("status")]
        public Output<Outputs.MarketplaceGalleryImageStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the gallery image version that you want to create or update.
        /// </summary>
        [Output("version")]
        public Output<Outputs.GalleryImageVersionResponse?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Marketplacegalleryimage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Marketplacegalleryimage(string name, MarketplacegalleryimageArgs args, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci/v20210901preview:Marketplacegalleryimage", name, args ?? new MarketplacegalleryimageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Marketplacegalleryimage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:azurestackhci/v20210901preview:Marketplacegalleryimage", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20210901preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20221215preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230701preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230701preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230901preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20230901preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240101:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240101:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240201preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240201preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240501preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240501preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240715preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240715preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240801preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20240801preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20241001preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20241001preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250201preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250201preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250401preview:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci/v20250401preview:marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci:Marketplacegalleryimage" },
                    new global::Pulumi.Alias { Type = "azure-native:azurestackhci:marketplacegalleryimage" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Marketplacegalleryimage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Marketplacegalleryimage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Marketplacegalleryimage(name, id, options);
        }
    }

    public sealed class MarketplacegalleryimageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Datasource for the gallery image when provisioning with cloud-init [Azure, NoCloud]
        /// </summary>
        [Input("cloudInitDataSource")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.V20210901Preview.CloudInitDataSource>? CloudInitDataSource { get; set; }

        /// <summary>
        /// Container Name for storage container
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The hypervisor generation of the Virtual Machine [V1, V2]
        /// </summary>
        [Input("hyperVGeneration")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.V20210901Preview.HyperVGeneration>? HyperVGeneration { get; set; }

        /// <summary>
        /// This is the gallery image definition identifier.
        /// </summary>
        [Input("identifier")]
        public Input<Inputs.GalleryImageIdentifierArgs>? Identifier { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the marketplace gallery image
        /// </summary>
        [Input("marketplacegalleryimagesName")]
        public Input<string>? MarketplacegalleryimagesName { get; set; }

        /// <summary>
        /// operating system type that the gallery image uses. Expected to be linux or windows
        /// </summary>
        [Input("osType")]
        public Input<Pulumi.AzureNative.AzureStackHCI.V20210901Preview.OperatingSystemTypes>? OsType { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// name of the object to be used in moc
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies information about the gallery image version that you want to create or update.
        /// </summary>
        [Input("version")]
        public Input<Inputs.GalleryImageVersionArgs>? Version { get; set; }

        public MarketplacegalleryimageArgs()
        {
        }
        public static new MarketplacegalleryimageArgs Empty => new MarketplacegalleryimageArgs();
    }
}
