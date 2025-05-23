// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// The ip configurations for the private link service.
    /// </summary>
    public sealed class IPConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of private link IP configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Indicates whether this IP configuration is primary for the corresponding NIC.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("privateIPAddress")]
        public Input<string>? PrivateIPAddress { get; set; }

        /// <summary>
        /// The method that private IP address is allocated.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public InputUnion<string, Pulumi.AzureNative.HDInsight.PrivateIPAllocationMethod>? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// The subnet resource id.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.ResourceIdArgs>? Subnet { get; set; }

        public IPConfigurationArgs()
        {
        }
        public static new IPConfigurationArgs Empty => new IPConfigurationArgs();
    }
}
