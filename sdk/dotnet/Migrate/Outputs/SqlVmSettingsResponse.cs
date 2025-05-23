// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// SQL VM assessment settings.
    /// </summary>
    [OutputType]
    public sealed class SqlVmSettingsResponse
    {
        /// <summary>
        /// Gets or sets the Azure VM families (calling instance series to keep it
        /// consistent with other targets).
        /// </summary>
        public readonly ImmutableArray<string> InstanceSeries;

        [OutputConstructor]
        private SqlVmSettingsResponse(ImmutableArray<string> instanceSeries)
        {
            InstanceSeries = instanceSeries;
        }
    }
}
