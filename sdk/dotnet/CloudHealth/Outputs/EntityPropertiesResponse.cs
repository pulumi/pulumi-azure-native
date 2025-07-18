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
    /// Properties which are common across all kinds of entities
    /// </summary>
    [OutputType]
    public sealed class EntityPropertiesResponse
    {
        /// <summary>
        /// Alert configuration for this entity
        /// </summary>
        public readonly Outputs.EntityAlertsResponse? Alerts;
        /// <summary>
        /// Positioning of the entity on the model canvas
        /// </summary>
        public readonly Outputs.EntityCoordinatesResponse? CanvasPosition;
        /// <summary>
        /// Date when the entity was (soft-)deleted
        /// </summary>
        public readonly string DeletionDate;
        /// <summary>
        /// Discovered by which discovery rule. If set, the entity cannot be deleted manually.
        /// </summary>
        public readonly string DiscoveredBy;
        /// <summary>
        /// Display name
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Health objective as a percentage of time the entity should be healthy.
        /// </summary>
        public readonly double? HealthObjective;
        /// <summary>
        /// Health state of this entity
        /// </summary>
        public readonly string HealthState;
        /// <summary>
        /// Visual icon definition. If not set, a default icon is used.
        /// </summary>
        public readonly Outputs.IconDefinitionResponse? Icon;
        /// <summary>
        /// Impact of the entity in health state propagation
        /// </summary>
        public readonly string? Impact;
        /// <summary>
        /// Entity kind
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Optional set of labels (key-value pairs)
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Signal groups which are assigned to this entity
        /// </summary>
        public readonly Outputs.SignalGroupResponse? Signals;

        [OutputConstructor]
        private EntityPropertiesResponse(
            Outputs.EntityAlertsResponse? alerts,

            Outputs.EntityCoordinatesResponse? canvasPosition,

            string deletionDate,

            string discoveredBy,

            string? displayName,

            double? healthObjective,

            string healthState,

            Outputs.IconDefinitionResponse? icon,

            string? impact,

            string? kind,

            ImmutableDictionary<string, string>? labels,

            string provisioningState,

            Outputs.SignalGroupResponse? signals)
        {
            Alerts = alerts;
            CanvasPosition = canvasPosition;
            DeletionDate = deletionDate;
            DiscoveredBy = discoveredBy;
            DisplayName = displayName;
            HealthObjective = healthObjective;
            HealthState = healthState;
            Icon = icon;
            Impact = impact;
            Kind = kind;
            Labels = labels;
            ProvisioningState = provisioningState;
            Signals = signals;
        }
    }
}
