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
    /// Definition of ClusterStateChangeReasonCodeEnumValue
    /// </summary>
    public sealed class ClusterStateChangeReasonCodeEnumValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property value
        /// </summary>
        [Input("value")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.ClusterStateChangeReasonCode>? Value { get; set; }

        public ClusterStateChangeReasonCodeEnumValueArgs()
        {
        }
        public static new ClusterStateChangeReasonCodeEnumValueArgs Empty => new ClusterStateChangeReasonCodeEnumValueArgs();
    }
}
