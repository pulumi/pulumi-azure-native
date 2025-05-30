// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Script block of scripts.
    /// </summary>
    public sealed class ScriptActivityScriptBlockArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameters")]
        private InputList<Inputs.ScriptActivityParameterArgs>? _parameters;

        /// <summary>
        /// Array of script parameters. Type: array.
        /// </summary>
        public InputList<Inputs.ScriptActivityParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ScriptActivityParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The query text. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("text", required: true)]
        public Input<object> Text { get; set; } = null!;

        /// <summary>
        /// The type of the query. Please refer to the ScriptType for valid options. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("type", required: true)]
        public Input<object> Type { get; set; } = null!;

        public ScriptActivityScriptBlockArgs()
        {
        }
        public static new ScriptActivityScriptBlockArgs Empty => new ScriptActivityScriptBlockArgs();
    }
}
