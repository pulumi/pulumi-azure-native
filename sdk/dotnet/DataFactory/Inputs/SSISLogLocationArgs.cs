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
    /// SSIS package execution log location
    /// </summary>
    public sealed class SSISLogLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The package execution log access credential.
        /// </summary>
        [Input("accessCredential")]
        public Input<Inputs.SSISAccessCredentialArgs>? AccessCredential { get; set; }

        /// <summary>
        /// The SSIS package execution log path. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("logPath", required: true)]
        public Input<object> LogPath { get; set; } = null!;

        /// <summary>
        /// Specifies the interval to refresh log. The default interval is 5 minutes. Type: string (or Expression with resultType string), pattern: ((\d+)\.)?(\d\d):(60|([0-5][0-9])):(60|([0-5][0-9])).
        /// </summary>
        [Input("logRefreshInterval")]
        public Input<object>? LogRefreshInterval { get; set; }

        /// <summary>
        /// The type of SSIS log location.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.SsisLogLocationType> Type { get; set; } = null!;

        public SSISLogLocationArgs()
        {
        }
        public static new SSISLogLocationArgs Empty => new SSISLogLocationArgs();
    }
}
