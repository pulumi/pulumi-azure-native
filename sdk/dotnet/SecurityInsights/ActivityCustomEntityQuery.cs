// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    /// <summary>
    /// Represents Activity entity query.
    /// 
    /// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights:ActivityCustomEntityQuery")]
    public partial class ActivityCustomEntityQuery : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The entity query content to display in timeline
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// The time the activity was created
        /// </summary>
        [Output("createdTimeUtc")]
        public Output<string> CreatedTimeUtc { get; private set; } = null!;

        /// <summary>
        /// The entity query description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Determines whether this activity is enabled or disabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The query applied only to entities matching to all filters
        /// </summary>
        [Output("entitiesFilter")]
        public Output<ImmutableDictionary<string, ImmutableArray<string>>?> EntitiesFilter { get; private set; } = null!;

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The type of the query's source entity
        /// </summary>
        [Output("inputEntityType")]
        public Output<string?> InputEntityType { get; private set; } = null!;

        /// <summary>
        /// The kind of the entity query
        /// Expected value is 'Activity'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The last time the activity was updated
        /// </summary>
        [Output("lastModifiedTimeUtc")]
        public Output<string> LastModifiedTimeUtc { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Activity query definitions
        /// </summary>
        [Output("queryDefinitions")]
        public Output<Outputs.ActivityEntityQueriesPropertiesResponseQueryDefinitions?> QueryDefinitions { get; private set; } = null!;

        /// <summary>
        /// List of the fields of the source entity that are required to run the query
        /// </summary>
        [Output("requiredInputFieldsSets")]
        public Output<ImmutableArray<ImmutableArray<string>>> RequiredInputFieldsSets { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The template id this activity was created from
        /// </summary>
        [Output("templateName")]
        public Output<string?> TemplateName { get; private set; } = null!;

        /// <summary>
        /// The entity query title
        /// </summary>
        [Output("title")]
        public Output<string?> Title { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ActivityCustomEntityQuery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActivityCustomEntityQuery(string name, ActivityCustomEntityQueryArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:ActivityCustomEntityQuery", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private ActivityCustomEntityQuery(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:ActivityCustomEntityQuery", name, null, MakeResourceOptions(options, id))
        {
        }

        private static ActivityCustomEntityQueryArgs MakeArgs(ActivityCustomEntityQueryArgs args)
        {
            args ??= new ActivityCustomEntityQueryArgs();
            args.Kind = "Activity";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210301preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20210901preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20211001preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220101preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220401preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220501preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220601preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220701preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220801preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20220901preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221001preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221101preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20221201preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230201preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230301preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250101preview:ActivityCustomEntityQuery" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250401preview:ActivityCustomEntityQuery" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActivityCustomEntityQuery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActivityCustomEntityQuery Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ActivityCustomEntityQuery(name, id, options);
        }
    }

    public sealed class ActivityCustomEntityQueryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The entity query content to display in timeline
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The entity query description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Determines whether this activity is enabled or disabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("entitiesFilter")]
        private InputMap<ImmutableArray<string>>? _entitiesFilter;

        /// <summary>
        /// The query applied only to entities matching to all filters
        /// </summary>
        public InputMap<ImmutableArray<string>> EntitiesFilter
        {
            get => _entitiesFilter ?? (_entitiesFilter = new InputMap<ImmutableArray<string>>());
            set => _entitiesFilter = value;
        }

        /// <summary>
        /// entity query ID
        /// </summary>
        [Input("entityQueryId")]
        public Input<string>? EntityQueryId { get; set; }

        /// <summary>
        /// The type of the query's source entity
        /// </summary>
        [Input("inputEntityType")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.EntityType>? InputEntityType { get; set; }

        /// <summary>
        /// The kind of the entity query that supports put request.
        /// Expected value is 'Activity'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The Activity query definitions
        /// </summary>
        [Input("queryDefinitions")]
        public Input<Inputs.ActivityEntityQueriesPropertiesQueryDefinitionsArgs>? QueryDefinitions { get; set; }

        [Input("requiredInputFieldsSets")]
        private InputList<ImmutableArray<string>>? _requiredInputFieldsSets;

        /// <summary>
        /// List of the fields of the source entity that are required to run the query
        /// </summary>
        public InputList<ImmutableArray<string>> RequiredInputFieldsSets
        {
            get => _requiredInputFieldsSets ?? (_requiredInputFieldsSets = new InputList<ImmutableArray<string>>());
            set => _requiredInputFieldsSets = value;
        }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The template id this activity was created from
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// The entity query title
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public ActivityCustomEntityQueryArgs()
        {
        }
        public static new ActivityCustomEntityQueryArgs Empty => new ActivityCustomEntityQueryArgs();
    }
}
