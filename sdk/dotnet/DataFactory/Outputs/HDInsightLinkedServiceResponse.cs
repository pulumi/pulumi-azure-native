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
    /// HDInsight linked service.
    /// </summary>
    [OutputType]
    public sealed class HDInsightLinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<object> Annotations;
        /// <summary>
        /// HDInsight cluster URI. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object ClusterUri;
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
        /// Specify the FileSystem if the main storage for the HDInsight is ADLS Gen2. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? FileSystem;
        /// <summary>
        /// A reference to the Azure SQL linked service that points to the HCatalog database.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? HcatalogLinkedServiceName;
        /// <summary>
        /// Specify if the HDInsight is created with ESP (Enterprise Security Package). Type: Boolean.
        /// </summary>
        public readonly object? IsEspEnabled;
        /// <summary>
        /// The Azure Storage linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? LinkedServiceName;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// HDInsight cluster password.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'HDInsight'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// HDInsight cluster user name. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? UserName;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private HDInsightLinkedServiceResponse(
            ImmutableArray<object> annotations,

            object clusterUri,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            string? encryptedCredential,

            object? fileSystem,

            Outputs.LinkedServiceReferenceResponse? hcatalogLinkedServiceName,

            object? isEspEnabled,

            Outputs.LinkedServiceReferenceResponse? linkedServiceName,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            string type,

            object? userName,

            string? version)
        {
            Annotations = annotations;
            ClusterUri = clusterUri;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            FileSystem = fileSystem;
            HcatalogLinkedServiceName = hcatalogLinkedServiceName;
            IsEspEnabled = isEspEnabled;
            LinkedServiceName = linkedServiceName;
            Parameters = parameters;
            Password = password;
            Type = type;
            UserName = userName;
            Version = version;
        }
    }
}
