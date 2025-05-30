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
    /// The configuration settings of the login flow of the custom Open ID Connect provider.
    /// </summary>
    public sealed class OpenIdConnectLoginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the claim that contains the users name.
        /// </summary>
        [Input("nameClaimType")]
        public Input<string>? NameClaimType { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// A list of the scopes that should be requested while authenticating.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public OpenIdConnectLoginArgs()
        {
        }
        public static new OpenIdConnectLoginArgs Empty => new OpenIdConnectLoginArgs();
    }
}
