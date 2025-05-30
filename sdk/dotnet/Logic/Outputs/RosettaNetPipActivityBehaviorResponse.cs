// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.Outputs
{

    /// <summary>
    /// The integration account RosettaNet ProcessConfiguration activity behavior.
    /// </summary>
    [OutputType]
    public sealed class RosettaNetPipActivityBehaviorResponse
    {
        /// <summary>
        /// The value indicating whether the RosettaNet PIP is used for a single action.
        /// </summary>
        public readonly string ActionType;
        /// <summary>
        /// The value indicating whether authorization is required.
        /// </summary>
        public readonly bool IsAuthorizationRequired;
        /// <summary>
        /// The value indicating whether secured transport is required.
        /// </summary>
        public readonly bool IsSecuredTransportRequired;
        /// <summary>
        /// The value indicating whether non-repudiation is for origin and content.
        /// </summary>
        public readonly bool NonRepudiationOfOriginAndContent;
        /// <summary>
        /// The persistent confidentiality encryption scope.
        /// </summary>
        public readonly string PersistentConfidentialityScope;
        /// <summary>
        /// The value indicating whether the RosettaNet PIP communication is synchronous.
        /// </summary>
        public readonly string ResponseType;
        /// <summary>
        /// The value indicating retry count.
        /// </summary>
        public readonly int RetryCount;
        /// <summary>
        /// The time to perform in seconds.
        /// </summary>
        public readonly int TimeToPerformInSeconds;

        [OutputConstructor]
        private RosettaNetPipActivityBehaviorResponse(
            string actionType,

            bool isAuthorizationRequired,

            bool isSecuredTransportRequired,

            bool nonRepudiationOfOriginAndContent,

            string persistentConfidentialityScope,

            string responseType,

            int retryCount,

            int timeToPerformInSeconds)
        {
            ActionType = actionType;
            IsAuthorizationRequired = isAuthorizationRequired;
            IsSecuredTransportRequired = isSecuredTransportRequired;
            NonRepudiationOfOriginAndContent = nonRepudiationOfOriginAndContent;
            PersistentConfidentialityScope = persistentConfidentialityScope;
            ResponseType = responseType;
            RetryCount = retryCount;
            TimeToPerformInSeconds = timeToPerformInSeconds;
        }
    }
}
