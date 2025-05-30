// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// The type of disk sku. For example, Standard_LRS, Standard_ZRS, Premium_LRS, Premium_ZRS.
    /// </summary>
    [OutputType]
    public sealed class DiskSkuResponse
    {
        /// <summary>
        /// Defines the disk sku name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private DiskSkuResponse(string? name)
        {
            Name = name;
        }
    }
}
