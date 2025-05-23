// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// A linked service for an SSH File Transfer Protocol (SFTP) server. 
    /// </summary>
    [OutputType]
    public sealed class SftpServerLinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The authentication type to be used to connect to the FTP server.
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponse? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string.
        /// </summary>
        public readonly string? EncryptedCredential;
        /// <summary>
        /// The SFTP server host name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Host;
        /// <summary>
        /// The host key finger-print of the SFTP server. When SkipHostKeyValidation is false, HostKeyFingerprint should be specified. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? HostKeyFingerprint;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The password to decrypt the SSH private key if the SSH private key is encrypted.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? PassPhrase;
        /// <summary>
        /// Password to logon the SFTP server for Basic authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// The TCP port number that the SFTP server uses to listen for client connections. Default value is 22. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public readonly object? Port;
        /// <summary>
        /// Base64 encoded SSH private key content for SshPublicKey authentication. For on-premises copy with SshPublicKey authentication, either PrivateKeyPath or PrivateKeyContent should be specified. SSH private key should be OpenSSH format.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? PrivateKeyContent;
        /// <summary>
        /// The SSH private key file path for SshPublicKey authentication. Only valid for on-premises copy. For on-premises copy with SshPublicKey authentication, either PrivateKeyPath or PrivateKeyContent should be specified. SSH private key should be OpenSSH format. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? PrivateKeyPath;
        /// <summary>
        /// If true, skip the SSH host key validation. Default value is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? SkipHostKeyValidation;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'Sftp'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The username used to log on to the SFTP server. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? UserName;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private SftpServerLinkedServiceResponse(
            ImmutableArray<object> annotations,

            string? authenticationType,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            string? encryptedCredential,

            object host,

            object? hostKeyFingerprint,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? passPhrase,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            object? port,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? privateKeyContent,

            object? privateKeyPath,

            object? skipHostKeyValidation,

            string type,

            object? userName,

            string? version)
        {
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Host = host;
            HostKeyFingerprint = hostKeyFingerprint;
            Parameters = parameters;
            PassPhrase = passPhrase;
            Password = password;
            Port = port;
            PrivateKeyContent = privateKeyContent;
            PrivateKeyPath = privateKeyPath;
            SkipHostKeyValidation = skipHostKeyValidation;
            Type = type;
            UserName = userName;
            Version = version;
        }
    }
}
