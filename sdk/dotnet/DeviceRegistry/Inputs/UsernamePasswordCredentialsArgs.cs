// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceRegistry.Inputs
{

    /// <summary>
    /// The credentials for authentication mode UsernamePassword.
    /// </summary>
    public sealed class UsernamePasswordCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the secret containing the password.
        /// </summary>
        [Input("passwordSecretName", required: true)]
        public Input<string> PasswordSecretName { get; set; } = null!;

        /// <summary>
        /// The name of the secret containing the username.
        /// </summary>
        [Input("usernameSecretName", required: true)]
        public Input<string> UsernameSecretName { get; set; } = null!;

        public UsernamePasswordCredentialsArgs()
        {
        }
        public static new UsernamePasswordCredentialsArgs Empty => new UsernamePasswordCredentialsArgs();
    }
}
