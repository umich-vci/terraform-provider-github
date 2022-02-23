package github

// GenerateOAuthAppToken generates a GitHub OAuth App access token from a set of valid GitHub OAuth App credentials.
func GenerateOAuthAppToken(baseURL, clientID, clientSecret string) (string, error) {

	token, err := getInstallationAccessToken(baseURL, clientID, clientSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
