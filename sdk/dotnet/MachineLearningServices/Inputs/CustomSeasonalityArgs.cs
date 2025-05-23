// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class CustomSeasonalityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Forecasting seasonality mode.
        /// Expected value is 'Custom'.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// [Required] Seasonality value.
        /// </summary>
        [Input("value", required: true)]
        public Input<int> Value { get; set; } = null!;

        public CustomSeasonalityArgs()
        {
        }
        public static new CustomSeasonalityArgs Empty => new CustomSeasonalityArgs();
    }
}
