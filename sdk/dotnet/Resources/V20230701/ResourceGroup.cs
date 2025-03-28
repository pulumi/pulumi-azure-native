// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.V20230701
{
    /// <summary>
    /// Resource group information.
    /// </summary>
    [AzureNativeResourceType("azure-native:resources/v20230701:ResourceGroup")]
    public partial class ResourceGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource that manages this resource group.
        /// </summary>
        [Output("managedBy")]
        public Output<string?> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource group properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ResourceGroupPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The tags attached to the resource group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource group.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceGroup(string name, ResourceGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("azure-native:resources/v20230701:ResourceGroup", name, args ?? new ResourceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:resources/v20230701:ResourceGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20151101:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20160201:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20160701:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20160901:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20170510:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20180201:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20180501:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20190301:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20190501:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20190510:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20190701:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20190801:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20191001:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20200601:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20200801:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20201001:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20210101:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20210401:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20220901:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20240301:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20240701:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources/v20241101:ResourceGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:resources:ResourceGroup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceGroup(name, id, options);
        }
    }

    public sealed class ResourceGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the resource that manages this resource group.
        /// </summary>
        [Input("managedBy")]
        public Input<string>? ManagedBy { get; set; }

        /// <summary>
        /// The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags attached to the resource group.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResourceGroupArgs()
        {
        }
        public static new ResourceGroupArgs Empty => new ResourceGroupArgs();
    }
}
