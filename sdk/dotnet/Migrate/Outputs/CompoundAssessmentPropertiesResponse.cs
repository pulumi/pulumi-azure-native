// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// Properties of a compound assessment.
    /// </summary>
    [OutputType]
    public sealed class CompoundAssessmentPropertiesResponse
    {
        /// <summary>
        /// Details of the compound assessment.
        /// </summary>
        public readonly Outputs.CompoundAssessmentDetailsResponse Details;
        /// <summary>
        /// Fallback machine assessment ARM ID.
        /// </summary>
        public readonly string? FallbackMachineAssessmentArmId;
        /// <summary>
        /// The status of the last operation.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// ARM IDs of the target assessments.
        /// </summary>
        public readonly Outputs.TargetAssessmentArmIdsResponse TargetAssessmentArmIds;

        [OutputConstructor]
        private CompoundAssessmentPropertiesResponse(
            Outputs.CompoundAssessmentDetailsResponse details,

            string? fallbackMachineAssessmentArmId,

            string provisioningState,

            Outputs.TargetAssessmentArmIdsResponse targetAssessmentArmIds)
        {
            Details = details;
            FallbackMachineAssessmentArmId = fallbackMachineAssessmentArmId;
            ProvisioningState = provisioningState;
            TargetAssessmentArmIds = targetAssessmentArmIds;
        }
    }
}
