// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// Distcp settings.
    /// </summary>
    public sealed class DistcpSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Distcp options. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("distcpOptions")]
        public Input<object>? DistcpOptions { get; set; }

        /// <summary>
        /// Specifies the Yarn ResourceManager endpoint. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("resourceManagerEndpoint", required: true)]
        public Input<object> ResourceManagerEndpoint { get; set; } = null!;

        /// <summary>
        /// Specifies an existing folder path which will be used to store temp Distcp command script. The script file is generated by ADF and will be removed after Copy job finished. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("tempScriptPath", required: true)]
        public Input<object> TempScriptPath { get; set; } = null!;

        public DistcpSettingsArgs()
        {
        }
        public static new DistcpSettingsArgs Empty => new DistcpSettingsArgs();
    }
}
