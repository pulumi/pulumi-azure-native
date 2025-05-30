// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// Active Directory domain details
    /// </summary>
    [OutputType]
    public sealed class ActiveDirectoryConnectorDomainDetailsResponse
    {
        /// <summary>
        /// null
        /// </summary>
        public readonly Outputs.ActiveDirectoryDomainControllersResponse? DomainControllers;
        /// <summary>
        /// NETBIOS name of the Active Directory domain.
        /// </summary>
        public readonly string? NetbiosDomainName;
        /// <summary>
        /// The distinguished name of the Active Directory Organizational Unit.
        /// </summary>
        public readonly string? OuDistinguishedName;
        /// <summary>
        /// Name (uppercase) of the Active Directory domain that this AD connector will be associated with.
        /// </summary>
        public readonly string Realm;
        /// <summary>
        /// The service account provisioning mode for this Active Directory connector.
        /// </summary>
        public readonly string? ServiceAccountProvisioning;

        [OutputConstructor]
        private ActiveDirectoryConnectorDomainDetailsResponse(
            Outputs.ActiveDirectoryDomainControllersResponse? domainControllers,

            string? netbiosDomainName,

            string? ouDistinguishedName,

            string realm,

            string? serviceAccountProvisioning)
        {
            DomainControllers = domainControllers;
            NetbiosDomainName = netbiosDomainName;
            OuDistinguishedName = ouDistinguishedName;
            Realm = realm;
            ServiceAccountProvisioning = serviceAccountProvisioning;
        }
    }
}
