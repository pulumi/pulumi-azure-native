// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.Inputs
{

    /// <summary>
    /// Credentials for a user on a lab VM.
    /// </summary>
    public sealed class CredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password for the user. This is required for the TemplateVM createOption.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The username to use when signing in to lab VMs.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public CredentialsArgs()
        {
        }
        public static new CredentialsArgs Empty => new CredentialsArgs();
    }
}
