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
    /// A copy activity Oracle Service Cloud source.
    /// </summary>
    public sealed class OracleServiceCloudSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the additional columns to be added to source data. Type: array of objects(AdditionalColumns) (or Expression with resultType array of objects).
        /// </summary>
        [Input("additionalColumns")]
        public Input<object>? AdditionalColumns { get; set; }

        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        [Input("disableMetricsCollection")]
        public Input<object>? DisableMetricsCollection { get; set; }

        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        [Input("maxConcurrentConnections")]
        public Input<object>? MaxConcurrentConnections { get; set; }

        /// <summary>
        /// A query to retrieve data from source. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("query")]
        public Input<object>? Query { get; set; }

        /// <summary>
        /// Query timeout. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        [Input("queryTimeout")]
        public Input<object>? QueryTimeout { get; set; }

        /// <summary>
        /// Source retry count. Type: integer (or Expression with resultType integer).
        /// </summary>
        [Input("sourceRetryCount")]
        public Input<object>? SourceRetryCount { get; set; }

        /// <summary>
        /// Source retry wait. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        [Input("sourceRetryWait")]
        public Input<object>? SourceRetryWait { get; set; }

        /// <summary>
        /// Copy source type.
        /// Expected value is 'OracleServiceCloudSource'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public OracleServiceCloudSourceArgs()
        {
        }
        public static new OracleServiceCloudSourceArgs Empty => new OracleServiceCloudSourceArgs();
    }
}
