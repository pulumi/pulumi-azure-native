// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    /// <summary>
    /// Source control configuration for an app.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:web:WebAppSourceControl")]
    public partial class WebAppSourceControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Name of branch to use for deployment.
        /// </summary>
        [Output("branch")]
        public Output<string?> Branch { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable deployment rollback; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Output("deploymentRollbackEnabled")]
        public Output<bool?> DeploymentRollbackEnabled { get; private set; } = null!;

        /// <summary>
        /// If GitHub Action is selected, than the associated configuration.
        /// </summary>
        [Output("gitHubActionConfiguration")]
        public Output<Outputs.GitHubActionConfigurationResponse?> GitHubActionConfiguration { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if this is deployed via GitHub action.
        /// </summary>
        [Output("isGitHubAction")]
        public Output<bool?> IsGitHubAction { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to limit to manual integration; &lt;code&gt;false&lt;/code&gt; to enable continuous integration (which configures webhooks into online repos like GitHub).
        /// </summary>
        [Output("isManualIntegration")]
        public Output<bool?> IsManualIntegration { get; private set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; for a Mercurial repository; &lt;code&gt;false&lt;/code&gt; for a Git repository.
        /// </summary>
        [Output("isMercurial")]
        public Output<bool?> IsMercurial { get; private set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Repository or source control URL.
        /// </summary>
        [Output("repoUrl")]
        public Output<string?> RepoUrl { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppSourceControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppSourceControl(string name, WebAppSourceControlArgs args, CustomResourceOptions? options = null)
            : base("azure-native:web:WebAppSourceControl", name, args ?? new WebAppSourceControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppSourceControl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:web:WebAppSourceControl", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:web/v20150801:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20160801:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20180201:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20181101:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20190801:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200601:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200901:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201001:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201201:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210101:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210115:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210201:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210301:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220301:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220901:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20230101:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20231201:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20240401:WebAppSourceControl" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20241101:WebAppSourceControl" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAppSourceControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppSourceControl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppSourceControl(name, id, options);
        }
    }

    public sealed class WebAppSourceControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of branch to use for deployment.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable deployment rollback; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("deploymentRollbackEnabled")]
        public Input<bool>? DeploymentRollbackEnabled { get; set; }

        /// <summary>
        /// If GitHub Action is selected, than the associated configuration.
        /// </summary>
        [Input("gitHubActionConfiguration")]
        public Input<Inputs.GitHubActionConfigurationArgs>? GitHubActionConfiguration { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if this is deployed via GitHub action.
        /// </summary>
        [Input("isGitHubAction")]
        public Input<bool>? IsGitHubAction { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to limit to manual integration; &lt;code&gt;false&lt;/code&gt; to enable continuous integration (which configures webhooks into online repos like GitHub).
        /// </summary>
        [Input("isManualIntegration")]
        public Input<bool>? IsManualIntegration { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; for a Mercurial repository; &lt;code&gt;false&lt;/code&gt; for a Git repository.
        /// </summary>
        [Input("isMercurial")]
        public Input<bool>? IsMercurial { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Repository or source control URL.
        /// </summary>
        [Input("repoUrl")]
        public Input<string>? RepoUrl { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public WebAppSourceControlArgs()
        {
        }
        public static new WebAppSourceControlArgs Empty => new WebAppSourceControlArgs();
    }
}
