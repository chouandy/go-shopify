package goshopify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const shopifyChecksumHeader = "X-Shopify-Hmac-Sha256"

// AuthorizeURL returns a Shopify oauth authorization url for the given shopname and state.
//
// State is a unique value that can be used to check the authenticity during a
// callback from Shopify.
func (app App) AuthorizeURL(shopName string, state string) string {
	shopURL, _ := url.Parse(ShopBaseURL(shopName))
	shopURL.Path = "/admin/oauth/authorize"
	query := shopURL.Query()
	query.Set("client_id", app.APIKey)
	query.Set("redirect_uri", app.RedirectURL)
	query.Set("scope", app.Scope)
	query.Set("state", state)
	shopURL.RawQuery = query.Encode()
	return shopURL.String()
}

// GetAccessToken get access token
func (app App) GetAccessToken(shopName string, code string) (string, error) {
	type Token struct {
		Token string `json:"access_token"`
	}

	data := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}{
		ClientID:     app.APIKey,
		ClientSecret: app.APISecret,
		Code:         code,
	}

	client := NewClient(app, shopName, "")
	req, err := client.NewRequest("POST", "admin/oauth/access_token", data, nil)

	token := new(Token)
	err = client.Do(req, token)
	return token.Token, err
}

// VerifyMessage verify a message against a message HMAC
func (app App) VerifyMessage(message, messageMAC string) bool {
	mac := hmac.New(sha256.New, []byte(app.APISecret))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)

	// shopify HMAC is in hex so it needs to be decoded
	actualMac, _ := hex.DecodeString(messageMAC)

	return hmac.Equal(actualMac, expectedMAC)
}

// VerifyAuthorizationURL verifying URL callback parameters.
func (app App) VerifyAuthorizationURL(u *url.URL) (bool, error) {
	q := u.Query()
	messageMAC := q.Get("hmac")

	// Remove hmac and signature and leave the rest of the parameters alone.
	q.Del("hmac")
	q.Del("signature")

	message, err := url.QueryUnescape(q.Encode())

	return app.VerifyMessage(message, messageMAC), err
}

// VerifyWebhookRequest verifies a webhook http request, sent by Shopify.
// The body of the request is still readable after invoking the method.
func (app App) VerifyWebhookRequest(httpRequest *http.Request) bool {
	shopifySha256 := httpRequest.Header.Get(shopifyChecksumHeader)
	actualMac := []byte(shopifySha256)

	mac := hmac.New(sha256.New, []byte(app.APISecret))
	requestBody, _ := ioutil.ReadAll(httpRequest.Body)
	httpRequest.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	mac.Write(requestBody)
	macSum := mac.Sum(nil)
	expectedMac := []byte(base64.StdEncoding.EncodeToString(macSum))

	return hmac.Equal(actualMac, expectedMac)
}

// VerifyWebhookRequestVerbose verifies a webhook http request, sent by Shopify.
// The body of the request is still readable after invoking the method.
// This method has more verbose error output which is useful for debugging.
func (app App) VerifyWebhookRequestVerbose(httpRequest *http.Request) (bool, error) {
	if app.APISecret == "" {
		return false, errors.New("APISecret is empty")
	}

	shopifySha256 := httpRequest.Header.Get(shopifyChecksumHeader)
	if shopifySha256 == "" {
		return false, fmt.Errorf("header %s not set", shopifyChecksumHeader)
	}

	decodedReceivedHMAC, err := base64.StdEncoding.DecodeString(shopifySha256)
	if err != nil {
		return false, err
	}
	if len(decodedReceivedHMAC) != 32 {
		return false, fmt.Errorf("received HMAC is not of length 32, it is of length %d", len(decodedReceivedHMAC))
	}

	mac := hmac.New(sha256.New, []byte(app.APISecret))
	requestBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		return false, err
	}

	httpRequest.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	if len(requestBody) == 0 {
		return false, errors.New("request body is empty")
	}

	// Sha256 write doesn't actually return an error
	mac.Write(requestBody)

	computedHMAC := mac.Sum(nil)
	HMACSame := hmac.Equal(decodedReceivedHMAC, computedHMAC)
	if !HMACSame {
		return HMACSame, fmt.Errorf("expected hash %x does not equal %x", computedHMAC, decodedReceivedHMAC)
	}

	return HMACSame, nil
}
