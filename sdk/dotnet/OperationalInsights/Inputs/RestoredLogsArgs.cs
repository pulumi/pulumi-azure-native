// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OperationalInsights.Inputs
{

    /// <summary>
    /// Restore parameters.
    /// </summary>
    public sealed class RestoredLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp to end the restore by (UTC).
        /// </summary>
        [Input("endRestoreTime")]
        public Input<string>? EndRestoreTime { get; set; }

        /// <summary>
        /// The table to restore data from.
        /// </summary>
        [Input("sourceTable")]
        public Input<string>? SourceTable { get; set; }

        /// <summary>
        /// The timestamp to start the restore from (UTC).
        /// </summary>
        [Input("startRestoreTime")]
        public Input<string>? StartRestoreTime { get; set; }

        public RestoredLogsArgs()
        {
        }
        public static new RestoredLogsArgs Empty => new RestoredLogsArgs();
    }
}
