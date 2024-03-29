// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview.Inputs
{

    /// <summary>
    /// The license profile of the provisioned cluster.
    /// </summary>
    public sealed class ProvisionedClusterLicenseProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether Azure Hybrid Benefit is opted in
        /// </summary>
        [Input("azureHybridBenefit")]
        public InputUnion<string, Pulumi.AzureNative.HybridContainerService.V20231115Preview.AzureHybridBenefit>? AzureHybridBenefit { get; set; }

        public ProvisionedClusterLicenseProfileArgs()
        {
            AzureHybridBenefit = "NotApplicable";
        }
        public static new ProvisionedClusterLicenseProfileArgs Empty => new ProvisionedClusterLicenseProfileArgs();
    }
}
