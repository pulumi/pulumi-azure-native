// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Inputs
{

    /// <summary>
    /// DiskBackedMessageBufferSettings properties
    /// </summary>
    public sealed class DiskBackedMessageBufferSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use the specified persistent volume claim template to mount a "generic ephemeral volume" for the message buffer. See &lt;https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/#generic-ephemeral-volumes&gt; for details.
        /// </summary>
        [Input("ephemeralVolumeClaimSpec")]
        public Input<Inputs.VolumeClaimSpecArgs>? EphemeralVolumeClaimSpec { get; set; }

        /// <summary>
        /// The max size of the message buffer on disk. If a PVC template is specified using one of ephemeralVolumeClaimSpec or persistentVolumeClaimSpec, then this size is used as the request and limit sizes of that template. If neither ephemeralVolumeClaimSpec nor persistentVolumeClaimSpec are specified, then an emptyDir volume is mounted with this size as its limit. See &lt;https://kubernetes.io/docs/concepts/storage/volumes/#emptydir&gt; for details.
        /// </summary>
        [Input("maxSize", required: true)]
        public Input<string> MaxSize { get; set; } = null!;

        /// <summary>
        /// Use the specified persistent volume claim template to mount a persistent volume for the message buffer.
        /// </summary>
        [Input("persistentVolumeClaimSpec")]
        public Input<Inputs.VolumeClaimSpecArgs>? PersistentVolumeClaimSpec { get; set; }

        public DiskBackedMessageBufferSettingsArgs()
        {
        }
        public static new DiskBackedMessageBufferSettingsArgs Empty => new DiskBackedMessageBufferSettingsArgs();
    }
}
