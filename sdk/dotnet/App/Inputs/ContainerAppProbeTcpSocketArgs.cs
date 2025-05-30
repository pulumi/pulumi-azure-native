// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported.
    /// </summary>
    public sealed class ContainerAppProbeTcpSocketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional: Host name to connect to, defaults to the pod IP.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public ContainerAppProbeTcpSocketArgs()
        {
        }
        public static new ContainerAppProbeTcpSocketArgs Empty => new ContainerAppProbeTcpSocketArgs();
    }
}
