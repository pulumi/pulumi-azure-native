// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManufacturingPlatform.Outputs
{

    /// <summary>
    /// The properties related to Azure Data Explorer (Adx) Resource
    /// </summary>
    [OutputType]
    public sealed class AdxProfileResponse
    {
        /// <summary>
        /// Data Ingestion Uri of Adx Resource
        /// </summary>
        public readonly string DataIngestionUri;
        /// <summary>
        /// Resource Id of Adx Resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Uri of Adx Resource
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private AdxProfileResponse(
            string dataIngestionUri,

            string id,

            string uri)
        {
            DataIngestionUri = dataIngestionUri;
            Id = id;
            Uri = uri;
        }
    }
}
