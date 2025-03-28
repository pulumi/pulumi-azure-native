// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230701Preview.Outputs
{

    /// <summary>
    /// Google Cloud Platform auth section properties.
    /// </summary>
    [OutputType]
    public sealed class GCPAuthPropertiesResponse
    {
        /// <summary>
        /// The GCP project number.
        /// </summary>
        public readonly string ProjectNumber;
        /// <summary>
        /// The service account that is used to access the GCP project.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// The workload identity provider id that is used to gain access to the GCP project.
        /// </summary>
        public readonly string WorkloadIdentityProviderId;

        [OutputConstructor]
        private GCPAuthPropertiesResponse(
            string projectNumber,

            string serviceAccountEmail,

            string workloadIdentityProviderId)
        {
            ProjectNumber = projectNumber;
            ServiceAccountEmail = serviceAccountEmail;
            WorkloadIdentityProviderId = workloadIdentityProviderId;
        }
    }
}
