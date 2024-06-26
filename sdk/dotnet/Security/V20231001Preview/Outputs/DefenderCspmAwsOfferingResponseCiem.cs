// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20231001Preview.Outputs
{

    /// <summary>
    /// Defenders CSPM Cloud infrastructure entitlement management (CIEM) offering configurations
    /// </summary>
    [OutputType]
    public sealed class DefenderCspmAwsOfferingResponseCiem
    {
        /// <summary>
        /// Defender CSPM CIEM discovery configuration
        /// </summary>
        public readonly Outputs.DefenderCspmAwsOfferingResponseCiemDiscovery? CiemDiscovery;
        /// <summary>
        /// Defender CSPM CIEM AWS OIDC (open id connect) configuration
        /// </summary>
        public readonly Outputs.DefenderCspmAwsOfferingResponseCiemOidc? CiemOidc;

        [OutputConstructor]
        private DefenderCspmAwsOfferingResponseCiem(
            Outputs.DefenderCspmAwsOfferingResponseCiemDiscovery? ciemDiscovery,

            Outputs.DefenderCspmAwsOfferingResponseCiemOidc? ciemOidc)
        {
            CiemDiscovery = ciemDiscovery;
            CiemOidc = ciemOidc;
        }
    }
}
