// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AgriculturePlatform.Inputs
{

    /// <summary>
    /// The properties related to an AgriService data connector.
    /// </summary>
    public sealed class DataConnectorCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client Id associated with the provider, if type of credentials is OAuthClientCredentials.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Name of the key vault key.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Uri of the key vault
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        /// <summary>
        /// Version of the key vault key.
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        /// <summary>
        /// Type of credential.
        /// </summary>
        [Input("kind")]
        public InputUnion<string, Pulumi.AzureNative.AgriculturePlatform.AuthCredentialsKind>? Kind { get; set; }

        public DataConnectorCredentialsArgs()
        {
        }
        public static new DataConnectorCredentialsArgs Empty => new DataConnectorCredentialsArgs();
    }
}
