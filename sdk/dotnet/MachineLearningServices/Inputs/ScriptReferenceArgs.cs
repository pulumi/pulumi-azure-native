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
    /// Script reference
    /// </summary>
    public sealed class ScriptReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional command line arguments passed to the script to run.
        /// </summary>
        [Input("scriptArguments")]
        public Input<string>? ScriptArguments { get; set; }

        /// <summary>
        /// The location of scripts in the mounted volume.
        /// </summary>
        [Input("scriptData")]
        public Input<string>? ScriptData { get; set; }

        /// <summary>
        /// The storage source of the script: workspace.
        /// </summary>
        [Input("scriptSource")]
        public Input<string>? ScriptSource { get; set; }

        /// <summary>
        /// Optional time period passed to timeout command.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ScriptReferenceArgs()
        {
        }
        public static new ScriptReferenceArgs Empty => new ScriptReferenceArgs();
    }
}
