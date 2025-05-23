// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceLinker.Outputs
{

    /// <summary>
    /// The dryrun parameters for creation or update a linker
    /// </summary>
    [OutputType]
    public sealed class CreateOrUpdateDryrunParametersResponse
    {
        /// <summary>
        /// The name of action for you dryrun job.
        /// Expected value is 'createOrUpdate'.
        /// </summary>
        public readonly string ActionName;
        /// <summary>
        /// The authentication type.
        /// </summary>
        public readonly object? AuthInfo;
        /// <summary>
        /// The application client type
        /// </summary>
        public readonly string? ClientType;
        /// <summary>
        /// The connection information consumed by applications, including secrets, connection strings.
        /// </summary>
        public readonly Outputs.ConfigurationInfoResponse? ConfigurationInfo;
        /// <summary>
        /// The provisioning state. 
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The network solution.
        /// </summary>
        public readonly Outputs.PublicNetworkSolutionResponse? PublicNetworkSolution;
        /// <summary>
        /// connection scope in source service.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// An option to store secret value in secure place
        /// </summary>
        public readonly Outputs.SecretStoreResponse? SecretStore;
        /// <summary>
        /// The target service properties
        /// </summary>
        public readonly object? TargetService;
        /// <summary>
        /// The VNet solution.
        /// </summary>
        public readonly Outputs.VNetSolutionResponse? VNetSolution;

        [OutputConstructor]
        private CreateOrUpdateDryrunParametersResponse(
            string actionName,

            object? authInfo,

            string? clientType,

            Outputs.ConfigurationInfoResponse? configurationInfo,

            string provisioningState,

            Outputs.PublicNetworkSolutionResponse? publicNetworkSolution,

            string? scope,

            Outputs.SecretStoreResponse? secretStore,

            object? targetService,

            Outputs.VNetSolutionResponse? vNetSolution)
        {
            ActionName = actionName;
            AuthInfo = authInfo;
            ClientType = clientType;
            ConfigurationInfo = configurationInfo;
            ProvisioningState = provisioningState;
            PublicNetworkSolution = publicNetworkSolution;
            Scope = scope;
            SecretStore = secretStore;
            TargetService = targetService;
            VNetSolution = vNetSolution;
        }
    }
}
