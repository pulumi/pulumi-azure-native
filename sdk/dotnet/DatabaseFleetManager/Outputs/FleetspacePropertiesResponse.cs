// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DatabaseFleetManager.Outputs
{

    /// <summary>
    /// A Fleetspace properties.
    /// </summary>
    [OutputType]
    public sealed class FleetspacePropertiesResponse
    {
        /// <summary>
        /// Maximum number of vCores database fleet manager is allowed to provision in the fleetspace.
        /// </summary>
        public readonly int? CapacityMax;
        /// <summary>
        /// Main Microsoft Entra ID principal that has admin access to all databases in the fleetspace.
        /// </summary>
        public readonly Outputs.MainPrincipalResponse? MainPrincipal;
        /// <summary>
        /// Fleetspace state.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private FleetspacePropertiesResponse(
            int? capacityMax,

            Outputs.MainPrincipalResponse? mainPrincipal,

            string provisioningState)
        {
            CapacityMax = capacityMax;
            MainPrincipal = mainPrincipal;
            ProvisioningState = provisioningState;
        }
    }
}
