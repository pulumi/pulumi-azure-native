// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceLinker.V20221101Preview.Outputs
{

    /// <summary>
    /// The authentication info when authType is servicePrincipal secret
    /// </summary>
    [OutputType]
    public sealed class ServicePrincipalSecretAuthInfoResponse
    {
        /// <summary>
        /// The authentication type.
        /// Expected value is 'servicePrincipalSecret'.
        /// </summary>
        public readonly string AuthType;
        /// <summary>
        /// ServicePrincipal application clientId for servicePrincipal auth.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Indicates whether to clean up previous operation when Linker is updating or deleting
        /// </summary>
        public readonly string? DeleteOrUpdateBehavior;
        /// <summary>
        /// Principal Id for servicePrincipal auth.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// Optional, this value specifies the Azure roles to be assigned. Automatically 
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// Secret for servicePrincipal auth.
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// Username created in the database which is mapped to a user in AAD.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private ServicePrincipalSecretAuthInfoResponse(
            string authType,

            string clientId,

            string? deleteOrUpdateBehavior,

            string principalId,

            ImmutableArray<string> roles,

            string secret,

            string? userName)
        {
            AuthType = authType;
            ClientId = clientId;
            DeleteOrUpdateBehavior = deleteOrUpdateBehavior;
            PrincipalId = principalId;
            Roles = roles;
            Secret = secret;
            UserName = userName;
        }
    }
}
