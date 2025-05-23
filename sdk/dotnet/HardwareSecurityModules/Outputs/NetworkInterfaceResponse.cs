// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HardwareSecurityModules.Outputs
{

    /// <summary>
    /// The network interface definition.
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceResponse
    {
        /// <summary>
        /// Private Ip address of the interface
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// The Azure resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
        /// </summary>
        public readonly string ResourceId;

        [OutputConstructor]
        private NetworkInterfaceResponse(
            string? privateIpAddress,

            string resourceId)
        {
            PrivateIpAddress = privateIpAddress;
            ResourceId = resourceId;
        }
    }
}
