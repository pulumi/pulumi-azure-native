// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Binding for a web application.
    /// </summary>
    public sealed class BindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// WebApplication certificate.
        /// </summary>
        [Input("cert")]
        public Input<Inputs.CertArgs>? Cert { get; set; }

        /// <summary>
        /// Gets or sets the binding host name.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// Gets or sets the IP Address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Gets or sets the application port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// WebApplication port mapping.
        /// </summary>
        [Input("portMapping")]
        public Input<Inputs.PortMappingArgs>? PortMapping { get; set; }

        /// <summary>
        /// Gets or sets the protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public BindingArgs()
        {
        }
        public static new BindingArgs Empty => new BindingArgs();
    }
}
