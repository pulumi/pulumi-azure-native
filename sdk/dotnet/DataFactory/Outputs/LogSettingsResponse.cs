// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Log settings.
    /// </summary>
    [OutputType]
    public sealed class LogSettingsResponse
    {
        /// <summary>
        /// Specifies settings for copy activity log.
        /// </summary>
        public readonly Outputs.CopyActivityLogSettingsResponse? CopyActivityLogSettings;
        /// <summary>
        /// Specifies whether to enable copy activity log. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? EnableCopyActivityLog;
        /// <summary>
        /// Log location settings customer needs to provide when enabling log.
        /// </summary>
        public readonly Outputs.LogLocationSettingsResponse LogLocationSettings;

        [OutputConstructor]
        private LogSettingsResponse(
            Outputs.CopyActivityLogSettingsResponse? copyActivityLogSettings,

            object? enableCopyActivityLog,

            Outputs.LogLocationSettingsResponse logLocationSettings)
        {
            CopyActivityLogSettings = copyActivityLogSettings;
            EnableCopyActivityLog = enableCopyActivityLog;
            LogLocationSettings = logLocationSettings;
        }
    }
}
