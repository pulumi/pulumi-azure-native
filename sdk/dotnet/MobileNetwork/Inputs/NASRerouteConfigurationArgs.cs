// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Inputs
{

    /// <summary>
    /// Configuration enabling NAS reroute.
    /// </summary>
    public sealed class NASRerouteConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The macro network's MME group ID. This is where unknown UEs are sent to via NAS reroute.
        /// </summary>
        [Input("macroMmeGroupId", required: true)]
        public Input<int> MacroMmeGroupId { get; set; } = null!;

        public NASRerouteConfigurationArgs()
        {
        }
        public static new NASRerouteConfigurationArgs Empty => new NASRerouteConfigurationArgs();
    }
}
