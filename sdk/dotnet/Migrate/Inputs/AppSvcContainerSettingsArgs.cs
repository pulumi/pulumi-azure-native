// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// App service container settings.
    /// </summary>
    public sealed class AppSvcContainerSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the isolation required.
        /// </summary>
        [Input("isolationRequired", required: true)]
        public Input<bool> IsolationRequired { get; set; } = null!;

        public AppSvcContainerSettingsArgs()
        {
        }
        public static new AppSvcContainerSettingsArgs Empty => new AppSvcContainerSettingsArgs();
    }
}
