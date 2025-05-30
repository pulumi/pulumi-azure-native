// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConfidentialLedger.Outputs
{

    /// <summary>
    /// Additional Managed CCF properties.
    /// </summary>
    [OutputType]
    public sealed class ManagedCCFPropertiesResponse
    {
        /// <summary>
        /// Unique name for the Managed CCF.
        /// </summary>
        public readonly string AppName;
        /// <summary>
        /// Endpoint for calling Managed CCF Service.
        /// </summary>
        public readonly string AppUri;
        /// <summary>
        /// Deployment Type of Managed CCF
        /// </summary>
        public readonly Outputs.DeploymentTypeResponse? DeploymentType;
        /// <summary>
        /// Endpoint for accessing network identity.
        /// </summary>
        public readonly string IdentityServiceUri;
        /// <summary>
        /// List of member identity certificates for  Managed CCF
        /// </summary>
        public readonly ImmutableArray<Outputs.MemberIdentityCertificateResponse> MemberIdentityCertificates;
        /// <summary>
        /// Number of CCF nodes in the Managed CCF.
        /// </summary>
        public readonly int? NodeCount;
        /// <summary>
        /// Provisioning state of Managed CCF Resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Object representing RunningState for Managed CCF.
        /// </summary>
        public readonly string? RunningState;

        [OutputConstructor]
        private ManagedCCFPropertiesResponse(
            string appName,

            string appUri,

            Outputs.DeploymentTypeResponse? deploymentType,

            string identityServiceUri,

            ImmutableArray<Outputs.MemberIdentityCertificateResponse> memberIdentityCertificates,

            int? nodeCount,

            string provisioningState,

            string? runningState)
        {
            AppName = appName;
            AppUri = appUri;
            DeploymentType = deploymentType;
            IdentityServiceUri = identityServiceUri;
            MemberIdentityCertificates = memberIdentityCertificates;
            NodeCount = nodeCount;
            ProvisioningState = provisioningState;
            RunningState = runningState;
        }
    }
}
