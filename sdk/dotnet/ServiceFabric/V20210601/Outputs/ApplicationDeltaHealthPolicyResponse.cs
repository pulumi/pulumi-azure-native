// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20210601.Outputs
{

    /// <summary>
    /// Defines a delta health policy used to evaluate the health of an application or one of its child entities when upgrading the cluster.
    /// </summary>
    [OutputType]
    public sealed class ApplicationDeltaHealthPolicyResponse
    {
        /// <summary>
        /// The delta health policy used by default to evaluate the health of a service type when upgrading the cluster.
        /// </summary>
        public readonly Outputs.ServiceTypeDeltaHealthPolicyResponse? DefaultServiceTypeDeltaHealthPolicy;
        /// <summary>
        /// The map with service type delta health policy per service type name. The map is empty by default.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ServiceTypeDeltaHealthPolicyResponse>? ServiceTypeDeltaHealthPolicies;

        [OutputConstructor]
        private ApplicationDeltaHealthPolicyResponse(
            Outputs.ServiceTypeDeltaHealthPolicyResponse? defaultServiceTypeDeltaHealthPolicy,

            ImmutableDictionary<string, Outputs.ServiceTypeDeltaHealthPolicyResponse>? serviceTypeDeltaHealthPolicies)
        {
            DefaultServiceTypeDeltaHealthPolicy = defaultServiceTypeDeltaHealthPolicy;
            ServiceTypeDeltaHealthPolicies = serviceTypeDeltaHealthPolicies;
        }
    }
}
