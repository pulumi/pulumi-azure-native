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
    /// Windows Server licensing settings.
    /// </summary>
    [OutputType]
    public sealed class WindowsServerLicensingSettingsResponse
    {
        /// <summary>
        /// Licence Cost.
        /// </summary>
        public readonly double LicenseCost;
        /// <summary>
        /// Licenses per core.
        /// </summary>
        public readonly int LicensesPerCore;
        /// <summary>
        /// Software assurance (SA) cost.
        /// </summary>
        public readonly double SoftwareAssuranceCost;

        [OutputConstructor]
        private WindowsServerLicensingSettingsResponse(
            double licenseCost,

            int licensesPerCore,

            double softwareAssuranceCost)
        {
            LicenseCost = licenseCost;
            LicensesPerCore = licensesPerCore;
            SoftwareAssuranceCost = softwareAssuranceCost;
        }
    }
}
