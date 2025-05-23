// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class ActualCapacityInfoResponse
    {
        /// <summary>
        /// Gets or sets the number of instances (scale units) which have Failed provisioning state and have target group payload.
        /// </summary>
        public readonly int? Failed;
        /// <summary>
        /// Gets or sets the number of instances (scale units) which have Failed provisioning state but do not have target group payload.
        /// </summary>
        public readonly int? OutdatedFailed;
        /// <summary>
        /// Gets or sets the number of instances (scale units) which have Succeeded provisioning state but do not have target group payload.
        /// </summary>
        public readonly int? OutdatedSucceeded;
        /// <summary>
        /// Gets or sets the number of instances (scale units) which have Succeeded provisioning state and target group payload.
        /// </summary>
        public readonly int? Succeeded;
        /// <summary>
        /// Gets or sets the total number of instances (scale units) regardless of provisioning state or whether current group payload version matches the target group payload.
        /// </summary>
        public readonly int? Total;

        [OutputConstructor]
        private ActualCapacityInfoResponse(
            int? failed,

            int? outdatedFailed,

            int? outdatedSucceeded,

            int? succeeded,

            int? total)
        {
            Failed = failed;
            OutdatedFailed = outdatedFailed;
            OutdatedSucceeded = outdatedSucceeded;
            Succeeded = succeeded;
            Total = total;
        }
    }
}
