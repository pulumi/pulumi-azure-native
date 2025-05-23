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
    /// Office365 linked service.
    /// </summary>
    public sealed class Office365LinkedServiceArgs : global::Pulumi.ResourceArgs
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
        /// Azure tenant ID to which the Office 365 account belongs. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("office365TenantId", required: true)]
        public Input<object> Office365TenantId { get; set; } = null!;

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
        /// The service principal credential type for authentication.'ServicePrincipalKey' for key/secret, 'ServicePrincipalCert' for certificate. If not specified, 'ServicePrincipalKey' is in use. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalCredentialType")]
        public Input<object>? ServicePrincipalCredentialType { get; set; }

        /// <summary>
        /// Specify the base64 encoded certificate of your application registered in Azure Active Directory. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalEmbeddedCert")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalEmbeddedCert { get; set; }

        /// <summary>
        /// Specify the password of your certificate if your certificate has a password and you are using AadServicePrincipal authentication. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalEmbeddedCertPassword")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? ServicePrincipalEmbeddedCertPassword { get; set; }

        /// <summary>
        /// Specify the application's client ID. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalId", required: true)]
        public Input<object> ServicePrincipalId { get; set; } = null!;

        /// <summary>
        /// Specify the application's key.
        /// </summary>
        [Input("servicePrincipalKey", required: true)]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs> ServicePrincipalKey { get; set; } = null!;

        /// <summary>
        /// Specify the tenant information under which your Azure AD web application resides. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("servicePrincipalTenantId", required: true)]
        public Input<object> ServicePrincipalTenantId { get; set; } = null!;

        /// <summary>
        /// Type of linked service.
        /// Expected value is 'Office365'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Version of the linked service.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public Office365LinkedServiceArgs()
        {
        }
        public static new Office365LinkedServiceArgs Empty => new Office365LinkedServiceArgs();
    }
}
