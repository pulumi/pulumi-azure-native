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
    /// Definition of LoggingProperties
    /// </summary>
    public sealed class LoggingPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property bucketName
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Property s3KeyPrefix
        /// </summary>
        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        public LoggingPropertiesArgs()
        {
        }
        public static new LoggingPropertiesArgs Empty => new LoggingPropertiesArgs();
    }
}
