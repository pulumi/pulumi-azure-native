// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// Represents the reported properties of a remote support job.
    /// </summary>
    [OutputType]
    public sealed class RemoteSupportJobReportedPropertiesResponse
    {
        /// <summary>
        /// Deployment status of job.
        /// </summary>
        public readonly Outputs.EceActionStatusResponse DeploymentStatus;
        /// <summary>
        /// Optional settings for configuring the node for remote support.
        /// </summary>
        public readonly Outputs.RemoteSupportJobNodeSettingsResponse NodeSettings;
        /// <summary>
        /// The percentage of the job that is complete.
        /// </summary>
        public readonly int PercentComplete;
        /// <summary>
        /// Details of the remote support session.
        /// </summary>
        public readonly ImmutableArray<Outputs.RemoteSupportSessionResponse> SessionDetails;
        /// <summary>
        /// Validation status of job.
        /// </summary>
        public readonly Outputs.EceActionStatusResponse ValidationStatus;

        [OutputConstructor]
        private RemoteSupportJobReportedPropertiesResponse(
            Outputs.EceActionStatusResponse deploymentStatus,

            Outputs.RemoteSupportJobNodeSettingsResponse nodeSettings,

            int percentComplete,

            ImmutableArray<Outputs.RemoteSupportSessionResponse> sessionDetails,

            Outputs.EceActionStatusResponse validationStatus)
        {
            DeploymentStatus = deploymentStatus;
            NodeSettings = nodeSettings;
            PercentComplete = percentComplete;
            SessionDetails = sessionDetails;
            ValidationStatus = validationStatus;
        }
    }
}
