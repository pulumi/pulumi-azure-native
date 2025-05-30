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
    /// Dataflow BuiltIn Transformation dataset properties
    /// </summary>
    [OutputType]
    public sealed class DataflowBuiltInTransformationDatasetResponse
    {
        /// <summary>
        /// A user provided optional description of the dataset.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Condition to enrich data from Broker State Store. Example: $1 &lt; 0 || $1 &gt; $2 (Assuming inputs section $1 and $2 are provided)
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// List of fields for enriching from the Broker State Store.
        /// </summary>
        public readonly ImmutableArray<string> Inputs;
        /// <summary>
        /// The key of the dataset.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The reference to the schema that describes the dataset. Allowed: JSON Schema/draft-7.
        /// </summary>
        public readonly string? SchemaRef;

        [OutputConstructor]
        private DataflowBuiltInTransformationDatasetResponse(
            string? description,

            string? expression,

            ImmutableArray<string> inputs,

            string key,

            string? schemaRef)
        {
            Description = description;
            Expression = expression;
            Inputs = inputs;
            Key = key;
            SchemaRef = schemaRef;
        }
    }
}
