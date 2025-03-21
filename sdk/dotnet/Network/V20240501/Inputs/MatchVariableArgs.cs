// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20240501.Inputs
{

    /// <summary>
    /// Define match variables.
    /// </summary>
    public sealed class MatchVariableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The selector of match variable.
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        /// <summary>
        /// Match Variable.
        /// </summary>
        [Input("variableName", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20240501.WebApplicationFirewallMatchVariable> VariableName { get; set; } = null!;

        public MatchVariableArgs()
        {
        }
        public static new MatchVariableArgs Empty => new MatchVariableArgs();
    }
}
