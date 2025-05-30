// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Outputs
{

    /// <summary>
    /// Billing plan information.
    /// </summary>
    [OutputType]
    public sealed class PlanDataResponse
    {
        /// <summary>
        /// different billing cycles like MONTHLY/WEEKLY
        /// </summary>
        public readonly string BillingCycle;
        /// <summary>
        /// date when plan was applied
        /// </summary>
        public readonly string EffectiveDate;
        /// <summary>
        /// plan id as published by Liftr.PAN
        /// </summary>
        public readonly string PlanId;
        /// <summary>
        /// different usage type like PAYG/COMMITTED
        /// </summary>
        public readonly string? UsageType;

        [OutputConstructor]
        private PlanDataResponse(
            string billingCycle,

            string effectiveDate,

            string planId,

            string? usageType)
        {
            BillingCycle = billingCycle;
            EffectiveDate = effectiveDate;
            PlanId = planId;
            UsageType = usageType;
        }
    }
}
