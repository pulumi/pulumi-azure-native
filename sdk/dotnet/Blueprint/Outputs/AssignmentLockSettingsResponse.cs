// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint.Outputs
{

    /// <summary>
    /// Defines how resources deployed by a blueprint assignment are locked.
    /// </summary>
    [OutputType]
    public sealed class AssignmentLockSettingsResponse
    {
        /// <summary>
        /// List of management operations that are excluded from blueprint locks. Up to 200 actions are permitted. If the lock mode is set to 'AllResourcesReadOnly', then the following actions are automatically appended to 'excludedActions': '*/read', 'Microsoft.Network/virtualNetworks/subnets/join/action' and 'Microsoft.Authorization/locks/delete'. If the lock mode is set to 'AllResourcesDoNotDelete', then the following actions are automatically appended to 'excludedActions': 'Microsoft.Authorization/locks/delete'. Duplicate actions will get removed.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedActions;
        /// <summary>
        /// List of AAD principals excluded from blueprint locks. Up to 5 principals are permitted.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedPrincipals;
        /// <summary>
        /// Lock mode.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private AssignmentLockSettingsResponse(
            ImmutableArray<string> excludedActions,

            ImmutableArray<string> excludedPrincipals,

            string? mode)
        {
            ExcludedActions = excludedActions;
            ExcludedPrincipals = excludedPrincipals;
            Mode = mode;
        }
    }
}
