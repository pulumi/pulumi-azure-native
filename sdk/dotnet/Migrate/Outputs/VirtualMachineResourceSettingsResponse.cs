// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// Gets or sets the virtual machine resource settings.
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineResourceSettingsResponse
    {
        /// <summary>
        /// The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        /// Expected value is 'Microsoft.Compute/virtualMachines'.
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// Gets or sets the Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets or sets the target availability set id for virtual machines not in an availability set at source.
        /// </summary>
        public readonly string? TargetAvailabilitySetId;
        /// <summary>
        /// Gets or sets the target availability zone.
        /// </summary>
        public readonly string? TargetAvailabilityZone;
        /// <summary>
        /// Gets or sets the target resource group name.
        /// </summary>
        public readonly string? TargetResourceGroupName;
        /// <summary>
        /// Gets or sets the target Resource name.
        /// </summary>
        public readonly string? TargetResourceName;
        /// <summary>
        /// Gets or sets the target virtual machine size.
        /// </summary>
        public readonly string? TargetVmSize;
        /// <summary>
        /// Gets or sets user-managed identities
        /// </summary>
        public readonly ImmutableArray<string> UserManagedIdentities;

        [OutputConstructor]
        private VirtualMachineResourceSettingsResponse(
            string resourceType,

            ImmutableDictionary<string, string>? tags,

            string? targetAvailabilitySetId,

            string? targetAvailabilityZone,

            string? targetResourceGroupName,

            string? targetResourceName,

            string? targetVmSize,

            ImmutableArray<string> userManagedIdentities)
        {
            ResourceType = resourceType;
            Tags = tags;
            TargetAvailabilitySetId = targetAvailabilitySetId;
            TargetAvailabilityZone = targetAvailabilityZone;
            TargetResourceGroupName = targetResourceGroupName;
            TargetResourceName = targetResourceName;
            TargetVmSize = targetVmSize;
            UserManagedIdentities = userManagedIdentities;
        }
    }
}
