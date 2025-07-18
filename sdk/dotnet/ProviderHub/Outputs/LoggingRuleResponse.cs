// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    [OutputType]
    public sealed class LoggingRuleResponse
    {
        /// <summary>
        /// The action.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The detail level.
        /// </summary>
        public readonly string DetailLevel;
        /// <summary>
        /// The direction.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The hidden property paths.
        /// </summary>
        public readonly Outputs.LoggingRuleResponseHiddenPropertyPaths? HiddenPropertyPaths;

        [OutputConstructor]
        private LoggingRuleResponse(
            string action,

            string detailLevel,

            string direction,

            Outputs.LoggingRuleResponseHiddenPropertyPaths? hiddenPropertyPaths)
        {
            Action = action;
            DetailLevel = detailLevel;
            Direction = direction;
            HiddenPropertyPaths = hiddenPropertyPaths;
        }
    }
}
