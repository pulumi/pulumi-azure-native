// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppComplianceAutomation.Outputs
{

    /// <summary>
    /// A class represent the certification record synchronized from app compliance.
    /// </summary>
    [OutputType]
    public sealed class CertSyncRecordResponse
    {
        /// <summary>
        /// Indicates the status of certification process.
        /// </summary>
        public readonly string? CertificationStatus;
        /// <summary>
        /// The control records list to be synchronized.
        /// </summary>
        public readonly ImmutableArray<Outputs.ControlSyncRecordResponse> Controls;
        /// <summary>
        /// Indicates the status of compliance process.
        /// </summary>
        public readonly string? IngestionStatus;
        /// <summary>
        /// The offerGuid which mapping to the reports.
        /// </summary>
        public readonly string? OfferGuid;

        [OutputConstructor]
        private CertSyncRecordResponse(
            string? certificationStatus,

            ImmutableArray<Outputs.ControlSyncRecordResponse> controls,

            string? ingestionStatus,

            string? offerGuid)
        {
            CertificationStatus = certificationStatus;
            Controls = controls;
            IngestionStatus = ingestionStatus;
            OfferGuid = offerGuid;
        }
    }
}
