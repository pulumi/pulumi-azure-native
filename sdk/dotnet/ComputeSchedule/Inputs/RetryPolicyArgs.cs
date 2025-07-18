// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ComputeSchedule.Inputs
{

    /// <summary>
    /// The retry policy for the user request
    /// </summary>
    public sealed class RetryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retry count for user request
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// Retry window in minutes for user request
        /// </summary>
        [Input("retryWindowInMinutes")]
        public Input<int>? RetryWindowInMinutes { get; set; }

        public RetryPolicyArgs()
        {
        }
        public static new RetryPolicyArgs Empty => new RetryPolicyArgs();
    }
}
