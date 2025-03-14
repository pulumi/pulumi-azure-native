// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logz.V20220101Preview.Outputs
{

    [OutputType]
    public sealed class PlanDataResponse
    {
        /// <summary>
        /// different billing cycles like MONTHLY/WEEKLY. this could be enum
        /// </summary>
        public readonly string? BillingCycle;
        /// <summary>
        /// date when plan was applied
        /// </summary>
        public readonly string? EffectiveDate;
        /// <summary>
        /// plan id as published by Logz
        /// </summary>
        public readonly string? PlanDetails;
        /// <summary>
        /// different usage type like PAYG/COMMITTED. this could be enum
        /// </summary>
        public readonly string? UsageType;

        [OutputConstructor]
        private PlanDataResponse(
            string? billingCycle,

            string? effectiveDate,

            string? planDetails,

            string? usageType)
        {
            BillingCycle = billingCycle;
            EffectiveDate = effectiveDate;
            PlanDetails = planDetails;
            UsageType = usageType;
        }
    }
}
