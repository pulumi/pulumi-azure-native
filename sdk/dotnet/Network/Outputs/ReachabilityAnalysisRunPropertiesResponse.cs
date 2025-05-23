// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Represents the Reachability Analysis Run properties.
    /// </summary>
    [OutputType]
    public sealed class ReachabilityAnalysisRunPropertiesResponse
    {
        public readonly string AnalysisResult;
        public readonly string? Description;
        public readonly string ErrorMessage;
        /// <summary>
        /// Intent information.
        /// </summary>
        public readonly Outputs.IntentContentResponse IntentContent;
        /// <summary>
        /// Id of the intent resource to run analysis on.
        /// </summary>
        public readonly string IntentId;
        /// <summary>
        /// Provisioning states of a resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private ReachabilityAnalysisRunPropertiesResponse(
            string analysisResult,

            string? description,

            string errorMessage,

            Outputs.IntentContentResponse intentContent,

            string intentId,

            string provisioningState)
        {
            AnalysisResult = analysisResult;
            Description = description;
            ErrorMessage = errorMessage;
            IntentContent = intentContent;
            IntentId = intentId;
            ProvisioningState = provisioningState;
        }
    }
}
