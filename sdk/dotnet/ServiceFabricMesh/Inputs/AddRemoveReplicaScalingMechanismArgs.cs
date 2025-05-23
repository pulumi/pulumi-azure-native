// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabricMesh.Inputs
{

    /// <summary>
    /// Describes the horizontal auto scaling mechanism that adds or removes replicas (containers or container groups).
    /// </summary>
    public sealed class AddRemoveReplicaScalingMechanismArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enumerates the mechanisms for auto scaling.
        /// Expected value is 'AddRemoveReplica'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Maximum number of containers (scale up won't be performed above this number).
        /// </summary>
        [Input("maxCount", required: true)]
        public Input<int> MaxCount { get; set; } = null!;

        /// <summary>
        /// Minimum number of containers (scale down won't be performed below this number).
        /// </summary>
        [Input("minCount", required: true)]
        public Input<int> MinCount { get; set; } = null!;

        /// <summary>
        /// Each time auto scaling is performed, this number of containers will be added or removed.
        /// </summary>
        [Input("scaleIncrement", required: true)]
        public Input<int> ScaleIncrement { get; set; } = null!;

        public AddRemoveReplicaScalingMechanismArgs()
        {
        }
        public static new AddRemoveReplicaScalingMechanismArgs Empty => new AddRemoveReplicaScalingMechanismArgs();
    }
}
