// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Model for API authentication with session cookie.
    /// </summary>
    public sealed class SessionAuthModelArgs : global::Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// HTTP request headers to session service endpoint.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Indicating whether API key is set in HTTP POST payload.
        /// </summary>
        [Input("isPostPayloadJson")]
        public Input<bool>? IsPostPayloadJson { get; set; }

        [Input("password", required: true)]
        private InputMap<string>? _password;

        /// <summary>
        /// The password attribute name.
        /// </summary>
        public InputMap<string> Password
        {
            get => _password ?? (_password = new InputMap<string>());
            set => _password = value;
        }

        /// <summary>
        /// Query parameters to session service endpoint.
        /// </summary>
        [Input("queryParameters")]
        public Input<object>? QueryParameters { get; set; }

        /// <summary>
        /// Session id attribute name from HTTP response header.
        /// </summary>
        [Input("sessionIdName")]
        public Input<string>? SessionIdName { get; set; }

        /// <summary>
        /// HTTP request URL to session service endpoint.
        /// </summary>
        [Input("sessionLoginRequestUri")]
        public Input<string>? SessionLoginRequestUri { get; set; }

        /// <summary>
        /// Session timeout in minutes.
        /// </summary>
        [Input("sessionTimeoutInMinutes")]
        public Input<int>? SessionTimeoutInMinutes { get; set; }

        /// <summary>
        /// Type of paging
        /// Expected value is 'Session'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userName", required: true)]
        private InputMap<string>? _userName;

        /// <summary>
        /// The user name attribute key value.
        /// </summary>
        public InputMap<string> UserName
        {
            get => _userName ?? (_userName = new InputMap<string>());
            set => _userName = value;
        }

        public SessionAuthModelArgs()
        {
        }
        public static new SessionAuthModelArgs Empty => new SessionAuthModelArgs();
    }
}
