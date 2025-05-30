// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// An Incident receiver.
    /// </summary>
    public sealed class IncidentReceiverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The incident service connection
        /// </summary>
        [Input("connection", required: true)]
        public Input<Inputs.IncidentServiceConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The incident management service type
        /// </summary>
        [Input("incidentManagementService", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Monitor.IncidentManagementService> IncidentManagementService { get; set; } = null!;

        [Input("mappings", required: true)]
        private InputMap<string>? _mappings;

        /// <summary>
        /// Field mappings for the incident service
        /// </summary>
        public InputMap<string> Mappings
        {
            get => _mappings ?? (_mappings = new InputMap<string>());
            set => _mappings = value;
        }

        /// <summary>
        /// The name of the Incident receiver. Names must be unique across all receivers within an action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public IncidentReceiverArgs()
        {
        }
        public static new IncidentReceiverArgs Empty => new IncidentReceiverArgs();
    }
}
