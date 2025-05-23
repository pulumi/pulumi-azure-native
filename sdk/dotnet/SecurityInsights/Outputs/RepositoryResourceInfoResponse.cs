// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// Resources created in user's repository for the source-control.
    /// </summary>
    [OutputType]
    public sealed class RepositoryResourceInfoResponse
    {
        /// <summary>
        /// Resources created in Azure DevOps for this source-control.
        /// </summary>
        public readonly Outputs.AzureDevOpsResourceInfoResponse? AzureDevOpsResourceInfo;
        /// <summary>
        /// Resources created in GitHub for this source-control.
        /// </summary>
        public readonly Outputs.GitHubResourceInfoResponse? GitHubResourceInfo;
        /// <summary>
        /// The webhook object created for the source-control.
        /// </summary>
        public readonly Outputs.WebhookResponse? Webhook;

        [OutputConstructor]
        private RepositoryResourceInfoResponse(
            Outputs.AzureDevOpsResourceInfoResponse? azureDevOpsResourceInfo,

            Outputs.GitHubResourceInfoResponse? gitHubResourceInfo,

            Outputs.WebhookResponse? webhook)
        {
            AzureDevOpsResourceInfo = azureDevOpsResourceInfo;
            GitHubResourceInfo = gitHubResourceInfo;
            Webhook = webhook;
        }
    }
}
