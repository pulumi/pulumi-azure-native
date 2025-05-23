// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of ScalingConfiguration
    /// </summary>
    public sealed class ScalingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value that indicates whether to allow or disallow automatic pause for an Aurora DB cluster in serverless DB engine mode. A DB cluster can be paused only when it's idle (it has no connections).
        /// </summary>
        [Input("autoPause")]
        public Input<bool>? AutoPause { get; set; }

        /// <summary>
        /// The maximum capacity for an Aurora DB cluster in serverless DB engine mode.For Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.For Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.The maximum capacity must be greater than or equal to the minimum capacity.
        /// </summary>
        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        /// <summary>
        /// The minimum capacity for an Aurora DB cluster in serverless DB engine mode.For Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.For Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.The minimum capacity must be less than or equal to the maximum capacity.
        /// </summary>
        [Input("minCapacity")]
        public Input<int>? MinCapacity { get; set; }

        /// <summary>
        /// The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.The default is 300.
        /// </summary>
        [Input("secondsBeforeTimeout")]
        public Input<int>? SecondsBeforeTimeout { get; set; }

        /// <summary>
        /// The time, in seconds, before an Aurora DB cluster in serverless mode is paused.
        /// </summary>
        [Input("secondsUntilAutoPause")]
        public Input<int>? SecondsUntilAutoPause { get; set; }

        /// <summary>
        /// The action to take when the timeout is reached, either ForceApplyCapacityChange or RollbackCapacityChange.ForceApplyCapacityChange sets the capacity to the specified value as soon as possible.RollbackCapacityChange, the default, ignores the capacity change if a scaling point isn't found in the timeout period.For more information, see Autoscaling for Aurora Serverless v1 in the Amazon Aurora User Guide.
        /// </summary>
        [Input("timeoutAction")]
        public Input<string>? TimeoutAction { get; set; }

        public ScalingConfigurationArgs()
        {
        }
        public static new ScalingConfigurationArgs Empty => new ScalingConfigurationArgs();
    }
}
