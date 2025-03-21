// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Outputs
{

    /// <summary>
    /// The properties of a vSphere Replication (VR) addon
    /// </summary>
    [OutputType]
    public sealed class AddonVrPropertiesResponse
    {
        /// <summary>
        /// The type of private cloud addon
        /// Expected value is 'VR'.
        /// </summary>
        public readonly string AddonType;
        /// <summary>
        /// The state of the addon provisioning
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The vSphere Replication Server (VRS) count
        /// </summary>
        public readonly int VrsCount;

        [OutputConstructor]
        private AddonVrPropertiesResponse(
            string addonType,

            string provisioningState,

            int vrsCount)
        {
            AddonType = addonType;
            ProvisioningState = provisioningState;
            VrsCount = vrsCount;
        }
    }
}
