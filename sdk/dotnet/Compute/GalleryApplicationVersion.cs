// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute
{
    /// <summary>
    /// Specifies information about the gallery Application Version that you want to create or update.
    /// 
    /// Uses Azure REST API version 2024-03-03. In version 2.x of the Azure Native provider, it used API version 2022-03-03.
    /// 
    /// Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:compute:GalleryApplicationVersion")]
    public partial class GalleryApplicationVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

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
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The publishing profile of a gallery image version.
        /// </summary>
        [Output("publishingProfile")]
        public Output<Outputs.GalleryApplicationVersionPublishingProfileResponse> PublishingProfile { get; private set; } = null!;

        /// <summary>
        /// This is the replication status of the gallery image version.
        /// </summary>
        [Output("replicationStatus")]
        public Output<Outputs.ReplicationStatusResponse> ReplicationStatus { get; private set; } = null!;

        /// <summary>
        /// The safety profile of the Gallery Application Version.
        /// </summary>
        [Output("safetyProfile")]
        public Output<Outputs.GalleryApplicationVersionSafetyProfileResponse?> SafetyProfile { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
        /// Create a GalleryApplicationVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GalleryApplicationVersion(string name, GalleryApplicationVersionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:compute:GalleryApplicationVersion", name, args ?? new GalleryApplicationVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GalleryApplicationVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:compute:GalleryApplicationVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20190301:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20190701:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20191201:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20200930:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20210701:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20211001:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20220103:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20220303:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20220803:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20230703:GalleryApplicationVersion" },
                    new global::Pulumi.Alias { Type = "azure-native:compute/v20240303:GalleryApplicationVersion" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GalleryApplicationVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GalleryApplicationVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GalleryApplicationVersion(name, id, options);
        }
    }

    public sealed class GalleryApplicationVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the gallery Application Definition to be retrieved.
        /// </summary>
        [Input("galleryApplicationName", required: true)]
        public Input<string> GalleryApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery Application Version to be retrieved.
        /// </summary>
        [Input("galleryApplicationVersionName")]
        public Input<string>? GalleryApplicationVersionName { get; set; }

        /// <summary>
        /// The name of the Shared Image Gallery.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The publishing profile of a gallery image version.
        /// </summary>
        [Input("publishingProfile", required: true)]
        public Input<Inputs.GalleryApplicationVersionPublishingProfileArgs> PublishingProfile { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The safety profile of the Gallery Application Version.
        /// </summary>
        [Input("safetyProfile")]
        public Input<Inputs.GalleryApplicationVersionSafetyProfileArgs>? SafetyProfile { get; set; }

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

        public GalleryApplicationVersionArgs()
        {
        }
        public static new GalleryApplicationVersionArgs Empty => new GalleryApplicationVersionArgs();
    }
}
