// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Configuration for static routes on this HubVnetConnectionConfiguration for static routes on this HubVnetConnection.
    /// </summary>
    public sealed class StaticRoutesConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter determining whether NVA in spoke vnet is bypassed for traffic with destination in spoke.
        /// </summary>
        [Input("vnetLocalRouteOverrideCriteria")]
        public InputUnion<string, Pulumi.AzureNative.Network.VnetLocalRouteOverrideCriteria>? VnetLocalRouteOverrideCriteria { get; set; }

        public StaticRoutesConfigArgs()
        {
        }
        public static new StaticRoutesConfigArgs Empty => new StaticRoutesConfigArgs();
    }
}
