// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    /// <summary>
    /// Managed resource group settings
    /// </summary>
    [OutputType]
    public sealed class ManagedResourceGroupSettingsResponse
    {
        /// <summary>
        /// List of assigned identities for the managed resource group
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedResourceGroupAssignedIdentitiesResponse> AssignedIdentities;

        [OutputConstructor]
        private ManagedResourceGroupSettingsResponse(ImmutableArray<Outputs.ManagedResourceGroupAssignedIdentitiesResponse> assignedIdentities)
        {
            AssignedIdentities = assignedIdentities;
        }
    }
}
