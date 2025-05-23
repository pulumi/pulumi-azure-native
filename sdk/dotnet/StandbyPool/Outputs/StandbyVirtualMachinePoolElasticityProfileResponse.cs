// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StandbyPool.Outputs
{

    /// <summary>
    /// Details of the elasticity profile.
    /// </summary>
    [OutputType]
    public sealed class StandbyVirtualMachinePoolElasticityProfileResponse
    {
        /// <summary>
        /// Specifies the maximum number of virtual machines in the standby virtual machine pool.
        /// </summary>
        public readonly double MaxReadyCapacity;
        /// <summary>
        /// Specifies the desired minimum number of virtual machines in the standby virtual machine pool. MinReadyCapacity cannot exceed MaxReadyCapacity.
        /// </summary>
        public readonly double? MinReadyCapacity;

        [OutputConstructor]
        private StandbyVirtualMachinePoolElasticityProfileResponse(
            double maxReadyCapacity,

            double? minReadyCapacity)
        {
            MaxReadyCapacity = maxReadyCapacity;
            MinReadyCapacity = minReadyCapacity;
        }
    }
}
