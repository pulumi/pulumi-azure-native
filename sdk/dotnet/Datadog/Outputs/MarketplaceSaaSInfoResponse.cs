// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Datadog.Outputs
{

    /// <summary>
    /// Marketplace SAAS Info of the resource.
    /// </summary>
    [OutputType]
    public sealed class MarketplaceSaaSInfoResponse
    {
        /// <summary>
        /// The Azure Subscription ID to which the Marketplace Subscription belongs and gets billed into.
        /// </summary>
        public readonly string? BilledAzureSubscriptionId;
        /// <summary>
        /// Marketplace Subscription Details: SAAS Name
        /// </summary>
        public readonly string? MarketplaceName;
        /// <summary>
        /// Marketplace Subscription Details: SaaS Subscription Status
        /// </summary>
        public readonly string? MarketplaceStatus;
        /// <summary>
        /// Marketplace Subscription Id. This is a GUID-formatted string.
        /// </summary>
        public readonly string? MarketplaceSubscriptionId;
        /// <summary>
        /// Flag specifying if the Marketplace status is subscribed or not.
        /// </summary>
        public readonly bool? Subscribed;

        [OutputConstructor]
        private MarketplaceSaaSInfoResponse(
            string? billedAzureSubscriptionId,

            string? marketplaceName,

            string? marketplaceStatus,

            string? marketplaceSubscriptionId,

            bool? subscribed)
        {
            BilledAzureSubscriptionId = billedAzureSubscriptionId;
            MarketplaceName = marketplaceName;
            MarketplaceStatus = marketplaceStatus;
            MarketplaceSubscriptionId = marketplaceSubscriptionId;
            Subscribed = subscribed;
        }
    }
}
