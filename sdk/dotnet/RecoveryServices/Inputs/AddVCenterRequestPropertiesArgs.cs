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
    /// The properties of an add vCenter request.
    /// </summary>
    public sealed class AddVCenterRequestPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The friendly name of the vCenter.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// The IP address of the vCenter to be discovered.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The port number for discovery.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The process server Id from where the discovery is orchestrated.
        /// </summary>
        [Input("processServerId")]
        public Input<string>? ProcessServerId { get; set; }

        /// <summary>
        /// The account Id which has privileges to discover the vCenter.
        /// </summary>
        [Input("runAsAccountId")]
        public Input<string>? RunAsAccountId { get; set; }

        public AddVCenterRequestPropertiesArgs()
        {
        }
        public static new AddVCenterRequestPropertiesArgs Empty => new AddVCenterRequestPropertiesArgs();
    }
}
