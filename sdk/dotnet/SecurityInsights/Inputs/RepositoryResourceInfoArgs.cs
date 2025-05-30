// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Resources created in user's repository for the source-control.
    /// </summary>
    public sealed class RepositoryResourceInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resources created in Azure DevOps for this source-control.
        /// </summary>
        [Input("azureDevOpsResourceInfo")]
        public Input<Inputs.AzureDevOpsResourceInfoArgs>? AzureDevOpsResourceInfo { get; set; }

        /// <summary>
        /// Resources created in GitHub for this source-control.
        /// </summary>
        [Input("gitHubResourceInfo")]
        public Input<Inputs.GitHubResourceInfoArgs>? GitHubResourceInfo { get; set; }

        /// <summary>
        /// The webhook object created for the source-control.
        /// </summary>
        [Input("webhook")]
        public Input<Inputs.WebhookArgs>? Webhook { get; set; }

        public RepositoryResourceInfoArgs()
        {
        }
        public static new RepositoryResourceInfoArgs Empty => new RepositoryResourceInfoArgs();
    }
}
