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
    /// Definition of TagProperty
    /// </summary>
    public sealed class TagPropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tag key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Set to ``true`` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group. Set to ``false`` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.
        /// </summary>
        [Input("propagateAtLaunch")]
        public Input<bool>? PropagateAtLaunch { get; set; }

        /// <summary>
        /// The tag value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TagPropertyArgs()
        {
        }
        public static new TagPropertyArgs Empty => new TagPropertyArgs();
    }
}
