// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// The details of the user assigned managed identity used by the Video Analyzer resource.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedManagedIdentityResponse
    {
        /// <summary>
        /// The client ID.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The principal ID.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private UserAssignedManagedIdentityResponse(
            string clientId,

            string principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
