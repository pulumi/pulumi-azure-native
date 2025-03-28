// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Outputs
{

    /// <summary>
    /// When enabling the operator, a set of AKS managed CRDs and controllers will be installed in the cluster. The operator automates the deployment of OSS models for inference and/or training purposes. It provides a set of preset models and enables distributed inference against them.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterAIToolchainOperatorProfileResponse
    {
        /// <summary>
        /// Indicates if AI toolchain operator  enabled or not.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private ManagedClusterAIToolchainOperatorProfileResponse(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
