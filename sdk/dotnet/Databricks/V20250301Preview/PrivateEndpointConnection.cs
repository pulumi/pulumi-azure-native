// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.V20250301Preview
{
    /// <summary>
    /// The private endpoint connection of a workspace
    /// </summary>
    [AzureNativeResourceType("azure-native:databricks/v20250301preview:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private endpoint connection properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointConnectionPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpointConnection(string name, PrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:databricks/v20250301preview:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:databricks/v20250301preview:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20210401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20220401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20230201:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20230915preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20240501:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks/v20240901preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:databricks:PrivateEndpointConnection" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateEndpointConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the private endpoint connection
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// The private endpoint connection properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.PrivateEndpointConnectionPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
