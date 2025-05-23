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
    /// Definition of ResourceSpec
    /// </summary>
    public sealed class ResourceSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type that the image version runs on.
        /// </summary>
        [Input("instanceType")]
        public InputUnion<string, Pulumi.AzureNative.AwsConnector.ResourceSpecInstanceType>? InstanceType { get; set; }

        /// <summary>
        /// The ARN of the SageMaker image that the image version belongs to.
        /// </summary>
        [Input("sageMakerImageArn")]
        public Input<string>? SageMakerImageArn { get; set; }

        /// <summary>
        /// The ARN of the image version created on the instance.
        /// </summary>
        [Input("sageMakerImageVersionArn")]
        public Input<string>? SageMakerImageVersionArn { get; set; }

        public ResourceSpecArgs()
        {
        }
        public static new ResourceSpecArgs Empty => new ResourceSpecArgs();
    }
}
