// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceLinker.Inputs
{

    /// <summary>
    /// The authentication info when authType is secret
    /// </summary>
    public sealed class SecretAuthInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Indicates how to configure authentication. If optInAllAuth, service linker configures authentication such as enabling identity on source resource and granting RBAC roles. If optOutAllAuth, opt out authentication setup. Default is optInAllAuth.
        /// </summary>
        [Input("authMode")]
        public InputUnion<string, Pulumi.AzureNative.ServiceLinker.AuthMode>? AuthMode { get; set; }

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
