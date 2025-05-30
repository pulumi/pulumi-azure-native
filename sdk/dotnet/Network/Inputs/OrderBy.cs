// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Describes a column to sort
    /// </summary>
    public sealed class OrderBy : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Describes the actual column name to sort by
        /// </summary>
        [Input("field")]
        public string? Field { get; set; }

        /// <summary>
        /// Describes if results should be in ascending/descending order
        /// </summary>
        [Input("order")]
        public Union<string, Pulumi.AzureNative.Network.FirewallPolicyIDPSQuerySortOrder>? Order { get; set; }

        public OrderBy()
        {
        }
        public static new OrderBy Empty => new OrderBy();
    }
}
