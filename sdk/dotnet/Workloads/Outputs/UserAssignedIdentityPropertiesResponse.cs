// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// User assigned managed identity properties.
    /// </summary>
    [OutputType]
    public sealed class UserAssignedIdentityPropertiesResponse
    {
        public readonly string? ClientId;
        public readonly string? PrincipalId;

        [OutputConstructor]
        private UserAssignedIdentityPropertiesResponse(
            string? clientId,

            string? principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
