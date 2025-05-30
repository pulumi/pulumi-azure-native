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
    /// Azure Databricks Delta Lake export command settings.
    /// </summary>
    public sealed class AzureDatabricksDeltaLakeExportCommandArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the date format for the csv in Azure Databricks Delta Lake Copy. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("dateFormat")]
        public Input<object>? DateFormat { get; set; }

        /// <summary>
        /// Specify the timestamp format for the csv in Azure Databricks Delta Lake Copy. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("timestampFormat")]
        public Input<object>? TimestampFormat { get; set; }

        /// <summary>
        /// The export setting type.
        /// Expected value is 'AzureDatabricksDeltaLakeExportCommand'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AzureDatabricksDeltaLakeExportCommandArgs()
        {
        }
        public static new AzureDatabricksDeltaLakeExportCommandArgs Empty => new AzureDatabricksDeltaLakeExportCommandArgs();
    }
}
