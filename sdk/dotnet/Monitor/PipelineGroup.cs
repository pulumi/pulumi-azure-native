// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor
{
    /// <summary>
    /// A pipeline group definition.
    /// 
    /// Uses Azure REST API version 2023-10-01-preview.
    /// 
    /// Other available API versions: 2024-10-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:monitor:PipelineGroup")]
    public partial class PipelineGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The exporters specified for a pipeline group instance.
        /// </summary>
        [Output("exporters")]
        public Output<ImmutableArray<Outputs.ExporterResponse>> Exporters { get; private set; } = null!;

        /// <summary>
        /// The extended location for given pipeline group.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.AzureResourceManagerCommonTypesExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

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
        /// Networking configurations for the pipeline group instance.
        /// </summary>
        [Output("networkingConfigurations")]
        public Output<ImmutableArray<Outputs.NetworkingConfigurationResponse>> NetworkingConfigurations { get; private set; } = null!;

        /// <summary>
        /// The processors specified for a pipeline group instance.
        /// </summary>
        [Output("processors")]
        public Output<ImmutableArray<Outputs.ProcessorResponse>> Processors { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of a pipeline group instance. Set to Succeeded if everything is healthy.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The receivers specified for a pipeline group instance.
        /// </summary>
        [Output("receivers")]
        public Output<ImmutableArray<Outputs.ReceiverResponse>> Receivers { get; private set; } = null!;

        /// <summary>
        /// Defines the amount of replicas of the pipeline group instance.
        /// </summary>
        [Output("replicas")]
        public Output<int?> Replicas { get; private set; } = null!;

        /// <summary>
        /// The service section for a given pipeline group instance.
        /// </summary>
        [Output("service")]
        public Output<Outputs.ServiceResponse> Service { get; private set; } = null!;

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
        /// Create a PipelineGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PipelineGroup(string name, PipelineGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-native:monitor:PipelineGroup", name, args ?? new PipelineGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PipelineGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:monitor:PipelineGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20231001preview:PipelineGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:monitor/v20241001preview:PipelineGroup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PipelineGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PipelineGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PipelineGroup(name, id, options);
        }
    }

    public sealed class PipelineGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("exporters", required: true)]
        private InputList<Inputs.ExporterArgs>? _exporters;

        /// <summary>
        /// The exporters specified for a pipeline group instance.
        /// </summary>
        public InputList<Inputs.ExporterArgs> Exporters
        {
            get => _exporters ?? (_exporters = new InputList<Inputs.ExporterArgs>());
            set => _exporters = value;
        }

        /// <summary>
        /// The extended location for given pipeline group.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.AzureResourceManagerCommonTypesExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("networkingConfigurations")]
        private InputList<Inputs.NetworkingConfigurationArgs>? _networkingConfigurations;

        /// <summary>
        /// Networking configurations for the pipeline group instance.
        /// </summary>
        public InputList<Inputs.NetworkingConfigurationArgs> NetworkingConfigurations
        {
            get => _networkingConfigurations ?? (_networkingConfigurations = new InputList<Inputs.NetworkingConfigurationArgs>());
            set => _networkingConfigurations = value;
        }

        /// <summary>
        /// The name of pipeline group. The name is case insensitive.
        /// </summary>
        [Input("pipelineGroupName")]
        public Input<string>? PipelineGroupName { get; set; }

        [Input("processors", required: true)]
        private InputList<Inputs.ProcessorArgs>? _processors;

        /// <summary>
        /// The processors specified for a pipeline group instance.
        /// </summary>
        public InputList<Inputs.ProcessorArgs> Processors
        {
            get => _processors ?? (_processors = new InputList<Inputs.ProcessorArgs>());
            set => _processors = value;
        }

        [Input("receivers", required: true)]
        private InputList<Inputs.ReceiverArgs>? _receivers;

        /// <summary>
        /// The receivers specified for a pipeline group instance.
        /// </summary>
        public InputList<Inputs.ReceiverArgs> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<Inputs.ReceiverArgs>());
            set => _receivers = value;
        }

        /// <summary>
        /// Defines the amount of replicas of the pipeline group instance.
        /// </summary>
        [Input("replicas")]
        public Input<int>? Replicas { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The service section for a given pipeline group instance.
        /// </summary>
        [Input("service", required: true)]
        public Input<Inputs.ServiceArgs> Service { get; set; } = null!;

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

        public PipelineGroupArgs()
        {
        }
        public static new PipelineGroupArgs Empty => new PipelineGroupArgs();
    }
}
