// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard.Inputs
{

    /// <summary>
    /// Grafana security settings
    /// </summary>
    public sealed class SecurityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true to execute the CSRF check even if the login cookie is not in a request (default false).
        /// </summary>
        [Input("csrfAlwaysCheck")]
        public Input<bool>? CsrfAlwaysCheck { get; set; }

        public SecurityArgs()
        {
        }
        public static new SecurityArgs Empty => new SecurityArgs();
    }
}
