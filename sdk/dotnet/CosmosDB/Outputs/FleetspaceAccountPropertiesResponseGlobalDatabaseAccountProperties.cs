// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Outputs
{

    /// <summary>
    /// Configuration for fleetspace Account in the fleetspace.
    /// </summary>
    [OutputType]
    public sealed class FleetspaceAccountPropertiesResponseGlobalDatabaseAccountProperties
    {
        /// <summary>
        /// The location of  global database account in the Fleetspace Account.
        /// </summary>
        public readonly string? ArmLocation;
        /// <summary>
        /// The resource identifier of global database account in the Fleetspace Account.
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private FleetspaceAccountPropertiesResponseGlobalDatabaseAccountProperties(
            string? armLocation,

            string? resourceId)
        {
            ArmLocation = armLocation;
            ResourceId = resourceId;
        }
    }
}
