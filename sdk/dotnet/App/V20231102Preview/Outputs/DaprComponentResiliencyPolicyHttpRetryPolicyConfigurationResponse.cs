// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20231102Preview.Outputs
{

    /// <summary>
    /// Dapr Component Resiliency Policy HTTP Retry Policy Configuration.
    /// </summary>
    [OutputType]
    public sealed class DaprComponentResiliencyPolicyHttpRetryPolicyConfigurationResponse
    {
        /// <summary>
        /// The optional maximum number of retries
        /// </summary>
        public readonly int? MaxRetries;
        /// <summary>
        /// The optional retry backoff configuration
        /// </summary>
        public readonly Outputs.DaprComponentResiliencyPolicyHttpRetryBackOffConfigurationResponse? RetryBackOff;

        [OutputConstructor]
        private DaprComponentResiliencyPolicyHttpRetryPolicyConfigurationResponse(
            int? maxRetries,

            Outputs.DaprComponentResiliencyPolicyHttpRetryBackOffConfigurationResponse? retryBackOff)
        {
            MaxRetries = maxRetries;
            RetryBackOff = retryBackOff;
        }
    }
}
