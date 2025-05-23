// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsDataProcessor.Inputs
{

    /// <summary>
    /// Extended location is an extension of Azure locations. They provide a way to use their Azure ARC enabled Kubernetes clusters as target locations for deploying Azure services instances.
    /// </summary>
    public sealed class ExtendedLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the extended location.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of the extended location.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ExtendedLocationArgs()
        {
        }
        public static new ExtendedLocationArgs Empty => new ExtendedLocationArgs();
    }
}
