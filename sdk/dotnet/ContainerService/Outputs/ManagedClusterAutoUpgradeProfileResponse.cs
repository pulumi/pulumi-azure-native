// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Auto upgrade profile for a managed cluster.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterAutoUpgradeProfileResponse
    {
        /// <summary>
        /// Manner in which the OS on your nodes is updated. The default is NodeImage.
        /// </summary>
        public readonly string? NodeOSUpgradeChannel;
        /// <summary>
        /// For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel).
        /// </summary>
        public readonly string? UpgradeChannel;

        [OutputConstructor]
        private ManagedClusterAutoUpgradeProfileResponse(
            string? nodeOSUpgradeChannel,

            string? upgradeChannel)
        {
            NodeOSUpgradeChannel = nodeOSUpgradeChannel;
            UpgradeChannel = upgradeChannel;
        }
    }
}
