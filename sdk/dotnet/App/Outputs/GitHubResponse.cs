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
    /// The configuration settings of the GitHub provider.
    /// </summary>
    [OutputType]
    public sealed class GitHubResponse
    {
        /// <summary>
        /// &lt;code&gt;false&lt;/code&gt; if the GitHub provider should not be enabled despite the set registration; otherwise, &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The configuration settings of the login flow.
        /// </summary>
        public readonly Outputs.LoginScopesResponse? Login;
        /// <summary>
        /// The configuration settings of the app registration for the GitHub provider.
        /// </summary>
        public readonly Outputs.ClientRegistrationResponse? Registration;

        [OutputConstructor]
        private GitHubResponse(
            bool? enabled,

            Outputs.LoginScopesResponse? login,

            Outputs.ClientRegistrationResponse? registration)
        {
            Enabled = enabled;
            Login = login;
            Registration = registration;
        }
    }
}
