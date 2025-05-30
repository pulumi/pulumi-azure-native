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
    /// Defines the move collection properties.
    /// </summary>
    [OutputType]
    public sealed class MoveCollectionPropertiesResponse
    {
        /// <summary>
        /// Defines the move collection errors.
        /// </summary>
        public readonly Outputs.MoveCollectionPropertiesResponseErrors Errors;
        /// <summary>
        /// Gets or sets the move region which indicates the region where the VM Regional to Zonal move will be conducted.
        /// </summary>
        public readonly string? MoveRegion;
        /// <summary>
        /// Defines the MoveType.
        /// </summary>
        public readonly string? MoveType;
        /// <summary>
        /// Defines the provisioning states.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Gets or sets the source region.
        /// </summary>
        public readonly string? SourceRegion;
        /// <summary>
        /// Gets or sets the target region.
        /// </summary>
        public readonly string? TargetRegion;
        /// <summary>
        /// Gets or sets the version of move collection.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private MoveCollectionPropertiesResponse(
            Outputs.MoveCollectionPropertiesResponseErrors errors,

            string? moveRegion,

            string? moveType,

            string provisioningState,

            string? sourceRegion,

            string? targetRegion,

            string? version)
        {
            Errors = errors;
            MoveRegion = moveRegion;
            MoveType = moveType;
            ProvisioningState = provisioningState;
            SourceRegion = sourceRegion;
            TargetRegion = targetRegion;
            Version = version;
        }
    }
}
