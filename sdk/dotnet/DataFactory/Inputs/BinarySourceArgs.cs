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
    /// A copy activity Binary source.
    /// </summary>
    public sealed class BinarySourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, disable data store metrics collection. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        [Input("disableMetricsCollection")]
        public Input<object>? DisableMetricsCollection { get; set; }

        /// <summary>
        /// Binary format settings.
        /// </summary>
        [Input("formatSettings")]
        public Input<Inputs.BinaryReadSettingsArgs>? FormatSettings { get; set; }

        /// <summary>
        /// The maximum concurrent connection count for the source data store. Type: integer (or Expression with resultType integer).
        /// </summary>
        [Input("maxConcurrentConnections")]
        public Input<object>? MaxConcurrentConnections { get; set; }

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
        /// Binary store settings.
        /// </summary>
        [Input("storeSettings")]
        public object? StoreSettings { get; set; }

        /// <summary>
        /// Copy source type.
        /// Expected value is 'BinarySource'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BinarySourceArgs()
        {
        }
        public static new BinarySourceArgs Empty => new BinarySourceArgs();
    }
}
