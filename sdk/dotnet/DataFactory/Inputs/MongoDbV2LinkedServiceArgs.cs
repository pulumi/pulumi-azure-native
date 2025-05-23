// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Linked service for MongoDB data source.
    /// </summary>
    public sealed class MongoDbV2LinkedServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<object>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public InputList<object> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<object>());
            set => _annotations = value;
        }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// The MongoDB connection string. Type: string, SecureString or AzureKeyVaultSecretReference. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        [Input("connectionString", required: true)]
        public Input<object> ConnectionString { get; set; } = null!;

        /// <summary>
        /// The name of the MongoDB database that you want to access. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("database", required: true)]
        public Input<object> Database { get; set; } = null!;

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("parameters")]
        private InputMap<Inputs.ParameterSpecificationArgs>? _parameters;

        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public InputMap<Inputs.ParameterSpecificationArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterSpecificationArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Type of linked service.
        /// Expected value is 'MongoDbV2'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MongoDbV2LinkedServiceArgs()
        {
        }
        public static new MongoDbV2LinkedServiceArgs Empty => new MongoDbV2LinkedServiceArgs();
    }
}
