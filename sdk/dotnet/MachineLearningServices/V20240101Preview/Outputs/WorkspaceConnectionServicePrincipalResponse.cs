// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240101Preview.Outputs
{

    [OutputType]
    public sealed class WorkspaceConnectionServicePrincipalResponse
    {
        public readonly string? ClientId;
        public readonly string? ClientSecret;
        public readonly string? TenantId;

        [OutputConstructor]
        private WorkspaceConnectionServicePrincipalResponse(
            string? clientId,

            string? clientSecret,

            string? tenantId)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            TenantId = tenantId;
        }
    }
}
