// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20241001Preview.Outputs
{

    /// <summary>
    /// The export dataset configuration.
    /// </summary>
    [OutputType]
    public sealed class ExportDatasetConfigurationResponse
    {
        /// <summary>
        /// This is on path to deprecation and will not be supported going forward.
        /// </summary>
        public readonly ImmutableArray<string> Columns;
        /// <summary>
        /// The data version for the selected for the export. If not provided then the export will default to latest data version.
        /// </summary>
        public readonly string? DataVersion;
        /// <summary>
        /// Filters associated with the data sets.
        /// </summary>
        public readonly ImmutableArray<Outputs.FilterItemsResponse> Filters;

        [OutputConstructor]
        private ExportDatasetConfigurationResponse(
            ImmutableArray<string> columns,

            string? dataVersion,

            ImmutableArray<Outputs.FilterItemsResponse> filters)
        {
            Columns = columns;
            DataVersion = dataVersion;
            Filters = filters;
        }
    }
}
