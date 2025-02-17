// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20230601Preview.Outputs
{

    /// <summary>
    /// CA certificate subject distinguished name information used by service to authenticate clients.
    /// For more information, see https://docs.microsoft.com/en-us/dotnet/api/system.security.cryptography.x509certificates.x500distinguishedname?view=net-6.0#remarks
    /// </summary>
    [OutputType]
    public sealed class ClientCertificateSubjectDistinguishedNameResponse
    {
        /// <summary>
        /// The common name field in the subject name. The allowed limit is 64 characters and it should be specified.
        /// </summary>
        public readonly string? CommonName;
        /// <summary>
        /// The country code field in the subject name. If present, the country code should be represented by two-letter code defined in ISO 2166-1 (alpha-2). For example: 'US'.
        /// </summary>
        public readonly string? CountryCode;
        /// <summary>
        /// The organization field in the subject name. If present, the allowed limit is 64 characters.
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// The organization unit field in the subject name. If present, the allowed limit is 32 characters.
        /// </summary>
        public readonly string? OrganizationUnit;

        [OutputConstructor]
        private ClientCertificateSubjectDistinguishedNameResponse(
            string? commonName,

            string? countryCode,

            string? organization,

            string? organizationUnit)
        {
            CommonName = commonName;
            CountryCode = countryCode;
            Organization = organization;
            OrganizationUnit = organizationUnit;
        }
    }
}
