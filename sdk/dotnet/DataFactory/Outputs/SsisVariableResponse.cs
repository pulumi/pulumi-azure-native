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
    /// Ssis variable.
    /// </summary>
    [OutputType]
    public sealed class SsisVariableResponse
    {
        /// <summary>
        /// Variable type.
        /// </summary>
        public readonly string? DataType;
        /// <summary>
        /// Variable description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Variable id.
        /// </summary>
        public readonly double? Id;
        /// <summary>
        /// Variable name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Whether variable is sensitive.
        /// </summary>
        public readonly bool? Sensitive;
        /// <summary>
        /// Variable sensitive value.
        /// </summary>
        public readonly string? SensitiveValue;
        /// <summary>
        /// Variable value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private SsisVariableResponse(
            string? dataType,

            string? description,

            double? id,

            string? name,

            bool? sensitive,

            string? sensitiveValue,

            string? value)
        {
            DataType = dataType;
            Description = description;
            Id = id;
            Name = name;
            Sensitive = sensitive;
            SensitiveValue = sensitiveValue;
            Value = value;
        }
    }
}
