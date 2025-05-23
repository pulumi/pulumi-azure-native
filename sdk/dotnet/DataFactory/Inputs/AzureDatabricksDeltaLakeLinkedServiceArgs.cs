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
    /// Azure Databricks Delta Lake linked service.
    /// </summary>
    public sealed class AzureDatabricksDeltaLakeLinkedServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access token for databricks REST API. Refer to https://docs.azuredatabricks.net/api/latest/authentication.html. Type: string, SecureString or AzureKeyVaultSecretReference.
        /// </summary>
        [Input("accessToken")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? AccessToken { get; set; }

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
        /// The id of an existing interactive cluster that will be used for all runs of this job. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("clusterId")]
        public Input<object>? ClusterId { get; set; }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// The credential reference containing authentication information.
        /// </summary>
        [Input("credential")]
        public Input<Inputs.CredentialReferenceArgs>? Credential { get; set; }

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// &lt;REGION&gt;.azuredatabricks.net, domain name of your Databricks deployment. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("domain", required: true)]
        public Input<object> Domain { get; set; } = null!;

        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string.
        /// </summary>
        [Input("encryptedCredential")]
        public Input<string>? EncryptedCredential { get; set; }

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
        /// Expected value is 'AzureDatabricksDeltaLake'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Workspace resource id for databricks REST API. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("workspaceResourceId")]
        public Input<object>? WorkspaceResourceId { get; set; }

        public AzureDatabricksDeltaLakeLinkedServiceArgs()
        {
        }
        public static new AzureDatabricksDeltaLakeLinkedServiceArgs Empty => new AzureDatabricksDeltaLakeLinkedServiceArgs();
    }
}
