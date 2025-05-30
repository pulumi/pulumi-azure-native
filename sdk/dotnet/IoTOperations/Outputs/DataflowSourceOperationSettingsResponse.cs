// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// Dataflow Source Operation properties
    /// </summary>
    [OutputType]
    public sealed class DataflowSourceOperationSettingsResponse
    {
        /// <summary>
        /// Reference to the resource in Azure Device Registry where the data in the endpoint originates from.
        /// </summary>
        public readonly string? AssetRef;
        /// <summary>
        /// List of source locations. Can be Broker or Kafka topics. Supports wildcards # and +.
        /// </summary>
        public readonly ImmutableArray<string> DataSources;
        /// <summary>
        /// Reference to the Dataflow Endpoint resource. Can only be of Broker and Kafka type.
        /// </summary>
        public readonly string EndpointRef;
        /// <summary>
        /// Schema CR reference. Data will be deserialized according to the schema, and dropped if it doesn't match.
        /// </summary>
        public readonly string? SchemaRef;
        /// <summary>
        /// Content is a JSON Schema. Allowed: JSON Schema/draft-7.
        /// </summary>
        public readonly string? SerializationFormat;

        [OutputConstructor]
        private DataflowSourceOperationSettingsResponse(
            string? assetRef,

            ImmutableArray<string> dataSources,

            string endpointRef,

            string? schemaRef,

            string? serializationFormat)
        {
            AssetRef = assetRef;
            DataSources = dataSources;
            EndpointRef = endpointRef;
            SchemaRef = schemaRef;
            SerializationFormat = serializationFormat;
        }
    }
}
