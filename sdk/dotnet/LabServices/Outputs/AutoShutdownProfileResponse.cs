// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.Outputs
{

    /// <summary>
    /// Profile for how to handle shutting down virtual machines.
    /// </summary>
    [OutputType]
    public sealed class AutoShutdownProfileResponse
    {
        /// <summary>
        /// The amount of time a VM will stay running after a user disconnects if this behavior is enabled.
        /// </summary>
        public readonly string? DisconnectDelay;
        /// <summary>
        /// The amount of time a VM will idle before it is shutdown if this behavior is enabled.
        /// </summary>
        public readonly string? IdleDelay;
        /// <summary>
        /// The amount of time a VM will stay running before it is shutdown if no connection is made and this behavior is enabled.
        /// </summary>
        public readonly string? NoConnectDelay;
        /// <summary>
        /// Whether shutdown on disconnect is enabled
        /// </summary>
        public readonly string? ShutdownOnDisconnect;
        /// <summary>
        /// Whether a VM will get shutdown when it has idled for a period of time.
        /// </summary>
        public readonly string? ShutdownOnIdle;
        /// <summary>
        /// Whether a VM will get shutdown when it hasn't been connected to after a period of time.
        /// </summary>
        public readonly string? ShutdownWhenNotConnected;

        [OutputConstructor]
        private AutoShutdownProfileResponse(
            string? disconnectDelay,

            string? idleDelay,

            string? noConnectDelay,

            string? shutdownOnDisconnect,

            string? shutdownOnIdle,

            string? shutdownWhenNotConnected)
        {
            DisconnectDelay = disconnectDelay;
            IdleDelay = idleDelay;
            NoConnectDelay = noConnectDelay;
            ShutdownOnDisconnect = shutdownOnDisconnect;
            ShutdownOnIdle = shutdownOnIdle;
            ShutdownWhenNotConnected = shutdownWhenNotConnected;
        }
    }
}
