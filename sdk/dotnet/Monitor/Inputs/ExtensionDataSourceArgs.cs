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
    /// Definition of which data will be collected from a separate VM extension that integrates with the Azure Monitor Agent.
    /// Collected from either Windows and Linux machines, depending on which extension is defined.
    /// </summary>
    public sealed class ExtensionDataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VM extension.
        /// </summary>
        [Input("extensionName", required: true)]
        public Input<string> ExtensionName { get; set; } = null!;

        /// <summary>
        /// The extension settings. The format is specific for particular extension.
        /// </summary>
        [Input("extensionSettings")]
        public Input<object>? ExtensionSettings { get; set; }

        [Input("inputDataSources")]
        private InputList<string>? _inputDataSources;

        /// <summary>
        /// The list of data sources this extension needs data from.
        /// </summary>
        public InputList<string> InputDataSources
        {
            get => _inputDataSources ?? (_inputDataSources = new InputList<string>());
            set => _inputDataSources = value;
        }

        /// <summary>
        /// A friendly name for the data source. 
        /// This name should be unique across all data sources (regardless of type) within the data collection rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("streams")]
        private InputList<Union<string, Pulumi.AzureNative.Monitor.KnownExtensionDataSourceStreams>>? _streams;

        /// <summary>
        /// List of streams that this data source will be sent to.
        /// A stream indicates what schema will be used for this data and usually what table in Log Analytics the data will be sent to.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Monitor.KnownExtensionDataSourceStreams>> Streams
        {
            get => _streams ?? (_streams = new InputList<Union<string, Pulumi.AzureNative.Monitor.KnownExtensionDataSourceStreams>>());
            set => _streams = value;
        }

        public ExtensionDataSourceArgs()
        {
        }
        public static new ExtensionDataSourceArgs Empty => new ExtensionDataSourceArgs();
    }
}
