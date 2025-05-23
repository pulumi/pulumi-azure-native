// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    /// <summary>
    /// Table properties
    /// </summary>
    [OutputType]
    public sealed class DatabaseTableResponse
    {
        /// <summary>
        /// Indicates whether table is empty or not
        /// </summary>
        public readonly bool HasRows;
        /// <summary>
        /// Schema-qualified name of the table
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private DatabaseTableResponse(
            bool hasRows,

            string name)
        {
            HasRows = hasRows;
            Name = name;
        }
    }
}
