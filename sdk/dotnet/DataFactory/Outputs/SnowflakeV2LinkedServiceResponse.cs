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
    /// Snowflake linked service.
    /// </summary>
    [OutputType]
    public sealed class SnowflakeV2LinkedServiceResponse
    {
        /// <summary>
        /// The account identifier of your Snowflake account, e.g. xy12345.east-us-2.azure
        /// </summary>
        public readonly object AccountIdentifier;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The type used for authentication. Type: string.
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// The client ID of the application registered in Azure Active Directory for AADServicePrincipal authentication.
        /// </summary>
        public readonly object? ClientId;
        /// <summary>
        /// The Azure key vault secret reference of client secret for AADServicePrincipal authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ClientSecret;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponse? ConnectVia;
        /// <summary>
        /// The name of the Snowflake database.
        /// </summary>
        public readonly object Database;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string.
        /// </summary>
        public readonly string? EncryptedCredential;
        /// <summary>
        /// The host name of the Snowflake account. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Host;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The Azure key vault secret reference of password in connection string.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// The Azure key vault secret reference of privateKey for KeyPair auth.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? PrivateKey;
        /// <summary>
        /// The Azure key vault secret reference of private key password for KeyPair auth with encrypted private key.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? PrivateKeyPassphrase;
        /// <summary>
        /// The default access control role to use in the Snowflake session. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Role;
        /// <summary>
        /// Schema name for connection. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Schema;
        /// <summary>
        /// The scope of the application registered in Azure Active Directory for AADServicePrincipal authentication.
        /// </summary>
        public readonly object? Scope;
        /// <summary>
        /// The tenant ID of the application registered in Azure Active Directory for AADServicePrincipal authentication.
        /// </summary>
        public readonly object? TenantId;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'SnowflakeV2'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The name of the Snowflake user.
        /// </summary>
        public readonly object? User;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// The name of the Snowflake warehouse.
        /// </summary>
        public readonly object Warehouse;

        [OutputConstructor]
        private SnowflakeV2LinkedServiceResponse(
            object accountIdentifier,

            ImmutableArray<object> annotations,

            string? authenticationType,

            object? clientId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? clientSecret,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            object database,

            string? description,

            string? encryptedCredential,

            object? host,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? privateKey,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? privateKeyPassphrase,

            object? role,

            object? schema,

            object? scope,

            object? tenantId,

            string type,

            object? user,

            string? version,

            object warehouse)
        {
            AccountIdentifier = accountIdentifier;
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ClientId = clientId;
            ClientSecret = clientSecret;
            ConnectVia = connectVia;
            Database = database;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Host = host;
            Parameters = parameters;
            Password = password;
            PrivateKey = privateKey;
            PrivateKeyPassphrase = privateKeyPassphrase;
            Role = role;
            Schema = schema;
            Scope = scope;
            TenantId = tenantId;
            Type = type;
            User = user;
            Version = version;
            Warehouse = warehouse;
        }
    }
}
