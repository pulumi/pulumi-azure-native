// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
    /// </summary>
    [OutputType]
    public sealed class LoginResponse
    {
        /// <summary>
        /// External URLs that can be redirected to as part of logging in or logging out of the app. Note that the query string part of the URL is ignored.
        /// This is an advanced setting typically only needed by Windows Store application backends.
        /// Note that URLs within the current domain are always implicitly allowed.
        /// </summary>
        public readonly ImmutableArray<string> AllowedExternalRedirectUrls;
        /// <summary>
        /// The configuration settings of the session cookie's expiration.
        /// </summary>
        public readonly Outputs.CookieExpirationResponse? CookieExpiration;
        /// <summary>
        /// The configuration settings of the nonce used in the login flow.
        /// </summary>
        public readonly Outputs.NonceResponse? Nonce;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the fragments from the request are preserved after the login request is made; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? PreserveUrlFragmentsForLogins;
        /// <summary>
        /// The routes that specify the endpoints used for login and logout requests.
        /// </summary>
        public readonly Outputs.LoginRoutesResponse? Routes;
        /// <summary>
        /// The configuration settings of the token store.
        /// </summary>
        public readonly Outputs.TokenStoreResponse? TokenStore;

        [OutputConstructor]
        private LoginResponse(
            ImmutableArray<string> allowedExternalRedirectUrls,

            Outputs.CookieExpirationResponse? cookieExpiration,

            Outputs.NonceResponse? nonce,

            bool? preserveUrlFragmentsForLogins,

            Outputs.LoginRoutesResponse? routes,

            Outputs.TokenStoreResponse? tokenStore)
        {
            AllowedExternalRedirectUrls = allowedExternalRedirectUrls;
            CookieExpiration = cookieExpiration;
            Nonce = nonce;
            PreserveUrlFragmentsForLogins = preserveUrlFragmentsForLogins;
            Routes = routes;
            TokenStore = tokenStore;
        }
    }
}
