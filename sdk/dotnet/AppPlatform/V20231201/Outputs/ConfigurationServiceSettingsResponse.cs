// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.V20231201.Outputs
{

    /// <summary>
    /// The settings of Application Configuration Service.
    /// </summary>
    [OutputType]
    public sealed class ConfigurationServiceSettingsResponse
    {
        /// <summary>
        /// Property of git environment.
        /// </summary>
        public readonly Outputs.ConfigurationServiceGitPropertyResponse? GitProperty;

        [OutputConstructor]
        private ConfigurationServiceSettingsResponse(Outputs.ConfigurationServiceGitPropertyResponse? gitProperty)
        {
            GitProperty = gitProperty;
        }
    }
}
