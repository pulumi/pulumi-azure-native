// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Datadog.Outputs
{

    [OutputType]
    public sealed class DatadogHostMetadataResponse
    {
        /// <summary>
        /// The agent version.
        /// </summary>
        public readonly string? AgentVersion;
        public readonly Outputs.DatadogInstallMethodResponse? InstallMethod;
        public readonly Outputs.DatadogLogsAgentResponse? LogsAgent;

        [OutputConstructor]
        private DatadogHostMetadataResponse(
            string? agentVersion,

            Outputs.DatadogInstallMethodResponse? installMethod,

            Outputs.DatadogLogsAgentResponse? logsAgent)
        {
            AgentVersion = agentVersion;
            InstallMethod = installMethod;
            LogsAgent = logsAgent;
        }
    }
}
