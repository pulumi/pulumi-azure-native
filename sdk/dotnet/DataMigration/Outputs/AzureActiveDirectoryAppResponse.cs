// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    /// <summary>
    /// Azure Active Directory Application
    /// </summary>
    [OutputType]
    public sealed class AzureActiveDirectoryAppResponse
    {
        /// <summary>
        /// Key used to authenticate to the Azure Active Directory Application
        /// </summary>
        public readonly string? AppKey;
        /// <summary>
        /// Application ID of the Azure Active Directory Application
        /// </summary>
        public readonly string? ApplicationId;
        /// <summary>
        /// Ignore checking azure permissions on the AAD app
        /// </summary>
        public readonly bool? IgnoreAzurePermissions;
        /// <summary>
        /// Tenant id of the customer
        /// </summary>
        public readonly string? TenantId;

        [OutputConstructor]
        private AzureActiveDirectoryAppResponse(
            string? appKey,

            string? applicationId,

            bool? ignoreAzurePermissions,

            string? tenantId)
        {
            AppKey = appKey;
            ApplicationId = applicationId;
            IgnoreAzurePermissions = ignoreAzurePermissions;
            TenantId = tenantId;
        }
    }
}
