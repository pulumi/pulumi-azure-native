// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Outputs
{

    /// <summary>
    /// Runs the specified PowerShell script during the validation phase (Windows). Corresponds to Packer powershell provisioner. Exactly one of 'scriptUri' or 'inline' can be specified.
    /// </summary>
    [OutputType]
    public sealed class ImageTemplatePowerShellValidatorResponse
    {
        /// <summary>
        /// Array of PowerShell commands to execute
        /// </summary>
        public readonly ImmutableArray<string> Inline;
        /// <summary>
        /// Friendly Name to provide context on what this validation step does
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// If specified, the PowerShell script will be run with elevated privileges using the Local System user. Can only be true when the runElevated field above is set to true.
        /// </summary>
        public readonly bool? RunAsSystem;
        /// <summary>
        /// If specified, the PowerShell script will be run with elevated privileges
        /// </summary>
        public readonly bool? RunElevated;
        /// <summary>
        /// URI of the PowerShell script to be run for validation. It can be a github link, Azure Storage URI, etc
        /// </summary>
        public readonly string? ScriptUri;
        /// <summary>
        /// SHA256 checksum of the power shell script provided in the scriptUri field above
        /// </summary>
        public readonly string? Sha256Checksum;
        /// <summary>
        /// The type of validation you want to use on the Image. For example, "Shell" can be shell validation
        /// Expected value is 'PowerShell'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Valid exit codes for the PowerShell script. [Default: 0]
        /// </summary>
        public readonly ImmutableArray<int> ValidExitCodes;

        [OutputConstructor]
        private ImageTemplatePowerShellValidatorResponse(
            ImmutableArray<string> inline,

            string? name,

            bool? runAsSystem,

            bool? runElevated,

            string? scriptUri,

            string? sha256Checksum,

            string type,

            ImmutableArray<int> validExitCodes)
        {
            Inline = inline;
            Name = name;
            RunAsSystem = runAsSystem;
            RunElevated = runElevated;
            ScriptUri = scriptUri;
            Sha256Checksum = sha256Checksum;
            Type = type;
            ValidExitCodes = validExitCodes;
        }
    }
}
