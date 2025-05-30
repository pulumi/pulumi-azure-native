// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// The data stored in Parquet format.
    /// </summary>
    public sealed class ParquetFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deserializer. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("deserializer")]
        public Input<object>? Deserializer { get; set; }

        /// <summary>
        /// Serializer. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("serializer")]
        public Input<object>? Serializer { get; set; }

        /// <summary>
        /// Type of dataset storage format.
        /// Expected value is 'ParquetFormat'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ParquetFormatArgs()
        {
        }
        public static new ParquetFormatArgs Empty => new ParquetFormatArgs();
    }
}
