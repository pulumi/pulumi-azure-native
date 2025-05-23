// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Auto upgrade profile for a managed cluster.
    /// </summary>
    public sealed class ManagedClusterAutoUpgradeProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Manner in which the OS on your nodes is updated. The default is NodeImage.
        /// </summary>
        [Input("nodeOSUpgradeChannel")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.NodeOSUpgradeChannel>? NodeOSUpgradeChannel { get; set; }

        /// <summary>
        /// For more information see [setting the AKS cluster auto-upgrade channel](https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel).
        /// </summary>
        [Input("upgradeChannel")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.UpgradeChannel>? UpgradeChannel { get; set; }

        public ManagedClusterAutoUpgradeProfileArgs()
        {
        }
        public static new ManagedClusterAutoUpgradeProfileArgs Empty => new ManagedClusterAutoUpgradeProfileArgs();
    }
}
