// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Script block of scripts.
    /// </summary>
    [OutputType]
    public sealed class ScriptActivityScriptBlockResponse
    {
        /// <summary>
        /// Array of script parameters. Type: array.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScriptActivityParameterResponse> Parameters;
        /// <summary>
        /// The query text. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Text;
        /// <summary>
        /// The type of the query. Please refer to the ScriptType for valid options. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Type;

        [OutputConstructor]
        private ScriptActivityScriptBlockResponse(
            ImmutableArray<Outputs.ScriptActivityParameterResponse> parameters,

            object text,

            object type)
        {
            Parameters = parameters;
            Text = text;
            Type = type;
        }
    }
}
