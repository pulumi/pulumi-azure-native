// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.Outputs
{

    /// <summary>
    /// Set of rules for sending logs for the Monitor resource.
    /// </summary>
    [OutputType]
    public sealed class LogRulesResponse
    {
        /// <summary>
        /// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all resources will be captured. If only Exclude action is specified, the rules will apply to the list of all available resources. If Include actions are specified, the rules will only include resources with the associated tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.FilteringTagResponse> FilteringTags;
        /// <summary>
        /// Flag specifying if AAD logs should be sent for the Monitor resource.
        /// </summary>
        public readonly bool? SendAadLogs;
        /// <summary>
        /// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
        /// </summary>
        public readonly bool? SendActivityLogs;
        /// <summary>
        /// Flag specifying if subscription logs should be sent for the Monitor resource.
        /// </summary>
        public readonly bool? SendSubscriptionLogs;

        [OutputConstructor]
        private LogRulesResponse(
            ImmutableArray<Outputs.FilteringTagResponse> filteringTags,

            bool? sendAadLogs,

            bool? sendActivityLogs,

            bool? sendSubscriptionLogs)
        {
            FilteringTags = filteringTags;
            SendAadLogs = sendAadLogs;
            SendActivityLogs = sendActivityLogs;
            SendSubscriptionLogs = sendSubscriptionLogs;
        }
    }
}
