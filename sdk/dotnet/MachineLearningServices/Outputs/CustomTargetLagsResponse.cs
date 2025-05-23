// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class CustomTargetLagsResponse
    {
        /// <summary>
        /// Target lags selection modes.
        /// Expected value is 'Custom'.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// [Required] Set target lags values.
        /// </summary>
        public readonly ImmutableArray<int> Values;

        [OutputConstructor]
        private CustomTargetLagsResponse(
            string mode,

            ImmutableArray<int> values)
        {
            Mode = mode;
            Values = values;
        }
    }
}
