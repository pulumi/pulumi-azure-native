// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer.Inputs
{

    /// <summary>
    /// The schema object.
    /// </summary>
    public sealed class SchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection ID associated with this schema
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// Content of the schema
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The direction of the schema.
        /// </summary>
        [Input("direction")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaDirection>? Direction { get; set; }

        /// <summary>
        /// ID associated with this schema
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the schema
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Schema Type
        /// </summary>
        [Input("schemaType")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaType>? SchemaType { get; set; }

        /// <summary>
        /// Uri containing SAS token for the zipped schema
        /// </summary>
        [Input("schemaUri")]
        public Input<string>? SchemaUri { get; set; }

        /// <summary>
        /// Status of the schema
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.SchemaStatus>? Status { get; set; }

        public SchemaArgs()
        {
        }
        public static new SchemaArgs Empty => new SchemaArgs();
    }
}
