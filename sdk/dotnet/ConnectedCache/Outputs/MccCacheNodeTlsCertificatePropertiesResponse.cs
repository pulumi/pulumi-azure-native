// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedCache.Outputs
{

    /// <summary>
    /// Mcc cache node resource auto update properties.
    /// </summary>
    [OutputType]
    public sealed class MccCacheNodeTlsCertificatePropertiesResponse
    {
        /// <summary>
        /// Mcc cache node resource Id.
        /// </summary>
        public readonly string CacheNodeId;
        /// <summary>
        /// Mcc customer resource Id.
        /// </summary>
        public readonly string CustomerId;
        /// <summary>
        /// Cache node resource tls certificate history details.
        /// </summary>
        public readonly ImmutableArray<Outputs.MccCacheNodeTlsCertificateResponse> TlsCertificateHistory;

        [OutputConstructor]
        private MccCacheNodeTlsCertificatePropertiesResponse(
            string cacheNodeId,

            string customerId,

            ImmutableArray<Outputs.MccCacheNodeTlsCertificateResponse> tlsCertificateHistory)
        {
            CacheNodeId = cacheNodeId;
            CustomerId = customerId;
            TlsCertificateHistory = tlsCertificateHistory;
        }
    }
}
