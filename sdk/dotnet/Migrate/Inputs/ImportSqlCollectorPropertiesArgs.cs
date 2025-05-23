// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Import SQL Collector properties class.
    /// </summary>
    public sealed class ImportSqlCollectorPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sql db extended details.
        /// </summary>
        [Input("discoverySiteId")]
        public Input<string>? DiscoverySiteId { get; set; }

        public ImportSqlCollectorPropertiesArgs()
        {
        }
        public static new ImportSqlCollectorPropertiesArgs Empty => new ImportSqlCollectorPropertiesArgs();
    }
}
