// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// The complex type of the extended location.
    /// </summary>
    public sealed class ExtendedLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the extended location.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the extended location.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.ExtendedLocationTypes>? Type { get; set; }

        public ExtendedLocationArgs()
        {
        }
        public static new ExtendedLocationArgs Empty => new ExtendedLocationArgs();
    }
}
