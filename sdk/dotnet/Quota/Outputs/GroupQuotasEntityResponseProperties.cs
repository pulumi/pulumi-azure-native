// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Quota.Outputs
{

    [OutputType]
    public sealed class GroupQuotasEntityResponseProperties
    {
        /// <summary>
        /// Display name of the GroupQuota entity.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Provisioning state of the operation.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private GroupQuotasEntityResponseProperties(
            string? displayName,

            string provisioningState)
        {
            DisplayName = displayName;
            ProvisioningState = provisioningState;
        }
    }
}
