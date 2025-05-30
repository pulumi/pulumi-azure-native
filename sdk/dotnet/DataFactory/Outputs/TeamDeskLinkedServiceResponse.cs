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
    /// Linked service for TeamDesk.
    /// </summary>
    [OutputType]
    public sealed class TeamDeskLinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// The api token for the TeamDesk source.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? ApiToken;
        /// <summary>
        /// The authentication type to use.
        /// </summary>
        public readonly string AuthenticationType;
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
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// The password of the TeamDesk source.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'TeamDesk'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The url to connect TeamDesk source. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Url;
        /// <summary>
        /// The username of the TeamDesk source. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? UserName;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private TeamDeskLinkedServiceResponse(
            ImmutableArray<object> annotations,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? apiToken,

            string authenticationType,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            string? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            string type,

            object url,

            object? userName,

            string? version)
        {
            Annotations = annotations;
            ApiToken = apiToken;
            AuthenticationType = authenticationType;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            Password = password;
            Type = type;
            Url = url;
            UserName = userName;
            Version = version;
        }
    }
}
