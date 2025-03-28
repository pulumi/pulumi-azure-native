// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Edge.V20240201Preview.Outputs
{

    /// <summary>
    /// Site properties
    /// </summary>
    [OutputType]
    public sealed class SitePropertiesResponse
    {
        /// <summary>
        /// AddressResource ArmId of Site resource
        /// </summary>
        public readonly string? AddressResourceId;
        /// <summary>
        /// Description of Site resource
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// displayName of Site resource
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Provisioning state of last operation
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private SitePropertiesResponse(
            string? addressResourceId,

            string? description,

            string? displayName,

            string provisioningState)
        {
            AddressResourceId = addressResourceId;
            Description = description;
            DisplayName = displayName;
            ProvisioningState = provisioningState;
        }
    }
}
