// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OperationalInsights.Outputs
{

    /// <summary>
    /// Workspace features.
    /// </summary>
    [OutputType]
    public sealed class WorkspaceFeaturesResponse
    {
        /// <summary>
        /// Dedicated LA cluster resourceId that is linked to the workspaces.
        /// </summary>
        public readonly string? ClusterResourceId;
        /// <summary>
        /// Disable Non-AAD based Auth.
        /// </summary>
        public readonly bool? DisableLocalAuth;
        /// <summary>
        /// Flag that indicate if data should be exported.
        /// </summary>
        public readonly bool? EnableDataExport;
        /// <summary>
        /// Flag that indicate which permission to use - resource or workspace or both.
        /// </summary>
        public readonly bool? EnableLogAccessUsingOnlyResourcePermissions;
        /// <summary>
        /// Flag that describes if we want to remove the data after 30 days.
        /// </summary>
        public readonly bool? ImmediatePurgeDataOn30Days;
        /// <summary>
        /// An indication if the specify workspace is limited to sentinel's unified billing model only.
        /// </summary>
        public readonly bool UnifiedSentinelBillingOnly;

        [OutputConstructor]
        private WorkspaceFeaturesResponse(
            string? clusterResourceId,

            bool? disableLocalAuth,

            bool? enableDataExport,

            bool? enableLogAccessUsingOnlyResourcePermissions,

            bool? immediatePurgeDataOn30Days,

            bool unifiedSentinelBillingOnly)
        {
            ClusterResourceId = clusterResourceId;
            DisableLocalAuth = disableLocalAuth;
            EnableDataExport = enableDataExport;
            EnableLogAccessUsingOnlyResourcePermissions = enableLogAccessUsingOnlyResourcePermissions;
            ImmediatePurgeDataOn30Days = immediatePurgeDataOn30Days;
            UnifiedSentinelBillingOnly = unifiedSentinelBillingOnly;
        }
    }
}
