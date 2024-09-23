package guards

func UrlSyntacticallyValid(url string) {
	if url == "" {
		panic("Url cannot be empty")
	}

	if url[:7] != "http://" && url[:8] != "https://" {
		panic("Url must start with http:// or https://")
	}
}
