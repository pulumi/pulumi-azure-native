// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Inputs
{

    /// <summary>
    /// a plain text value execution parameter
    /// </summary>
    public sealed class ScriptStringExecutionParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameter name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// script execution parameter type
        /// Expected value is 'Value'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value for the passed parameter
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ScriptStringExecutionParameterArgs()
        {
        }
        public static new ScriptStringExecutionParameterArgs Empty => new ScriptStringExecutionParameterArgs();
    }
}
