// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// The UiConfig for 'Customizable' connector definition kind.
    /// </summary>
    public sealed class CustomizableConnectorUiConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The exposure status of the connector to the customers.
        /// </summary>
        [Input("availability")]
        public Input<Inputs.ConnectorDefinitionsAvailabilityArgs>? Availability { get; set; }

        [Input("connectivityCriteria", required: true)]
        private InputList<Inputs.ConnectivityCriterionArgs>? _connectivityCriteria;

        /// <summary>
        /// Gets or sets the way the connector checks whether the connector is connected.
        /// </summary>
        public InputList<Inputs.ConnectivityCriterionArgs> ConnectivityCriteria
        {
            get => _connectivityCriteria ?? (_connectivityCriteria = new InputList<Inputs.ConnectivityCriterionArgs>());
            set => _connectivityCriteria = value;
        }

        [Input("dataTypes", required: true)]
        private InputList<Inputs.ConnectorDataTypeArgs>? _dataTypes;

        /// <summary>
        /// Gets or sets the data types to check for last data received.
        /// </summary>
        public InputList<Inputs.ConnectorDataTypeArgs> DataTypes
        {
            get => _dataTypes ?? (_dataTypes = new InputList<Inputs.ConnectorDataTypeArgs>());
            set => _dataTypes = value;
        }

        /// <summary>
        /// Gets or sets the connector description in markdown format.
        /// </summary>
        [Input("descriptionMarkdown", required: true)]
        public Input<string> DescriptionMarkdown { get; set; } = null!;

        [Input("graphQueries", required: true)]
        private InputList<Inputs.GraphQueryArgs>? _graphQueries;

        /// <summary>
        /// Gets or sets the graph queries to show the current data volume over time.
        /// </summary>
        public InputList<Inputs.GraphQueryArgs> GraphQueries
        {
            get => _graphQueries ?? (_graphQueries = new InputList<Inputs.GraphQueryArgs>());
            set => _graphQueries = value;
        }

        /// <summary>
        /// Gets or sets custom connector id. optional field.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instructionSteps", required: true)]
        private InputList<Inputs.InstructionStepArgs>? _instructionSteps;

        /// <summary>
        /// Gets or sets the instruction steps to enable the connector.
        /// </summary>
        public InputList<Inputs.InstructionStepArgs> InstructionSteps
        {
            get => _instructionSteps ?? (_instructionSteps = new InputList<Inputs.InstructionStepArgs>());
            set => _instructionSteps = value;
        }

        /// <summary>
        /// Gets or sets a value indicating whether to use 'OR'(SOME) or 'AND' between ConnectivityCriteria items.
        /// </summary>
        [Input("isConnectivityCriteriasMatchSome")]
        public Input<bool>? IsConnectivityCriteriasMatchSome { get; set; }

        /// <summary>
        /// Gets or sets the connector logo to be used when displaying the connector within Azure Sentinel's connector's gallery.
        /// The logo value should be in SVG format.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The required Permissions for the connector.
        /// </summary>
        [Input("permissions", required: true)]
        public Input<Inputs.ConnectorDefinitionsPermissionsArgs> Permissions { get; set; } = null!;

        /// <summary>
        /// Gets or sets the connector publisher name.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        /// <summary>
        /// Gets or sets the connector blade title.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public CustomizableConnectorUiConfigArgs()
        {
        }
        public static new CustomizableConnectorUiConfigArgs Empty => new CustomizableConnectorUiConfigArgs();
    }
}
