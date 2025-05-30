// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The Microsoft Defender for Endpoint autoprovisioning configuration
    /// </summary>
    public sealed class DefenderForServersAwsOfferingMdeAutoProvisioningArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// configuration for Microsoft Defender for Endpoint autoprovisioning
        /// </summary>
        [Input("configuration")]
        public Input<object>? Configuration { get; set; }

        /// <summary>
        /// Is Microsoft Defender for Endpoint auto provisioning enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public DefenderForServersAwsOfferingMdeAutoProvisioningArgs()
        {
        }
        public static new DefenderForServersAwsOfferingMdeAutoProvisioningArgs Empty => new DefenderForServersAwsOfferingMdeAutoProvisioningArgs();
    }
}
