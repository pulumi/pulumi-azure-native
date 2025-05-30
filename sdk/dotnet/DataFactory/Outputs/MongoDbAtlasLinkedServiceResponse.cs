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
    /// Linked service for MongoDB Atlas data source.
    /// </summary>
    [OutputType]
    public sealed class MongoDbAtlasLinkedServiceResponse
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
        /// The MongoDB Atlas connection string. Type: string, SecureString or AzureKeyVaultSecretReference. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        public readonly object ConnectionString;
        /// <summary>
        /// The name of the MongoDB Atlas database that you want to access. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Database;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The driver version that you want to choose. Allowed value are v1 and v2. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? DriverVersion;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// Type of linked service.
        /// Expected value is 'MongoDbAtlas'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Version of the linked service.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private MongoDbAtlasLinkedServiceResponse(
            ImmutableArray<object> annotations,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            object connectionString,

            object database,

            string? description,

            object? driverVersion,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            string type,

            string? version)
        {
            Annotations = annotations;
            ConnectVia = connectVia;
            ConnectionString = connectionString;
            Database = database;
            Description = description;
            DriverVersion = driverVersion;
            Parameters = parameters;
            Type = type;
            Version = version;
        }
    }
}
