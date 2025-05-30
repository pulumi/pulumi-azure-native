# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from ._enums import *
from ._inputs import *

__all__ = ['AccessReviewHistoryDefinitionByIdArgs', 'AccessReviewHistoryDefinitionById']

@pulumi.input_type
class AccessReviewHistoryDefinitionByIdArgs:
    def __init__(__self__, *,
                 decisions: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 history_definition_id: Optional[pulumi.Input[builtins.str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewHistoryInstanceArgs']]]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 range: Optional[pulumi.Input['AccessReviewRecurrenceRangeArgs']] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewScopeArgs']]]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']]] = None):
        """
        The set of arguments for constructing a AccessReviewHistoryDefinitionById resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]] decisions: Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
        :param pulumi.Input[builtins.str] display_name: The display name for the history definition.
        :param pulumi.Input[builtins.str] history_definition_id: The id of the access review history definition.
        :param pulumi.Input[Sequence[pulumi.Input['AccessReviewHistoryInstanceArgs']]] instances: Set of access review history instances for this history definition.
        :param pulumi.Input[builtins.int] interval: The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        :param pulumi.Input['AccessReviewRecurrenceRangeArgs'] range: Access Review History Definition recurrence settings.
        :param pulumi.Input[Sequence[pulumi.Input['AccessReviewScopeArgs']]] scopes: A collection of scopes used when selecting review history data
        :param pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']] type: The recurrence type : weekly, monthly, etc.
        """
        if decisions is not None:
            pulumi.set(__self__, "decisions", decisions)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if history_definition_id is not None:
            pulumi.set(__self__, "history_definition_id", history_definition_id)
        if instances is not None:
            pulumi.set(__self__, "instances", instances)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def decisions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]]]:
        """
        Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
        """
        return pulumi.get(self, "decisions")

    @decisions.setter
    def decisions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]]]):
        pulumi.set(self, "decisions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The display name for the history definition.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="historyDefinitionId")
    def history_definition_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the access review history definition.
        """
        return pulumi.get(self, "history_definition_id")

    @history_definition_id.setter
    def history_definition_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "history_definition_id", value)

    @property
    @pulumi.getter
    def instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewHistoryInstanceArgs']]]]:
        """
        Set of access review history instances for this history definition.
        """
        return pulumi.get(self, "instances")

    @instances.setter
    def instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewHistoryInstanceArgs']]]]):
        pulumi.set(self, "instances", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input['AccessReviewRecurrenceRangeArgs']]:
        """
        Access Review History Definition recurrence settings.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input['AccessReviewRecurrenceRangeArgs']]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewScopeArgs']]]]:
        """
        A collection of scopes used when selecting review history data
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessReviewScopeArgs']]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']]]:
        """
        The recurrence type : weekly, monthly, etc.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("azure-native:authorization:AccessReviewHistoryDefinitionById")
class AccessReviewHistoryDefinitionById(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decisions: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 history_definition_id: Optional[pulumi.Input[builtins.str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewHistoryInstanceArgs', 'AccessReviewHistoryInstanceArgsDict']]]]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 range: Optional[pulumi.Input[Union['AccessReviewRecurrenceRangeArgs', 'AccessReviewRecurrenceRangeArgsDict']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewScopeArgs', 'AccessReviewScopeArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']]] = None,
                 __props__=None):
        """
        Access Review History Definition.

        Uses Azure REST API version 2021-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01-preview.

        Other available API versions: 2021-11-16-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]] decisions: Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
        :param pulumi.Input[builtins.str] display_name: The display name for the history definition.
        :param pulumi.Input[builtins.str] history_definition_id: The id of the access review history definition.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewHistoryInstanceArgs', 'AccessReviewHistoryInstanceArgsDict']]]] instances: Set of access review history instances for this history definition.
        :param pulumi.Input[builtins.int] interval: The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        :param pulumi.Input[Union['AccessReviewRecurrenceRangeArgs', 'AccessReviewRecurrenceRangeArgsDict']] range: Access Review History Definition recurrence settings.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewScopeArgs', 'AccessReviewScopeArgsDict']]]] scopes: A collection of scopes used when selecting review history data
        :param pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']] type: The recurrence type : weekly, monthly, etc.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AccessReviewHistoryDefinitionByIdArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Access Review History Definition.

        Uses Azure REST API version 2021-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01-preview.

        Other available API versions: 2021-11-16-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AccessReviewHistoryDefinitionByIdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessReviewHistoryDefinitionByIdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decisions: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AccessReviewResult']]]]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 history_definition_id: Optional[pulumi.Input[builtins.str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewHistoryInstanceArgs', 'AccessReviewHistoryInstanceArgsDict']]]]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 range: Optional[pulumi.Input[Union['AccessReviewRecurrenceRangeArgs', 'AccessReviewRecurrenceRangeArgsDict']]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AccessReviewScopeArgs', 'AccessReviewScopeArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[Union[builtins.str, 'AccessReviewRecurrencePatternType']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessReviewHistoryDefinitionByIdArgs.__new__(AccessReviewHistoryDefinitionByIdArgs)

            __props__.__dict__["decisions"] = decisions
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["history_definition_id"] = history_definition_id
            __props__.__dict__["instances"] = instances
            __props__.__dict__["interval"] = interval
            __props__.__dict__["range"] = range
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["type"] = type
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["created_date_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["principal_id"] = None
            __props__.__dict__["principal_name"] = None
            __props__.__dict__["principal_type"] = None
            __props__.__dict__["review_history_period_end_date_time"] = None
            __props__.__dict__["review_history_period_start_date_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["user_principal_name"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:authorization/v20211116preview:AccessReviewHistoryDefinitionById"), pulumi.Alias(type_="azure-native:authorization/v20211201preview:AccessReviewHistoryDefinitionById")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AccessReviewHistoryDefinitionById, __self__).__init__(
            'azure-native:authorization:AccessReviewHistoryDefinitionById',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessReviewHistoryDefinitionById':
        """
        Get an existing AccessReviewHistoryDefinitionById resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccessReviewHistoryDefinitionByIdArgs.__new__(AccessReviewHistoryDefinitionByIdArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["created_date_time"] = None
        __props__.__dict__["decisions"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["instances"] = None
        __props__.__dict__["interval"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["principal_id"] = None
        __props__.__dict__["principal_name"] = None
        __props__.__dict__["principal_type"] = None
        __props__.__dict__["range"] = None
        __props__.__dict__["review_history_period_end_date_time"] = None
        __props__.__dict__["review_history_period_start_date_time"] = None
        __props__.__dict__["scopes"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["user_principal_name"] = None
        return AccessReviewHistoryDefinitionById(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdDateTime")
    def created_date_time(self) -> pulumi.Output[builtins.str]:
        """
        Date time when history definition was created
        """
        return pulumi.get(self, "created_date_time")

    @property
    @pulumi.getter
    def decisions(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Collection of review decisions which the history data should be filtered on. For example if Approve and Deny are supplied the data will only contain review results in which the decision maker approved or denied a review request.
        """
        return pulumi.get(self, "decisions")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The display name for the history definition.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def instances(self) -> pulumi.Output[Optional[Sequence['outputs.AccessReviewHistoryInstanceResponse']]]:
        """
        Set of access review history instances for this history definition.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The access review history definition unique id.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[builtins.str]:
        """
        The identity id
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> pulumi.Output[builtins.str]:
        """
        The identity display name
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[builtins.str]:
        """
        The identity type : user/servicePrincipal
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter
    def range(self) -> pulumi.Output[Optional['outputs.AccessReviewRecurrenceRangeResponse']]:
        """
        Access Review History Definition recurrence settings.
        """
        return pulumi.get(self, "range")

    @property
    @pulumi.getter(name="reviewHistoryPeriodEndDateTime")
    def review_history_period_end_date_time(self) -> pulumi.Output[builtins.str]:
        """
        Date time used when selecting review data, all reviews included in data end on or before this date. For use only with one-time/non-recurring reports.
        """
        return pulumi.get(self, "review_history_period_end_date_time")

    @property
    @pulumi.getter(name="reviewHistoryPeriodStartDateTime")
    def review_history_period_start_date_time(self) -> pulumi.Output[builtins.str]:
        """
        Date time used when selecting review data, all reviews included in data start on or after this date. For use only with one-time/non-recurring reports.
        """
        return pulumi.get(self, "review_history_period_start_date_time")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence['outputs.AccessReviewScopeResponse']]]:
        """
        A collection of scopes used when selecting review history data
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        This read-only field specifies the of the requested review history data. This is either requested, in-progress, done or error.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userPrincipalName")
    def user_principal_name(self) -> pulumi.Output[builtins.str]:
        """
        The user principal name(if valid)
        """
        return pulumi.get(self, "user_principal_name")

