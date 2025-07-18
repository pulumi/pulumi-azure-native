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
    /// QuickBooks server linked service. This linked service has supported version property. The Version 1.0 is scheduled for deprecation while your pipeline will continue to run after EOL but without any bug fix or new features.
    /// </summary>
    public sealed class QuickBooksLinkedServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access token for OAuth 2.0 authentication.
        /// </summary>
        [Input("accessToken")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? AccessToken { get; set; }

        /// <summary>
        /// The access token secret is deprecated for OAuth 1.0 authentication. Only used for version 1.0.
        /// </summary>
        [Input("accessTokenSecret")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? AccessTokenSecret { get; set; }

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
        /// The company ID of the QuickBooks company to authorize.
        /// </summary>
        [Input("companyId")]
        public Input<object>? CompanyId { get; set; }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// Properties used to connect to QuickBooks. It is mutually exclusive with any other properties in the linked service. Type: object.
        /// </summary>
        [Input("connectionProperties")]
        public Input<object>? ConnectionProperties { get; set; }

        /// <summary>
        /// The consumer key for OAuth 2.0 authentication.
        /// </summary>
        [Input("consumerKey")]
        public Input<object>? ConsumerKey { get; set; }

        /// <summary>
        /// The consumer secret for OAuth 2.0 authentication.
        /// </summary>
        [Input("consumerSecret")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ConsumerSecret { get; set; }

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

        /// <summary>
        /// The endpoint of the QuickBooks server. (i.e. quickbooks.api.intuit.com)
        /// </summary>
        [Input("endpoint")]
        public Input<object>? Endpoint { get; set; }

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
        /// The refresh token for OAuth 2.0 authentication.
        /// </summary>
        [Input("refreshToken")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? RefreshToken { get; set; }

        /// <summary>
        /// Type of linked service.
        /// Expected value is 'QuickBooks'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Specifies whether the data source endpoints are encrypted using HTTPS. The default value is true. Only used for version 1.0.
        /// </summary>
        [Input("useEncryptedEndpoints")]
        public Input<object>? UseEncryptedEndpoints { get; set; }

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public QuickBooksLinkedServiceArgs()
        {
        }
        public static new QuickBooksLinkedServiceArgs Empty => new QuickBooksLinkedServiceArgs();
    }
}
