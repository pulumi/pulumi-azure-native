// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Data specific to GlusterFS.
    /// </summary>
    [OutputType]
    public sealed class GlusterFsSectionResponse
    {
        /// <summary>
        /// The server address of one of the servers that hosts the GlusterFS. Can be either the IP address or server name.
        /// </summary>
        public readonly string ServerAddress;
        /// <summary>
        /// The name of the created GlusterFS volume.
        /// </summary>
        public readonly string VolumeName;

        [OutputConstructor]
        private GlusterFsSectionResponse(
            string serverAddress,

            string volumeName)
        {
            ServerAddress = serverAddress;
            VolumeName = volumeName;
        }
    }
}
