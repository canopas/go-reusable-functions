# JWT

JWT is useful for authorizing resources or content of particular applications. 
Once the user is logged in, each subsequent request will include the JWT access token, allowing the user to access routes, services, and resources that are permitted with that token. 

Read [JWT Introduction](https://jwt.io/introduction) for more information.

Read a detailed article on [JWT authentication and its working](https://blog.canopas.com/jwt-in-golang-how-to-implement-token-based-authentication-298c89a26ffd).

## Generating secret keys

- For generating tokens, we need a separate secret key for both tokens (or the same key can also be used for both tokens). 
- You can generate it with HSA 256 OR RS256 encryption. It should at least be 32 characters long, but longer is better.
- Follow [How to generate secret keys for JWT?](https://mojitocoder.medium.com/generate-a-random-jwt-secret-22a89e8be00d) to create secret keys

## Example

Find a working example of JWT-reusable at [JWT-Example](https://github.com/canopas/go-reusables/blob/main/examples/jwtAuth.go).
