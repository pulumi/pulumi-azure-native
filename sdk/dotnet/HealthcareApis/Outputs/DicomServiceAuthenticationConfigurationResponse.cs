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
    /// Authentication configuration information
    /// </summary>
    [OutputType]
    public sealed class DicomServiceAuthenticationConfigurationResponse
    {
        /// <summary>
        /// The audiences for the service
        /// </summary>
        public readonly ImmutableArray<string> Audiences;
        /// <summary>
        /// The authority url for the service
        /// </summary>
        public readonly string Authority;

        [OutputConstructor]
        private DicomServiceAuthenticationConfigurationResponse(
            ImmutableArray<string> audiences,

            string authority)
        {
            Audiences = audiences;
            Authority = authority;
        }
    }
}
