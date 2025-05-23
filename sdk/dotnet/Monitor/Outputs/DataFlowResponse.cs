// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Definition of which streams are sent to which destinations.
    /// </summary>
    [OutputType]
    public sealed class DataFlowResponse
    {
        /// <summary>
        /// The builtIn transform to transform stream data
        /// </summary>
        public readonly string? BuiltInTransform;
        /// <summary>
        /// List of destinations for this data flow.
        /// </summary>
        public readonly ImmutableArray<string> Destinations;
        /// <summary>
        /// The output stream of the transform. Only required if the transform changes data to a different stream.
        /// </summary>
        public readonly string? OutputStream;
        /// <summary>
        /// List of streams for this data flow.
        /// </summary>
        public readonly ImmutableArray<string> Streams;
        /// <summary>
        /// The KQL query to transform stream data.
        /// </summary>
        public readonly string? TransformKql;

        [OutputConstructor]
        private DataFlowResponse(
            string? builtInTransform,

            ImmutableArray<string> destinations,

            string? outputStream,

            ImmutableArray<string> streams,

            string? transformKql)
        {
            BuiltInTransform = builtInTransform;
            Destinations = destinations;
            OutputStream = outputStream;
            Streams = streams;
            TransformKql = transformKql;
        }
    }
}
