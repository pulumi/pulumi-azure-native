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
    /// The configuration settings of the registration for the Azure Static Web Apps provider
    /// </summary>
    public sealed class AzureStaticWebAppsRegistrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of the app used for login.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        public AzureStaticWebAppsRegistrationArgs()
        {
        }
        public static new AzureStaticWebAppsRegistrationArgs Empty => new AzureStaticWebAppsRegistrationArgs();
    }
}
