// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Policy to set request timeouts
    /// </summary>
    public sealed class TimeoutPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Timeout, in seconds, for a request to initiate a connection
        /// </summary>
        [Input("connectionTimeoutInSeconds")]
        public Input<int>? ConnectionTimeoutInSeconds { get; set; }

        /// <summary>
        /// Timeout, in seconds, for a request to respond
        /// </summary>
        [Input("responseTimeoutInSeconds")]
        public Input<int>? ResponseTimeoutInSeconds { get; set; }

        public TimeoutPolicyArgs()
        {
        }
        public static new TimeoutPolicyArgs Empty => new TimeoutPolicyArgs();
    }
}
