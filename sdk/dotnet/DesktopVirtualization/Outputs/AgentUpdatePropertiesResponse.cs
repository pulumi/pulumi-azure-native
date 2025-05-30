// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.Outputs
{

    /// <summary>
    /// The session host configuration for updating agent, monitoring agent, and stack component.
    /// </summary>
    [OutputType]
    public sealed class AgentUpdatePropertiesResponse
    {
        /// <summary>
        /// Time zone for maintenance as defined in https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.findsystemtimezonebyid?view=net-5.0. Must be set if useLocalTime is true.
        /// </summary>
        public readonly string? MaintenanceWindowTimeZone;
        /// <summary>
        /// List of maintenance windows. Maintenance windows are 2 hours long.
        /// </summary>
        public readonly ImmutableArray<Outputs.MaintenanceWindowPropertiesResponse> MaintenanceWindows;
        /// <summary>
        /// The type of maintenance for session host components.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Whether to use localTime of the virtual machine.
        /// </summary>
        public readonly bool? UseSessionHostLocalTime;

        [OutputConstructor]
        private AgentUpdatePropertiesResponse(
            string? maintenanceWindowTimeZone,

            ImmutableArray<Outputs.MaintenanceWindowPropertiesResponse> maintenanceWindows,

            string? type,

            bool? useSessionHostLocalTime)
        {
            MaintenanceWindowTimeZone = maintenanceWindowTimeZone;
            MaintenanceWindows = maintenanceWindows;
            Type = type;
            UseSessionHostLocalTime = useSessionHostLocalTime;
        }
    }
}
