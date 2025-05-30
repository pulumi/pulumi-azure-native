// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Outputs
{

    /// <summary>
    /// a plain text value execution parameter
    /// </summary>
    [OutputType]
    public sealed class ScriptStringExecutionParameterResponse
    {
        /// <summary>
        /// The parameter name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// script execution parameter type
        /// Expected value is 'Value'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value for the passed parameter
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ScriptStringExecutionParameterResponse(
            string name,

            string type,

            string? value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
