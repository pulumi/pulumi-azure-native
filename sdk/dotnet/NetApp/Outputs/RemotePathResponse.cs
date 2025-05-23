// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.Outputs
{

    /// <summary>
    /// The full path to a volume that is to be migrated into ANF. Required for Migration volumes
    /// </summary>
    [OutputType]
    public sealed class RemotePathResponse
    {
        /// <summary>
        /// The Path to a ONTAP Host
        /// </summary>
        public readonly string ExternalHostName;
        /// <summary>
        /// The name of a server on the ONTAP Host
        /// </summary>
        public readonly string ServerName;
        /// <summary>
        /// The name of a volume on the server
        /// </summary>
        public readonly string VolumeName;

        [OutputConstructor]
        private RemotePathResponse(
            string externalHostName,

            string serverName,

            string volumeName)
        {
            ExternalHostName = externalHostName;
            ServerName = serverName;
            VolumeName = volumeName;
        }
    }
}
