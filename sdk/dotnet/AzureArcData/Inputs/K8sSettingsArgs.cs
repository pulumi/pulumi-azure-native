// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Inputs
{

    /// <summary>
    /// The kubernetes settings information.
    /// </summary>
    public sealed class K8sSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kubernetes network settings information.
        /// </summary>
        [Input("network")]
        public Input<Inputs.K8sNetworkSettingsArgs>? Network { get; set; }

        public K8sSettingsArgs()
        {
        }
        public static new K8sSettingsArgs Empty => new K8sSettingsArgs();
    }
}
