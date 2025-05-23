// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// Extended location of the resource.
    /// </summary>
    public sealed class RegisteredClusterNodesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The BIOS ID.
        /// </summary>
        [Input("biosId")]
        public Input<string>? BiosId { get; set; }

        /// <summary>
        /// The cluster node name.
        /// </summary>
        [Input("clusterNodeFqdn")]
        public Input<string>? ClusterNodeFqdn { get; set; }

        /// <summary>
        /// A value indicating whether this represents virtual entity hosting all the shared disks.
        /// </summary>
        [Input("isSharedDiskVirtualNode")]
        public Input<bool>? IsSharedDiskVirtualNode { get; set; }

        /// <summary>
        /// The machine ID.
        /// </summary>
        [Input("machineId")]
        public Input<string>? MachineId { get; set; }

        public RegisteredClusterNodesArgs()
        {
        }
        public static new RegisteredClusterNodesArgs Empty => new RegisteredClusterNodesArgs();
    }
}
