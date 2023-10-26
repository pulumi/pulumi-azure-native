package gen

func additionalRequiredContainers(resourceTok string) RequiredContainers {
	switch resourceTok {
	case "azure-native:app:ManagedEnvironment":
		return [][]string{{"properties"}}
	}
	return nil
}
