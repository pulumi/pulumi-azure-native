// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CloudHealth.Outputs
{

    /// <summary>
    /// Relationship properties
    /// </summary>
    [OutputType]
    public sealed class RelationshipPropertiesResponse
    {
        /// <summary>
        /// Resource name of the child entity
        /// </summary>
        public readonly string ChildEntityName;
        /// <summary>
        /// Date when the relationship was (soft-)deleted
        /// </summary>
        public readonly string DeletionDate;
        /// <summary>
        /// Discovered by which discovery rule. If set, the relationship cannot be deleted manually.
        /// </summary>
        public readonly string DiscoveredBy;
        /// <summary>
        /// Display name
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Optional set of labels (key-value pairs)
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Resource name of the parent entity
        /// </summary>
        public readonly string ParentEntityName;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private RelationshipPropertiesResponse(
            string childEntityName,

            string deletionDate,

            string discoveredBy,

            string? displayName,

            ImmutableDictionary<string, string>? labels,

            string parentEntityName,

            string provisioningState)
        {
            ChildEntityName = childEntityName;
            DeletionDate = deletionDate;
            DiscoveredBy = discoveredBy;
            DisplayName = displayName;
            Labels = labels;
            ParentEntityName = parentEntityName;
            ProvisioningState = provisioningState;
        }
    }
}
