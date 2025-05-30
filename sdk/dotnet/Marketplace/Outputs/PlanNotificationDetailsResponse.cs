// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Marketplace.Outputs
{

    /// <summary>
    /// Plan notification details
    /// </summary>
    [OutputType]
    public sealed class PlanNotificationDetailsResponse
    {
        /// <summary>
        /// Gets or sets the plan display name
        /// </summary>
        public readonly string? PlanDisplayName;
        /// <summary>
        /// Gets or sets the plan id
        /// </summary>
        public readonly string? PlanId;

        [OutputConstructor]
        private PlanNotificationDetailsResponse(
            string? planDisplayName,

            string? planId)
        {
            PlanDisplayName = planDisplayName;
            PlanId = planId;
        }
    }
}
