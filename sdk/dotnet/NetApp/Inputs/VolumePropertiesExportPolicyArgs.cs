// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.Inputs
{

    /// <summary>
    /// Set of export policy rules
    /// </summary>
    public sealed class VolumePropertiesExportPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.ExportPolicyRuleArgs>? _rules;

        /// <summary>
        /// Export policy rule
        /// </summary>
        public InputList<Inputs.ExportPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ExportPolicyRuleArgs>());
            set => _rules = value;
        }

        public VolumePropertiesExportPolicyArgs()
        {
        }
        public static new VolumePropertiesExportPolicyArgs Empty => new VolumePropertiesExportPolicyArgs();
    }
}
