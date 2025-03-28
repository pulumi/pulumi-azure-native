// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.V20250115Preview.Outputs
{

    /// <summary>
    /// Connected Partner Resource Properties
    /// </summary>
    [OutputType]
    public sealed class ConnectedPartnerResourcePropertiesResponse
    {
        /// <summary>
        /// The azure resource Id of the resource.
        /// </summary>
        public readonly string? AzureResourceId;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Elastic resource name
        /// </summary>
        public readonly string? PartnerDeploymentName;
        /// <summary>
        /// URL of the resource in Elastic cloud.
        /// </summary>
        public readonly string? PartnerDeploymentUri;
        /// <summary>
        /// The hosting type of the resource.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ConnectedPartnerResourcePropertiesResponse(
            string? azureResourceId,

            string? location,

            string? partnerDeploymentName,

            string? partnerDeploymentUri,

            string? type)
        {
            AzureResourceId = azureResourceId;
            Location = location;
            PartnerDeploymentName = partnerDeploymentName;
            PartnerDeploymentUri = partnerDeploymentUri;
            Type = type;
        }
    }
}
