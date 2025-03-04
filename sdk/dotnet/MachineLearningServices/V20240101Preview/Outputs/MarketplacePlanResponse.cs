// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240101Preview.Outputs
{

    [OutputType]
    public sealed class MarketplacePlanResponse
    {
        /// <summary>
        /// The Offer ID of the Marketplace Plan.
        /// </summary>
        public readonly string OfferId;
        /// <summary>
        /// The Plan ID of the Marketplace Plan.
        /// </summary>
        public readonly string PlanId;
        /// <summary>
        /// The Publisher ID of the Marketplace Plan.
        /// </summary>
        public readonly string PublisherId;

        [OutputConstructor]
        private MarketplacePlanResponse(
            string offerId,

            string planId,

            string publisherId)
        {
            OfferId = offerId;
            PlanId = planId;
            PublisherId = publisherId;
        }
    }
}
