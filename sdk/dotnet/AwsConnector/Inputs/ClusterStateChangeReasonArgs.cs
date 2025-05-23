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
    /// Definition of ClusterStateChangeReason
    /// </summary>
    public sealed class ClusterStateChangeReasonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The programmatic code for the state change reason.&lt;/p&gt;
        /// </summary>
        [Input("code")]
        public Input<Inputs.ClusterStateChangeReasonCodeEnumValueArgs>? Code { get; set; }

        /// <summary>
        /// &lt;p&gt;The descriptive message for the state change reason.&lt;/p&gt;
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public ClusterStateChangeReasonArgs()
        {
        }
        public static new ClusterStateChangeReasonArgs Empty => new ClusterStateChangeReasonArgs();
    }
}
