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
    /// The registry node that generated the event. Put differently, while the actor initiates the event, the source generates it.
    /// </summary>
    [OutputType]
    public sealed class SourceResponse
    {
        /// <summary>
        /// The IP or hostname and the port of the registry node that generated the event. Generally, this will be resolved by os.Hostname() along with the running port.
        /// </summary>
        public readonly string? Addr;
        /// <summary>
        /// The running instance of an application. Changes after each restart.
        /// </summary>
        public readonly string? InstanceID;

        [OutputConstructor]
        private SourceResponse(
            string? addr,

            string? instanceID)
        {
            Addr = addr;
            InstanceID = instanceID;
        }
    }
}
