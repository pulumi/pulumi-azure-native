// *** WARNING: this file was generated by pulumi. ***
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
        /// A reference to secret containing the password.
        /// </summary>
        [Input("passwordReference", required: true)]
        public Input<string> PasswordReference { get; set; } = null!;

        /// <summary>
        /// A reference to secret containing the username.
        /// </summary>
        [Input("usernameReference", required: true)]
        public Input<string> UsernameReference { get; set; } = null!;

        public UsernamePasswordCredentialsArgs()
        {
        }
        public static new UsernamePasswordCredentialsArgs Empty => new UsernamePasswordCredentialsArgs();
    }
}
