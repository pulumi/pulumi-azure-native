// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// PrivateEndpointConnection resource.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:PrivateLinkServicePrivateEndpointConnection")]
    public partial class PrivateLinkServicePrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The consumer link id.
        /// </summary>
        [Output("linkIdentifier")]
        public Output<string> LinkIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The resource of private end point.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.PrivateEndpointResponse> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// The location of the private endpoint.
        /// </summary>
        [Output("privateEndpointLocation")]
        public Output<string> PrivateEndpointLocation { get; private set; } = null!;

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.PrivateLinkServiceConnectionStateResponse?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the private endpoint connection resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLinkServicePrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkServicePrivateEndpointConnection(string name, PrivateLinkServicePrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:PrivateLinkServicePrivateEndpointConnection", name, args ?? new PrivateLinkServicePrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkServicePrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:PrivateLinkServicePrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:PrivateLinkServicePrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:PrivateLinkServicePrivateEndpointConnection" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateLinkServicePrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkServicePrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateLinkServicePrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateLinkServicePrivateEndpointConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the private end point connection.
        /// </summary>
        [Input("peConnectionName")]
        public Input<string>? PeConnectionName { get; set; }

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateLinkServiceConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the private link service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public PrivateLinkServicePrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateLinkServicePrivateEndpointConnectionArgs Empty => new PrivateLinkServicePrivateEndpointConnectionArgs();
    }
}
