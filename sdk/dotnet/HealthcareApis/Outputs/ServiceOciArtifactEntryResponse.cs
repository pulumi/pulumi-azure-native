// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Outputs
{

    /// <summary>
    /// An Open Container Initiative (OCI) artifact.
    /// </summary>
    [OutputType]
    public sealed class ServiceOciArtifactEntryResponse
    {
        /// <summary>
        /// The artifact digest.
        /// </summary>
        public readonly string? Digest;
        /// <summary>
        /// The artifact name.
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// The Azure Container Registry login server.
        /// </summary>
        public readonly string? LoginServer;

        [OutputConstructor]
        private ServiceOciArtifactEntryResponse(
            string? digest,

            string? imageName,

            string? loginServer)
        {
            Digest = digest;
            ImageName = imageName;
            LoginServer = loginServer;
        }
    }
}
