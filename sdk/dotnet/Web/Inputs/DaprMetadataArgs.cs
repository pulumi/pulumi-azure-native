// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Container App Dapr component metadata.
    /// </summary>
    public sealed class DaprMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Metadata property name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the Container App secret from which to pull the metadata property value.
        /// </summary>
        [Input("secretRef")]
        public Input<string>? SecretRef { get; set; }

        /// <summary>
        /// Metadata property value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DaprMetadataArgs()
        {
        }
        public static new DaprMetadataArgs Empty => new DaprMetadataArgs();
    }
}
