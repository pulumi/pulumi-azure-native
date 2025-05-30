// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.Inputs
{

    /// <summary>
    /// Describes how data from an input is serialized or how data is serialized when written to an output in Avro format.
    /// </summary>
    public sealed class AvroSerializationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the type of serialization that the input or output uses. Required on PUT (CreateOrReplace) requests.
        /// Expected value is 'Avro'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AvroSerializationArgs()
        {
        }
        public static new AvroSerializationArgs Empty => new AvroSerializationArgs();
    }
}
