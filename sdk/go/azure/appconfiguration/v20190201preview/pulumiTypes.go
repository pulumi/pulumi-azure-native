


package v20190201preview

type ApiKeyResponse struct {
	ConnectionString string `pulumi:"connectionString"`
	Id               string `pulumi:"id"`
	LastModified     string `pulumi:"lastModified"`
	Name             string `pulumi:"name"`
	ReadOnly         bool   `pulumi:"readOnly"`
	Value            string `pulumi:"value"`
}

func init() {
}
