/* Demonstration of the OAuth2 Client Credentials Grant performing
   (for demonstration purpose) a "list users" call to the Keycloak API
   * See also: https://oauth.net/2/grant-types/client-credentials/
   * See also: https://www.keycloak.org/docs-api/12.0/rest-api/#_users_resource */

package main

import(
  "io"
  "log"
  "os"
  "github.com/coreos/go-oidc/v3/oidc"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/clientcredentials"
  "example-service/ini"
)

var (
  clientName = "example-service"
  clientID = ""
  clientSecret = ""
  providerUrl = ""
  usersUrl = ""
)

func main() {
  arr := []ini.Ref{
    {"clientID", &clientID},
    {"clientSecret", &clientSecret},
    {"providerUrl", &providerUrl},
    {"usersUrl", &usersUrl}}

  err := ini.ReadIni(clientName, arr)

  if err != nil {
    log.Fatal()
    os.Exit(1)
  }

  ctx := context.Background()
  provider, err := oidc.NewProvider(ctx, providerUrl)

	if err != nil {
		log.Fatal(err)
	}

  config := clientcredentials.Config{
    ClientID: clientID,
    ClientSecret: clientSecret,
    TokenURL: provider.Endpoint().TokenURL,
    Scopes: []string{oidc.ScopeOpenID, "profile", "email", "roles"},
  }

  client := config.Client(ctx)
  resp, err := client.Get(usersUrl)
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  log.Printf("body = %s", body)
}
