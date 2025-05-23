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
    /// Model for API authentication with OAuth2.
    /// </summary>
    public sealed class OAuthModelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access token prepend. Default is 'Bearer'.
        /// </summary>
        [Input("accessTokenPrepend")]
        public Input<string>? AccessTokenPrepend { get; set; }

        /// <summary>
        /// The user's authorization code.
        /// </summary>
        [Input("authorizationCode")]
        public Input<string>? AuthorizationCode { get; set; }

        /// <summary>
        /// The authorization endpoint.
        /// </summary>
        [Input("authorizationEndpoint")]
        public Input<string>? AuthorizationEndpoint { get; set; }

        [Input("authorizationEndpointHeaders")]
        private InputMap<string>? _authorizationEndpointHeaders;

        /// <summary>
        /// The authorization endpoint headers.
        /// </summary>
        public InputMap<string> AuthorizationEndpointHeaders
        {
            get => _authorizationEndpointHeaders ?? (_authorizationEndpointHeaders = new InputMap<string>());
            set => _authorizationEndpointHeaders = value;
        }

        [Input("authorizationEndpointQueryParameters")]
        private InputMap<string>? _authorizationEndpointQueryParameters;

        /// <summary>
        /// The authorization endpoint query parameters.
        /// </summary>
        public InputMap<string> AuthorizationEndpointQueryParameters
        {
            get => _authorizationEndpointQueryParameters ?? (_authorizationEndpointQueryParameters = new InputMap<string>());
            set => _authorizationEndpointQueryParameters = value;
        }

        /// <summary>
        /// The Application (client) ID that the OAuth provider assigned to your app.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Application (client) secret that the OAuth provider assigned to your app.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// The grant type, usually will be 'authorization code'.
        /// </summary>
        [Input("grantType", required: true)]
        public Input<string> GrantType { get; set; } = null!;

        /// <summary>
        /// Indicating whether we want to send the clientId and clientSecret to token endpoint in the headers.
        /// </summary>
        [Input("isCredentialsInHeaders")]
        public Input<bool>? IsCredentialsInHeaders { get; set; }

        /// <summary>
        /// A value indicating whether it's a JWT flow.
        /// </summary>
        [Input("isJwtBearerFlow")]
        public Input<bool>? IsJwtBearerFlow { get; set; }

        /// <summary>
        /// The Application redirect url that the user config in the OAuth provider.
        /// </summary>
        [Input("redirectUri")]
        public Input<string>? RedirectUri { get; set; }

        /// <summary>
        /// The Application (client) Scope that the OAuth provider assigned to your app.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The token endpoint. Defines the OAuth2 refresh token.
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        [Input("tokenEndpointHeaders")]
        private InputMap<string>? _tokenEndpointHeaders;

        /// <summary>
        /// The token endpoint headers.
        /// </summary>
        public InputMap<string> TokenEndpointHeaders
        {
            get => _tokenEndpointHeaders ?? (_tokenEndpointHeaders = new InputMap<string>());
            set => _tokenEndpointHeaders = value;
        }

        [Input("tokenEndpointQueryParameters")]
        private InputMap<string>? _tokenEndpointQueryParameters;

        /// <summary>
        /// The token endpoint query parameters.
        /// </summary>
        public InputMap<string> TokenEndpointQueryParameters
        {
            get => _tokenEndpointQueryParameters ?? (_tokenEndpointQueryParameters = new InputMap<string>());
            set => _tokenEndpointQueryParameters = value;
        }

        /// <summary>
        /// Type of paging
        /// Expected value is 'OAuth2'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public OAuthModelArgs()
        {
            IsCredentialsInHeaders = false;
        }
        public static new OAuthModelArgs Empty => new OAuthModelArgs();
    }
}
