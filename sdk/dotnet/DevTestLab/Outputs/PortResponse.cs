// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Outputs
{

    /// <summary>
    /// Properties of a network port.
    /// </summary>
    [OutputType]
    public sealed class PortResponse
    {
        /// <summary>
        /// Backend port of the target virtual machine.
        /// </summary>
        public readonly int? BackendPort;
        /// <summary>
        /// Protocol type of the port.
        /// </summary>
        public readonly string? TransportProtocol;

        [OutputConstructor]
        private PortResponse(
            int? backendPort,

            string? transportProtocol)
        {
            BackendPort = backendPort;
            TransportProtocol = transportProtocol;
        }
    }
}
