// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.Inputs
{

    /// <summary>
    /// TLS settings for the resource
    /// </summary>
    public sealed class SignalRTlsSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Request client certificate during TLS handshake if enabled. Not supported for free tier. Any input will be ignored for free tier.
        /// </summary>
        [Input("clientCertEnabled")]
        public Input<bool>? ClientCertEnabled { get; set; }

        public SignalRTlsSettingsArgs()
        {
            ClientCertEnabled = false;
        }
        public static new SignalRTlsSettingsArgs Empty => new SignalRTlsSettingsArgs();
    }
}
