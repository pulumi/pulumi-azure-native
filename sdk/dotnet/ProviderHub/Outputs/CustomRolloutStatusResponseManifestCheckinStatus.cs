// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    /// <summary>
    /// The manifest checkin status.
    /// </summary>
    [OutputType]
    public sealed class CustomRolloutStatusResponseManifestCheckinStatus
    {
        /// <summary>
        /// The commit id.
        /// </summary>
        public readonly string? CommitId;
        /// <summary>
        /// Whether the manifest is checked in.
        /// </summary>
        public readonly bool IsCheckedIn;
        /// <summary>
        /// The pull request.
        /// </summary>
        public readonly string? PullRequest;
        /// <summary>
        /// The status message.
        /// </summary>
        public readonly string StatusMessage;

        [OutputConstructor]
        private CustomRolloutStatusResponseManifestCheckinStatus(
            string? commitId,

            bool isCheckedIn,

            string? pullRequest,

            string statusMessage)
        {
            CommitId = commitId;
            IsCheckedIn = isCheckedIn;
            PullRequest = pullRequest;
            StatusMessage = statusMessage;
        }
    }
}
