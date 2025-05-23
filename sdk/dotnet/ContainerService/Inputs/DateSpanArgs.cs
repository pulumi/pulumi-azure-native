// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// For example, between '2022-12-23' and '2023-01-05'.
    /// </summary>
    public sealed class DateSpanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end date of the date span.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// The start date of the date span.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public DateSpanArgs()
        {
        }
        public static new DateSpanArgs Empty => new DateSpanArgs();
    }
}
