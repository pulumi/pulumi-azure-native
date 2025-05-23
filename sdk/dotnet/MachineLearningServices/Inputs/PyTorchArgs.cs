// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// PyTorch distribution configuration.
    /// </summary>
    public sealed class PyTorchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to determine the job distribution type.
        /// Expected value is 'PyTorch'.
        /// </summary>
        [Input("distributionType", required: true)]
        public Input<string> DistributionType { get; set; } = null!;

        /// <summary>
        /// Number of processes per node.
        /// </summary>
        [Input("processCountPerInstance")]
        public Input<int>? ProcessCountPerInstance { get; set; }

        public PyTorchArgs()
        {
        }
        public static new PyTorchArgs Empty => new PyTorchArgs();
    }
}
