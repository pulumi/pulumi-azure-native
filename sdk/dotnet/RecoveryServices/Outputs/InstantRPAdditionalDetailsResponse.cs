// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    [OutputType]
    public sealed class InstantRPAdditionalDetailsResponse
    {
        public readonly string? AzureBackupRGNamePrefix;
        public readonly string? AzureBackupRGNameSuffix;

        [OutputConstructor]
        private InstantRPAdditionalDetailsResponse(
            string? azureBackupRGNamePrefix,

            string? azureBackupRGNameSuffix)
        {
            AzureBackupRGNamePrefix = azureBackupRGNamePrefix;
            AzureBackupRGNameSuffix = azureBackupRGNameSuffix;
        }
    }
}
