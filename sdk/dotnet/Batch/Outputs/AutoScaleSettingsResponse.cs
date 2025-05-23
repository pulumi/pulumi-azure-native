// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    [OutputType]
    public sealed class AutoScaleSettingsResponse
    {
        /// <summary>
        /// If omitted, the default value is 15 minutes (PT15M).
        /// </summary>
        public readonly string? EvaluationInterval;
        public readonly string Formula;

        [OutputConstructor]
        private AutoScaleSettingsResponse(
            string? evaluationInterval,

            string formula)
        {
            EvaluationInterval = evaluationInterval;
            Formula = formula;
        }
    }
}
