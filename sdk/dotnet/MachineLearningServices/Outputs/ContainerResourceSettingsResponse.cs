// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class ContainerResourceSettingsResponse
    {
        /// <summary>
        /// Number of vCPUs request/limit for container. More info:
        /// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
        /// </summary>
        public readonly string? Cpu;
        /// <summary>
        /// Number of Nvidia GPU cards request/limit for container. More info:
        /// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
        /// </summary>
        public readonly string? Gpu;
        /// <summary>
        /// Memory size request/limit for container. More info:
        /// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
        /// </summary>
        public readonly string? Memory;

        [OutputConstructor]
        private ContainerResourceSettingsResponse(
            string? cpu,

            string? gpu,

            string? memory)
        {
            Cpu = cpu;
            Gpu = gpu;
            Memory = memory;
        }
    }
}
