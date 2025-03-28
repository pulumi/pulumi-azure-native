# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetAccessReviewScheduleDefinitionByIdResult',
    'AwaitableGetAccessReviewScheduleDefinitionByIdResult',
    'get_access_review_schedule_definition_by_id',
    'get_access_review_schedule_definition_by_id_output',
]

@pulumi.output_type
class GetAccessReviewScheduleDefinitionByIdResult:
    """
    Access Review Schedule Definition.
    """
    def __init__(__self__, assignment_state=None, auto_apply_decisions_enabled=None, backup_reviewers=None, default_decision=None, default_decision_enabled=None, description_for_admins=None, description_for_reviewers=None, display_name=None, end_date=None, exclude_resource_id=None, exclude_role_definition_id=None, expand_nested_memberships=None, id=None, inactive_duration=None, include_access_below_resource=None, include_inherited_access=None, instance_duration_in_days=None, instances=None, interval=None, justification_required_on_approval=None, mail_notifications_enabled=None, name=None, number_of_occurrences=None, principal_id=None, principal_name=None, principal_type=None, recommendation_look_back_duration=None, recommendations_enabled=None, reminder_notifications_enabled=None, resource_id=None, reviewers=None, reviewers_type=None, role_definition_id=None, start_date=None, status=None, type=None, user_principal_name=None):
        if assignment_state and not isinstance(assignment_state, str):
            raise TypeError("Expected argument 'assignment_state' to be a str")
        pulumi.set(__self__, "assignment_state", assignment_state)
        if auto_apply_decisions_enabled and not isinstance(auto_apply_decisions_enabled, bool):
            raise TypeError("Expected argument 'auto_apply_decisions_enabled' to be a bool")
        pulumi.set(__self__, "auto_apply_decisions_enabled", auto_apply_decisions_enabled)
        if backup_reviewers and not isinstance(backup_reviewers, list):
            raise TypeError("Expected argument 'backup_reviewers' to be a list")
        pulumi.set(__self__, "backup_reviewers", backup_reviewers)
        if default_decision and not isinstance(default_decision, str):
            raise TypeError("Expected argument 'default_decision' to be a str")
        pulumi.set(__self__, "default_decision", default_decision)
        if default_decision_enabled and not isinstance(default_decision_enabled, bool):
            raise TypeError("Expected argument 'default_decision_enabled' to be a bool")
        pulumi.set(__self__, "default_decision_enabled", default_decision_enabled)
        if description_for_admins and not isinstance(description_for_admins, str):
            raise TypeError("Expected argument 'description_for_admins' to be a str")
        pulumi.set(__self__, "description_for_admins", description_for_admins)
        if description_for_reviewers and not isinstance(description_for_reviewers, str):
            raise TypeError("Expected argument 'description_for_reviewers' to be a str")
        pulumi.set(__self__, "description_for_reviewers", description_for_reviewers)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if exclude_resource_id and not isinstance(exclude_resource_id, str):
            raise TypeError("Expected argument 'exclude_resource_id' to be a str")
        pulumi.set(__self__, "exclude_resource_id", exclude_resource_id)
        if exclude_role_definition_id and not isinstance(exclude_role_definition_id, str):
            raise TypeError("Expected argument 'exclude_role_definition_id' to be a str")
        pulumi.set(__self__, "exclude_role_definition_id", exclude_role_definition_id)
        if expand_nested_memberships and not isinstance(expand_nested_memberships, bool):
            raise TypeError("Expected argument 'expand_nested_memberships' to be a bool")
        pulumi.set(__self__, "expand_nested_memberships", expand_nested_memberships)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inactive_duration and not isinstance(inactive_duration, str):
            raise TypeError("Expected argument 'inactive_duration' to be a str")
        pulumi.set(__self__, "inactive_duration", inactive_duration)
        if include_access_below_resource and not isinstance(include_access_below_resource, bool):
            raise TypeError("Expected argument 'include_access_below_resource' to be a bool")
        pulumi.set(__self__, "include_access_below_resource", include_access_below_resource)
        if include_inherited_access and not isinstance(include_inherited_access, bool):
            raise TypeError("Expected argument 'include_inherited_access' to be a bool")
        pulumi.set(__self__, "include_inherited_access", include_inherited_access)
        if instance_duration_in_days and not isinstance(instance_duration_in_days, int):
            raise TypeError("Expected argument 'instance_duration_in_days' to be a int")
        pulumi.set(__self__, "instance_duration_in_days", instance_duration_in_days)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if interval and not isinstance(interval, int):
            raise TypeError("Expected argument 'interval' to be a int")
        pulumi.set(__self__, "interval", interval)
        if justification_required_on_approval and not isinstance(justification_required_on_approval, bool):
            raise TypeError("Expected argument 'justification_required_on_approval' to be a bool")
        pulumi.set(__self__, "justification_required_on_approval", justification_required_on_approval)
        if mail_notifications_enabled and not isinstance(mail_notifications_enabled, bool):
            raise TypeError("Expected argument 'mail_notifications_enabled' to be a bool")
        pulumi.set(__self__, "mail_notifications_enabled", mail_notifications_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_occurrences and not isinstance(number_of_occurrences, int):
            raise TypeError("Expected argument 'number_of_occurrences' to be a int")
        pulumi.set(__self__, "number_of_occurrences", number_of_occurrences)
        if principal_id and not isinstance(principal_id, str):
            raise TypeError("Expected argument 'principal_id' to be a str")
        pulumi.set(__self__, "principal_id", principal_id)
        if principal_name and not isinstance(principal_name, str):
            raise TypeError("Expected argument 'principal_name' to be a str")
        pulumi.set(__self__, "principal_name", principal_name)
        if principal_type and not isinstance(principal_type, str):
            raise TypeError("Expected argument 'principal_type' to be a str")
        pulumi.set(__self__, "principal_type", principal_type)
        if recommendation_look_back_duration and not isinstance(recommendation_look_back_duration, str):
            raise TypeError("Expected argument 'recommendation_look_back_duration' to be a str")
        pulumi.set(__self__, "recommendation_look_back_duration", recommendation_look_back_duration)
        if recommendations_enabled and not isinstance(recommendations_enabled, bool):
            raise TypeError("Expected argument 'recommendations_enabled' to be a bool")
        pulumi.set(__self__, "recommendations_enabled", recommendations_enabled)
        if reminder_notifications_enabled and not isinstance(reminder_notifications_enabled, bool):
            raise TypeError("Expected argument 'reminder_notifications_enabled' to be a bool")
        pulumi.set(__self__, "reminder_notifications_enabled", reminder_notifications_enabled)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if reviewers and not isinstance(reviewers, list):
            raise TypeError("Expected argument 'reviewers' to be a list")
        pulumi.set(__self__, "reviewers", reviewers)
        if reviewers_type and not isinstance(reviewers_type, str):
            raise TypeError("Expected argument 'reviewers_type' to be a str")
        pulumi.set(__self__, "reviewers_type", reviewers_type)
        if role_definition_id and not isinstance(role_definition_id, str):
            raise TypeError("Expected argument 'role_definition_id' to be a str")
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_principal_name and not isinstance(user_principal_name, str):
            raise TypeError("Expected argument 'user_principal_name' to be a str")
        pulumi.set(__self__, "user_principal_name", user_principal_name)

    @property
    @pulumi.getter(name="assignmentState")
    def assignment_state(self) -> str:
        """
        The role assignment state eligible/active to review
        """
        return pulumi.get(self, "assignment_state")

    @property
    @pulumi.getter(name="autoApplyDecisionsEnabled")
    def auto_apply_decisions_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
        """
        return pulumi.get(self, "auto_apply_decisions_enabled")

    @property
    @pulumi.getter(name="backupReviewers")
    def backup_reviewers(self) -> Optional[Sequence['outputs.AccessReviewReviewerResponse']]:
        """
        This is the collection of backup reviewers.
        """
        return pulumi.get(self, "backup_reviewers")

    @property
    @pulumi.getter(name="defaultDecision")
    def default_decision(self) -> Optional[str]:
        """
        This specifies the behavior for the autoReview feature when an access review completes.
        """
        return pulumi.get(self, "default_decision")

    @property
    @pulumi.getter(name="defaultDecisionEnabled")
    def default_decision_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether reviewers are required to provide a justification when reviewing access.
        """
        return pulumi.get(self, "default_decision_enabled")

    @property
    @pulumi.getter(name="descriptionForAdmins")
    def description_for_admins(self) -> Optional[str]:
        """
        The description provided by the access review creator and visible to admins.
        """
        return pulumi.get(self, "description_for_admins")

    @property
    @pulumi.getter(name="descriptionForReviewers")
    def description_for_reviewers(self) -> Optional[str]:
        """
        The description provided by the access review creator to be shown to reviewers.
        """
        return pulumi.get(self, "description_for_reviewers")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name for the schedule definition.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[str]:
        """
        The DateTime when the review is scheduled to end. Required if type is endDate
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="excludeResourceId")
    def exclude_resource_id(self) -> Optional[str]:
        """
        This is used to indicate the resource id(s) to exclude
        """
        return pulumi.get(self, "exclude_resource_id")

    @property
    @pulumi.getter(name="excludeRoleDefinitionId")
    def exclude_role_definition_id(self) -> Optional[str]:
        """
        This is used to indicate the role definition id(s) to exclude
        """
        return pulumi.get(self, "exclude_role_definition_id")

    @property
    @pulumi.getter(name="expandNestedMemberships")
    def expand_nested_memberships(self) -> Optional[bool]:
        """
        Flag to indicate whether to expand nested memberships or not.
        """
        return pulumi.get(self, "expand_nested_memberships")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The access review schedule definition id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inactiveDuration")
    def inactive_duration(self) -> Optional[str]:
        """
        Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
        """
        return pulumi.get(self, "inactive_duration")

    @property
    @pulumi.getter(name="includeAccessBelowResource")
    def include_access_below_resource(self) -> Optional[bool]:
        """
        Flag to indicate whether to expand nested memberships or not.
        """
        return pulumi.get(self, "include_access_below_resource")

    @property
    @pulumi.getter(name="includeInheritedAccess")
    def include_inherited_access(self) -> Optional[bool]:
        """
        Flag to indicate whether to expand nested memberships or not.
        """
        return pulumi.get(self, "include_inherited_access")

    @property
    @pulumi.getter(name="instanceDurationInDays")
    def instance_duration_in_days(self) -> Optional[int]:
        """
        The duration in days for an instance.
        """
        return pulumi.get(self, "instance_duration_in_days")

    @property
    @pulumi.getter
    def instances(self) -> Optional[Sequence['outputs.AccessReviewInstanceResponse']]:
        """
        This is the collection of instances returned when one does an expand on it.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        """
        The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="justificationRequiredOnApproval")
    def justification_required_on_approval(self) -> Optional[bool]:
        """
        Flag to indicate whether the reviewer is required to pass justification when recording a decision.
        """
        return pulumi.get(self, "justification_required_on_approval")

    @property
    @pulumi.getter(name="mailNotificationsEnabled")
    def mail_notifications_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether sending mails to reviewers and the review creator is enabled.
        """
        return pulumi.get(self, "mail_notifications_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The access review schedule definition unique id.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfOccurrences")
    def number_of_occurrences(self) -> Optional[int]:
        """
        The number of times to repeat the access review. Required and must be positive if type is numbered.
        """
        return pulumi.get(self, "number_of_occurrences")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The identity id
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> str:
        """
        The identity display name
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> str:
        """
        The identity type user/servicePrincipal to review
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="recommendationLookBackDuration")
    def recommendation_look_back_duration(self) -> Optional[str]:
        """
        Recommendations for access reviews are calculated by looking back at 30 days of data(w.r.t the start date of the review) by default. However, in some scenarios, customers want to change how far back to look at and want to configure 60 days, 90 days, etc. instead. This setting allows customers to configure this duration. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
        """
        return pulumi.get(self, "recommendation_look_back_duration")

    @property
    @pulumi.getter(name="recommendationsEnabled")
    def recommendations_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether showing recommendations to reviewers is enabled.
        """
        return pulumi.get(self, "recommendations_enabled")

    @property
    @pulumi.getter(name="reminderNotificationsEnabled")
    def reminder_notifications_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether sending reminder emails to reviewers are enabled.
        """
        return pulumi.get(self, "reminder_notifications_enabled")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        ResourceId in which this review is getting created
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def reviewers(self) -> Optional[Sequence['outputs.AccessReviewReviewerResponse']]:
        """
        This is the collection of reviewers.
        """
        return pulumi.get(self, "reviewers")

    @property
    @pulumi.getter(name="reviewersType")
    def reviewers_type(self) -> str:
        """
        This field specifies the type of reviewers for a review. Usually for a review, reviewers are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen dynamically. For example managers review or self review.
        """
        return pulumi.get(self, "reviewers_type")

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> str:
        """
        This is used to indicate the role being reviewed
        """
        return pulumi.get(self, "role_definition_id")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[str]:
        """
        The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        This read-only field specifies the status of an accessReview.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userPrincipalName")
    def user_principal_name(self) -> str:
        """
        The user principal name(if valid)
        """
        return pulumi.get(self, "user_principal_name")


class AwaitableGetAccessReviewScheduleDefinitionByIdResult(GetAccessReviewScheduleDefinitionByIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessReviewScheduleDefinitionByIdResult(
            assignment_state=self.assignment_state,
            auto_apply_decisions_enabled=self.auto_apply_decisions_enabled,
            backup_reviewers=self.backup_reviewers,
            default_decision=self.default_decision,
            default_decision_enabled=self.default_decision_enabled,
            description_for_admins=self.description_for_admins,
            description_for_reviewers=self.description_for_reviewers,
            display_name=self.display_name,
            end_date=self.end_date,
            exclude_resource_id=self.exclude_resource_id,
            exclude_role_definition_id=self.exclude_role_definition_id,
            expand_nested_memberships=self.expand_nested_memberships,
            id=self.id,
            inactive_duration=self.inactive_duration,
            include_access_below_resource=self.include_access_below_resource,
            include_inherited_access=self.include_inherited_access,
            instance_duration_in_days=self.instance_duration_in_days,
            instances=self.instances,
            interval=self.interval,
            justification_required_on_approval=self.justification_required_on_approval,
            mail_notifications_enabled=self.mail_notifications_enabled,
            name=self.name,
            number_of_occurrences=self.number_of_occurrences,
            principal_id=self.principal_id,
            principal_name=self.principal_name,
            principal_type=self.principal_type,
            recommendation_look_back_duration=self.recommendation_look_back_duration,
            recommendations_enabled=self.recommendations_enabled,
            reminder_notifications_enabled=self.reminder_notifications_enabled,
            resource_id=self.resource_id,
            reviewers=self.reviewers,
            reviewers_type=self.reviewers_type,
            role_definition_id=self.role_definition_id,
            start_date=self.start_date,
            status=self.status,
            type=self.type,
            user_principal_name=self.user_principal_name)


def get_access_review_schedule_definition_by_id(schedule_definition_id: Optional[str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessReviewScheduleDefinitionByIdResult:
    """
    Get single access review definition

    Uses Azure REST API version 2021-12-01-preview.


    :param str schedule_definition_id: The id of the access review schedule definition.
    """
    __args__ = dict()
    __args__['scheduleDefinitionId'] = schedule_definition_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:authorization:getAccessReviewScheduleDefinitionById', __args__, opts=opts, typ=GetAccessReviewScheduleDefinitionByIdResult).value

    return AwaitableGetAccessReviewScheduleDefinitionByIdResult(
        assignment_state=pulumi.get(__ret__, 'assignment_state'),
        auto_apply_decisions_enabled=pulumi.get(__ret__, 'auto_apply_decisions_enabled'),
        backup_reviewers=pulumi.get(__ret__, 'backup_reviewers'),
        default_decision=pulumi.get(__ret__, 'default_decision'),
        default_decision_enabled=pulumi.get(__ret__, 'default_decision_enabled'),
        description_for_admins=pulumi.get(__ret__, 'description_for_admins'),
        description_for_reviewers=pulumi.get(__ret__, 'description_for_reviewers'),
        display_name=pulumi.get(__ret__, 'display_name'),
        end_date=pulumi.get(__ret__, 'end_date'),
        exclude_resource_id=pulumi.get(__ret__, 'exclude_resource_id'),
        exclude_role_definition_id=pulumi.get(__ret__, 'exclude_role_definition_id'),
        expand_nested_memberships=pulumi.get(__ret__, 'expand_nested_memberships'),
        id=pulumi.get(__ret__, 'id'),
        inactive_duration=pulumi.get(__ret__, 'inactive_duration'),
        include_access_below_resource=pulumi.get(__ret__, 'include_access_below_resource'),
        include_inherited_access=pulumi.get(__ret__, 'include_inherited_access'),
        instance_duration_in_days=pulumi.get(__ret__, 'instance_duration_in_days'),
        instances=pulumi.get(__ret__, 'instances'),
        interval=pulumi.get(__ret__, 'interval'),
        justification_required_on_approval=pulumi.get(__ret__, 'justification_required_on_approval'),
        mail_notifications_enabled=pulumi.get(__ret__, 'mail_notifications_enabled'),
        name=pulumi.get(__ret__, 'name'),
        number_of_occurrences=pulumi.get(__ret__, 'number_of_occurrences'),
        principal_id=pulumi.get(__ret__, 'principal_id'),
        principal_name=pulumi.get(__ret__, 'principal_name'),
        principal_type=pulumi.get(__ret__, 'principal_type'),
        recommendation_look_back_duration=pulumi.get(__ret__, 'recommendation_look_back_duration'),
        recommendations_enabled=pulumi.get(__ret__, 'recommendations_enabled'),
        reminder_notifications_enabled=pulumi.get(__ret__, 'reminder_notifications_enabled'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        reviewers=pulumi.get(__ret__, 'reviewers'),
        reviewers_type=pulumi.get(__ret__, 'reviewers_type'),
        role_definition_id=pulumi.get(__ret__, 'role_definition_id'),
        start_date=pulumi.get(__ret__, 'start_date'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'),
        user_principal_name=pulumi.get(__ret__, 'user_principal_name'))
def get_access_review_schedule_definition_by_id_output(schedule_definition_id: Optional[pulumi.Input[str]] = None,
                                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAccessReviewScheduleDefinitionByIdResult]:
    """
    Get single access review definition

    Uses Azure REST API version 2021-12-01-preview.


    :param str schedule_definition_id: The id of the access review schedule definition.
    """
    __args__ = dict()
    __args__['scheduleDefinitionId'] = schedule_definition_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:authorization:getAccessReviewScheduleDefinitionById', __args__, opts=opts, typ=GetAccessReviewScheduleDefinitionByIdResult)
    return __ret__.apply(lambda __response__: GetAccessReviewScheduleDefinitionByIdResult(
        assignment_state=pulumi.get(__response__, 'assignment_state'),
        auto_apply_decisions_enabled=pulumi.get(__response__, 'auto_apply_decisions_enabled'),
        backup_reviewers=pulumi.get(__response__, 'backup_reviewers'),
        default_decision=pulumi.get(__response__, 'default_decision'),
        default_decision_enabled=pulumi.get(__response__, 'default_decision_enabled'),
        description_for_admins=pulumi.get(__response__, 'description_for_admins'),
        description_for_reviewers=pulumi.get(__response__, 'description_for_reviewers'),
        display_name=pulumi.get(__response__, 'display_name'),
        end_date=pulumi.get(__response__, 'end_date'),
        exclude_resource_id=pulumi.get(__response__, 'exclude_resource_id'),
        exclude_role_definition_id=pulumi.get(__response__, 'exclude_role_definition_id'),
        expand_nested_memberships=pulumi.get(__response__, 'expand_nested_memberships'),
        id=pulumi.get(__response__, 'id'),
        inactive_duration=pulumi.get(__response__, 'inactive_duration'),
        include_access_below_resource=pulumi.get(__response__, 'include_access_below_resource'),
        include_inherited_access=pulumi.get(__response__, 'include_inherited_access'),
        instance_duration_in_days=pulumi.get(__response__, 'instance_duration_in_days'),
        instances=pulumi.get(__response__, 'instances'),
        interval=pulumi.get(__response__, 'interval'),
        justification_required_on_approval=pulumi.get(__response__, 'justification_required_on_approval'),
        mail_notifications_enabled=pulumi.get(__response__, 'mail_notifications_enabled'),
        name=pulumi.get(__response__, 'name'),
        number_of_occurrences=pulumi.get(__response__, 'number_of_occurrences'),
        principal_id=pulumi.get(__response__, 'principal_id'),
        principal_name=pulumi.get(__response__, 'principal_name'),
        principal_type=pulumi.get(__response__, 'principal_type'),
        recommendation_look_back_duration=pulumi.get(__response__, 'recommendation_look_back_duration'),
        recommendations_enabled=pulumi.get(__response__, 'recommendations_enabled'),
        reminder_notifications_enabled=pulumi.get(__response__, 'reminder_notifications_enabled'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        reviewers=pulumi.get(__response__, 'reviewers'),
        reviewers_type=pulumi.get(__response__, 'reviewers_type'),
        role_definition_id=pulumi.get(__response__, 'role_definition_id'),
        start_date=pulumi.get(__response__, 'start_date'),
        status=pulumi.get(__response__, 'status'),
        type=pulumi.get(__response__, 'type'),
        user_principal_name=pulumi.get(__response__, 'user_principal_name')))
