# apigateway
API Gateway service provides a unified web backend API across all cloudnativedesign microservices for both web and mobile clients using a GraphQL API

## Service configuration
```(shell)
git clone https://github.com/cloudnativedesign/apigateway.git

go mod download

go run server.go
```

## Architecture
The service implements the API gateway pattern to help us abstract the complex interactions across multiple backend microservices around a given POJO.

See [micoservice Pattern : API Gateway](https://microservices.io/patterns/apigateway.html) for more details on the pattern

### Implementation
The API is designed as a GraphQL API under a schema first approach using code generator frameworks. This allows us to use this touchpoint between microservice backend and UI to quickly iterate over user requirements while reducing the need to reimplement code in the resolvers. See [qlgen](https://github.com/99designs/gqlgen)



