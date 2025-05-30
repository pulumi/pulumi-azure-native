// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabricMesh.Inputs
{

    /// <summary>
    /// Describes the hostname properties for http routing.
    /// </summary>
    public sealed class HttpRouteConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes destination endpoint for routing traffic.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.GatewayDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Describes a rule for http route matching.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.HttpRouteMatchRuleArgs> Match { get; set; } = null!;

        /// <summary>
        /// http route name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public HttpRouteConfigArgs()
        {
        }
        public static new HttpRouteConfigArgs Empty => new HttpRouteConfigArgs();
    }
}
