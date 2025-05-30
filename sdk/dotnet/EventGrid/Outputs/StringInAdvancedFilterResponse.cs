// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Outputs
{

    /// <summary>
    /// StringIn Advanced Filter.
    /// </summary>
    [OutputType]
    public sealed class StringInAdvancedFilterResponse
    {
        /// <summary>
        /// The field/property in the event based on which you want to filter.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        /// Expected value is 'StringIn'.
        /// </summary>
        public readonly string OperatorType;
        /// <summary>
        /// The set of filter values.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private StringInAdvancedFilterResponse(
            string? key,

            string operatorType,

            ImmutableArray<string> values)
        {
            Key = key;
            OperatorType = operatorType;
            Values = values;
        }
    }
}
