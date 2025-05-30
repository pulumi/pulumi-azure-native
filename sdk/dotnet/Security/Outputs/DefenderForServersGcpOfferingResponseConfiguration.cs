// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// configuration for Vulnerability Assessment autoprovisioning
    /// </summary>
    [OutputType]
    public sealed class DefenderForServersGcpOfferingResponseConfiguration
    {
        /// <summary>
        /// The Vulnerability Assessment solution to be provisioned. Can be either 'TVM' or 'Qualys'
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DefenderForServersGcpOfferingResponseConfiguration(string? type)
        {
            Type = type;
        }
    }
}
