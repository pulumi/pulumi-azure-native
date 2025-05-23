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
    /// Resources created in GitHub repository.
    /// </summary>
    [OutputType]
    public sealed class GitHubResourceInfoResponse
    {
        /// <summary>
        /// GitHub application installation id.
        /// </summary>
        public readonly string? AppInstallationId;

        [OutputConstructor]
        private GitHubResourceInfoResponse(string? appInstallationId)
        {
            AppInstallationId = appInstallationId;
        }
    }
}
