import pulumi
import pulumi_random as random
from pulumi_azurerm import resources, storage

random_string = random.RandomString("random_string", length=12, special=False, upper=False)

resource_group = resources.v20200601.ResourceGroup('resource_group', resource_group_name=random_string.result, location='westus')

storage_account = storage.v20190601.StorageAccount('sa',
    account_name=random_string.result,
    resource_group_name=resource_group.name,
    location='westus2',
    sku={
        'name': 'Standard_LRS',
        'tier': 'Standard',
    },
    kind='StorageV2')
