// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Outputs
{

    [OutputType]
    public sealed class FeatureStoreSettingsResponse
    {
        public readonly Outputs.ComputeRuntimeDtoResponse? ComputeRuntime;
        public readonly string? OfflineStoreConnectionName;
        public readonly string? OnlineStoreConnectionName;

        [OutputConstructor]
        private FeatureStoreSettingsResponse(
            Outputs.ComputeRuntimeDtoResponse? computeRuntime,

            string? offlineStoreConnectionName,

            string? onlineStoreConnectionName)
        {
            ComputeRuntime = computeRuntime;
            OfflineStoreConnectionName = offlineStoreConnectionName;
            OnlineStoreConnectionName = onlineStoreConnectionName;
        }
    }
}
