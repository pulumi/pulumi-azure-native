// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// The GitHub action configuration.
    /// </summary>
    public sealed class GitHubActionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GitHub Action code configuration.
        /// </summary>
        [Input("codeConfiguration")]
        public Input<Inputs.GitHubActionCodeConfigurationArgs>? CodeConfiguration { get; set; }

        /// <summary>
        /// GitHub Action container configuration.
        /// </summary>
        [Input("containerConfiguration")]
        public Input<Inputs.GitHubActionContainerConfigurationArgs>? ContainerConfiguration { get; set; }

        /// <summary>
        /// Workflow option to determine whether the workflow file should be generated and written to the repository.
        /// </summary>
        [Input("generateWorkflowFile")]
        public Input<bool>? GenerateWorkflowFile { get; set; }

        /// <summary>
        /// This will help determine the workflow configuration to select.
        /// </summary>
        [Input("isLinux")]
        public Input<bool>? IsLinux { get; set; }

        public GitHubActionConfigurationArgs()
        {
        }
        public static new GitHubActionConfigurationArgs Empty => new GitHubActionConfigurationArgs();
    }
}
