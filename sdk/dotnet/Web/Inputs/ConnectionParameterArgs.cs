// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Connection provider parameters
    /// </summary>
    public sealed class ConnectionParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OAuth settings for the connection provider
        /// </summary>
        [Input("oAuthSettings")]
        public Input<Inputs.ApiOAuthSettingsArgs>? OAuthSettings { get; set; }

        /// <summary>
        /// Type of the parameter
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.Web.ConnectionParameterType>? Type { get; set; }

        public ConnectionParameterArgs()
        {
        }
        public static new ConnectionParameterArgs Empty => new ConnectionParameterArgs();
    }
}
