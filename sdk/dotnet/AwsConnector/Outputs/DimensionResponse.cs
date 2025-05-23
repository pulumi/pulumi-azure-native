// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of Dimension
    /// </summary>
    [OutputType]
    public sealed class DimensionResponse
    {
        /// <summary>
        /// The name for the CW metric dimension that the metric filter creates. Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:).
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The name of the dimension, from 1–255 characters in length. This dimension name must have been included when the metric was published.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Property value
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private DimensionResponse(
            string? key,

            string? name,

            string? value)
        {
            Key = key;
            Name = name;
            Value = value;
        }
    }
}
