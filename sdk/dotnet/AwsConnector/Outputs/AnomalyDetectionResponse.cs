// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of AnomalyDetection
    /// </summary>
    [OutputType]
    public sealed class AnomalyDetectionResponse
    {
        /// <summary>
        /// &lt;p&gt;Indicates whether anomaly mitigation is in progress.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.MitigationInEffectEnumEnumValueResponse? MitigationInEffect;
        /// <summary>
        /// &lt;p&gt;The latest anomaly detection result.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.AnomalyResultEnumEnumValueResponse? Result;

        [OutputConstructor]
        private AnomalyDetectionResponse(
            Outputs.MitigationInEffectEnumEnumValueResponse? mitigationInEffect,

            Outputs.AnomalyResultEnumEnumValueResponse? result)
        {
            MitigationInEffect = mitigationInEffect;
            Result = result;
        }
    }
}
