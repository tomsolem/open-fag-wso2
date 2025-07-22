package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	. "github.com/openfga/go-sdk/client"
)

var (
  // Do not do this in production.
  // In production, you would have the private key and public key pair generated
  // in advance. NEVER add a private key to any GitHub repo.
  privateKey *rsa.PrivateKey
)

func main() {
  app := fiber.New()

  // Just as a demo, generate a new private/public key pair on each run.
  rng := rand.Reader
  var err error
  privateKey, err = rsa.GenerateKey(rng, 2048)
  if err != nil {
    log.Fatalf("rsa.GenerateKey: %v", err)
  }
  // create a JWT token with the private key
  token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
    "name": "anne",
  })
  // Sign the token with the private key
  signedToken, err := token.SignedString(privateKey)
  if err != nil {
    log.Fatalf("token.SignedString: %v", err)
  }
  log.Printf("curl -H 'Authorization: Bearer %s' http://localhost:3000/read/1", signedToken)

  // JWT Middleware
  app.Use(jwtware.New(jwtware.Config{
    SigningMethod: "RS256",
    SigningKey:    privateKey.Public(),
  }))

  app.Use("/read/:document", preauthorize)

  app.Use(checkAuthorization)

  app.Get("/read/:document", read)

  app.Listen(":3000")
}

func read(c *fiber.Ctx) error {
  user := c.Locals("user").(*jwt.Token)
  claims := user.Claims.(jwt.MapClaims)
  name := claims["name"].(string)
  return c.SendString(name + " read " + c.Params("document"))
}

func preauthorize(c *fiber.Ctx) error {
  // get the user name from JWT
  user := c.Locals("user").(*jwt.Token)
  claims := user.Claims.(jwt.MapClaims)
  name := claims["name"].(string)
  c.Locals("username", fmt.Sprintf("user:%s",name))

  // parse the HTTP method
  switch (c.Method()) {
    case "GET":
      c.Locals("relation", "can_view")
    case "POST":
      c.Locals("relation", "can_write")
    case "DELETE":
      c.Locals("relation", "owner")
    default:
      c.Locals("relation", "owner")
  }

  // get the object name and prepend with type name "document:"
  c.Locals("object", "document:" + c.Params("document"))
  return c.Next()
}

// Middleware to check whether user is authorized to access document
func checkAuthorization(c *fiber.Ctx) error {
  fgaClient, err := NewSdkClient(&ClientConfiguration{
    ApiUrl:               os.Getenv("FGA_API_URL"), // required, e.g. are set in the docker compose file
    StoreId:        os.Getenv("FGA_STORE_ID"), // run this command to set new values:  source /workspace/open-fga/configure-model.sh 
    AuthorizationModelId: os.Getenv("FGA_MODEL_ID"),  // optional, can be overridden per request
  })

  if err != nil {
    return fiber.NewError(fiber.StatusServiceUnavailable, "Unable to build OpenFGA client")
  }

  body := ClientCheckRequest{
    User: c.Locals("username").(string),
    Relation: c.Locals("relation").(string),
    Object: c.Locals("object").(string),
  }
  data, err := fgaClient.Check(context.Background()).Body(body).Execute()

  if err != nil {
    return fiber.NewError(fiber.StatusServiceUnavailable, "Unable to check for authorization")
  }

  if !(*data.Allowed) {
    return fiber.NewError(fiber.StatusForbidden, "Forbidden to access document")
  }

  // Go to the next middleware
  return c.Next()
}