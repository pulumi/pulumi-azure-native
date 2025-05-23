// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Job endpoint definition
    /// </summary>
    public sealed class JobServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Url for endpoint.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Endpoint type.
        /// </summary>
        [Input("jobServiceType")]
        public Input<string>? JobServiceType { get; set; }

        /// <summary>
        /// Nodes that user would like to start the service on.
        /// If Nodes is not set or set to null, the service will only be started on leader node.
        /// </summary>
        [Input("nodes")]
        public Input<Inputs.AllNodesArgs>? Nodes { get; set; }

        /// <summary>
        /// Port for endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Additional properties to set on the endpoint.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        public JobServiceArgs()
        {
        }
        public static new JobServiceArgs Empty => new JobServiceArgs();
    }
}
