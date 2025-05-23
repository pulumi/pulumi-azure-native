// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Outputs
{

    /// <summary>
    /// Network and credentials configuration currently applied to terminal server.
    /// </summary>
    [OutputType]
    public sealed class TerminalServerConfigurationResponse
    {
        /// <summary>
        /// ARM Resource ID used for the NetworkDevice.
        /// </summary>
        public readonly string NetworkDeviceId;
        /// <summary>
        /// Password for the terminal server connection.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// IPv4 Address Prefix.
        /// </summary>
        public readonly string PrimaryIpv4Prefix;
        /// <summary>
        /// IPv6 Address Prefix.
        /// </summary>
        public readonly string? PrimaryIpv6Prefix;
        /// <summary>
        /// Secondary IPv4 Address Prefix.
        /// </summary>
        public readonly string SecondaryIpv4Prefix;
        /// <summary>
        /// Secondary IPv6 Address Prefix.
        /// </summary>
        public readonly string? SecondaryIpv6Prefix;
        /// <summary>
        /// Serial Number of Terminal server.
        /// </summary>
        public readonly string? SerialNumber;
        /// <summary>
        /// Username for the terminal server connection.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private TerminalServerConfigurationResponse(
            string networkDeviceId,

            string password,

            string primaryIpv4Prefix,

            string? primaryIpv6Prefix,

            string secondaryIpv4Prefix,

            string? secondaryIpv6Prefix,

            string? serialNumber,

            string username)
        {
            NetworkDeviceId = networkDeviceId;
            Password = password;
            PrimaryIpv4Prefix = primaryIpv4Prefix;
            PrimaryIpv6Prefix = primaryIpv6Prefix;
            SecondaryIpv4Prefix = secondaryIpv4Prefix;
            SecondaryIpv6Prefix = secondaryIpv6Prefix;
            SerialNumber = serialNumber;
            Username = username;
        }
    }
}
