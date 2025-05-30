// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    /// <summary>
    /// A single field mapping of the mapped entity
    /// </summary>
    [OutputType]
    public sealed class FieldMappingResponse
    {
        /// <summary>
        /// the column name to be mapped to the identifier
        /// </summary>
        public readonly string? ColumnName;
        /// <summary>
        /// the V3 identifier of the entity
        /// </summary>
        public readonly string? Identifier;

        [OutputConstructor]
        private FieldMappingResponse(
            string? columnName,

            string? identifier)
        {
            ColumnName = columnName;
            Identifier = identifier;
        }
    }
}
