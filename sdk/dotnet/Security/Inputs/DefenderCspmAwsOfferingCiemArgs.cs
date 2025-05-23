// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// Defenders CSPM Permissions Management offering configurations
    /// </summary>
    public sealed class DefenderCspmAwsOfferingCiemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defender CSPM Permissions Management discovery configuration
        /// </summary>
        [Input("ciemDiscovery")]
        public Input<Inputs.DefenderCspmAwsOfferingCiemDiscoveryArgs>? CiemDiscovery { get; set; }

        /// <summary>
        /// AWS Defender CSPM Permissions Management OIDC (open id connect) connection configurations
        /// </summary>
        [Input("ciemOidc")]
        public Input<Inputs.DefenderCspmAwsOfferingCiemOidcArgs>? CiemOidc { get; set; }

        public DefenderCspmAwsOfferingCiemArgs()
        {
        }
        public static new DefenderCspmAwsOfferingCiemArgs Empty => new DefenderCspmAwsOfferingCiemArgs();
    }
}
