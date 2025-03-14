// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20220201Preview.Outputs
{

    /// <summary>
    /// Network manager security group item.
    /// </summary>
    [OutputType]
    public sealed class NetworkManagerSecurityGroupItemResponse
    {
        /// <summary>
        /// Network manager group Id.
        /// </summary>
        public readonly string NetworkGroupId;

        [OutputConstructor]
        private NetworkManagerSecurityGroupItemResponse(string networkGroupId)
        {
            NetworkGroupId = networkGroupId;
        }
    }
}
