// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20231001Preview.Outputs
{

    [OutputType]
    public sealed class CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria
    {
        /// <summary>
        /// type of connectivity
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Queries for checking connectivity
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria(
            string? type,

            ImmutableArray<string> value)
        {
            Type = type;
            Value = value;
        }
    }
}
