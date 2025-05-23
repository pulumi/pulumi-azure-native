// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HardwareSecurityModules.Inputs
{

    /// <summary>
    /// Cloud Hsm Cluster SKU information
    /// </summary>
    public sealed class CloudHsmClusterSkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sku capacity
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Sku family of the Cloud HSM Cluster
        /// </summary>
        [Input("family", required: true)]
        public InputUnion<string, Pulumi.AzureNative.HardwareSecurityModules.CloudHsmClusterSkuFamily> Family { get; set; } = null!;

        /// <summary>
        /// Sku name of the Cloud HSM Cluster
        /// </summary>
        [Input("name", required: true)]
        public Input<Pulumi.AzureNative.HardwareSecurityModules.CloudHsmClusterSkuName> Name { get; set; } = null!;

        public CloudHsmClusterSkuArgs()
        {
        }
        public static new CloudHsmClusterSkuArgs Empty => new CloudHsmClusterSkuArgs();
    }
}
