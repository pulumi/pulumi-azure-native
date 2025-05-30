// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Labeling MLAssist configuration definition when MLAssist is disabled
    /// </summary>
    [OutputType]
    public sealed class MLAssistConfigurationDisabledResponse
    {
        /// <summary>
        /// 
        /// Expected value is 'Disabled'.
        /// </summary>
        public readonly string MlAssist;

        [OutputConstructor]
        private MLAssistConfigurationDisabledResponse(string mlAssist)
        {
            MlAssist = mlAssist;
        }
    }
}
