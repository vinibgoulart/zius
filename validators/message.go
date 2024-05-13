package validators

func MessageGet(message *string, fallback string) string {
	if *message == "" {
		return fallback
	}
	return *message
}
