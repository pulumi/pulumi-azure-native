// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover.Outputs
{

    /// <summary>
    /// The properties of NFS share endpoint.
    /// </summary>
    [OutputType]
    public sealed class NfsMountEndpointPropertiesResponse
    {
        /// <summary>
        /// A description for the Endpoint.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Endpoint resource type.
        /// Expected value is 'NfsMount'.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// The directory being exported from the server.
        /// </summary>
        public readonly string Export;
        /// <summary>
        /// The host name or IP address of the server exporting the file system.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The NFS protocol version.
        /// </summary>
        public readonly string? NfsVersion;
        /// <summary>
        /// The provisioning state of this resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private NfsMountEndpointPropertiesResponse(
            string? description,

            string endpointType,

            string export,

            string host,

            string? nfsVersion,

            string provisioningState)
        {
            Description = description;
            EndpointType = endpointType;
            Export = export;
            Host = host;
            NfsVersion = nfsVersion;
            ProvisioningState = provisioningState;
        }
    }
}
