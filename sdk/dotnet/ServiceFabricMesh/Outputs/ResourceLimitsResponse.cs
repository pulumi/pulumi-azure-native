// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabricMesh.Outputs
{

    /// <summary>
    /// This type describes the resource limits for a given container. It describes the most amount of resources a container is allowed to use before being restarted.
    /// </summary>
    [OutputType]
    public sealed class ResourceLimitsResponse
    {
        /// <summary>
        /// CPU limits in cores. At present, only full cores are supported.
        /// </summary>
        public readonly double? Cpu;
        /// <summary>
        /// The memory limit in GB.
        /// </summary>
        public readonly double? MemoryInGB;

        [OutputConstructor]
        private ResourceLimitsResponse(
            double? cpu,

            double? memoryInGB)
        {
            Cpu = cpu;
            MemoryInGB = memoryInGB;
        }
    }
}
