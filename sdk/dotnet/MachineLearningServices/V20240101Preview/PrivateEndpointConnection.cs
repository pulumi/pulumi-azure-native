// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240101Preview
{
    /// <summary>
    /// The Private Endpoint Connection resource.
    /// </summary>
    [AzureNativeResourceType("azure-native:machinelearningservices/v20240101preview:PrivateEndpointConnection")]
    public partial class PrivateEndpointConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Same as workspace location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Private Endpoint resource.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.WorkspacePrivateEndpointResourceResponse?> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// The connection state.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.PrivateLinkServiceConnectionStateResponse?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// The current provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Optional. This field is required to be implemented by the RP because AML is supporting more than one tier
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

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
            : base("azure-native:machinelearningservices/v20240101preview:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:machinelearningservices/v20240101preview:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
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
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240701preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250101preview:PrivateEndpointConnection" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices:PrivateEndpointConnection" },
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
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Same as workspace location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// NRP Private Endpoint Connection Name
        /// </summary>
        [Input("privateEndpointConnectionName")]
        public Input<string>? PrivateEndpointConnectionName { get; set; }

        /// <summary>
        /// The connection state.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateLinkServiceConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// The current provisioning state.
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20240101Preview.PrivateEndpointConnectionProvisioningState>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Optional. This field is required to be implemented by the RP because AML is supporting more than one tier
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Azure Machine Learning Workspace Name
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
        public static new PrivateEndpointConnectionArgs Empty => new PrivateEndpointConnectionArgs();
    }
}
