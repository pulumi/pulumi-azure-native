// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20220401Preview
{
    /// <summary>
    /// Defines the security user configuration
    /// </summary>
    [AzureNativeResourceType("azure-native:network/v20220401preview:SecurityUserConfiguration")]
    public partial class SecurityUserConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Flag if need to delete existing network security groups.
        /// </summary>
        [Output("deleteExistingNSGs")]
        public Output<string?> DeleteExistingNSGs { get; private set; } = null!;

        /// <summary>
        /// A description of the security user configuration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityUserConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityUserConfiguration(string name, SecurityUserConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network/v20220401preview:SecurityUserConfiguration", name, args ?? new SecurityUserConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityUserConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network/v20220401preview:SecurityUserConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201preview:SecurityUserConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501preview:SecurityUserConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220201preview:SecurityUserConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:SecurityUserConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:SecurityUserConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:network:SecurityUserConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityUserConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityUserConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityUserConfiguration(name, id, options);
        }
    }

    public sealed class SecurityUserConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network manager Security Configuration.
        /// </summary>
        [Input("configurationName")]
        public Input<string>? ConfigurationName { get; set; }

        /// <summary>
        /// Flag if need to delete existing network security groups.
        /// </summary>
        [Input("deleteExistingNSGs")]
        public InputUnion<string, Pulumi.AzureNative.Network.V20220401Preview.DeleteExistingNSGs>? DeleteExistingNSGs { get; set; }

        /// <summary>
        /// A description of the security user configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public Input<string> NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public SecurityUserConfigurationArgs()
        {
        }
        public static new SecurityUserConfigurationArgs Empty => new SecurityUserConfigurationArgs();
    }
}
