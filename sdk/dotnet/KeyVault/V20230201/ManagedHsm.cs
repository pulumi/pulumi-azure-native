// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.V20230201
{
    /// <summary>
    /// Resource information with extended details.
    /// </summary>
    [AzureNativeResourceType("azure-native:keyvault/v20230201:ManagedHsm")]
    public partial class ManagedHsm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The supported Azure location where the managed HSM Pool should be created.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the managed HSM Pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the managed HSM
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ManagedHsmPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// SKU details
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ManagedHsmSkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the key vault resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type of the managed HSM Pool.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedHsm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedHsm(string name, ManagedHsmArgs args, CustomResourceOptions? options = null)
            : base("azure-native:keyvault/v20230201:ManagedHsm", name, args ?? new ManagedHsmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedHsm(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:keyvault/v20230201:ManagedHsm", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20200401preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20210401preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20210601preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20211001:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20211101preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20220201preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20220701:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20221101:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20230701:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20240401preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20241101:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault/v20241201preview:ManagedHsm" },
                    new global::Pulumi.Alias { Type = "azure-native:keyvault:ManagedHsm" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedHsm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedHsm Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedHsm(name, id, options);
        }
    }

    public sealed class ManagedHsmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The supported Azure location where the managed HSM Pool should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the managed HSM Pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Properties of the managed HSM
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ManagedHsmPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group that contains the managed HSM pool.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// SKU details
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ManagedHsmSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ManagedHsmArgs()
        {
        }
        public static new ManagedHsmArgs Empty => new ManagedHsmArgs();
    }
}
