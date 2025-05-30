// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm.Inputs
{

    /// <summary>
    /// Username / Password Credentials to connect to guest.
    /// </summary>
    public sealed class GuestCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the password to connect with the guest.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Gets or sets username to connect with the guest.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GuestCredentialArgs()
        {
        }
        public static new GuestCredentialArgs Empty => new GuestCredentialArgs();
    }
}
