package gen

func additionalRequiredContainers(resourceTok string) []string {
	switch resourceTok {
	case "azure-native:app:ManagedEnvironment":
		return []string{"properties"}
	}
	return nil
}
