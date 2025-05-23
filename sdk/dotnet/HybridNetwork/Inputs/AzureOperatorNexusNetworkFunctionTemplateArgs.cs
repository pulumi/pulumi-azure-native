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
    /// Azure Operator Distributed Services network function template.
    /// </summary>
    public sealed class AzureOperatorNexusNetworkFunctionTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("networkFunctionApplications")]
        private InputList<Union<Inputs.AzureOperatorNexusNetworkFunctionArmTemplateApplicationArgs, Inputs.AzureOperatorNexusNetworkFunctionImageApplicationArgs>>? _networkFunctionApplications;

        /// <summary>
        /// Network function applications.
        /// </summary>
        public InputList<Union<Inputs.AzureOperatorNexusNetworkFunctionArmTemplateApplicationArgs, Inputs.AzureOperatorNexusNetworkFunctionImageApplicationArgs>> NetworkFunctionApplications
        {
            get => _networkFunctionApplications ?? (_networkFunctionApplications = new InputList<Union<Inputs.AzureOperatorNexusNetworkFunctionArmTemplateApplicationArgs, Inputs.AzureOperatorNexusNetworkFunctionImageApplicationArgs>>());
            set => _networkFunctionApplications = value;
        }

        /// <summary>
        /// The network function type.
        /// Expected value is 'AzureOperatorNexus'.
        /// </summary>
        [Input("nfviType", required: true)]
        public Input<string> NfviType { get; set; } = null!;

        public AzureOperatorNexusNetworkFunctionTemplateArgs()
        {
        }
        public static new AzureOperatorNexusNetworkFunctionTemplateArgs Empty => new AzureOperatorNexusNetworkFunctionTemplateArgs();
    }
}
