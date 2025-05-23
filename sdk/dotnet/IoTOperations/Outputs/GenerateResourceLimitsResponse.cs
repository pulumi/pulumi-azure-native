// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// GenerateResourceLimits properties
    /// </summary>
    [OutputType]
    public sealed class GenerateResourceLimitsResponse
    {
        /// <summary>
        /// The toggle to enable/disable cpu resource limits.
        /// </summary>
        public readonly string? Cpu;

        [OutputConstructor]
        private GenerateResourceLimitsResponse(string? cpu)
        {
            Cpu = cpu;
        }
    }
}
