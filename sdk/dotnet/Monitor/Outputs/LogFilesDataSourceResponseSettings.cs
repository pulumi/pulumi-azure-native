// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// The log files specific settings.
    /// </summary>
    [OutputType]
    public sealed class LogFilesDataSourceResponseSettings
    {
        /// <summary>
        /// Text settings
        /// </summary>
        public readonly Outputs.LogFileSettingsResponseText? Text;

        [OutputConstructor]
        private LogFilesDataSourceResponseSettings(Outputs.LogFileSettingsResponseText? text)
        {
            Text = text;
        }
    }
}
