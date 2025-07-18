// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    /// <summary>
    /// Uses Azure REST API version 2025-02-15. In version 2.x of the Azure Native provider, it used API version 2022-06-15.
    /// 
    /// Other available API versions: 2022-06-15, 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:eventgrid:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<string>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.PrivateEndpointResponse?> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.ConnectionStateResponse?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
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
            : base("azure-native:eventgrid:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:eventgrid:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20200401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20200601:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20201015preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20210601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20211015preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20211201:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20220615:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20230601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20231215preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20240601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20241215preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20250215:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:eventgrid/v20250401preview:PrivateEndpointConnection" },
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
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or namespace name).
        /// </summary>
        [Input("parentName", required: true)]
        public Input<string> ParentName { get; set; } = null!;

        /// <summary>
        /// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
        /// </summary>
        [Input("parentType", required: true)]
        public Input<string> ParentType { get; set; } = null!;

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// The name of the private endpoint connection connection.
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.ConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNative.EventGrid.ResourceProvisioningState>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
