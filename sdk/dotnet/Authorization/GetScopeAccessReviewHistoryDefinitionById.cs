// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization
{
    public static class GetScopeAccessReviewHistoryDefinitionById
    {
        /// <summary>
        /// Get access review history definition by definition Id
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Task<GetScopeAccessReviewHistoryDefinitionByIdResult> InvokeAsync(GetScopeAccessReviewHistoryDefinitionByIdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScopeAccessReviewHistoryDefinitionByIdResult>("azure-native:authorization:getScopeAccessReviewHistoryDefinitionById", args ?? new GetScopeAccessReviewHistoryDefinitionByIdArgs(), options.WithDefaults());

        /// <summary>
        /// Get access review history definition by definition Id
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Output<GetScopeAccessReviewHistoryDefinitionByIdResult> Invoke(GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScopeAccessReviewHistoryDefinitionByIdResult>("azure-native:authorization:getScopeAccessReviewHistoryDefinitionById", args ?? new GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get access review history definition by definition Id
        /// 
        /// Uses Azure REST API version 2021-12-01-preview.
        /// </summary>
        public static Output<GetScopeAccessReviewHistoryDefinitionByIdResult> Invoke(GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetScopeAccessReviewHistoryDefinitionByIdResult>("azure-native:authorization:getScopeAccessReviewHistoryDefinitionById", args ?? new GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScopeAccessReviewHistoryDefinitionByIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the access review history definition.
        /// </summary>
        [Input("historyDefinitionId", required: true)]
        public string HistoryDefinitionId { get; set; } = null!;

        /// <summary>
        /// The scope of the resource.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetScopeAccessReviewHistoryDefinitionByIdArgs()
        {
        }
        public static new GetScopeAccessReviewHistoryDefinitionByIdArgs Empty => new GetScopeAccessReviewHistoryDefinitionByIdArgs();
    }

    public sealed class GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the access review history definition.
        /// </summary>
        [Input("historyDefinitionId", required: true)]
        public Input<string> HistoryDefinitionId { get; set; } = null!;

        /// <summary>
        /// The scope of the resource.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs()
        {
        }
        public static new GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs Empty => new GetScopeAccessReviewHistoryDefinitionByIdInvokeArgs();
    }


    [OutputType]
    public sealed class GetScopeAccessReviewHistoryDefinitionByIdResult
    {
        /// <summary>
        /// Date time when history definition was created
        /// </summary>
        public readonly string CreatedDateTime;
        /// <summary>
        /// Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
        /// </summary>
        public readonly ImmutableArray<string> Decisions;
        /// <summary>
        /// The display name for the history definition.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The DateTime when the review is scheduled to end. Required if type is endDate
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// The access review history definition id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of access review history instances for this history definition.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessReviewHistoryInstanceResponse> Instances;
        /// <summary>
        /// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// The access review history definition unique id.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of times to repeat the access review. Required and must be positive if type is numbered.
        /// </summary>
        public readonly int? NumberOfOccurrences;
        /// <summary>
        /// The identity id
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The identity display name
        /// </summary>
        public readonly string PrincipalName;
        /// <summary>
        /// The identity type : user/servicePrincipal
        /// </summary>
        public readonly string PrincipalType;
        /// <summary>
        /// Date time used when selecting review data, all reviews included in data end on or before this date. For use only with one-time/non-recurring reports.
        /// </summary>
        public readonly string ReviewHistoryPeriodEndDateTime;
        /// <summary>
        /// Date time used when selecting review data, all reviews included in data start on or after this date. For use only with one-time/non-recurring reports.
        /// </summary>
        public readonly string ReviewHistoryPeriodStartDateTime;
        /// <summary>
        /// A collection of scopes used when selecting review history data
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessReviewScopeResponse> Scopes;
        /// <summary>
        /// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// This read-only field specifies the of the requested review history data. This is either requested, in-progress, done or error.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user principal name(if valid)
        /// </summary>
        public readonly string UserPrincipalName;

        [OutputConstructor]
        private GetScopeAccessReviewHistoryDefinitionByIdResult(
            string createdDateTime,

            ImmutableArray<string> decisions,

            string? displayName,

            string? endDate,

            string id,

            ImmutableArray<Outputs.AccessReviewHistoryInstanceResponse> instances,

            int? interval,

            string name,

            int? numberOfOccurrences,

            string principalId,

            string principalName,

            string principalType,

            string reviewHistoryPeriodEndDateTime,

            string reviewHistoryPeriodStartDateTime,

            ImmutableArray<Outputs.AccessReviewScopeResponse> scopes,

            string? startDate,

            string status,

            string type,

            string userPrincipalName)
        {
            CreatedDateTime = createdDateTime;
            Decisions = decisions;
            DisplayName = displayName;
            EndDate = endDate;
            Id = id;
            Instances = instances;
            Interval = interval;
            Name = name;
            NumberOfOccurrences = numberOfOccurrences;
            PrincipalId = principalId;
            PrincipalName = principalName;
            PrincipalType = principalType;
            ReviewHistoryPeriodEndDateTime = reviewHistoryPeriodEndDateTime;
            ReviewHistoryPeriodStartDateTime = reviewHistoryPeriodStartDateTime;
            Scopes = scopes;
            StartDate = startDate;
            Status = status;
            Type = type;
            UserPrincipalName = userPrincipalName;
        }
    }
}
