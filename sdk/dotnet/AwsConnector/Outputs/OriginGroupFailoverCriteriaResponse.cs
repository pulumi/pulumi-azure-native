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
    /// Definition of OriginGroupFailoverCriteria
    /// </summary>
    [OutputType]
    public sealed class OriginGroupFailoverCriteriaResponse
    {
        /// <summary>
        /// The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin. A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.
        /// </summary>
        public readonly Outputs.StatusCodesResponse? StatusCodes;

        [OutputConstructor]
        private OriginGroupFailoverCriteriaResponse(Outputs.StatusCodesResponse? statusCodes)
        {
            StatusCodes = statusCodes;
        }
    }
}
