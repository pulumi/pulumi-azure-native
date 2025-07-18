// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Outputs
{

    /// <summary>
    /// Compute information of a flexible server.
    /// </summary>
    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// Name by which is known a given compute size assigned to a flexible server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Tier of the compute assigned to a flexible server.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private SkuResponse(
            string name,

            string tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
