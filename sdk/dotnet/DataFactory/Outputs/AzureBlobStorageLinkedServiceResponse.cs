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
    /// The azure blob storage linked service.
    /// </summary>
    [OutputType]
    public sealed class AzureBlobStorageLinkedServiceResponse
    {
        /// <summary>
        /// The Azure key vault secret reference of accountKey in connection string.
        /// </summary>
        public readonly Outputs.AzureKeyVaultSecretReferenceResponse? AccountKey;
        /// <summary>
        /// Specify the kind of your storage account. Allowed values are: Storage (general purpose v1), StorageV2 (general purpose v2), BlobStorage, or BlockBlobStorage. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? AccountKind;
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The type used for authentication. Type: string.
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// Indicates the azure cloud type of the service principle auth. Allowed values are AzurePublic, AzureChina, AzureUsGovernment, AzureGermany. Default value is the data factory regions’ cloud type. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? AzureCloudType;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponse? ConnectVia;
        /// <summary>
        /// The connection string. It is mutually exclusive with sasUri, serviceEndpoint property. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        public readonly object? ConnectionString;
        /// <summary>
        /// Container uri of the Azure Blob Storage resource only support for anonymous access. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ContainerUri;
        /// <summary>
        /// The credential reference containing authentication information.
        /// </summary>
        public readonly Outputs.CredentialReferenceResponse? Credential;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string.
        /// </summary>
        public readonly string? EncryptedCredential;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The Azure key vault secret reference of sasToken in sas uri.
        /// </summary>
        public readonly Outputs.AzureKeyVaultSecretReferenceResponse? SasToken;
        /// <summary>
        /// SAS URI of the Azure Blob Storage resource. It is mutually exclusive with connectionString, serviceEndpoint property. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        public readonly object? SasUri;
        /// <summary>
        /// Blob service endpoint of the Azure Blob Storage resource. It is mutually exclusive with connectionString, sasUri property.
        /// </summary>
        public readonly object? ServiceEndpoint;
        /// <summary>
        /// The ID of the service principal used to authenticate against Azure SQL Data Warehouse. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ServicePrincipalId;
        /// <summary>
        /// The key of the service principal used to authenticate against Azure SQL Data Warehouse.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ServicePrincipalKey;
        /// <summary>
        /// The name or ID of the tenant to which the service principal belongs. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? Tenant;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'AzureBlobStorage'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private AzureBlobStorageLinkedServiceResponse(
            Outputs.AzureKeyVaultSecretReferenceResponse? accountKey,

            object? accountKind,

            ImmutableArray<object> annotations,

            string? authenticationType,

            object? azureCloudType,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            object? connectionString,

            object? containerUri,

            Outputs.CredentialReferenceResponse? credential,

            string? description,

            string? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Outputs.AzureKeyVaultSecretReferenceResponse? sasToken,

            object? sasUri,

            object? serviceEndpoint,

            object? servicePrincipalId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? servicePrincipalKey,

            object? tenant,

            string type,

            string? version)
        {
            AccountKey = accountKey;
            AccountKind = accountKind;
            Annotations = annotations;
            AuthenticationType = authenticationType;
            AzureCloudType = azureCloudType;
            ConnectVia = connectVia;
            ConnectionString = connectionString;
            ContainerUri = containerUri;
            Credential = credential;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            SasToken = sasToken;
            SasUri = sasUri;
            ServiceEndpoint = serviceEndpoint;
            ServicePrincipalId = servicePrincipalId;
            ServicePrincipalKey = servicePrincipalKey;
            Tenant = tenant;
            Type = type;
            Version = version;
        }
    }
}
