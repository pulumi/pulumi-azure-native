// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// The ARM ID for a Network Interface.
    /// </summary>
    public sealed class NetworkInterfaceArmReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARM ID for a Network Interface.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public NetworkInterfaceArmReferenceArgs()
        {
        }
        public static new NetworkInterfaceArmReferenceArgs Empty => new NetworkInterfaceArmReferenceArgs();
    }
}
