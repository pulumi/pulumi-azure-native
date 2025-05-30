// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Inputs
{

    /// <summary>
    /// The key authorization type integration runtime.
    /// </summary>
    public sealed class LinkedIntegrationRuntimeKeyAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization type for integration runtime sharing.
        /// Expected value is 'Key'.
        /// </summary>
        [Input("authorizationType", required: true)]
        public Input<string> AuthorizationType { get; set; } = null!;

        /// <summary>
        /// The key used for authorization.
        /// </summary>
        [Input("key", required: true)]
        public Input<Inputs.SecureStringArgs> Key { get; set; } = null!;

        public LinkedIntegrationRuntimeKeyAuthorizationArgs()
        {
        }
        public static new LinkedIntegrationRuntimeKeyAuthorizationArgs Empty => new LinkedIntegrationRuntimeKeyAuthorizationArgs();
    }
}
