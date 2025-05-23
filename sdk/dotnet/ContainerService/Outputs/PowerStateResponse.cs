// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Describes the Power State of the cluster
    /// </summary>
    [OutputType]
    public sealed class PowerStateResponse
    {
        /// <summary>
        /// Tells whether the cluster is Running or Stopped
        /// </summary>
        public readonly string? Code;

        [OutputConstructor]
        private PowerStateResponse(string? code)
        {
            Code = code;
        }
    }
}
