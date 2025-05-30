// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Inputs
{

    /// <summary>
    /// Unsecured endpoint describes an endpoint that the pipeline can connect to over clear transport (no encryption in transit).
    /// </summary>
    public sealed class UnsecuredEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Credentials to be presented to the endpoint.
        /// </summary>
        [Input("credentials", required: true)]
        public Input<Inputs.UsernamePasswordCredentialsArgs> Credentials { get; set; } = null!;

        /// <summary>
        /// Describes the tunnel through which Video Analyzer can connect to the endpoint URL. This is an optional property, typically used when the endpoint is behind a firewall.
        /// </summary>
        [Input("tunnel")]
        public Input<Inputs.SecureIotDeviceRemoteTunnelArgs>? Tunnel { get; set; }

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.VideoAnalyzer.UnsecuredEndpoint'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The endpoint URL for Video Analyzer to connect to.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public UnsecuredEndpointArgs()
        {
        }
        public static new UnsecuredEndpointArgs Empty => new UnsecuredEndpointArgs();
    }
}
