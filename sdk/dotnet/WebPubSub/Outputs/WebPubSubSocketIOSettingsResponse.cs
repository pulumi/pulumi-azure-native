// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WebPubSub.Outputs
{

    /// <summary>
    /// SocketIO settings for the resource
    /// </summary>
    [OutputType]
    public sealed class WebPubSubSocketIOSettingsResponse
    {
        /// <summary>
        /// The service mode of Web PubSub for Socket.IO. Values allowed: 
        /// "Default": have your own backend Socket.IO server
        /// "Serverless": your application doesn't have a backend server
        /// </summary>
        public readonly string? ServiceMode;

        [OutputConstructor]
        private WebPubSubSocketIOSettingsResponse(string? serviceMode)
        {
            ServiceMode = serviceMode;
        }
    }
}
