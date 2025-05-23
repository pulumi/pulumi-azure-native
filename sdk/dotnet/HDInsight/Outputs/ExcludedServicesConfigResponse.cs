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
    /// The configuration that services will be excluded when creating cluster.
    /// </summary>
    [OutputType]
    public sealed class ExcludedServicesConfigResponse
    {
        /// <summary>
        /// The config id of excluded services.
        /// </summary>
        public readonly string? ExcludedServicesConfigId;
        /// <summary>
        /// The list of excluded services.
        /// </summary>
        public readonly string? ExcludedServicesList;

        [OutputConstructor]
        private ExcludedServicesConfigResponse(
            string? excludedServicesConfigId,

            string? excludedServicesList)
        {
            ExcludedServicesConfigId = excludedServicesConfigId;
            ExcludedServicesList = excludedServicesList;
        }
    }
}
