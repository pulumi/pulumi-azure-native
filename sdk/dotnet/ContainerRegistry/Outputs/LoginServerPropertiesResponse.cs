// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The login server properties of the connected registry.
    /// </summary>
    [OutputType]
    public sealed class LoginServerPropertiesResponse
    {
        /// <summary>
        /// The host of the connected registry. Can be FQDN or IP.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The TLS properties of the connected registry login server.
        /// </summary>
        public readonly Outputs.TlsPropertiesResponse Tls;

        [OutputConstructor]
        private LoginServerPropertiesResponse(
            string host,

            Outputs.TlsPropertiesResponse tls)
        {
            Host = host;
            Tls = tls;
        }
    }
}
