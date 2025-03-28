// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.V20241101Preview.Inputs
{

    /// <summary>
    /// The object that contains a reference to a Container Group Profile and it's other related properties.
    /// </summary>
    public sealed class ContainerGroupProfileStubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        ///  Container Group properties which can be set while creating or updating the NGroups.
        /// </summary>
        [Input("containerGroupProperties")]
        public Input<Inputs.NGroupContainerGroupPropertiesArgs>? ContainerGroupProperties { get; set; }

        /// <summary>
        /// A network profile for network settings of a ContainerGroupProfile.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.NetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// A reference to the container group profile ARM resource hosted in ACI RP.
        /// </summary>
        [Input("resource")]
        public Input<Inputs.ApiEntityReferenceArgs>? Resource { get; set; }

        /// <summary>
        /// The revision of the CG profile is an optional property. If customer does not to provide a revision then NGroups will pickup the latest revision of CGProfile.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        /// <summary>
        /// Storage profile for storage related settings of a container group profile.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        public ContainerGroupProfileStubArgs()
        {
        }
        public static new ContainerGroupProfileStubArgs Empty => new ContainerGroupProfileStubArgs();
    }
}
