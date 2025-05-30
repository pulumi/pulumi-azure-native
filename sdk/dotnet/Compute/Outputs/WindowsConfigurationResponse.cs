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
    /// Specifies Windows operating system settings on the virtual machine.
    /// </summary>
    [OutputType]
    public sealed class WindowsConfigurationResponse
    {
        /// <summary>
        /// Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdditionalUnattendContentResponse> AdditionalUnattendContent;
        /// <summary>
        /// Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is true. For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.
        /// </summary>
        public readonly bool? EnableAutomaticUpdates;
        /// <summary>
        /// Indicates whether VMAgent Platform Updates are enabled for the Windows Virtual Machine.
        /// </summary>
        public readonly bool EnableVMAgentPlatformUpdates;
        /// <summary>
        /// [Preview Feature] Specifies settings related to VM Guest Patching on Windows.
        /// </summary>
        public readonly Outputs.PatchSettingsResponse? PatchSettings;
        /// <summary>
        /// Indicates whether virtual machine agent should be provisioned on the virtual machine. When this property is not specified in the request body, it is set to true by default. This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
        /// </summary>
        public readonly bool? ProvisionVMAgent;
        /// <summary>
        /// Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time". Possible values can be [TimeZoneInfo.Id](https://docs.microsoft.com/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id) value from time zones returned by [TimeZoneInfo.GetSystemTimeZones](https://docs.microsoft.com/dotnet/api/system.timezoneinfo.getsystemtimezones).
        /// </summary>
        public readonly string? TimeZone;
        /// <summary>
        /// Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
        /// </summary>
        public readonly Outputs.WinRMConfigurationResponse? WinRM;

        [OutputConstructor]
        private WindowsConfigurationResponse(
            ImmutableArray<Outputs.AdditionalUnattendContentResponse> additionalUnattendContent,

            bool? enableAutomaticUpdates,

            bool enableVMAgentPlatformUpdates,

            Outputs.PatchSettingsResponse? patchSettings,

            bool? provisionVMAgent,

            string? timeZone,

            Outputs.WinRMConfigurationResponse? winRM)
        {
            AdditionalUnattendContent = additionalUnattendContent;
            EnableAutomaticUpdates = enableAutomaticUpdates;
            EnableVMAgentPlatformUpdates = enableVMAgentPlatformUpdates;
            PatchSettings = patchSettings;
            ProvisionVMAgent = provisionVMAgent;
            TimeZone = timeZone;
            WinRM = winRM;
        }
    }
}
