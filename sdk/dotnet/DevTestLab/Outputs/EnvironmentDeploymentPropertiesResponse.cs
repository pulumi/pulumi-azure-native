// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Outputs
{

    /// <summary>
    /// Properties of an environment deployment.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentDeploymentPropertiesResponse
    {
        /// <summary>
        /// The Azure Resource Manager template's identifier.
        /// </summary>
        public readonly string? ArmTemplateId;
        /// <summary>
        /// The parameters of the Azure Resource Manager template.
        /// </summary>
        public readonly ImmutableArray<Outputs.ArmTemplateParameterPropertiesResponse> Parameters;

        [OutputConstructor]
        private EnvironmentDeploymentPropertiesResponse(
            string? armTemplateId,

            ImmutableArray<Outputs.ArmTemplateParameterPropertiesResponse> parameters)
        {
            ArmTemplateId = armTemplateId;
            Parameters = parameters;
        }
    }
}
