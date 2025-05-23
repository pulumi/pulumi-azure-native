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
    /// Azure Data Lake Store linked service.
    /// </summary>
    public sealed class AzureDataLakeStoreLinkedServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data Lake Store account name. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("accountName")]
        public Input<object>? AccountName { get; set; }

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
        /// Indicates the azure cloud type of the service principle auth. Allowed values are AzurePublic, AzureChina, AzureUsGovernment, AzureGermany. Default value is the data factory regions’ cloud type. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("azureCloudType")]
        public Input<object>? AzureCloudType { get; set; }

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
        /// Data Lake Store service URI. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("dataLakeStoreUri", required: true)]
        public Input<object> DataLakeStoreUri { get; set; } = null!;

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Data Lake Store account resource group name (if different from Data Factory account). Type: string (or Expression with resultType string).
        /// </summary>
        [Input("resourceGroupName")]
        public Input<object>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ID of the application used to authenticate against the Azure Data Lake Store account. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<object>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The Key of the application used to authenticate against the Azure Data Lake Store account.
        /// </summary>
        [Input("servicePrincipalKey")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalKey { get; set; }

        /// <summary>
        /// Data Lake Store account subscription ID (if different from Data Factory account). Type: string (or Expression with resultType string).
        /// </summary>
        [Input("subscriptionId")]
        public Input<object>? SubscriptionId { get; set; }

        /// <summary>
        /// The name or ID of the tenant to which the service principal belongs. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("tenant")]
        public Input<object>? Tenant { get; set; }

        /// <summary>
        /// Type of linked service.
        /// Expected value is 'AzureDataLakeStore'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AzureDataLakeStoreLinkedServiceArgs()
        {
        }
        public static new AzureDataLakeStoreLinkedServiceArgs Empty => new AzureDataLakeStoreLinkedServiceArgs();
    }
}
