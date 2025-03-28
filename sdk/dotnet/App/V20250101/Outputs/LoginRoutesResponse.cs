// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Outputs
{

    /// <summary>
    /// The routes that specify the endpoints used for login and logout requests.
    /// </summary>
    [OutputType]
    public sealed class LoginRoutesResponse
    {
        /// <summary>
        /// The endpoint at which a logout request should be made.
        /// </summary>
        public readonly string? LogoutEndpoint;

        [OutputConstructor]
        private LoginRoutesResponse(string? logoutEndpoint)
        {
            LogoutEndpoint = logoutEndpoint;
        }
    }
}
