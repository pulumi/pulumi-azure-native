// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// Discovered entity light summary.
    /// </summary>
    [OutputType]
    public sealed class DiscoveredEntityLightSummaryResponse
    {
        /// <summary>
        /// Gets or sets the number of machines.
        /// </summary>
        public readonly int NumberOfMachines;
        /// <summary>
        /// Gets or sets the number of servers.
        /// </summary>
        public readonly int NumberOfServers;
        /// <summary>
        /// Gets or sets the number of web apps.
        /// </summary>
        public readonly int NumberOfWebApps;

        [OutputConstructor]
        private DiscoveredEntityLightSummaryResponse(
            int numberOfMachines,

            int numberOfServers,

            int numberOfWebApps)
        {
            NumberOfMachines = numberOfMachines;
            NumberOfServers = numberOfServers;
            NumberOfWebApps = numberOfWebApps;
        }
    }
}
