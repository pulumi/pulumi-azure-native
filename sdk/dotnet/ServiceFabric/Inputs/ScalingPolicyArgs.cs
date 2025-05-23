// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.Inputs
{

    /// <summary>
    /// Specifies a metric to load balance a service during runtime.
    /// </summary>
    public sealed class ScalingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the mechanism associated with this scaling policy
        /// </summary>
        [Input("scalingMechanism", required: true)]
        public InputUnion<Inputs.AddRemoveIncrementalNamedPartitionScalingMechanismArgs, Inputs.PartitionInstanceCountScaleMechanismArgs> ScalingMechanism { get; set; } = null!;

        /// <summary>
        /// Specifies the trigger associated with this scaling policy.
        /// </summary>
        [Input("scalingTrigger", required: true)]
        public InputUnion<Inputs.AveragePartitionLoadScalingTriggerArgs, Inputs.AverageServiceLoadScalingTriggerArgs> ScalingTrigger { get; set; } = null!;

        public ScalingPolicyArgs()
        {
        }
        public static new ScalingPolicyArgs Empty => new ScalingPolicyArgs();
    }
}
