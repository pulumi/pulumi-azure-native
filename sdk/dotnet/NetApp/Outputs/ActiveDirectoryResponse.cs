// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.Outputs
{

    /// <summary>
    /// Active Directory
    /// </summary>
    [OutputType]
    public sealed class ActiveDirectoryResponse
    {
        /// <summary>
        /// Id of the Active Directory
        /// </summary>
        public readonly string? ActiveDirectoryId;
        /// <summary>
        /// Name of the active directory machine. This optional parameter is used only while creating kerberos volume
        /// </summary>
        public readonly string? AdName;
        /// <summary>
        /// Users to be added to the Built-in Administrators active directory group. A list of unique usernames without domain specifier
        /// </summary>
        public readonly ImmutableArray<string> Administrators;
        /// <summary>
        /// If enabled, AES encryption will be enabled for SMB communication.
        /// </summary>
        public readonly bool? AesEncryption;
        /// <summary>
        ///  If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.
        /// </summary>
        public readonly bool? AllowLocalNfsUsersWithLdap;
        /// <summary>
        /// Users to be added to the Built-in Backup Operator active directory group. A list of unique usernames without domain specifier
        /// </summary>
        public readonly ImmutableArray<string> BackupOperators;
        /// <summary>
        /// Comma separated list of DNS server IP addresses (IPv4 only) for the Active Directory domain
        /// </summary>
        public readonly string? Dns;
        /// <summary>
        /// Name of the Active Directory domain
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// If enabled, Traffic between the SMB server to Domain Controller (DC) will be encrypted.
        /// </summary>
        public readonly bool? EncryptDCConnections;
        /// <summary>
        /// kdc server IP address for the active directory machine. This optional parameter is used only while creating kerberos volume.
        /// </summary>
        public readonly string? KdcIP;
        /// <summary>
        /// Specifies whether or not the LDAP traffic needs to be secured via TLS.
        /// </summary>
        public readonly bool? LdapOverTLS;
        /// <summary>
        /// LDAP Search scope options
        /// </summary>
        public readonly Outputs.LdapSearchScopeOptResponse? LdapSearchScope;
        /// <summary>
        /// Specifies whether or not the LDAP traffic needs to be signed.
        /// </summary>
        public readonly bool? LdapSigning;
        /// <summary>
        /// The Organizational Unit (OU) within the Windows Active Directory
        /// </summary>
        public readonly string? OrganizationalUnit;
        /// <summary>
        /// Plain text password of Active Directory domain administrator, value is masked in the response
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Comma separated list of IPv4 addresses of preferred servers for LDAP client. At most two comma separated IPv4 addresses can be passed.
        /// </summary>
        public readonly string? PreferredServersForLdapClient;
        /// <summary>
        /// Domain Users in the Active directory to be given SeSecurityPrivilege privilege (Needed for SMB Continuously available shares for SQL). A list of unique usernames without domain specifier
        /// </summary>
        public readonly ImmutableArray<string> SecurityOperators;
        /// <summary>
        /// When LDAP over SSL/TLS is enabled, the LDAP client is required to have base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes.
        /// </summary>
        public readonly string? ServerRootCACertificate;
        /// <summary>
        /// The Active Directory site the service will limit Domain Controller discovery to
        /// </summary>
        public readonly string? Site;
        /// <summary>
        /// NetBIOS name of the SMB server. This name will be registered as a computer account in the AD and used to mount volumes
        /// </summary>
        public readonly string? SmbServerName;
        /// <summary>
        /// Status of the Active Directory
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Any details in regards to the Status of the Active Directory
        /// </summary>
        public readonly string StatusDetails;
        /// <summary>
        /// A domain user account with permission to create machine accounts
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ActiveDirectoryResponse(
            string? activeDirectoryId,

            string? adName,

            ImmutableArray<string> administrators,

            bool? aesEncryption,

            bool? allowLocalNfsUsersWithLdap,

            ImmutableArray<string> backupOperators,

            string? dns,

            string? domain,

            bool? encryptDCConnections,

            string? kdcIP,

            bool? ldapOverTLS,

            Outputs.LdapSearchScopeOptResponse? ldapSearchScope,

            bool? ldapSigning,

            string? organizationalUnit,

            string? password,

            string? preferredServersForLdapClient,

            ImmutableArray<string> securityOperators,

            string? serverRootCACertificate,

            string? site,

            string? smbServerName,

            string status,

            string statusDetails,

            string? username)
        {
            ActiveDirectoryId = activeDirectoryId;
            AdName = adName;
            Administrators = administrators;
            AesEncryption = aesEncryption;
            AllowLocalNfsUsersWithLdap = allowLocalNfsUsersWithLdap;
            BackupOperators = backupOperators;
            Dns = dns;
            Domain = domain;
            EncryptDCConnections = encryptDCConnections;
            KdcIP = kdcIP;
            LdapOverTLS = ldapOverTLS;
            LdapSearchScope = ldapSearchScope;
            LdapSigning = ldapSigning;
            OrganizationalUnit = organizationalUnit;
            Password = password;
            PreferredServersForLdapClient = preferredServersForLdapClient;
            SecurityOperators = securityOperators;
            ServerRootCACertificate = serverRootCACertificate;
            Site = site;
            SmbServerName = smbServerName;
            Status = status;
            StatusDetails = statusDetails;
            Username = username;
        }
    }
}
