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
    /// Resources created in GitHub repository.
    /// </summary>
    public sealed class GitHubResourceInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GitHub application installation id.
        /// </summary>
        [Input("appInstallationId")]
        public Input<string>? AppInstallationId { get; set; }

        public GitHubResourceInfoArgs()
        {
        }
        public static new GitHubResourceInfoArgs Empty => new GitHubResourceInfoArgs();
    }
}
