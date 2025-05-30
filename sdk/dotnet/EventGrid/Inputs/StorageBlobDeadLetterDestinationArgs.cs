// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Inputs
{

    /// <summary>
    /// Information about the storage blob based dead letter destination.
    /// </summary>
    public sealed class StorageBlobDeadLetterDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Storage blob container that is the destination of the deadletter events
        /// </summary>
        [Input("blobContainerName")]
        public Input<string>? BlobContainerName { get; set; }

        /// <summary>
        /// Type of the endpoint for the dead letter destination
        /// Expected value is 'StorageBlob'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The Azure Resource ID of the storage account that is the destination of the deadletter events
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public StorageBlobDeadLetterDestinationArgs()
        {
        }
        public static new StorageBlobDeadLetterDestinationArgs Empty => new StorageBlobDeadLetterDestinationArgs();
    }
}
