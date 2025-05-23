// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Inputs
{

    /// <summary>
    /// IsNullOrUndefined Advanced Filter.
    /// </summary>
    public sealed class IsNullOrUndefinedAdvancedFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field/property in the event based on which you want to filter.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        /// Expected value is 'IsNullOrUndefined'.
        /// </summary>
        [Input("operatorType", required: true)]
        public Input<string> OperatorType { get; set; } = null!;

        public IsNullOrUndefinedAdvancedFilterArgs()
        {
        }
        public static new IsNullOrUndefinedAdvancedFilterArgs Empty => new IsNullOrUndefinedAdvancedFilterArgs();
    }
}
