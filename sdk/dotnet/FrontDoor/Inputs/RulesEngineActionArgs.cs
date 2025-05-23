// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.FrontDoor.Inputs
{

    /// <summary>
    /// One or more actions that will execute, modifying the request and/or response.
    /// </summary>
    public sealed class RulesEngineActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("requestHeaderActions")]
        private InputList<Inputs.HeaderActionArgs>? _requestHeaderActions;

        /// <summary>
        /// A list of header actions to apply from the request from AFD to the origin.
        /// </summary>
        public InputList<Inputs.HeaderActionArgs> RequestHeaderActions
        {
            get => _requestHeaderActions ?? (_requestHeaderActions = new InputList<Inputs.HeaderActionArgs>());
            set => _requestHeaderActions = value;
        }

        [Input("responseHeaderActions")]
        private InputList<Inputs.HeaderActionArgs>? _responseHeaderActions;

        /// <summary>
        /// A list of header actions to apply from the response from AFD to the client.
        /// </summary>
        public InputList<Inputs.HeaderActionArgs> ResponseHeaderActions
        {
            get => _responseHeaderActions ?? (_responseHeaderActions = new InputList<Inputs.HeaderActionArgs>());
            set => _responseHeaderActions = value;
        }

        /// <summary>
        /// Override the route configuration.
        /// </summary>
        [Input("routeConfigurationOverride")]
        public InputUnion<Inputs.ForwardingConfigurationArgs, Inputs.RedirectConfigurationArgs>? RouteConfigurationOverride { get; set; }

        public RulesEngineActionArgs()
        {
        }
        public static new RulesEngineActionArgs Empty => new RulesEngineActionArgs();
    }
}
