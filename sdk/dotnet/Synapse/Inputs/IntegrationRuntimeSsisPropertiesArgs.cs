// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Inputs
{

    /// <summary>
    /// SSIS properties for managed integration runtime.
    /// </summary>
    public sealed class IntegrationRuntimeSsisPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Catalog information for managed dedicated integration runtime.
        /// </summary>
        [Input("catalogInfo")]
        public Input<Inputs.IntegrationRuntimeSsisCatalogInfoArgs>? CatalogInfo { get; set; }

        /// <summary>
        /// Custom setup script properties for a managed dedicated integration runtime.
        /// </summary>
        [Input("customSetupScriptProperties")]
        public Input<Inputs.IntegrationRuntimeCustomSetupScriptPropertiesArgs>? CustomSetupScriptProperties { get; set; }

        /// <summary>
        /// Data proxy properties for a managed dedicated integration runtime.
        /// </summary>
        [Input("dataProxyProperties")]
        public Input<Inputs.IntegrationRuntimeDataProxyPropertiesArgs>? DataProxyProperties { get; set; }

        /// <summary>
        /// The edition for the SSIS Integration Runtime
        /// </summary>
        [Input("edition")]
        public InputUnion<string, Pulumi.AzureNative.Synapse.IntegrationRuntimeEdition>? Edition { get; set; }

        [Input("expressCustomSetupProperties")]
        private InputList<object>? _expressCustomSetupProperties;

        /// <summary>
        /// Custom setup without script properties for a SSIS integration runtime.
        /// </summary>
        public InputList<object> ExpressCustomSetupProperties
        {
            get => _expressCustomSetupProperties ?? (_expressCustomSetupProperties = new InputList<object>());
            set => _expressCustomSetupProperties = value;
        }

        /// <summary>
        /// License type for bringing your own license scenario.
        /// </summary>
        [Input("licenseType")]
        public InputUnion<string, Pulumi.AzureNative.Synapse.IntegrationRuntimeLicenseType>? LicenseType { get; set; }

        public IntegrationRuntimeSsisPropertiesArgs()
        {
        }
        public static new IntegrationRuntimeSsisPropertiesArgs Empty => new IntegrationRuntimeSsisPropertiesArgs();
    }
}
