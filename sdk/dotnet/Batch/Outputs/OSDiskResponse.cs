// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    [OutputType]
    public sealed class OSDiskResponse
    {
        public readonly Outputs.DiffDiskSettingsResponse? EphemeralOSDiskSettings;

        [OutputConstructor]
        private OSDiskResponse(Outputs.DiffDiskSettingsResponse? ephemeralOSDiskSettings)
        {
            EphemeralOSDiskSettings = ephemeralOSDiskSettings;
        }
    }
}
