// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration
{
    /// <summary>
    /// A project resource
    /// 
    /// Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-30.
    /// 
    /// Other available API versions: 2021-06-30, 2021-10-30-preview, 2022-01-30-preview, 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:datamigration:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Field that defines the Azure active directory application info, used to connect to the target Azure resource
        /// </summary>
        [Output("azureAuthenticationInfo")]
        public Output<Outputs.AzureActiveDirectoryAppResponse?> AzureAuthenticationInfo { get; private set; } = null!;

        /// <summary>
        /// UTC Date and time when project was created
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// List of DatabaseInfo
        /// </summary>
        [Output("databasesInfo")]
        public Output<ImmutableArray<Outputs.DatabaseInfoResponse>> DatabasesInfo { get; private set; } = null!;

        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project's provisioning state
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Information for connecting to source
        /// </summary>
        [Output("sourceConnectionInfo")]
        public Output<object?> SourceConnectionInfo { get; private set; } = null!;

        /// <summary>
        /// Source platform for the project
        /// </summary>
        [Output("sourcePlatform")]
        public Output<string> SourcePlatform { get; private set; } = null!;

        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Information for connecting to target
        /// </summary>
        [Output("targetConnectionInfo")]
        public Output<object?> TargetConnectionInfo { get; private set; } = null!;

        /// <summary>
        /// Target platform for the project
        /// </summary>
        [Output("targetPlatform")]
        public Output<string> TargetPlatform { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("azure-native:datamigration:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:datamigration:Project", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20171115preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20180315preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20180331preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20180419:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20180715preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20210630:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20211030preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20220130preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20220330preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20230715preview:Project" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20250315preview:Project" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Project(name, id, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field that defines the Azure active directory application info, used to connect to the target Azure resource
        /// </summary>
        [Input("azureAuthenticationInfo")]
        public Input<Inputs.AzureActiveDirectoryAppArgs>? AzureAuthenticationInfo { get; set; }

        [Input("databasesInfo")]
        private InputList<Inputs.DatabaseInfoArgs>? _databasesInfo;

        /// <summary>
        /// List of DatabaseInfo
        /// </summary>
        public InputList<Inputs.DatabaseInfoArgs> DatabasesInfo
        {
            get => _databasesInfo ?? (_databasesInfo = new InputList<Inputs.DatabaseInfoArgs>());
            set => _databasesInfo = value;
        }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Information for connecting to source
        /// </summary>
        [Input("sourceConnectionInfo")]
        public object? SourceConnectionInfo { get; set; }

        /// <summary>
        /// Source platform for the project
        /// </summary>
        [Input("sourcePlatform", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataMigration.ProjectSourcePlatform> SourcePlatform { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Information for connecting to target
        /// </summary>
        [Input("targetConnectionInfo")]
        public object? TargetConnectionInfo { get; set; }

        /// <summary>
        /// Target platform for the project
        /// </summary>
        [Input("targetPlatform", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataMigration.ProjectTargetPlatform> TargetPlatform { get; set; } = null!;

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }
}
