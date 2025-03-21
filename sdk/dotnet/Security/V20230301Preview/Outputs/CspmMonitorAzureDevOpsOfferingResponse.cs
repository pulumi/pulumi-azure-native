// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20230301Preview.Outputs
{

    /// <summary>
    /// The CSPM monitoring for AzureDevOps offering
    /// </summary>
    [OutputType]
    public sealed class CspmMonitorAzureDevOpsOfferingResponse
    {
        /// <summary>
        /// The offering description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The type of the security offering.
        /// Expected value is 'CspmMonitorAzureDevOps'.
        /// </summary>
        public readonly string OfferingType;

        [OutputConstructor]
        private CspmMonitorAzureDevOpsOfferingResponse(
            string description,

            string offeringType)
        {
            Description = description;
            OfferingType = offeringType;
        }
    }
}
