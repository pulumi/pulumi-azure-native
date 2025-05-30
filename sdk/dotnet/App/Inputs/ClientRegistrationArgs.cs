// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// The configuration settings of the app registration for providers that have client ids and client secrets
    /// </summary>
    public sealed class ClientRegistrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of the app used for login.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The app setting name that contains the client secret.
        /// </summary>
        [Input("clientSecretSettingName")]
        public Input<string>? ClientSecretSettingName { get; set; }

        public ClientRegistrationArgs()
        {
        }
        public static new ClientRegistrationArgs Empty => new ClientRegistrationArgs();
    }
}
