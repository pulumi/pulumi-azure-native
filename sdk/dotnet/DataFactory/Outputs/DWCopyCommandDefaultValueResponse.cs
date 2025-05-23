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
    /// Default value.
    /// </summary>
    [OutputType]
    public sealed class DWCopyCommandDefaultValueResponse
    {
        /// <summary>
        /// Column name. Type: object (or Expression with resultType string).
        /// </summary>
        public readonly object? ColumnName;
        /// <summary>
        /// The default value of the column. Type: object (or Expression with resultType string).
        /// </summary>
        public readonly object? DefaultValue;

        [OutputConstructor]
        private DWCopyCommandDefaultValueResponse(
            object? columnName,

            object? defaultValue)
        {
            ColumnName = columnName;
            DefaultValue = defaultValue;
        }
    }
}
