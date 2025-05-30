// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class JobResourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Extra arguments to pass to the Docker run command. This would override any parameters that have already been set by the system, or in this section. This parameter is only supported for Azure ML compute types.
        /// </summary>
        [Input("dockerArgs")]
        public Input<string>? DockerArgs { get; set; }

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

        /// <summary>
        /// Size of the docker container's shared memory block. This should be in the format of (number)(unit) where number as to be greater than 0 and the unit can be one of b(bytes), k(kilobytes), m(megabytes), or g(gigabytes).
        /// </summary>
        [Input("shmSize")]
        public Input<string>? ShmSize { get; set; }

        public JobResourceConfigurationArgs()
        {
            InstanceCount = 1;
            ShmSize = "2g";
        }
        public static new JobResourceConfigurationArgs Empty => new JobResourceConfigurationArgs();
    }
}
