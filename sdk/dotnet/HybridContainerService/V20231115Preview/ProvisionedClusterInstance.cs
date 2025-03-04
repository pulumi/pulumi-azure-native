// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview
{
    /// <summary>
    /// The provisionedClusterInstances resource definition.
    /// </summary>
    [AzureNativeResourceType("azure-native:hybridcontainerservice/v20231115preview:ProvisionedClusterInstance")]
    public partial class ProvisionedClusterInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Extended Location definition
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// All properties of the provisioned cluster
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ProvisionedClusterPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProvisionedClusterInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProvisionedClusterInstance(string name, ProvisionedClusterInstanceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:hybridcontainerservice/v20231115preview:ProvisionedClusterInstance", name, args ?? new ProvisionedClusterInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProvisionedClusterInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:hybridcontainerservice/v20231115preview:ProvisionedClusterInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20231115preview:provisionedClusterInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20240101:ProvisionedClusterInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:hybridcontainerservice/v20240101:provisionedClusterInstance" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProvisionedClusterInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProvisionedClusterInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProvisionedClusterInstance(name, id, options);
        }
    }

    public sealed class ProvisionedClusterInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the connected cluster resource.
        /// </summary>
        [Input("connectedClusterResourceUri", required: true)]
        public Input<string> ConnectedClusterResourceUri { get; set; } = null!;

        /// <summary>
        /// Extended Location definition
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// All properties of the provisioned cluster
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ProvisionedClusterPropertiesArgs>? Properties { get; set; }

        public ProvisionedClusterInstanceArgs()
        {
        }
        public static new ProvisionedClusterInstanceArgs Empty => new ProvisionedClusterInstanceArgs();
    }
}
