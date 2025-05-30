// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover.Inputs
{

    /// <summary>
    /// The properties of SMB share endpoint.
    /// </summary>
    public sealed class SmbMountEndpointPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Key Vault secret URIs which store the required credentials to access the SMB share.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.AzureKeyVaultSmbCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// A description for the Endpoint.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Endpoint resource type.
        /// Expected value is 'SmbMount'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The host name or IP address of the server exporting the file system.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The name of the SMB share being exported from the server.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        public SmbMountEndpointPropertiesArgs()
        {
        }
        public static new SmbMountEndpointPropertiesArgs Empty => new SmbMountEndpointPropertiesArgs();
    }
}
