// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// The helm deployment install options
    /// </summary>
    public sealed class HelmUpgradeOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The helm deployment atomic options
        /// </summary>
        [Input("atomic")]
        public Input<string>? Atomic { get; set; }

        /// <summary>
        /// The helm deployment timeout options
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// The helm deployment wait options
        /// </summary>
        [Input("wait")]
        public Input<string>? Wait { get; set; }

        public HelmUpgradeOptionsArgs()
        {
        }
        public static new HelmUpgradeOptionsArgs Empty => new HelmUpgradeOptionsArgs();
    }
}
