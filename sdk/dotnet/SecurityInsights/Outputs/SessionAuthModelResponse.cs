// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// Model for API authentication with session cookie.
    /// </summary>
    [OutputType]
    public sealed class SessionAuthModelResponse
    {
        /// <summary>
        /// HTTP request headers to session service endpoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Headers;
        /// <summary>
        /// Indicating whether API key is set in HTTP POST payload.
        /// </summary>
        public readonly bool? IsPostPayloadJson;
        /// <summary>
        /// The password attribute name.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Password;
        /// <summary>
        /// Query parameters to session service endpoint.
        /// </summary>
        public readonly object? QueryParameters;
        /// <summary>
        /// Session id attribute name from HTTP response header.
        /// </summary>
        public readonly string? SessionIdName;
        /// <summary>
        /// HTTP request URL to session service endpoint.
        /// </summary>
        public readonly string? SessionLoginRequestUri;
        /// <summary>
        /// Session timeout in minutes.
        /// </summary>
        public readonly int? SessionTimeoutInMinutes;
        /// <summary>
        /// Type of paging
        /// Expected value is 'Session'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user name attribute key value.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UserName;

        [OutputConstructor]
        private SessionAuthModelResponse(
            ImmutableDictionary<string, string>? headers,

            bool? isPostPayloadJson,

            ImmutableDictionary<string, string> password,

            object? queryParameters,

            string? sessionIdName,

            string? sessionLoginRequestUri,

            int? sessionTimeoutInMinutes,

            string type,

            ImmutableDictionary<string, string> userName)
        {
            Headers = headers;
            IsPostPayloadJson = isPostPayloadJson;
            Password = password;
            QueryParameters = queryParameters;
            SessionIdName = sessionIdName;
            SessionLoginRequestUri = sessionLoginRequestUri;
            SessionTimeoutInMinutes = sessionTimeoutInMinutes;
            Type = type;
            UserName = userName;
        }
    }
}
