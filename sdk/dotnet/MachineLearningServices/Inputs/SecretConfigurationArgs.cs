// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Secret Configuration definition.
    /// </summary>
    public sealed class SecretConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Secret Uri.
        /// Sample Uri : https://myvault.vault.azure.net/secrets/mysecretname/secretversion
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        /// <summary>
        /// Name of secret in workspace key vault.
        /// </summary>
        [Input("workspaceSecretName")]
        public Input<string>? WorkspaceSecretName { get; set; }

        public SecretConfigurationArgs()
        {
        }
        public static new SecretConfigurationArgs Empty => new SecretConfigurationArgs();
    }
}
