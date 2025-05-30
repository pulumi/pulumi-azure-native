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
    'GetWorkflowResult',
    'AwaitableGetWorkflowResult',
    'get_workflow',
    'get_workflow_output',
]

@pulumi.output_type
class GetWorkflowResult:
    """
    The workflow type.
    """
    def __init__(__self__, access_control=None, access_endpoint=None, azure_api_version=None, changed_time=None, created_time=None, definition=None, endpoints_configuration=None, id=None, identity=None, integration_account=None, integration_service_environment=None, location=None, name=None, parameters=None, provisioning_state=None, sku=None, state=None, tags=None, type=None, version=None):
        if access_control and not isinstance(access_control, dict):
            raise TypeError("Expected argument 'access_control' to be a dict")
        pulumi.set(__self__, "access_control", access_control)
        if access_endpoint and not isinstance(access_endpoint, str):
            raise TypeError("Expected argument 'access_endpoint' to be a str")
        pulumi.set(__self__, "access_endpoint", access_endpoint)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if changed_time and not isinstance(changed_time, str):
            raise TypeError("Expected argument 'changed_time' to be a str")
        pulumi.set(__self__, "changed_time", changed_time)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if definition and not isinstance(definition, dict):
            raise TypeError("Expected argument 'definition' to be a dict")
        pulumi.set(__self__, "definition", definition)
        if endpoints_configuration and not isinstance(endpoints_configuration, dict):
            raise TypeError("Expected argument 'endpoints_configuration' to be a dict")
        pulumi.set(__self__, "endpoints_configuration", endpoints_configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if integration_account and not isinstance(integration_account, dict):
            raise TypeError("Expected argument 'integration_account' to be a dict")
        pulumi.set(__self__, "integration_account", integration_account)
        if integration_service_environment and not isinstance(integration_service_environment, dict):
            raise TypeError("Expected argument 'integration_service_environment' to be a dict")
        pulumi.set(__self__, "integration_service_environment", integration_service_environment)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="accessControl")
    def access_control(self) -> Optional['outputs.FlowAccessControlConfigurationResponse']:
        """
        The access control configuration.
        """
        return pulumi.get(self, "access_control")

    @property
    @pulumi.getter(name="accessEndpoint")
    def access_endpoint(self) -> builtins.str:
        """
        Gets the access endpoint.
        """
        return pulumi.get(self, "access_endpoint")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="changedTime")
    def changed_time(self) -> builtins.str:
        """
        Gets the changed time.
        """
        return pulumi.get(self, "changed_time")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> builtins.str:
        """
        Gets the created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def definition(self) -> Optional[Any]:
        """
        The definition.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="endpointsConfiguration")
    def endpoints_configuration(self) -> Optional['outputs.FlowEndpointsConfigurationResponse']:
        """
        The endpoints configuration.
        """
        return pulumi.get(self, "endpoints_configuration")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The resource id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedServiceIdentityResponse']:
        """
        Managed service identity properties.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="integrationAccount")
    def integration_account(self) -> Optional['outputs.ResourceReferenceResponse']:
        """
        The integration account.
        """
        return pulumi.get(self, "integration_account")

    @property
    @pulumi.getter(name="integrationServiceEnvironment")
    def integration_service_environment(self) -> Optional['outputs.ResourceReferenceResponse']:
        """
        The integration service environment.
        """
        return pulumi.get(self, "integration_service_environment")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Gets the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.WorkflowParameterResponse']]:
        """
        The parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Gets the provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        The sku.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def state(self) -> Optional[builtins.str]:
        """
        The state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Gets the resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        Gets the version.
        """
        return pulumi.get(self, "version")


class AwaitableGetWorkflowResult(GetWorkflowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkflowResult(
            access_control=self.access_control,
            access_endpoint=self.access_endpoint,
            azure_api_version=self.azure_api_version,
            changed_time=self.changed_time,
            created_time=self.created_time,
            definition=self.definition,
            endpoints_configuration=self.endpoints_configuration,
            id=self.id,
            identity=self.identity,
            integration_account=self.integration_account,
            integration_service_environment=self.integration_service_environment,
            location=self.location,
            name=self.name,
            parameters=self.parameters,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            state=self.state,
            tags=self.tags,
            type=self.type,
            version=self.version)


def get_workflow(resource_group_name: Optional[builtins.str] = None,
                 workflow_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkflowResult:
    """
    Gets a workflow.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2015-02-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The resource group name.
    :param builtins.str workflow_name: The workflow name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workflowName'] = workflow_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:logic:getWorkflow', __args__, opts=opts, typ=GetWorkflowResult).value

    return AwaitableGetWorkflowResult(
        access_control=pulumi.get(__ret__, 'access_control'),
        access_endpoint=pulumi.get(__ret__, 'access_endpoint'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        changed_time=pulumi.get(__ret__, 'changed_time'),
        created_time=pulumi.get(__ret__, 'created_time'),
        definition=pulumi.get(__ret__, 'definition'),
        endpoints_configuration=pulumi.get(__ret__, 'endpoints_configuration'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        integration_account=pulumi.get(__ret__, 'integration_account'),
        integration_service_environment=pulumi.get(__ret__, 'integration_service_environment'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        parameters=pulumi.get(__ret__, 'parameters'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sku=pulumi.get(__ret__, 'sku'),
        state=pulumi.get(__ret__, 'state'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        version=pulumi.get(__ret__, 'version'))
def get_workflow_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        workflow_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkflowResult]:
    """
    Gets a workflow.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2015-02-01-preview, 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The resource group name.
    :param builtins.str workflow_name: The workflow name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workflowName'] = workflow_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:logic:getWorkflow', __args__, opts=opts, typ=GetWorkflowResult)
    return __ret__.apply(lambda __response__: GetWorkflowResult(
        access_control=pulumi.get(__response__, 'access_control'),
        access_endpoint=pulumi.get(__response__, 'access_endpoint'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        changed_time=pulumi.get(__response__, 'changed_time'),
        created_time=pulumi.get(__response__, 'created_time'),
        definition=pulumi.get(__response__, 'definition'),
        endpoints_configuration=pulumi.get(__response__, 'endpoints_configuration'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        integration_account=pulumi.get(__response__, 'integration_account'),
        integration_service_environment=pulumi.get(__response__, 'integration_service_environment'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        parameters=pulumi.get(__response__, 'parameters'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sku=pulumi.get(__response__, 'sku'),
        state=pulumi.get(__response__, 'state'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        version=pulumi.get(__response__, 'version')))
