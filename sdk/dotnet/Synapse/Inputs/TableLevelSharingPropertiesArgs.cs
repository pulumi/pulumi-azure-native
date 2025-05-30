// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Inputs
{

    /// <summary>
    /// Tables that will be included and excluded in the follower database
    /// </summary>
    public sealed class TableLevelSharingPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("externalTablesToExclude")]
        private InputList<string>? _externalTablesToExclude;

        /// <summary>
        /// List of external tables exclude from the follower database
        /// </summary>
        public InputList<string> ExternalTablesToExclude
        {
            get => _externalTablesToExclude ?? (_externalTablesToExclude = new InputList<string>());
            set => _externalTablesToExclude = value;
        }

        [Input("externalTablesToInclude")]
        private InputList<string>? _externalTablesToInclude;

        /// <summary>
        /// List of external tables to include in the follower database
        /// </summary>
        public InputList<string> ExternalTablesToInclude
        {
            get => _externalTablesToInclude ?? (_externalTablesToInclude = new InputList<string>());
            set => _externalTablesToInclude = value;
        }

        [Input("materializedViewsToExclude")]
        private InputList<string>? _materializedViewsToExclude;

        /// <summary>
        /// List of materialized views exclude from the follower database
        /// </summary>
        public InputList<string> MaterializedViewsToExclude
        {
            get => _materializedViewsToExclude ?? (_materializedViewsToExclude = new InputList<string>());
            set => _materializedViewsToExclude = value;
        }

        [Input("materializedViewsToInclude")]
        private InputList<string>? _materializedViewsToInclude;

        /// <summary>
        /// List of materialized views to include in the follower database
        /// </summary>
        public InputList<string> MaterializedViewsToInclude
        {
            get => _materializedViewsToInclude ?? (_materializedViewsToInclude = new InputList<string>());
            set => _materializedViewsToInclude = value;
        }

        [Input("tablesToExclude")]
        private InputList<string>? _tablesToExclude;

        /// <summary>
        /// List of tables to exclude from the follower database
        /// </summary>
        public InputList<string> TablesToExclude
        {
            get => _tablesToExclude ?? (_tablesToExclude = new InputList<string>());
            set => _tablesToExclude = value;
        }

        [Input("tablesToInclude")]
        private InputList<string>? _tablesToInclude;

        /// <summary>
        /// List of tables to include in the follower database
        /// </summary>
        public InputList<string> TablesToInclude
        {
            get => _tablesToInclude ?? (_tablesToInclude = new InputList<string>());
            set => _tablesToInclude = value;
        }

        public TableLevelSharingPropertiesArgs()
        {
        }
        public static new TableLevelSharingPropertiesArgs Empty => new TableLevelSharingPropertiesArgs();
    }
}
