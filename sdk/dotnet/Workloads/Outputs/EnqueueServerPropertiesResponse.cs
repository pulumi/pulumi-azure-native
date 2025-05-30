// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Defines the SAP Enqueue Server properties.
    /// </summary>
    [OutputType]
    public sealed class EnqueueServerPropertiesResponse
    {
        /// <summary>
        /// Defines the health of SAP Instances.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// Enqueue Server SAP Hostname.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// Enqueue Server SAP IP Address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Enqueue Server Port.
        /// </summary>
        public readonly double Port;

        [OutputConstructor]
        private EnqueueServerPropertiesResponse(
            string health,

            string hostname,

            string ipAddress,

            double port)
        {
            Health = health;
            Hostname = hostname;
            IpAddress = ipAddress;
            Port = port;
        }
    }
}
