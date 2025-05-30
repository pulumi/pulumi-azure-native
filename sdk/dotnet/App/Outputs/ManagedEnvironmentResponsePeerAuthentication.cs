// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Peer authentication settings for the Managed Environment
    /// </summary>
    [OutputType]
    public sealed class ManagedEnvironmentResponsePeerAuthentication
    {
        /// <summary>
        /// Mutual TLS authentication settings for the Managed Environment
        /// </summary>
        public readonly Outputs.MtlsResponse? Mtls;

        [OutputConstructor]
        private ManagedEnvironmentResponsePeerAuthentication(Outputs.MtlsResponse? mtls)
        {
            Mtls = mtls;
        }
    }
}
