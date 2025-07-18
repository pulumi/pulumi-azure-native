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

__all__ = [
    'GetGroupsOperationResult',
    'AwaitableGetGroupsOperationResult',
    'get_groups_operation',
    'get_groups_operation_output',
]

@pulumi.output_type
class GetGroupsOperationResult:
    """
    Group resource.
    """
    def __init__(__self__, are_assessments_running=None, assessments=None, azure_api_version=None, created_timestamp=None, group_status=None, group_type=None, id=None, machine_count=None, name=None, provisioning_state=None, supported_assessment_types=None, system_data=None, type=None, updated_timestamp=None):
        if are_assessments_running and not isinstance(are_assessments_running, bool):
            raise TypeError("Expected argument 'are_assessments_running' to be a bool")
        pulumi.set(__self__, "are_assessments_running", are_assessments_running)
        if assessments and not isinstance(assessments, list):
            raise TypeError("Expected argument 'assessments' to be a list")
        pulumi.set(__self__, "assessments", assessments)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created_timestamp and not isinstance(created_timestamp, str):
            raise TypeError("Expected argument 'created_timestamp' to be a str")
        pulumi.set(__self__, "created_timestamp", created_timestamp)
        if group_status and not isinstance(group_status, str):
            raise TypeError("Expected argument 'group_status' to be a str")
        pulumi.set(__self__, "group_status", group_status)
        if group_type and not isinstance(group_type, str):
            raise TypeError("Expected argument 'group_type' to be a str")
        pulumi.set(__self__, "group_type", group_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if machine_count and not isinstance(machine_count, int):
            raise TypeError("Expected argument 'machine_count' to be a int")
        pulumi.set(__self__, "machine_count", machine_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if supported_assessment_types and not isinstance(supported_assessment_types, list):
            raise TypeError("Expected argument 'supported_assessment_types' to be a list")
        pulumi.set(__self__, "supported_assessment_types", supported_assessment_types)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_timestamp and not isinstance(updated_timestamp, str):
            raise TypeError("Expected argument 'updated_timestamp' to be a str")
        pulumi.set(__self__, "updated_timestamp", updated_timestamp)

    @property
    @pulumi.getter(name="areAssessmentsRunning")
    def are_assessments_running(self) -> builtins.bool:
        """
        If the assessments are in running state.
        """
        return pulumi.get(self, "are_assessments_running")

    @property
    @pulumi.getter
    def assessments(self) -> Sequence[builtins.str]:
        """
        List of References to Assessments created on this group.
        """
        return pulumi.get(self, "assessments")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdTimestamp")
    def created_timestamp(self) -> builtins.str:
        """
        Time when this group was created. Date-Time represented in ISO-8601 format.
        """
        return pulumi.get(self, "created_timestamp")

    @property
    @pulumi.getter(name="groupStatus")
    def group_status(self) -> builtins.str:
        """
        Whether the group has been created and is valid.
        """
        return pulumi.get(self, "group_status")

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> Optional[builtins.str]:
        """
        The type of group.
        """
        return pulumi.get(self, "group_type")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="machineCount")
    def machine_count(self) -> builtins.int:
        """
        Number of machines part of this group.
        """
        return pulumi.get(self, "machine_count")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[builtins.str]:
        """
        The status of the last operation.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="supportedAssessmentTypes")
    def supported_assessment_types(self) -> Optional[Sequence[builtins.str]]:
        """
        List of assessment types supported on this group.
        """
        return pulumi.get(self, "supported_assessment_types")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedTimestamp")
    def updated_timestamp(self) -> builtins.str:
        """
        Time when this group was last updated. Date-Time represented in ISO-8601 format.
        """
        return pulumi.get(self, "updated_timestamp")


class AwaitableGetGroupsOperationResult(GetGroupsOperationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsOperationResult(
            are_assessments_running=self.are_assessments_running,
            assessments=self.assessments,
            azure_api_version=self.azure_api_version,
            created_timestamp=self.created_timestamp,
            group_status=self.group_status,
            group_type=self.group_type,
            id=self.id,
            machine_count=self.machine_count,
            name=self.name,
            provisioning_state=self.provisioning_state,
            supported_assessment_types=self.supported_assessment_types,
            system_data=self.system_data,
            type=self.type,
            updated_timestamp=self.updated_timestamp)


def get_groups_operation(group_name: Optional[builtins.str] = None,
                         project_name: Optional[builtins.str] = None,
                         resource_group_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsOperationResult:
    """
    Get a Group

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str group_name: Group ARM name
    :param builtins.str project_name: Assessment Project Name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:migrate:getGroupsOperation', __args__, opts=opts, typ=GetGroupsOperationResult).value

    return AwaitableGetGroupsOperationResult(
        are_assessments_running=pulumi.get(__ret__, 'are_assessments_running'),
        assessments=pulumi.get(__ret__, 'assessments'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created_timestamp=pulumi.get(__ret__, 'created_timestamp'),
        group_status=pulumi.get(__ret__, 'group_status'),
        group_type=pulumi.get(__ret__, 'group_type'),
        id=pulumi.get(__ret__, 'id'),
        machine_count=pulumi.get(__ret__, 'machine_count'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        supported_assessment_types=pulumi.get(__ret__, 'supported_assessment_types'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        updated_timestamp=pulumi.get(__ret__, 'updated_timestamp'))
def get_groups_operation_output(group_name: Optional[pulumi.Input[builtins.str]] = None,
                                project_name: Optional[pulumi.Input[builtins.str]] = None,
                                resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupsOperationResult]:
    """
    Get a Group

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str group_name: Group ARM name
    :param builtins.str project_name: Assessment Project Name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:migrate:getGroupsOperation', __args__, opts=opts, typ=GetGroupsOperationResult)
    return __ret__.apply(lambda __response__: GetGroupsOperationResult(
        are_assessments_running=pulumi.get(__response__, 'are_assessments_running'),
        assessments=pulumi.get(__response__, 'assessments'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created_timestamp=pulumi.get(__response__, 'created_timestamp'),
        group_status=pulumi.get(__response__, 'group_status'),
        group_type=pulumi.get(__response__, 'group_type'),
        id=pulumi.get(__response__, 'id'),
        machine_count=pulumi.get(__response__, 'machine_count'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        supported_assessment_types=pulumi.get(__response__, 'supported_assessment_types'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        updated_timestamp=pulumi.get(__response__, 'updated_timestamp')))
