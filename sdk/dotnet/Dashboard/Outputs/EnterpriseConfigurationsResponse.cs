// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard.Outputs
{

    /// <summary>
    /// Enterprise settings of a Grafana instance
    /// </summary>
    [OutputType]
    public sealed class EnterpriseConfigurationsResponse
    {
        /// <summary>
        /// The AutoRenew setting of the Enterprise subscription
        /// </summary>
        public readonly string? MarketplaceAutoRenew;
        /// <summary>
        /// The Plan Id of the Azure Marketplace subscription for the Enterprise plugins
        /// </summary>
        public readonly string? MarketplacePlanId;

        [OutputConstructor]
        private EnterpriseConfigurationsResponse(
            string? marketplaceAutoRenew,

            string? marketplacePlanId)
        {
            MarketplaceAutoRenew = marketplaceAutoRenew;
            MarketplacePlanId = marketplacePlanId;
        }
    }
}
