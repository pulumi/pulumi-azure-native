// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kusto.Outputs
{

    /// <summary>
    /// Tables that will be included and excluded in the follower database
    /// </summary>
    [OutputType]
    public sealed class TableLevelSharingPropertiesResponse
    {
        /// <summary>
        /// List of external tables to exclude from the follower database
        /// </summary>
        public readonly ImmutableArray<string> ExternalTablesToExclude;
        /// <summary>
        /// List of external tables to include in the follower database
        /// </summary>
        public readonly ImmutableArray<string> ExternalTablesToInclude;
        /// <summary>
        /// List of functions to exclude from the follower database
        /// </summary>
        public readonly ImmutableArray<string> FunctionsToExclude;
        /// <summary>
        /// List of functions to include in the follower database
        /// </summary>
        public readonly ImmutableArray<string> FunctionsToInclude;
        /// <summary>
        /// List of materialized views to exclude from the follower database
        /// </summary>
        public readonly ImmutableArray<string> MaterializedViewsToExclude;
        /// <summary>
        /// List of materialized views to include in the follower database
        /// </summary>
        public readonly ImmutableArray<string> MaterializedViewsToInclude;
        /// <summary>
        /// List of tables to exclude from the follower database
        /// </summary>
        public readonly ImmutableArray<string> TablesToExclude;
        /// <summary>
        /// List of tables to include in the follower database
        /// </summary>
        public readonly ImmutableArray<string> TablesToInclude;

        [OutputConstructor]
        private TableLevelSharingPropertiesResponse(
            ImmutableArray<string> externalTablesToExclude,

            ImmutableArray<string> externalTablesToInclude,

            ImmutableArray<string> functionsToExclude,

            ImmutableArray<string> functionsToInclude,

            ImmutableArray<string> materializedViewsToExclude,

            ImmutableArray<string> materializedViewsToInclude,

            ImmutableArray<string> tablesToExclude,

            ImmutableArray<string> tablesToInclude)
        {
            ExternalTablesToExclude = externalTablesToExclude;
            ExternalTablesToInclude = externalTablesToInclude;
            FunctionsToExclude = functionsToExclude;
            FunctionsToInclude = functionsToInclude;
            MaterializedViewsToExclude = materializedViewsToExclude;
            MaterializedViewsToInclude = materializedViewsToInclude;
            TablesToExclude = tablesToExclude;
            TablesToInclude = tablesToInclude;
        }
    }
}
