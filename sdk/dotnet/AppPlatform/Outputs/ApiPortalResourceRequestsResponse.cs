// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// Resource requests of the API portal
    /// </summary>
    [OutputType]
    public sealed class ApiPortalResourceRequestsResponse
    {
        /// <summary>
        /// Cpu allocated to each API portal instance
        /// </summary>
        public readonly string Cpu;
        /// <summary>
        /// Memory allocated to each API portal instance
        /// </summary>
        public readonly string Memory;

        [OutputConstructor]
        private ApiPortalResourceRequestsResponse(
            string cpu,

            string memory)
        {
            Cpu = cpu;
            Memory = memory;
        }
    }
}
