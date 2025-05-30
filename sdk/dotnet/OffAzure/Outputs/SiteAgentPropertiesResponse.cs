// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OffAzure.Outputs
{

    /// <summary>
    /// Class for site agent properties.
    /// </summary>
    [OutputType]
    public sealed class SiteAgentPropertiesResponse
    {
        /// <summary>
        /// Gets the ID of the agent.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets the key vault ARM Id.
        /// </summary>
        public readonly string? KeyVaultId;
        /// <summary>
        /// Gets or sets the key vault URI.
        /// </summary>
        public readonly string? KeyVaultUri;
        /// <summary>
        /// Gets the last heartbeat time of the agent in UTC.
        /// </summary>
        public readonly string LastHeartBeatUtc;
        /// <summary>
        /// Gets the version of the agent.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private SiteAgentPropertiesResponse(
            string id,

            string? keyVaultId,

            string? keyVaultUri,

            string lastHeartBeatUtc,

            string version)
        {
            Id = id;
            KeyVaultId = keyVaultId;
            KeyVaultUri = keyVaultUri;
            LastHeartBeatUtc = lastHeartBeatUtc;
            Version = version;
        }
    }
}
