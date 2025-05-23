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
    /// Cluster service configs.
    /// </summary>
    [OutputType]
    public sealed class ClusterServiceConfigsProfileResponse
    {
        /// <summary>
        /// List of service configs.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterServiceConfigResponse> Configs;
        /// <summary>
        /// Name of the service the configurations should apply to.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private ClusterServiceConfigsProfileResponse(
            ImmutableArray<Outputs.ClusterServiceConfigResponse> configs,

            string serviceName)
        {
            Configs = configs;
            ServiceName = serviceName;
        }
    }
}
