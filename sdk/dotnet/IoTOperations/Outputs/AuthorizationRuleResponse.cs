// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.Outputs
{

    /// <summary>
    /// AuthorizationConfig Rule Properties
    /// </summary>
    [OutputType]
    public sealed class AuthorizationRuleResponse
    {
        /// <summary>
        /// Give access to Broker methods and topics.
        /// </summary>
        public readonly ImmutableArray<Outputs.BrokerResourceRuleResponse> BrokerResources;
        /// <summary>
        /// Give access to clients based on the following properties.
        /// </summary>
        public readonly Outputs.PrincipalDefinitionResponse Principals;
        /// <summary>
        /// Give access to state store resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.StateStoreResourceRuleResponse> StateStoreResources;

        [OutputConstructor]
        private AuthorizationRuleResponse(
            ImmutableArray<Outputs.BrokerResourceRuleResponse> brokerResources,

            Outputs.PrincipalDefinitionResponse principals,

            ImmutableArray<Outputs.StateStoreResourceRuleResponse> stateStoreResources)
        {
            BrokerResources = brokerResources;
            Principals = principals;
            StateStoreResources = stateStoreResources;
        }
    }
}
