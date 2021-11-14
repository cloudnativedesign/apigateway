# apigateway
API Gateway service provides a unified web backend API across all cloudnativedesign microservices for both web and mobile clients using a GraphQL API

It allows :
* schema first design of the API
* effective migration managment for iterative design
* optimized caching of the response objects 
* integrated queries across multiple providers
* schema validation on item creation

## Service usage
### Setting up the service
Download the repository code, initiate the required go module dependencies locally and start the server to set up your local development environment. 

```(shell)
git clone https://github.com/cloudnativedesign/apigateway.git

go mod download

go run server.go
```

### Changing the schema
Start by adapting the Types in `schema.graphqls` along with the required mutations and inputs.

Then compile to generate the `model files` and the `resolver stumps`
```shell
go run github.com/99designs/gqlgen generate
```

Complete the code in `schema.resolver.go` with the custom logic to resolve the request to the abstracted microservice APIs using `gRPC` calls for internal resolution on service-to-service communication.


## Architecture
The service implements the API gateway pattern to help us abstract the complex interactions across multiple backend microservices around a given POJO.

See [micoservice Pattern : API Gateway](https://microservices.io/patterns/apigateway.html) for more details on the pattern

### Implementation
The API is designed as a GraphQL API under a schema first approach using code generator frameworks. This allows us to use this touchpoint between microservice backend and UI to quickly iterate over user requirements while reducing the need to reimplement code in the resolvers. See [qlgen](https://github.com/99designs/gqlgen)



