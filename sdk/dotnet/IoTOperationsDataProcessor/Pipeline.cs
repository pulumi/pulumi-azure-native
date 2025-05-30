// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsDataProcessor
{
    /// <summary>
    /// A Pipeline resource belonging to an Instance resource.
    /// 
    /// Uses Azure REST API version 2023-10-04-preview. In version 2.x of the Azure Native provider, it used API version 2023-10-04-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:iotoperationsdataprocessor:Pipeline")]
    public partial class Pipeline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Detailed description of the Pipeline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Flag indicating whether the pipeline should be running or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Edge location of the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// Information about where to pull input data from.
        /// </summary>
        [Output("input")]
        public Output<Outputs.PipelineInputResponse> Input { get; private set; } = null!;

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
        /// The status of the last operation.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Map of stage ids to stage configurations for all pipeline processing and output stages.
        /// </summary>
        [Output("stages")]
        public Output<ImmutableDictionary<string, Outputs.PipelineStageResponse>> Stages { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
        /// Create a Pipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipeline(string name, PipelineArgs args, CustomResourceOptions? options = null)
            : base("azure-native:iotoperationsdataprocessor:Pipeline", name, args ?? new PipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipeline(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:iotoperationsdataprocessor:Pipeline", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:iotoperationsdataprocessor/v20231004preview:Pipeline" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipeline Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Pipeline(name, id, options);
        }
    }

    public sealed class PipelineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Detailed description of the Pipeline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Flag indicating whether the pipeline should be running or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Edge location of the resource.
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// Information about where to pull input data from.
        /// </summary>
        [Input("input", required: true)]
        public Input<Inputs.PipelineInputArgs> Input { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of pipeline
        /// </summary>
        [Input("pipelineName")]
        public Input<string>? PipelineName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("stages", required: true)]
        private InputMap<Inputs.PipelineStageArgs>? _stages;

        /// <summary>
        /// Map of stage ids to stage configurations for all pipeline processing and output stages.
        /// </summary>
        public InputMap<Inputs.PipelineStageArgs> Stages
        {
            get => _stages ?? (_stages = new InputMap<Inputs.PipelineStageArgs>());
            set => _stages = value;
        }

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

        public PipelineArgs()
        {
        }
        public static new PipelineArgs Empty => new PipelineArgs();
    }
}
