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
    /// Execution policy for an activity.
    /// </summary>
    public sealed class RetryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum ordinary retry attempts. Default is 0. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        [Input("count")]
        public Input<object>? Count { get; set; }

        /// <summary>
        /// Interval between retries in seconds. Default is 30.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        public RetryPolicyArgs()
        {
        }
        public static new RetryPolicyArgs Empty => new RetryPolicyArgs();
    }
}
