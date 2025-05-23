// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// Specifications of pull based data sources
    /// </summary>
    public sealed class DataSourcesSpecDataImportsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Definition of Event Hub configuration.
        /// </summary>
        [Input("eventHub")]
        public Input<Inputs.DataImportSourcesEventHubArgs>? EventHub { get; set; }

        public DataSourcesSpecDataImportsArgs()
        {
        }
        public static new DataSourcesSpecDataImportsArgs Empty => new DataSourcesSpecDataImportsArgs();
    }
}
