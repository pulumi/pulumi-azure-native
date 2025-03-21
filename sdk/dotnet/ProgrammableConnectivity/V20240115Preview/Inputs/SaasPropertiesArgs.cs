// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProgrammableConnectivity.V20240115Preview.Inputs
{

    /// <summary>
    /// Details about the SaaS offer purchased from the marketplace.
    /// </summary>
    public sealed class SaasPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of the SaaS offer purchased from the marketplace.
        /// </summary>
        [Input("saasResourceId")]
        public Input<string>? SaasResourceId { get; set; }

        /// <summary>
        /// Subscription ID of the SaaS offer purchased from the marketplace.
        /// </summary>
        [Input("saasSubscriptionId")]
        public Input<string>? SaasSubscriptionId { get; set; }

        public SaasPropertiesArgs()
        {
        }
        public static new SaasPropertiesArgs Empty => new SaasPropertiesArgs();
    }
}
