// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    [OutputType]
    public sealed class MarketplaceSubscriptionResponse
    {
        /// <summary>
        /// Marketplace Plan associated with the Marketplace Subscription.
        /// </summary>
        public readonly Outputs.MarketplacePlanResponse MarketplacePlan;
        /// <summary>
        /// Current status of the Marketplace Subscription.
        /// </summary>
        public readonly string MarketplaceSubscriptionStatus;
        /// <summary>
        /// [Required] Target Marketplace Model ID to create a Marketplace Subscription for.
        /// </summary>
        public readonly string ModelId;
        /// <summary>
        /// Provisioning State of the Marketplace Subscription.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private MarketplaceSubscriptionResponse(
            Outputs.MarketplacePlanResponse marketplacePlan,

            string marketplaceSubscriptionStatus,

            string modelId,

            string provisioningState)
        {
            MarketplacePlan = marketplacePlan;
            MarketplaceSubscriptionStatus = marketplaceSubscriptionStatus;
            ModelId = modelId;
            ProvisioningState = provisioningState;
        }
    }
}
