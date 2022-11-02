# example-service

## Installation

```bash
git clone https://github.com/b1-systems/example-service.git
cd example-service
go build
```

## Configuration

1. In Keycloak, create an openid-connect client, assigning a client ID, for example "example-service".
   * Set Access Type to "confidential".
   * Set Service Accounts Enabled to "ON".
   * Make sure that Standard Flow Enabled is set to "ON" (the default, other flows can be disabled).
   * Save the client.
   * Change to tab "Credentials" and retrieve the generated client secret.

2. In Keycloak, assign service account roles to the client.
   * Change to Clients -> "example-service" -> Tab "Service Account Roles".
   * To perform the example action of this client (listing users):
      - assign to it the client "realm-management" roles "query-users" and "view-users".

3. Create a configuration file `example-service.ini`:

```bash
cp example-service.ini.sample example-service.ini
```

Example `example-service.ini`:

```
[example-service]
clientID = example-service
clientSecret = secret
providerUrl = https://www.example.test/keycloak/realms/golang-oidc
usersUrl = https://www.example.test/keycloak/admin/realms/golang-oidc/users
```

# Usage

```bash
./example-service
```

