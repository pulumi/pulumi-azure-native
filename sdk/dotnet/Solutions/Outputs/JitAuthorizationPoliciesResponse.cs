// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions.Outputs
{

    /// <summary>
    /// The JIT authorization policies.
    /// </summary>
    [OutputType]
    public sealed class JitAuthorizationPoliciesResponse
    {
        /// <summary>
        /// The the principal id that will be granted JIT access.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The role definition id that will be granted to the Principal.
        /// </summary>
        public readonly string RoleDefinitionId;

        [OutputConstructor]
        private JitAuthorizationPoliciesResponse(
            string principalId,

            string roleDefinitionId)
        {
            PrincipalId = principalId;
            RoleDefinitionId = roleDefinitionId;
        }
    }
}
