// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240401.Outputs
{

    /// <summary>
    /// Forecast horizon determined automatically by system.
    /// </summary>
    [OutputType]
    public sealed class AutoForecastHorizonResponse
    {
        /// <summary>
        /// Enum to determine forecast horizon selection mode.
        /// Expected value is 'Auto'.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private AutoForecastHorizonResponse(string mode)
        {
            Mode = mode;
        }
    }
}
