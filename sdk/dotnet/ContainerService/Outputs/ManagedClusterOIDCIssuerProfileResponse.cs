// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// The OIDC issuer profile of the Managed Cluster.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterOIDCIssuerProfileResponse
    {
        /// <summary>
        /// Whether the OIDC issuer is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The OIDC issuer url of the Managed Cluster.
        /// </summary>
        public readonly string IssuerURL;

        [OutputConstructor]
        private ManagedClusterOIDCIssuerProfileResponse(
            bool? enabled,

            string issuerURL)
        {
            Enabled = enabled;
            IssuerURL = issuerURL;
        }
    }
}
