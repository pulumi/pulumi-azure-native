// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevOpsInfrastructure.V20240326Preview.Outputs
{

    /// <summary>
    /// Stateful profile meaning that the machines will be returned to the pool after running a job.
    /// </summary>
    [OutputType]
    public sealed class StatefulResponse
    {
        /// <summary>
        /// Discriminator property for AgentProfile.
        /// Expected value is 'Stateful'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// How long should stateful machines be kept around. The maximum is one week.
        /// </summary>
        public readonly string MaxAgentLifetime;
        /// <summary>
        /// Defines pool buffer.
        /// </summary>
        public readonly object? ResourcePredictions;

        [OutputConstructor]
        private StatefulResponse(
            string kind,

            string maxAgentLifetime,

            object? resourcePredictions)
        {
            Kind = kind;
            MaxAgentLifetime = maxAgentLifetime;
            ResourcePredictions = resourcePredictions;
        }
    }
}
