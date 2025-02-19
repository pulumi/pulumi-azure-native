// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppConfiguration.V20230901Preview.Inputs
{

    /// <summary>
    /// The data plane proxy settings for a configuration store.
    /// </summary>
    public sealed class DataPlaneProxyPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data plane proxy authentication mode. This property manages the authentication mode of request to the data plane resources.
        /// </summary>
        [Input("authenticationMode")]
        public InputUnion<string, Pulumi.AzureNative.AppConfiguration.V20230901Preview.AuthenticationMode>? AuthenticationMode { get; set; }

        /// <summary>
        /// The data plane proxy private link delegation. This property manages if a request from delegated ARM private link is allowed when the data plane resource requires private link.
        /// </summary>
        [Input("privateLinkDelegation")]
        public InputUnion<string, Pulumi.AzureNative.AppConfiguration.V20230901Preview.PrivateLinkDelegation>? PrivateLinkDelegation { get; set; }

        public DataPlaneProxyPropertiesArgs()
        {
        }
        public static new DataPlaneProxyPropertiesArgs Empty => new DataPlaneProxyPropertiesArgs();
    }
}
