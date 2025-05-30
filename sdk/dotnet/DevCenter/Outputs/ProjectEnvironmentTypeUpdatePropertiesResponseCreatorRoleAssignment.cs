// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.Outputs
{

    /// <summary>
    /// The role definition assigned to the environment creator on backing resources.
    /// </summary>
    [OutputType]
    public sealed class ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment
    {
        /// <summary>
        /// A map of roles to assign to the environment creator.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.EnvironmentRoleResponse>? Roles;

        [OutputConstructor]
        private ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignment(ImmutableDictionary<string, Outputs.EnvironmentRoleResponse>? roles)
        {
            Roles = roles;
        }
    }
}
