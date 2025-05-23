// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// Log analytics workspace id and primary key
    /// </summary>
    [OutputType]
    public sealed class LogAnalyticsWorkspaceConfigResponse
    {
        /// <summary>
        /// Azure Log Analytics workspace ID
        /// </summary>
        public readonly string? WorkspaceId;

        [OutputConstructor]
        private LogAnalyticsWorkspaceConfigResponse(string? workspaceId)
        {
            WorkspaceId = workspaceId;
        }
    }
}
