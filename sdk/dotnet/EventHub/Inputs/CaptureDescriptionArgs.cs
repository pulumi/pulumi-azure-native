// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.Inputs
{

    /// <summary>
    /// Properties to configure capture description for eventhub
    /// </summary>
    public sealed class CaptureDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Properties of Destination where capture will be stored. (Storage Account, Blob Names)
        /// </summary>
        [Input("destination")]
        public Input<Inputs.DestinationArgs>? Destination { get; set; }

        /// <summary>
        /// A value that indicates whether capture description is enabled. 
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in New API Version
        /// </summary>
        [Input("encoding")]
        public Input<Pulumi.AzureNative.EventHub.EncodingCaptureDescription>? Encoding { get; set; }

        /// <summary>
        /// The time window allows you to set the frequency with which the capture to Azure Blobs will happen, value should between 60 to 900 seconds
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// The size window defines the amount of data built up in your Event Hub before an capture operation, value should be between 10485760 to 524288000 bytes
        /// </summary>
        [Input("sizeLimitInBytes")]
        public Input<int>? SizeLimitInBytes { get; set; }

        /// <summary>
        /// A value that indicates whether to Skip Empty Archives
        /// </summary>
        [Input("skipEmptyArchives")]
        public Input<bool>? SkipEmptyArchives { get; set; }

        public CaptureDescriptionArgs()
        {
        }
        public static new CaptureDescriptionArgs Empty => new CaptureDescriptionArgs();
    }
}
