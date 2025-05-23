// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard.Outputs
{

    /// <summary>
    /// Integrations for Azure Monitor Workspace.
    /// </summary>
    [OutputType]
    public sealed class AzureMonitorWorkspaceIntegrationResponse
    {
        /// <summary>
        /// The resource Id of the connected Azure Monitor Workspace.
        /// </summary>
        public readonly string? AzureMonitorWorkspaceResourceId;

        [OutputConstructor]
        private AzureMonitorWorkspaceIntegrationResponse(string? azureMonitorWorkspaceResourceId)
        {
            AzureMonitorWorkspaceResourceId = azureMonitorWorkspaceResourceId;
        }
    }
}
