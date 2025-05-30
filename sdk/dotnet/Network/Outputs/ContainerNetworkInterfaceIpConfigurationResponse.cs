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
    /// The ip configuration for a container network interface.
    /// </summary>
    [OutputType]
    public sealed class ContainerNetworkInterfaceIpConfigurationResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the resource. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the container network interface IP configuration resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Sub Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ContainerNetworkInterfaceIpConfigurationResponse(
            string etag,

            string? name,

            string provisioningState,

            string type)
        {
            Etag = etag;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
