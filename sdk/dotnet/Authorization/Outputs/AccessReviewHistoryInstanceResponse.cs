// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Outputs
{

    /// <summary>
    /// Access Review History Definition Instance.
    /// </summary>
    [OutputType]
    public sealed class AccessReviewHistoryInstanceResponse
    {
        /// <summary>
        /// The display name for the parent history definition.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Uri which can be used to retrieve review history data. To generate this Uri, generateDownloadUri() must be called for a specific accessReviewHistoryDefinitionInstance. The link expires after a 24 hour period. Callers can see the expiration date time by looking at the 'se' parameter in the generated uri.
        /// </summary>
        public readonly string DownloadUri;
        /// <summary>
        /// Date time when history data report expires and the associated data is deleted.
        /// </summary>
        public readonly string? Expiration;
        /// <summary>
        /// Date time when the history data report is scheduled to be generated.
        /// </summary>
        public readonly string? FulfilledDateTime;
        /// <summary>
        /// The access review history definition instance id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The access review history definition instance unique id.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Date time used when selecting review data, all reviews included in data end on or before this date. For use only with one-time/non-recurring reports.
        /// </summary>
        public readonly string? ReviewHistoryPeriodEndDateTime;
        /// <summary>
        /// Date time used when selecting review data, all reviews included in data start on or after this date. For use only with one-time/non-recurring reports.
        /// </summary>
        public readonly string? ReviewHistoryPeriodStartDateTime;
        /// <summary>
        /// Date time when the history data report is scheduled to be generated.
        /// </summary>
        public readonly string? RunDateTime;
        /// <summary>
        /// Status of the requested review history instance data. This is either requested, in-progress, done or error. The state transitions are as follows - Requested -&gt; InProgress -&gt; Done -&gt; Expired
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccessReviewHistoryInstanceResponse(
            string? displayName,

            string downloadUri,

            string? expiration,

            string? fulfilledDateTime,

            string id,

            string name,

            string? reviewHistoryPeriodEndDateTime,

            string? reviewHistoryPeriodStartDateTime,

            string? runDateTime,

            string status,

            string type)
        {
            DisplayName = displayName;
            DownloadUri = downloadUri;
            Expiration = expiration;
            FulfilledDateTime = fulfilledDateTime;
            Id = id;
            Name = name;
            ReviewHistoryPeriodEndDateTime = reviewHistoryPeriodEndDateTime;
            ReviewHistoryPeriodStartDateTime = reviewHistoryPeriodStartDateTime;
            RunDateTime = runDateTime;
            Status = status;
            Type = type;
        }
    }
}
