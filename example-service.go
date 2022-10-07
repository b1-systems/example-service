/* Demonstration of the OAuth2 Client Credentials Grant
   see also: https://oauth.net/2/grant-types/client-credentials/ */

package main

import(
  "io"
  "log"
  "os"
  "path/filepath"
  "github.com/coreos/go-oidc/v3/oidc"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/clientcredentials"
  "gopkg.in/ini.v1"
)

var (
  clientName = "example-service"
  clientID = ""
  clientSecret = ""
  providerUrl = ""
  usersUrl = ""
)

func readIni() {
  ex, err := os.Executable()

  if err != nil {
    panic(err)
  }

  cfg, err := ini.Load(filepath.Join(filepath.Dir(ex), clientName + ".ini"))

  if err != nil {
    panic(err)
  }

  cs := cfg.Section(clientName)

  clientID = cs.Key("clientID").String()

  if clientID == "" {
    log.Fatal(clientName + ".ini does not specify clientID")
    os.Exit(1)
  }

  clientSecret = cs.Key("clientSecret").String()

  if clientSecret == "" {
    log.Fatal(clientName + ".ini does not specify clientSecret")
    os.Exit(1)
  }

  providerUrl = cs.Key("providerUrl").String()

  if providerUrl == "" {
    log.Fatal(clientName + ".ini does not specify providerUrl")
    os.Exit(1)
  }

  usersUrl = cs.Key("usersUrl").String()

  if usersUrl == "" {
    log.Fatal(clientName + ".ini does not specify usersUrl")
    os.Exit(1)
  }

  log.Printf(
    "Read configuration:\n" +
    " clientID = %s\n" +
    " clientSecret = %s\n" +
    " providerUrl = %s\n" +
    " usersUrl = %s\n",
    clientID,
    "*REDACTED*",
    providerUrl,
    usersUrl,
  )
}

func main() {
  readIni()

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

  /* -DEBUG-
  token, err := config.Token(ctx)

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  } else {
    log.Printf("token = %s\n", token)
  }
  */

  client := config.Client(ctx)
  resp, err := client.Get(usersUrl)
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  log.Printf("body = %s", body)
}
