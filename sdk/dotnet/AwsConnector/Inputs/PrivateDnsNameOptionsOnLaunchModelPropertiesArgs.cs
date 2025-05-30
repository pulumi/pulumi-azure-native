// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of PrivateDnsNameOptionsOnLaunchModelProperties
    /// </summary>
    public sealed class PrivateDnsNameOptionsOnLaunchModelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property enableResourceNameDnsAAAARecord
        /// </summary>
        [Input("enableResourceNameDnsAAAARecord")]
        public Input<bool>? EnableResourceNameDnsAAAARecord { get; set; }

        /// <summary>
        /// Property enableResourceNameDnsARecord
        /// </summary>
        [Input("enableResourceNameDnsARecord")]
        public Input<bool>? EnableResourceNameDnsARecord { get; set; }

        /// <summary>
        /// Property hostnameType
        /// </summary>
        [Input("hostnameType")]
        public Input<string>? HostnameType { get; set; }

        public PrivateDnsNameOptionsOnLaunchModelPropertiesArgs()
        {
        }
        public static new PrivateDnsNameOptionsOnLaunchModelPropertiesArgs Empty => new PrivateDnsNameOptionsOnLaunchModelPropertiesArgs();
    }
}
