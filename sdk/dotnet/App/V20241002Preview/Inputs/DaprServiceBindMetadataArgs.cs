// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20241002Preview.Inputs
{

    /// <summary>
    /// Dapr component metadata.
    /// </summary>
    public sealed class DaprServiceBindMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service bind metadata property name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Service bind metadata property value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DaprServiceBindMetadataArgs()
        {
        }
        public static new DaprServiceBindMetadataArgs Empty => new DaprServiceBindMetadataArgs();
    }
}
