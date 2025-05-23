// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Inputs
{

    /// <summary>
    /// Infra network profile for HCI platform
    /// </summary>
    public sealed class VirtualNetworksPropertiesHciArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource group in MOC(Microsoft On-premises Cloud)
        /// </summary>
        [Input("mocGroup")]
        public Input<string>? MocGroup { get; set; }

        /// <summary>
        /// Location in MOC(Microsoft On-premises Cloud)
        /// </summary>
        [Input("mocLocation")]
        public Input<string>? MocLocation { get; set; }

        /// <summary>
        /// Virtual Network name in MOC(Microsoft On-premises Cloud)
        /// </summary>
        [Input("mocVnetName")]
        public Input<string>? MocVnetName { get; set; }

        public VirtualNetworksPropertiesHciArgs()
        {
        }
        public static new VirtualNetworksPropertiesHciArgs Empty => new VirtualNetworksPropertiesHciArgs();
    }
}
