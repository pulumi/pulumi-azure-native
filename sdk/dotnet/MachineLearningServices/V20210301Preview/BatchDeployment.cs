// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210301Preview
{
    [AzureNativeResourceType("azure-native:machinelearningservices/v20210301preview:BatchDeployment")]
    public partial class BatchDeployment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Service identity associated with a resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ResourceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.BatchDeploymentResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// System data associated with resource provider
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BatchDeployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BatchDeployment(string name, BatchDeploymentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:machinelearningservices/v20210301preview:BatchDeployment", name, args ?? new BatchDeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BatchDeployment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:machinelearningservices/v20210301preview:BatchDeployment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220201preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220501:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20220601preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221001:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221001preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20221201preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230201preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230401:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230401preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230601preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20230801preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20231001:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240101preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240401preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20240701preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20241001preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices/v20250101preview:BatchDeployment" },
                    new global::Pulumi.Alias { Type = "azure-native:machinelearningservices:BatchDeployment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BatchDeployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BatchDeployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BatchDeployment(name, id, options);
        }
    }

    public sealed class BatchDeploymentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the Batch inference deployment.
        /// </summary>
        [Input("deploymentName")]
        public Input<string>? DeploymentName { get; set; }

        /// <summary>
        /// Inference endpoint name
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// Service identity associated with a resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// [Required] Additional attributes of the entity.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.BatchDeploymentArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
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

        public BatchDeploymentArgs()
        {
        }
        public static new BatchDeploymentArgs Empty => new BatchDeploymentArgs();
    }
}
