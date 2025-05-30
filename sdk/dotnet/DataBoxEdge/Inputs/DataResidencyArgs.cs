// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Inputs
{

    /// <summary>
    /// Wraps data-residency related information for edge-resource and this should be used with ARM layer.
    /// </summary>
    public sealed class DataResidencyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DataResidencyType enum
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.DataBoxEdge.DataResidencyType>? Type { get; set; }

        public DataResidencyArgs()
        {
        }
        public static new DataResidencyArgs Empty => new DataResidencyArgs();
    }
}
