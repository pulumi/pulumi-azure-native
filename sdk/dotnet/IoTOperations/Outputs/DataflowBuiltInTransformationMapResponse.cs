// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// Dataflow BuiltIn Transformation map properties
    /// </summary>
    [OutputType]
    public sealed class DataflowBuiltInTransformationMapResponse
    {
        /// <summary>
        /// A user provided optional description of the mapping function.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Modify the inputs field(s) to the final output field. Example: $1 * 2.2 (Assuming inputs section $1 is provided)
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// List of fields for mapping in JSON path expression.
        /// </summary>
        public readonly ImmutableArray<string> Inputs;
        /// <summary>
        /// Where and how the input fields to be organized in the output record.
        /// </summary>
        public readonly string Output;
        /// <summary>
        /// Type of transformation.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DataflowBuiltInTransformationMapResponse(
            string? description,

            string? expression,

            ImmutableArray<string> inputs,

            string output,

            string? type)
        {
            Description = description;
            Expression = expression;
            Inputs = inputs;
            Output = output;
            Type = type;
        }
    }
}
