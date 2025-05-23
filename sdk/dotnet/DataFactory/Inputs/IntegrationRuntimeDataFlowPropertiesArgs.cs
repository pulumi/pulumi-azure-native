// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Data flow properties for managed integration runtime.
    /// </summary>
    public sealed class IntegrationRuntimeDataFlowPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster will not be recycled and it will be used in next data flow activity run until TTL (time to live) is reached if this is set as false. Default is true.
        /// </summary>
        [Input("cleanup")]
        public Input<bool>? Cleanup { get; set; }

        /// <summary>
        /// Compute type of the cluster which will execute data flow job.
        /// </summary>
        [Input("computeType")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.DataFlowComputeType>? ComputeType { get; set; }

        /// <summary>
        /// Core count of the cluster which will execute data flow job. Supported values are: 8, 16, 32, 48, 80, 144 and 272.
        /// </summary>
        [Input("coreCount")]
        public Input<int>? CoreCount { get; set; }

        [Input("customProperties")]
        private InputList<Inputs.IntegrationRuntimeDataFlowPropertiesCustomPropertiesArgs>? _customProperties;

        /// <summary>
        /// Custom properties are used to tune the data flow runtime performance.
        /// </summary>
        public InputList<Inputs.IntegrationRuntimeDataFlowPropertiesCustomPropertiesArgs> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputList<Inputs.IntegrationRuntimeDataFlowPropertiesCustomPropertiesArgs>());
            set => _customProperties = value;
        }

        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job.
        /// </summary>
        [Input("timeToLive")]
        public Input<int>? TimeToLive { get; set; }

        public IntegrationRuntimeDataFlowPropertiesArgs()
        {
        }
        public static new IntegrationRuntimeDataFlowPropertiesArgs Empty => new IntegrationRuntimeDataFlowPropertiesArgs();
    }
}
