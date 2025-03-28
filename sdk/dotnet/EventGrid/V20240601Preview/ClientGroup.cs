// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20240601Preview
{
    /// <summary>
    /// The Client group resource.
    /// </summary>
    [AzureNativeResourceType("azure-native:eventgrid/v20240601preview:ClientGroup")]
    public partial class ClientGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the Client Group resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the ClientGroup resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The grouping query for the clients.
        /// Example : attributes.keyName IN ['a', 'b', 'c'].
        /// </summary>
        [Output("query")]
        public Output<string?> Query { get; private set; } = null!;

        /// <summary>
        /// The system metadata relating to the ClientGroup resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ClientGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientGroup(string name, ClientGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-native:eventgrid/v20240601preview:ClientGroup", name, args ?? new ClientGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:eventgrid/v20240601preview:ClientGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20230601preview:ClientGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20231215preview:ClientGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20241215preview:ClientGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20250215:ClientGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid:ClientGroup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClientGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClientGroup(name, id, options);
        }
    }

    public sealed class ClientGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client group name.
        /// </summary>
        [Input("clientGroupName")]
        public Input<string>? ClientGroupName { get; set; }

        /// <summary>
        /// Description for the Client Group resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The grouping query for the clients.
        /// Example : attributes.keyName IN ['a', 'b', 'c'].
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ClientGroupArgs()
        {
        }
        public static new ClientGroupArgs Empty => new ClientGroupArgs();
    }
}
