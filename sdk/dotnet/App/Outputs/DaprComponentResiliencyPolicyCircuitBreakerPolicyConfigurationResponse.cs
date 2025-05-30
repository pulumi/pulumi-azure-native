// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Dapr Component Resiliency Policy Circuit Breaker Policy Configuration.
    /// </summary>
    [OutputType]
    public sealed class DaprComponentResiliencyPolicyCircuitBreakerPolicyConfigurationResponse
    {
        /// <summary>
        /// The number of consecutive errors before the circuit is opened.
        /// </summary>
        public readonly int? ConsecutiveErrors;
        /// <summary>
        /// The optional interval in seconds after which the error count resets to 0. An interval of 0 will never reset. If not specified, the timeoutInSeconds value will be used.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// The interval in seconds until a retry attempt is made after the circuit is opened.
        /// </summary>
        public readonly int? TimeoutInSeconds;

        [OutputConstructor]
        private DaprComponentResiliencyPolicyCircuitBreakerPolicyConfigurationResponse(
            int? consecutiveErrors,

            int? intervalInSeconds,

            int? timeoutInSeconds)
        {
            ConsecutiveErrors = consecutiveErrors;
            IntervalInSeconds = intervalInSeconds;
            TimeoutInSeconds = timeoutInSeconds;
        }
    }
}
