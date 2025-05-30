// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// Parameter example.
    /// </summary>
    [OutputType]
    public sealed class ParameterExampleContractResponse
    {
        /// <summary>
        /// Long description for the example
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A URL that points to the literal example
        /// </summary>
        public readonly string? ExternalValue;
        /// <summary>
        /// Short description for the example
        /// </summary>
        public readonly string? Summary;
        /// <summary>
        /// Example value. May be a primitive value, or an object.
        /// </summary>
        public readonly object? Value;

        [OutputConstructor]
        private ParameterExampleContractResponse(
            string? description,

            string? externalValue,

            string? summary,

            object? value)
        {
            Description = description;
            ExternalValue = externalValue;
            Summary = summary;
            Value = value;
        }
    }
}
