// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Specifies additional XML formatted information that can be included in the
    /// Unattend.xml file, which is used by Windows Setup. Contents are defined by
    /// setting name, component name, and the pass in which the content is applied.
    /// </summary>
    [OutputType]
    public sealed class AdditionalUnattendContentResponse
    {
        /// <summary>
        /// The component name. Currently, the only allowable value is
        /// Microsoft-Windows-Shell-Setup.
        /// </summary>
        public readonly string? ComponentName;
        /// <summary>
        /// The pass name. Currently, the only allowable value is OobeSystem.
        /// </summary>
        public readonly string? PassName;
        /// <summary>
        /// Specifies the name of the setting to which the content applies. Possible values
        /// are: FirstLogonCommands and AutoLogon.
        /// </summary>
        public readonly string? SettingName;

        [OutputConstructor]
        private AdditionalUnattendContentResponse(
            string? componentName,

            string? passName,

            string? settingName)
        {
            ComponentName = componentName;
            PassName = passName;
            SettingName = settingName;
        }
    }
}
