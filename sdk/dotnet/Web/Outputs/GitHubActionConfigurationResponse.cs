// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// The GitHub action configuration.
    /// </summary>
    [OutputType]
    public sealed class GitHubActionConfigurationResponse
    {
        /// <summary>
        /// GitHub Action code configuration.
        /// </summary>
        public readonly Outputs.GitHubActionCodeConfigurationResponse? CodeConfiguration;
        /// <summary>
        /// GitHub Action container configuration.
        /// </summary>
        public readonly Outputs.GitHubActionContainerConfigurationResponse? ContainerConfiguration;
        /// <summary>
        /// Workflow option to determine whether the workflow file should be generated and written to the repository.
        /// </summary>
        public readonly bool? GenerateWorkflowFile;
        /// <summary>
        /// This will help determine the workflow configuration to select.
        /// </summary>
        public readonly bool? IsLinux;

        [OutputConstructor]
        private GitHubActionConfigurationResponse(
            Outputs.GitHubActionCodeConfigurationResponse? codeConfiguration,

            Outputs.GitHubActionContainerConfigurationResponse? containerConfiguration,

            bool? generateWorkflowFile,

            bool? isLinux)
        {
            CodeConfiguration = codeConfiguration;
            ContainerConfiguration = containerConfiguration;
            GenerateWorkflowFile = generateWorkflowFile;
            IsLinux = isLinux;
        }
    }
}
