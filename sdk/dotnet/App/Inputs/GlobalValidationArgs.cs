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
    /// The configuration settings that determines the validation flow of users using ContainerApp Service Authentication/Authorization.
    /// </summary>
    public sealed class GlobalValidationArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedPaths")]
        private InputList<string>? _excludedPaths;

        /// <summary>
        /// The paths for which unauthenticated flow would not be redirected to the login page.
        /// </summary>
        public InputList<string> ExcludedPaths
        {
            get => _excludedPaths ?? (_excludedPaths = new InputList<string>());
            set => _excludedPaths = value;
        }

        /// <summary>
        /// The default authentication provider to use when multiple providers are configured.
        /// This setting is only needed if multiple providers are configured and the unauthenticated client
        /// action is set to "RedirectToLoginPage".
        /// </summary>
        [Input("redirectToProvider")]
        public Input<string>? RedirectToProvider { get; set; }

        /// <summary>
        /// The action to take when an unauthenticated client attempts to access the app.
        /// </summary>
        [Input("unauthenticatedClientAction")]
        public Input<Pulumi.AzureNative.App.UnauthenticatedClientActionV2>? UnauthenticatedClientAction { get; set; }

        public GlobalValidationArgs()
        {
        }
        public static new GlobalValidationArgs Empty => new GlobalValidationArgs();
    }
}
