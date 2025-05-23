// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.Outputs
{

    /// <summary>
    /// The Azure SKU of the machines in the pool.
    /// </summary>
    [OutputType]
    public sealed class DevOpsAzureSkuResponse
    {
        /// <summary>
        /// The Azure SKU name of the machines in the pool.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private DevOpsAzureSkuResponse(string name)
        {
            Name = name;
        }
    }
}
