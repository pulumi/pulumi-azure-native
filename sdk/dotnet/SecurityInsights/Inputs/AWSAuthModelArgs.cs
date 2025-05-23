// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Model for API authentication with AWS.
    /// </summary>
    public sealed class AWSAuthModelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS STS assume role external ID. This is used to prevent the confused deputy problem: 'https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html'
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// AWS STS assume role ARN
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Type of paging
        /// Expected value is 'AWS'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AWSAuthModelArgs()
        {
        }
        public static new AWSAuthModelArgs Empty => new AWSAuthModelArgs();
    }
}
