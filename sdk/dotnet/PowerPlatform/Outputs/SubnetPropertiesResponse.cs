// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerPlatform.Outputs
{

    /// <summary>
    /// Properties of a subnet.
    /// </summary>
    [OutputType]
    public sealed class SubnetPropertiesResponse
    {
        /// <summary>
        /// Subnet name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SubnetPropertiesResponse(string? name)
        {
            Name = name;
        }
    }
}
