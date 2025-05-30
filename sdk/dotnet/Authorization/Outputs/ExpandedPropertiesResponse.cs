// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Outputs
{

    /// <summary>
    /// Expanded info of resource, role and principal
    /// </summary>
    [OutputType]
    public sealed class ExpandedPropertiesResponse
    {
        /// <summary>
        /// Details of the principal
        /// </summary>
        public readonly Outputs.ExpandedPropertiesResponsePrincipal? Principal;
        /// <summary>
        /// Details of role definition
        /// </summary>
        public readonly Outputs.ExpandedPropertiesResponseRoleDefinition? RoleDefinition;
        /// <summary>
        /// Details of the resource scope
        /// </summary>
        public readonly Outputs.ExpandedPropertiesResponseScope? Scope;

        [OutputConstructor]
        private ExpandedPropertiesResponse(
            Outputs.ExpandedPropertiesResponsePrincipal? principal,

            Outputs.ExpandedPropertiesResponseRoleDefinition? roleDefinition,

            Outputs.ExpandedPropertiesResponseScope? scope)
        {
            Principal = principal;
            RoleDefinition = roleDefinition;
            Scope = scope;
        }
    }
}
