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
    /// Office365 linked service.
    /// </summary>
    [OutputType]
    public sealed class Office365LinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
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
        /// Azure tenant ID to which the Office 365 account belongs. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Office365TenantId;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The service principal credential type for authentication.'ServicePrincipalKey' for key/secret, 'ServicePrincipalCert' for certificate. If not specified, 'ServicePrincipalKey' is in use. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ServicePrincipalCredentialType;
        /// <summary>
        /// Specify the base64 encoded certificate of your application registered in Azure Active Directory. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ServicePrincipalEmbeddedCert;
        /// <summary>
        /// Specify the password of your certificate if your certificate has a password and you are using AadServicePrincipal authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ServicePrincipalEmbeddedCertPassword;
        /// <summary>
        /// Specify the application's client ID. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object ServicePrincipalId;
        /// <summary>
        /// Specify the application's key.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse> ServicePrincipalKey;
        /// <summary>
        /// Specify the tenant information under which your Azure AD web application resides. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object ServicePrincipalTenantId;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'Office365'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private Office365LinkedServiceResponse(
            ImmutableArray<object> annotations,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            string? encryptedCredential,

            object office365TenantId,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            object? servicePrincipalCredentialType,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? servicePrincipalEmbeddedCert,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? servicePrincipalEmbeddedCertPassword,

            object servicePrincipalId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse> servicePrincipalKey,

            object servicePrincipalTenantId,

            string type,

            string? version)
        {
            Annotations = annotations;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Office365TenantId = office365TenantId;
            Parameters = parameters;
            ServicePrincipalCredentialType = servicePrincipalCredentialType;
            ServicePrincipalEmbeddedCert = servicePrincipalEmbeddedCert;
            ServicePrincipalEmbeddedCertPassword = servicePrincipalEmbeddedCertPassword;
            ServicePrincipalId = servicePrincipalId;
            ServicePrincipalKey = servicePrincipalKey;
            ServicePrincipalTenantId = servicePrincipalTenantId;
            Type = type;
            Version = version;
        }
    }
}
