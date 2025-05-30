// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class CollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The msi client id used to collect logging to blob storage. If it's null,backend will pick a registered endpoint identity to auth.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Enable or disable data collection.
        /// </summary>
        [Input("dataCollectionMode")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.DataCollectionMode>? DataCollectionMode { get; set; }

        /// <summary>
        /// The data asset arm resource id. Client side will ensure data asset is pointing to the blob storage, and backend will collect data to the blob storage.
        /// </summary>
        [Input("dataId")]
        public Input<string>? DataId { get; set; }

        /// <summary>
        /// The sampling rate for collection. Sampling rate 1.0 means we collect 100% of data by default.
        /// </summary>
        [Input("samplingRate")]
        public Input<double>? SamplingRate { get; set; }

        public CollectionArgs()
        {
            DataCollectionMode = "Disabled";
            SamplingRate = 1;
        }
        public static new CollectionArgs Empty => new CollectionArgs();
    }
}
