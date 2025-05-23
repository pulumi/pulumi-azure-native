// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Storage properties
    /// </summary>
    public sealed class ConnectedEnvironmentStoragePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure file properties
        /// </summary>
        [Input("azureFile")]
        public Input<Inputs.AzureFilePropertiesArgs>? AzureFile { get; set; }

        public ConnectedEnvironmentStoragePropertiesArgs()
        {
        }
        public static new ConnectedEnvironmentStoragePropertiesArgs Empty => new ConnectedEnvironmentStoragePropertiesArgs();
    }
}
