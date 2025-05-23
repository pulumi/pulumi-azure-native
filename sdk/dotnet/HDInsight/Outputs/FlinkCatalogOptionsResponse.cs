// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Flink cluster catalog options.
    /// </summary>
    [OutputType]
    public sealed class FlinkCatalogOptionsResponse
    {
        /// <summary>
        /// Hive Catalog Option for Flink cluster.
        /// </summary>
        public readonly Outputs.FlinkHiveCatalogOptionResponse? Hive;

        [OutputConstructor]
        private FlinkCatalogOptionsResponse(Outputs.FlinkHiveCatalogOptionResponse? hive)
        {
            Hive = hive;
        }
    }
}
