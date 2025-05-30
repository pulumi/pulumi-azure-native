// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
    /// </summary>
    public sealed class RecommendedMachineConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the resource range.
        /// </summary>
        [Input("memory")]
        public Input<Inputs.ResourceRangeArgs>? Memory { get; set; }

        /// <summary>
        /// Describes the resource range.
        /// </summary>
        [Input("vCPUs")]
        public Input<Inputs.ResourceRangeArgs>? VCPUs { get; set; }

        public RecommendedMachineConfigurationArgs()
        {
        }
        public static new RecommendedMachineConfigurationArgs Empty => new RecommendedMachineConfigurationArgs();
    }
}
