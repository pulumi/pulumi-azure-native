// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.V20220829.Inputs
{

    /// <summary>
    /// MarketplaceDetails of PAN Firewall resource
    /// </summary>
    public sealed class MarketplaceDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Marketplace Subscription Status
        /// </summary>
        [Input("marketplaceSubscriptionStatus")]
        public InputUnion<string, Pulumi.AzureNative.Cloudngfw.V20220829.MarketplaceSubscriptionStatus>? MarketplaceSubscriptionStatus { get; set; }

        /// <summary>
        /// Offer Id
        /// </summary>
        [Input("offerId", required: true)]
        public Input<string> OfferId { get; set; } = null!;

        /// <summary>
        /// Publisher Id
        /// </summary>
        [Input("publisherId", required: true)]
        public Input<string> PublisherId { get; set; } = null!;

        public MarketplaceDetailsArgs()
        {
        }
        public static new MarketplaceDetailsArgs Empty => new MarketplaceDetailsArgs();
    }
}
