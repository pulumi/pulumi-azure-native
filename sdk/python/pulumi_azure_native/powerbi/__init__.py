# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .get_private_endpoint_connection import *
from .get_workspace_collection import *
from .list_workspace_collection_access_keys import *
from .power_bi_resource import *
from .private_endpoint_connection import *
from .workspace_collection import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.powerbi.v20160129 as __v20160129
    v20160129 = __v20160129
    import pulumi_azure_native.powerbi.v20200601 as __v20200601
    v20200601 = __v20200601
else:
    v20160129 = _utilities.lazy_import('pulumi_azure_native.powerbi.v20160129')
    v20200601 = _utilities.lazy_import('pulumi_azure_native.powerbi.v20200601')

