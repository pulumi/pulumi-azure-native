// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Impact.Outputs
{

    /// <summary>
    /// Information about the impacted workload
    /// </summary>
    [OutputType]
    public sealed class WorkloadResponse
    {
        /// <summary>
        /// the scenario for the workload
        /// </summary>
        public readonly string? Context;
        /// <summary>
        /// Tool used to interact with Azure. SDK, AzPortal, etc.., Other
        /// </summary>
        public readonly string? Toolset;

        [OutputConstructor]
        private WorkloadResponse(
            string? context,

            string? toolset)
        {
            Context = context;
            Toolset = toolset;
        }
    }
}
