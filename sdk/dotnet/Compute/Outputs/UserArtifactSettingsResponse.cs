// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Additional settings for the VM app that contains the target package and config file name when it is deployed to target VM or VM scale set.
    /// </summary>
    [OutputType]
    public sealed class UserArtifactSettingsResponse
    {
        /// <summary>
        /// Optional. The name to assign the downloaded config file on the VM. This is limited to 4096 characters. If not specified, the config file will be named the Gallery Application name appended with "_config".
        /// </summary>
        public readonly string? ConfigFileName;
        /// <summary>
        /// Optional. The name to assign the downloaded package file on the VM. This is limited to 4096 characters. If not specified, the package file will be named the same as the Gallery Application name.
        /// </summary>
        public readonly string? PackageFileName;
        /// <summary>
        /// Optional. The action to be taken with regards to install/update/remove of the gallery application in the event of a reboot.
        /// </summary>
        public readonly string? ScriptBehaviorAfterReboot;

        [OutputConstructor]
        private UserArtifactSettingsResponse(
            string? configFileName,

            string? packageFileName,

            string? scriptBehaviorAfterReboot)
        {
            ConfigFileName = configFileName;
            PackageFileName = packageFileName;
            ScriptBehaviorAfterReboot = scriptBehaviorAfterReboot;
        }
    }
}
