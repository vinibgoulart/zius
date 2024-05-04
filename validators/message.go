package validators

func GetMessage(message *string, fallback string) string {
	if *message == "" {
		return fallback
	}
	return *message
}
