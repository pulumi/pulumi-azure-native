// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of Csv
    /// </summary>
    public sealed class CsvArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The delimiter used for separating items in the CSV file being imported.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("headerList")]
        private InputList<string>? _headerList;

        /// <summary>
        /// List of the headers used to specify a common header for all source CSV files being imported. If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.
        /// </summary>
        public InputList<string> HeaderList
        {
            get => _headerList ?? (_headerList = new InputList<string>());
            set => _headerList = value;
        }

        public CsvArgs()
        {
        }
        public static new CsvArgs Empty => new CsvArgs();
    }
}
