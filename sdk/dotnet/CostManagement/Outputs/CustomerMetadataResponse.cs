// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    /// <summary>
    /// The customer billing metadata
    /// </summary>
    [OutputType]
    public sealed class CustomerMetadataResponse
    {
        /// <summary>
        /// Customer billing account id
        /// </summary>
        public readonly string BillingAccountId;
        /// <summary>
        /// Customer billing profile id
        /// </summary>
        public readonly string BillingProfileId;

        [OutputConstructor]
        private CustomerMetadataResponse(
            string billingAccountId,

            string billingProfileId)
        {
            BillingAccountId = billingAccountId;
            BillingProfileId = billingProfileId;
        }
    }
}
