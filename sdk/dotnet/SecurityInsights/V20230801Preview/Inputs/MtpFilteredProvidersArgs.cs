// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230801Preview.Inputs
{

    /// <summary>
    /// Represents the connector's Filtered providers
    /// </summary>
    public sealed class MtpFilteredProvidersArgs : global::Pulumi.ResourceArgs
    {
        [Input("alerts", required: true)]
        private InputList<Union<string, Pulumi.AzureNative.SecurityInsights.V20230801Preview.MtpProvider>>? _alerts;

        /// <summary>
        /// Alerts filtered providers. When filters are not applied, all alerts will stream through the MTP pipeline, still in private preview for all products EXCEPT MDA and MDI, which are in GA state.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.SecurityInsights.V20230801Preview.MtpProvider>> Alerts
        {
            get => _alerts ?? (_alerts = new InputList<Union<string, Pulumi.AzureNative.SecurityInsights.V20230801Preview.MtpProvider>>());
            set => _alerts = value;
        }

        public MtpFilteredProvidersArgs()
        {
        }
        public static new MtpFilteredProvidersArgs Empty => new MtpFilteredProvidersArgs();
    }
}
