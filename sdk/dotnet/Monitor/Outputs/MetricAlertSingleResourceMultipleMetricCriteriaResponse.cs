// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Specifies the metric alert criteria for a single resource that has multiple metric criteria.
    /// </summary>
    [OutputType]
    public sealed class MetricAlertSingleResourceMultipleMetricCriteriaResponse
    {
        /// <summary>
        /// The list of metric criteria for this 'all of' operation. 
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricCriteriaResponse> AllOf;
        /// <summary>
        /// specifies the type of the alert criteria.
        /// Expected value is 'Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria'.
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private MetricAlertSingleResourceMultipleMetricCriteriaResponse(
            ImmutableArray<Outputs.MetricCriteriaResponse> allOf,

            string odataType)
        {
            AllOf = allOf;
            OdataType = odataType;
        }
    }
}
