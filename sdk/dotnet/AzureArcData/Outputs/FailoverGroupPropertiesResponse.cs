// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// The properties of a failover group resource.
    /// </summary>
    [OutputType]
    public sealed class FailoverGroupPropertiesResponse
    {
        /// <summary>
        /// The resource ID of the partner SQL managed instance.
        /// </summary>
        public readonly string PartnerManagedInstanceId;
        /// <summary>
        /// The provisioning state of the failover group resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The specifications of the failover group resource.
        /// </summary>
        public readonly Outputs.FailoverGroupSpecResponse Spec;
        /// <summary>
        /// The status of the failover group custom resource.
        /// </summary>
        public readonly object? Status;

        [OutputConstructor]
        private FailoverGroupPropertiesResponse(
            string partnerManagedInstanceId,

            string provisioningState,

            Outputs.FailoverGroupSpecResponse spec,

            object? status)
        {
            PartnerManagedInstanceId = partnerManagedInstanceId;
            ProvisioningState = provisioningState;
            Spec = spec;
            Status = status;
        }
    }
}
