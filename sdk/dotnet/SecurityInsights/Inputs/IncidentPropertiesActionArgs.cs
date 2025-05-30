// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    public sealed class IncidentPropertiesActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reason the incident was closed
        /// </summary>
        [Input("classification")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.IncidentClassification>? Classification { get; set; }

        /// <summary>
        /// Describes the reason the incident was closed.
        /// </summary>
        [Input("classificationComment")]
        public Input<string>? ClassificationComment { get; set; }

        /// <summary>
        /// The classification reason the incident was closed with
        /// </summary>
        [Input("classificationReason")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.IncidentClassificationReason>? ClassificationReason { get; set; }

        [Input("labels")]
        private InputList<Inputs.IncidentLabelArgs>? _labels;

        /// <summary>
        /// List of labels to add to the incident.
        /// </summary>
        public InputList<Inputs.IncidentLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.IncidentLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Information on the user an incident is assigned to
        /// </summary>
        [Input("owner")]
        public Input<Inputs.IncidentOwnerInfoArgs>? Owner { get; set; }

        /// <summary>
        /// The severity of the incident
        /// </summary>
        [Input("severity")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.IncidentSeverity>? Severity { get; set; }

        /// <summary>
        /// The status of the incident
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.IncidentStatus>? Status { get; set; }

        public IncidentPropertiesActionArgs()
        {
        }
        public static new IncidentPropertiesActionArgs Empty => new IncidentPropertiesActionArgs();
    }
}
