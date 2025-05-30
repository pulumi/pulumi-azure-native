// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of FunctionAssociation
    /// </summary>
    [OutputType]
    public sealed class FunctionAssociationResponse
    {
        /// <summary>
        /// The event type of the function, either ``viewer-request`` or ``viewer-response``. You cannot use origin-facing event types (``origin-request`` and ``origin-response``) with a CloudFront function.
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the function.
        /// </summary>
        public readonly string? FunctionARN;

        [OutputConstructor]
        private FunctionAssociationResponse(
            string? eventType,

            string? functionARN)
        {
            EventType = eventType;
            FunctionARN = functionARN;
        }
    }
}
