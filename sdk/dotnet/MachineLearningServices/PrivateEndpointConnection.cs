// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices
{
    /// <summary>
    /// The Private Endpoint Connection resource.
    /// 
    /// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
    /// 
    /// Other available API versions: 2020-05-01-preview, 2020-05-15-preview, 2020-06-01, 2020-08-01, 2020-09-01-preview, 2021-01-01, 2021-03-01-preview, 2021-04-01, 2021-07-01, 2022-01-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:machinelearningservices:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource of private end point.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.PrivateEndpointResponse?> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.PrivateLinkServiceConnectionStateResponse> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the private endpoint connection resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
            : base("azure-native:machinelearningservices:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:machinelearningservices:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200101:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200218preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200301:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200501preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200515preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200601:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200801:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20200901preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20210101:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20210301preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20210401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20210701:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220101preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220201preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220501:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221001preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221201preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230201preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230601preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230801preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20231001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240101preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240701preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250101preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250601:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250701preview:PrivateEndpointConnection" },
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
        /// The identity of the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the private endpoint connection associated with the workspace
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Input("privateLinkServiceConnectionState", required: true)]
        public Input<Inputs.PrivateLinkServiceConnectionStateArgs> PrivateLinkServiceConnectionState { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
