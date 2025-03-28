// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.V20250301Preview.Outputs
{

    /// <summary>
    /// Azure container registry configuration information
    /// </summary>
    [OutputType]
    public sealed class FhirServiceAcrConfigurationResponse
    {
        /// <summary>
        /// The list of the Azure container registry login servers.
        /// </summary>
        public readonly ImmutableArray<string> LoginServers;
        /// <summary>
        /// The list of Open Container Initiative (OCI) artifacts.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceOciArtifactEntryResponse> OciArtifacts;

        [OutputConstructor]
        private FhirServiceAcrConfigurationResponse(
            ImmutableArray<string> loginServers,

            ImmutableArray<Outputs.ServiceOciArtifactEntryResponse> ociArtifacts)
        {
            LoginServers = loginServers;
            OciArtifacts = ociArtifacts;
        }
    }
}
