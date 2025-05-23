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
    /// ServiceNowV2 server linked service.
    /// </summary>
    [OutputType]
    public sealed class ServiceNowV2LinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The authentication type to use.
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// The client id for OAuth2 authentication.
        /// </summary>
        public readonly object? ClientId;
        /// <summary>
        /// The client secret for OAuth2 authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ClientSecret;
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
        /// The endpoint of the ServiceNowV2 server. (i.e. &lt;instance&gt;.service-now.com)
        /// </summary>
        public readonly object Endpoint;
        /// <summary>
        /// GrantType for OAuth2 authentication. Default value is password.
        /// </summary>
        public readonly object? GrantType;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The password corresponding to the user name for Basic and OAuth2 authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'ServiceNowV2'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user name used to connect to the ServiceNowV2 server for Basic and OAuth2 authentication.
        /// </summary>
        public readonly object? Username;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ServiceNowV2LinkedServiceResponse(
            ImmutableArray<object> annotations,

            string authenticationType,

            object? clientId,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? clientSecret,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            string? encryptedCredential,

            object endpoint,

            object? grantType,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            string type,

            object? username,

            string? version)
        {
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ClientId = clientId;
            ClientSecret = clientSecret;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Endpoint = endpoint;
            GrantType = grantType;
            Parameters = parameters;
            Password = password;
            Type = type;
            Username = username;
            Version = version;
        }
    }
}
