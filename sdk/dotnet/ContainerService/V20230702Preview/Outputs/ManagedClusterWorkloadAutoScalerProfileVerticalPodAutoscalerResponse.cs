// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20230702Preview.Outputs
{

    [OutputType]
    public sealed class ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscalerResponse
    {
        /// <summary>
        /// Whether to enable VPA. Default value is false.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscalerResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
