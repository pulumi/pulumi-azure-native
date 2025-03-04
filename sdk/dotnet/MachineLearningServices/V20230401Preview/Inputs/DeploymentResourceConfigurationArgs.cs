// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230401Preview.Inputs
{

    public sealed class DeploymentResourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional number of instances or nodes used by the compute target.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// Optional type of VM used as supported by the compute target.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// Locations where the job can run.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// Optional max allowed number of instances or nodes to be used by the compute target.
        /// For use with elastic training, currently supported by PyTorch distribution type only.
        /// </summary>
        [Input("maxInstanceCount")]
        public Input<int>? MaxInstanceCount { get; set; }

        [Input("properties")]
        private InputMap<object>? _properties;

        /// <summary>
        /// Additional properties bag.
        /// </summary>
        public InputMap<object> Properties
        {
            get => _properties ?? (_properties = new InputMap<object>());
            set => _properties = value;
        }

        public DeploymentResourceConfigurationArgs()
        {
            InstanceCount = 1;
        }
        public static new DeploymentResourceConfigurationArgs Empty => new DeploymentResourceConfigurationArgs();
    }
}
