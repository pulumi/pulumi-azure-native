// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// Cluster pool network profile.
    /// </summary>
    [OutputType]
    public sealed class ClusterPoolResourcePropertiesResponseNetworkProfile
    {
        /// <summary>
        /// Cluster pool subnet resource id.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private ClusterPoolResourcePropertiesResponseNetworkProfile(string subnetId)
        {
            SubnetId = subnetId;
        }
    }
}
