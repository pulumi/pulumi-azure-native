// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Inputs
{

    /// <summary>
    /// The settings of config server.
    /// </summary>
    public sealed class ConfigServerSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property of git environment.
        /// </summary>
        [Input("gitProperty")]
        public Input<Inputs.ConfigServerGitPropertyArgs>? GitProperty { get; set; }

        public ConfigServerSettingsArgs()
        {
        }
        public static new ConfigServerSettingsArgs Empty => new ConfigServerSettingsArgs();
    }
}
