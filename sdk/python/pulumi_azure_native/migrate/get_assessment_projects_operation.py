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
    'GetAssessmentProjectsOperationResult',
    'AwaitableGetAssessmentProjectsOperationResult',
    'get_assessment_projects_operation',
    'get_assessment_projects_operation_output',
]

@pulumi.output_type
class GetAssessmentProjectsOperationResult:
    """
    An Assessment project site resource.
    """
    def __init__(__self__, assessment_solution_id=None, azure_api_version=None, created_timestamp=None, customer_storage_account_arm_id=None, customer_workspace_id=None, customer_workspace_location=None, id=None, location=None, name=None, private_endpoint_connections=None, project_status=None, provisioning_state=None, public_network_access=None, service_endpoint=None, system_data=None, tags=None, type=None, updated_timestamp=None):
        if assessment_solution_id and not isinstance(assessment_solution_id, str):
            raise TypeError("Expected argument 'assessment_solution_id' to be a str")
        pulumi.set(__self__, "assessment_solution_id", assessment_solution_id)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if created_timestamp and not isinstance(created_timestamp, str):
            raise TypeError("Expected argument 'created_timestamp' to be a str")
        pulumi.set(__self__, "created_timestamp", created_timestamp)
        if customer_storage_account_arm_id and not isinstance(customer_storage_account_arm_id, str):
            raise TypeError("Expected argument 'customer_storage_account_arm_id' to be a str")
        pulumi.set(__self__, "customer_storage_account_arm_id", customer_storage_account_arm_id)
        if customer_workspace_id and not isinstance(customer_workspace_id, str):
            raise TypeError("Expected argument 'customer_workspace_id' to be a str")
        pulumi.set(__self__, "customer_workspace_id", customer_workspace_id)
        if customer_workspace_location and not isinstance(customer_workspace_location, str):
            raise TypeError("Expected argument 'customer_workspace_location' to be a str")
        pulumi.set(__self__, "customer_workspace_location", customer_workspace_location)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if project_status and not isinstance(project_status, str):
            raise TypeError("Expected argument 'project_status' to be a str")
        pulumi.set(__self__, "project_status", project_status)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_network_access and not isinstance(public_network_access, str):
            raise TypeError("Expected argument 'public_network_access' to be a str")
        pulumi.set(__self__, "public_network_access", public_network_access)
        if service_endpoint and not isinstance(service_endpoint, str):
            raise TypeError("Expected argument 'service_endpoint' to be a str")
        pulumi.set(__self__, "service_endpoint", service_endpoint)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_timestamp and not isinstance(updated_timestamp, str):
            raise TypeError("Expected argument 'updated_timestamp' to be a str")
        pulumi.set(__self__, "updated_timestamp", updated_timestamp)

    @property
    @pulumi.getter(name="assessmentSolutionId")
    def assessment_solution_id(self) -> Optional[builtins.str]:
        """
        Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
        """
        return pulumi.get(self, "assessment_solution_id")

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
        Time when this project was created. Date-Time represented in ISO-8601 format.
        """
        return pulumi.get(self, "created_timestamp")

    @property
    @pulumi.getter(name="customerStorageAccountArmId")
    def customer_storage_account_arm_id(self) -> Optional[builtins.str]:
        """
        The ARM id of the storage account used for interactions when public access is
        disabled.
        """
        return pulumi.get(self, "customer_storage_account_arm_id")

    @property
    @pulumi.getter(name="customerWorkspaceId")
    def customer_workspace_id(self) -> Optional[builtins.str]:
        """
        The ARM id of service map workspace created by customer.
        """
        return pulumi.get(self, "customer_workspace_id")

    @property
    @pulumi.getter(name="customerWorkspaceLocation")
    def customer_workspace_location(self) -> Optional[builtins.str]:
        """
        Location of service map workspace created by customer.
        """
        return pulumi.get(self, "customer_workspace_location")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Sequence['outputs.PrivateEndpointConnectionResponse']:
        """
        The list of private endpoint connections to the project.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="projectStatus")
    def project_status(self) -> Optional[builtins.str]:
        """
        Assessment project status.
        """
        return pulumi.get(self, "project_status")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[builtins.str]:
        """
        The status of the last operation.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[builtins.str]:
        """
        This value can be set to 'enabled' to avoid breaking changes on existing
        customer resources and templates. If set to 'disabled', traffic over public
        interface is not allowed, and private endpoint connections would be the
        exclusive access method.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="serviceEndpoint")
    def service_endpoint(self) -> builtins.str:
        """
        Endpoint at which the collector agent can call agent REST API.
        """
        return pulumi.get(self, "service_endpoint")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

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
        Time when this project was last updated. Date-Time represented in ISO-8601
        format.
        """
        return pulumi.get(self, "updated_timestamp")


class AwaitableGetAssessmentProjectsOperationResult(GetAssessmentProjectsOperationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssessmentProjectsOperationResult(
            assessment_solution_id=self.assessment_solution_id,
            azure_api_version=self.azure_api_version,
            created_timestamp=self.created_timestamp,
            customer_storage_account_arm_id=self.customer_storage_account_arm_id,
            customer_workspace_id=self.customer_workspace_id,
            customer_workspace_location=self.customer_workspace_location,
            id=self.id,
            location=self.location,
            name=self.name,
            private_endpoint_connections=self.private_endpoint_connections,
            project_status=self.project_status,
            provisioning_state=self.provisioning_state,
            public_network_access=self.public_network_access,
            service_endpoint=self.service_endpoint,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            updated_timestamp=self.updated_timestamp)


def get_assessment_projects_operation(project_name: Optional[builtins.str] = None,
                                      resource_group_name: Optional[builtins.str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssessmentProjectsOperationResult:
    """
    Get a AssessmentProject

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str project_name: Assessment Project Name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:migrate:getAssessmentProjectsOperation', __args__, opts=opts, typ=GetAssessmentProjectsOperationResult).value

    return AwaitableGetAssessmentProjectsOperationResult(
        assessment_solution_id=pulumi.get(__ret__, 'assessment_solution_id'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        created_timestamp=pulumi.get(__ret__, 'created_timestamp'),
        customer_storage_account_arm_id=pulumi.get(__ret__, 'customer_storage_account_arm_id'),
        customer_workspace_id=pulumi.get(__ret__, 'customer_workspace_id'),
        customer_workspace_location=pulumi.get(__ret__, 'customer_workspace_location'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        private_endpoint_connections=pulumi.get(__ret__, 'private_endpoint_connections'),
        project_status=pulumi.get(__ret__, 'project_status'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        public_network_access=pulumi.get(__ret__, 'public_network_access'),
        service_endpoint=pulumi.get(__ret__, 'service_endpoint'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        updated_timestamp=pulumi.get(__ret__, 'updated_timestamp'))
def get_assessment_projects_operation_output(project_name: Optional[pulumi.Input[builtins.str]] = None,
                                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAssessmentProjectsOperationResult]:
    """
    Get a AssessmentProject

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str project_name: Assessment Project Name
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:migrate:getAssessmentProjectsOperation', __args__, opts=opts, typ=GetAssessmentProjectsOperationResult)
    return __ret__.apply(lambda __response__: GetAssessmentProjectsOperationResult(
        assessment_solution_id=pulumi.get(__response__, 'assessment_solution_id'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        created_timestamp=pulumi.get(__response__, 'created_timestamp'),
        customer_storage_account_arm_id=pulumi.get(__response__, 'customer_storage_account_arm_id'),
        customer_workspace_id=pulumi.get(__response__, 'customer_workspace_id'),
        customer_workspace_location=pulumi.get(__response__, 'customer_workspace_location'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        private_endpoint_connections=pulumi.get(__response__, 'private_endpoint_connections'),
        project_status=pulumi.get(__response__, 'project_status'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        public_network_access=pulumi.get(__response__, 'public_network_access'),
        service_endpoint=pulumi.get(__response__, 'service_endpoint'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        updated_timestamp=pulumi.get(__response__, 'updated_timestamp')))
