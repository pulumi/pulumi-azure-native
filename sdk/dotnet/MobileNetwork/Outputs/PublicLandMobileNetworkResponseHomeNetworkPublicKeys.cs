// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Outputs
{

    /// <summary>
    /// Configuration relating to SUPI concealment.
    /// </summary>
    [OutputType]
    public sealed class PublicLandMobileNetworkResponseHomeNetworkPublicKeys
    {
        /// <summary>
        /// This provides a mapping to identify which public key has been used for SUPI concealment using the Profile A Protection Scheme.
        /// </summary>
        public readonly ImmutableArray<Outputs.HomeNetworkPublicKeyResponse> ProfileA;
        /// <summary>
        /// This provides a mapping to identify which public key has been used for SUPI concealment using the Profile B Protection Scheme.
        /// </summary>
        public readonly ImmutableArray<Outputs.HomeNetworkPublicKeyResponse> ProfileB;

        [OutputConstructor]
        private PublicLandMobileNetworkResponseHomeNetworkPublicKeys(
            ImmutableArray<Outputs.HomeNetworkPublicKeyResponse> profileA,

            ImmutableArray<Outputs.HomeNetworkPublicKeyResponse> profileB)
        {
            ProfileA = profileA;
            ProfileB = profileB;
        }
    }
}
