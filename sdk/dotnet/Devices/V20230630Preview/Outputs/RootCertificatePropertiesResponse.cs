// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20230630Preview.Outputs
{

    /// <summary>
    /// This property store root certificate related information
    /// </summary>
    [OutputType]
    public sealed class RootCertificatePropertiesResponse
    {
        /// <summary>
        /// This property when set to true, hub will use G2 cert; while it's set to false, hub uses Baltimore Cert.
        /// </summary>
        public readonly bool? EnableRootCertificateV2;
        /// <summary>
        /// the last update time to root certificate flag.
        /// </summary>
        public readonly string LastUpdatedTimeUtc;

        [OutputConstructor]
        private RootCertificatePropertiesResponse(
            bool? enableRootCertificateV2,

            string lastUpdatedTimeUtc)
        {
            EnableRootCertificateV2 = enableRootCertificateV2;
            LastUpdatedTimeUtc = lastUpdatedTimeUtc;
        }
    }
}
