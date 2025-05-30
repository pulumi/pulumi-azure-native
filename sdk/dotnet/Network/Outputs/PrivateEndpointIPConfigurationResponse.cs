// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// An IP Configuration of the private endpoint.
    /// </summary>
    [OutputType]
    public sealed class PrivateEndpointIPConfigurationResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The ID of a group obtained from the remote resource that this private endpoint should connect to.
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// The member name of a group obtained from the remote resource that this private endpoint should connect to.
        /// </summary>
        public readonly string? MemberName;
        /// <summary>
        /// The name of the resource that is unique within a resource group.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A private ip address obtained from the private endpoint's subnet.
        /// </summary>
        public readonly string? PrivateIPAddress;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointIPConfigurationResponse(
            string etag,

            string? groupId,

            string? memberName,

            string? name,

            string? privateIPAddress,

            string type)
        {
            Etag = etag;
            GroupId = groupId;
            MemberName = memberName;
            Name = name;
            PrivateIPAddress = privateIPAddress;
            Type = type;
        }
    }
}
