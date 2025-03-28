// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Inputs
{

    /// <summary>
    /// Auto upgrade profile for a managed cluster.
    /// </summary>
    public sealed class ManagedClusterAutoUpgradeProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.
        /// </summary>
        [Input("nodeOSUpgradeChannel")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20241002Preview.NodeOSUpgradeChannel>? NodeOSUpgradeChannel { get; set; }

        /// <summary>
        /// For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel).
        /// </summary>
        [Input("upgradeChannel")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20241002Preview.UpgradeChannel>? UpgradeChannel { get; set; }

        public ManagedClusterAutoUpgradeProfileArgs()
        {
        }
        public static new ManagedClusterAutoUpgradeProfileArgs Empty => new ManagedClusterAutoUpgradeProfileArgs();
    }
}
