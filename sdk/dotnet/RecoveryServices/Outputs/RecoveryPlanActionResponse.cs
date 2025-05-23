// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Recovery plan action details.
    /// </summary>
    [OutputType]
    public sealed class RecoveryPlanActionResponse
    {
        /// <summary>
        /// The action name.
        /// </summary>
        public readonly string ActionName;
        /// <summary>
        /// The custom details.
        /// </summary>
        public readonly object CustomDetails;
        /// <summary>
        /// The list of failover directions.
        /// </summary>
        public readonly ImmutableArray<string> FailoverDirections;
        /// <summary>
        /// The list of failover types.
        /// </summary>
        public readonly ImmutableArray<string> FailoverTypes;

        [OutputConstructor]
        private RecoveryPlanActionResponse(
            string actionName,

            object customDetails,

            ImmutableArray<string> failoverDirections,

            ImmutableArray<string> failoverTypes)
        {
            ActionName = actionName;
            CustomDetails = customDetails;
            FailoverDirections = failoverDirections;
            FailoverTypes = failoverTypes;
        }
    }
}
