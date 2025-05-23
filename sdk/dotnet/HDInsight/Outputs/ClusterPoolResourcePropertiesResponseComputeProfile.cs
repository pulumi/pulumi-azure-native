// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// CLuster pool compute profile.
    /// </summary>
    [OutputType]
    public sealed class ClusterPoolResourcePropertiesResponseComputeProfile
    {
        /// <summary>
        /// The list of Availability zones to use for AKS VMSS nodes.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// The number of virtual machines.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The virtual machine SKU.
        /// </summary>
        public readonly string VmSize;

        [OutputConstructor]
        private ClusterPoolResourcePropertiesResponseComputeProfile(
            ImmutableArray<string> availabilityZones,

            int count,

            string vmSize)
        {
            AvailabilityZones = availabilityZones;
            Count = count;
            VmSize = vmSize;
        }
    }
}
