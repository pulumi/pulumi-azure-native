// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// Unsecured endpoint describes an endpoint that the pipeline can connect to over clear transport (no encryption in transit).
    /// </summary>
    [OutputType]
    public sealed class UnsecuredEndpointResponse
    {
        /// <summary>
        /// Credentials to be presented to the endpoint.
        /// </summary>
        public readonly Outputs.UsernamePasswordCredentialsResponse Credentials;
        /// <summary>
        /// Describes the tunnel through which Video Analyzer can connect to the endpoint URL. This is an optional property, typically used when the endpoint is behind a firewall.
        /// </summary>
        public readonly Outputs.SecureIotDeviceRemoteTunnelResponse? Tunnel;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.VideoAnalyzer.UnsecuredEndpoint'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The endpoint URL for Video Analyzer to connect to.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private UnsecuredEndpointResponse(
            Outputs.UsernamePasswordCredentialsResponse credentials,

            Outputs.SecureIotDeviceRemoteTunnelResponse? tunnel,

            string type,

            string url)
        {
            Credentials = credentials;
            Tunnel = tunnel;
            Type = type;
            Url = url;
        }
    }
}
