// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceLinker.V20221101Preview.Inputs
{

    /// <summary>
    /// The authentication info when authType is secret
    /// </summary>
    public sealed class SecretAuthInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication type.
        /// Expected value is 'secret'.
        /// </summary>
        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        /// <summary>
        /// Username or account name for secret auth.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password or key vault secret for secret auth.
        /// </summary>
        [Input("secretInfo")]
        public object? SecretInfo { get; set; }

        public SecretAuthInfoArgs()
        {
        }
        public static new SecretAuthInfoArgs Empty => new SecretAuthInfoArgs();
    }
}
