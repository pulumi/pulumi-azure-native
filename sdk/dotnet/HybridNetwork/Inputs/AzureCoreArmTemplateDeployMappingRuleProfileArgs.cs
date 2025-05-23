// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Azure template deploy mapping rule profile.
    /// </summary>
    public sealed class AzureCoreArmTemplateDeployMappingRuleProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application enablement.
        /// </summary>
        [Input("applicationEnablement")]
        public InputUnion<string, Pulumi.AzureNative.HybridNetwork.ApplicationEnablement>? ApplicationEnablement { get; set; }

        /// <summary>
        /// The template mapping rule profile.
        /// </summary>
        [Input("templateMappingRuleProfile")]
        public Input<Inputs.ArmTemplateMappingRuleProfileArgs>? TemplateMappingRuleProfile { get; set; }

        public AzureCoreArmTemplateDeployMappingRuleProfileArgs()
        {
        }
        public static new AzureCoreArmTemplateDeployMappingRuleProfileArgs Empty => new AzureCoreArmTemplateDeployMappingRuleProfileArgs();
    }
}
