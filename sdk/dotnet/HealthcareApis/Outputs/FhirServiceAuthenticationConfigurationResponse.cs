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
    public sealed class FhirServiceAuthenticationConfigurationResponse
    {
        /// <summary>
        /// The audience url for the service
        /// </summary>
        public readonly string? Audience;
        /// <summary>
        /// The authority url for the service
        /// </summary>
        public readonly string? Authority;
        /// <summary>
        /// The array of identity provider configurations for SMART on FHIR authentication.
        /// </summary>
        public readonly ImmutableArray<Outputs.SmartIdentityProviderConfigurationResponse> SmartIdentityProviders;
        /// <summary>
        /// If the SMART on FHIR proxy is enabled
        /// </summary>
        public readonly bool? SmartProxyEnabled;

        [OutputConstructor]
        private FhirServiceAuthenticationConfigurationResponse(
            string? audience,

            string? authority,

            ImmutableArray<Outputs.SmartIdentityProviderConfigurationResponse> smartIdentityProviders,

            bool? smartProxyEnabled)
        {
            Audience = audience;
            Authority = authority;
            SmartIdentityProviders = smartIdentityProviders;
            SmartProxyEnabled = smartProxyEnabled;
        }
    }
}
