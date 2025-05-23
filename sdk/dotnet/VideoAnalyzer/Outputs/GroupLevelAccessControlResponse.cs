// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// Group level network access control.
    /// </summary>
    [OutputType]
    public sealed class GroupLevelAccessControlResponse
    {
        /// <summary>
        /// Whether or not public network access is allowed for specified resources under the Video Analyzer account.
        /// </summary>
        public readonly string? PublicNetworkAccess;

        [OutputConstructor]
        private GroupLevelAccessControlResponse(string? publicNetworkAccess)
        {
            PublicNetworkAccess = publicNetworkAccess;
        }
    }
}
