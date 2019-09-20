package config

type AuthConfig struct {
	settings *AuthConfigSettings
}

type AuthConfigSettings struct {
	hmacSecret string
}

func GetAuthConfig() *AuthConfig {
	return &AuthConfig{
		settings: &AuthConfigSettings{
			// https://mkjwk.org
			hmacSecret: "z5qkREmStJx-IFI96YH8V7jivw78FTk6ZCtDCtq-RsYik6WopxYODgZTlWxUd_jjs6-X041ZdDn4p27lWuY3fQ",
		},
	}
}
