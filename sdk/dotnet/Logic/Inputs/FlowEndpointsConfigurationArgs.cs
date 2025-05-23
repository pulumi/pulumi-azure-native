// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.Inputs
{

    /// <summary>
    /// The endpoints configuration.
    /// </summary>
    public sealed class FlowEndpointsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connector endpoints.
        /// </summary>
        [Input("connector")]
        public Input<Inputs.FlowEndpointsArgs>? Connector { get; set; }

        /// <summary>
        /// The workflow endpoints.
        /// </summary>
        [Input("workflow")]
        public Input<Inputs.FlowEndpointsArgs>? Workflow { get; set; }

        public FlowEndpointsConfigurationArgs()
        {
        }
        public static new FlowEndpointsConfigurationArgs Empty => new FlowEndpointsConfigurationArgs();
    }
}
