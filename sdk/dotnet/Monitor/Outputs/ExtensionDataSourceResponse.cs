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
    /// Definition of which data will be collected from a separate VM extension that integrates with the Azure Monitor Agent.
    /// Collected from either Windows and Linux machines, depending on which extension is defined.
    /// </summary>
    [OutputType]
    public sealed class ExtensionDataSourceResponse
    {
        /// <summary>
        /// The name of the VM extension.
        /// </summary>
        public readonly string ExtensionName;
        /// <summary>
        /// The extension settings. The format is specific for particular extension.
        /// </summary>
        public readonly object? ExtensionSettings;
        /// <summary>
        /// The list of data sources this extension needs data from.
        /// </summary>
        public readonly ImmutableArray<string> InputDataSources;
        /// <summary>
        /// A friendly name for the data source. 
        /// This name should be unique across all data sources (regardless of type) within the data collection rule.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// List of streams that this data source will be sent to.
        /// A stream indicates what schema will be used for this data and usually what table in Log Analytics the data will be sent to.
        /// </summary>
        public readonly ImmutableArray<string> Streams;

        [OutputConstructor]
        private ExtensionDataSourceResponse(
            string extensionName,

            object? extensionSettings,

            ImmutableArray<string> inputDataSources,

            string? name,

            ImmutableArray<string> streams)
        {
            ExtensionName = extensionName;
            ExtensionSettings = extensionSettings;
            InputDataSources = inputDataSources;
            Name = name;
            Streams = streams;
        }
    }
}
