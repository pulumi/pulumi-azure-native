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
    /// Dataset reference type.
    /// </summary>
    public sealed class DatasetReferenceArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// Arguments for dataset.
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        /// <summary>
        /// Reference dataset name.
        /// </summary>
        [Input("referenceName", required: true)]
        public Input<string> ReferenceName { get; set; } = null!;

        /// <summary>
        /// Dataset reference type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatasetReferenceArgs()
        {
        }
        public static new DatasetReferenceArgs Empty => new DatasetReferenceArgs();
    }
}
