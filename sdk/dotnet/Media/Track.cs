// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media
{
    /// <summary>
    /// An Asset Track resource.
    /// 
    /// Uses Azure REST API version 2023-01-01. In version 2.x of the Azure Native provider, it used API version 2023-01-01.
    /// 
    /// Other available API versions: 2021-11-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:media:Track")]
    public partial class Track : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the asset track.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Detailed information about a track in the asset.
        /// </summary>
        [Output("track")]
        public Output<object?> TrackValue { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Track resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Track(string name, TrackArgs args, CustomResourceOptions? options = null)
            : base("azure-native:media:Track", name, args ?? new TrackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Track(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:media:Track", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:media/v20211101:Track" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20220801:Track" },
                    new global::Pulumi.Alias { Type = "azure-native:media/v20230101:Track" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Track resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Track Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Track(name, id, options);
        }
    }

    public sealed class TrackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The Asset name.
        /// </summary>
        [Input("assetName", required: true)]
        public Input<string> AssetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Detailed information about a track in the asset.
        /// </summary>
        [Input("track")]
        public object? Track { get; set; }

        /// <summary>
        /// The Asset Track name.
        /// </summary>
        [Input("trackName")]
        public Input<string>? TrackName { get; set; }

        public TrackArgs()
        {
        }
        public static new TrackArgs Empty => new TrackArgs();
    }
}
