// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// Represents security alert timeline item.
    /// </summary>
    [OutputType]
    public sealed class SecurityAlertTimelineItemResponse
    {
        /// <summary>
        /// The name of the alert type.
        /// </summary>
        public readonly string AlertType;
        /// <summary>
        /// The alert azure resource id.
        /// </summary>
        public readonly string AzureResourceId;
        /// <summary>
        /// The alert description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The alert name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The alert end time.
        /// </summary>
        public readonly string EndTimeUtc;
        /// <summary>
        /// The intent of the alert.
        /// </summary>
        public readonly string Intent;
        /// <summary>
        /// The entity query kind
        /// Expected value is 'SecurityAlert'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The alert product name.
        /// </summary>
        public readonly string? ProductName;
        /// <summary>
        /// The alert severity.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// The alert start time.
        /// </summary>
        public readonly string StartTimeUtc;
        /// <summary>
        /// The techniques of the alert.
        /// </summary>
        public readonly ImmutableArray<string> Techniques;
        /// <summary>
        /// The alert generated time.
        /// </summary>
        public readonly string TimeGenerated;

        [OutputConstructor]
        private SecurityAlertTimelineItemResponse(
            string alertType,

            string azureResourceId,

            string? description,

            string displayName,

            string endTimeUtc,

            string intent,

            string kind,

            string? productName,

            string severity,

            string startTimeUtc,

            ImmutableArray<string> techniques,

            string timeGenerated)
        {
            AlertType = alertType;
            AzureResourceId = azureResourceId;
            Description = description;
            DisplayName = displayName;
            EndTimeUtc = endTimeUtc;
            Intent = intent;
            Kind = kind;
            ProductName = productName;
            Severity = severity;
            StartTimeUtc = startTimeUtc;
            Techniques = techniques;
            TimeGenerated = timeGenerated;
        }
    }
}
