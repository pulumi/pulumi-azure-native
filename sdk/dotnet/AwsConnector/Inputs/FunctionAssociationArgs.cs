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
    /// Definition of FunctionAssociation
    /// </summary>
    public sealed class FunctionAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The event type of the function, either ``viewer-request`` or ``viewer-response``. You cannot use origin-facing event types (``origin-request`` and ``origin-response``) with a CloudFront function.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the function.
        /// </summary>
        [Input("functionARN")]
        public Input<string>? FunctionARN { get; set; }

        public FunctionAssociationArgs()
        {
        }
        public static new FunctionAssociationArgs Empty => new FunctionAssociationArgs();
    }
}
