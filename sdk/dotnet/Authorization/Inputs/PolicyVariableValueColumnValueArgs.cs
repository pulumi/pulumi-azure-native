// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Inputs
{

    /// <summary>
    /// The name value tuple for this variable value column.
    /// </summary>
    public sealed class PolicyVariableValueColumnValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Column name for the variable value
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        /// <summary>
        /// Column value for the variable value; this can be an integer, double, boolean, null or a string.
        /// </summary>
        [Input("columnValue", required: true)]
        public Input<object> ColumnValue { get; set; } = null!;

        public PolicyVariableValueColumnValueArgs()
        {
        }
        public static new PolicyVariableValueColumnValueArgs Empty => new PolicyVariableValueColumnValueArgs();
    }
}
