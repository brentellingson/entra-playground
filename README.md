# entra-playground

Azure Entra OAuth 2.0 and GraphQL Tests



## Flows

This application demonstrates two ways to call the Microsoft GraphQL from a WebAPI

### Client Credential Flow

The client credential flow demonstrates system-to-system calls to the Microsoft GraphQL

1.  user call Web API anonymously
2.  WebAPI requests a system token from Entra using the Client Credential flow
3.  WebAPI calls GraphQL using system token

```mermaid
flowchart LR
    Browser
    API["`Web API
        /graphql/client-credentials`"]
    Entra
    GraphQL

    Browser -->|1 unauthenticated| API
    API -->|2 client-credential-grant| Entra
    API -->|3 client token| GraphQL
```

### On-Behalf-Of Flow

The on-behalf-of flow demonstrates a mediated user call to the Microsoft GraphQL

1. user requests a user token froom Entra using Authorization Code flow
2. user calls web API using user token
3. WebAPI requests an impersonation token from Entra using on-behalf-of flow
4. WebAPI calls GraphQL using impersonation token

```mermaid
flowchart LR
    Browser
    API["`Web API
        /graphql/on-behalf-of`"]
    Entra
    GraphQL

    Browser -->|1 authorization code| Entra
    Browser -->|2 user token| API
    API -->|3 on-behalf-of| Entra
    API -->|4 impersonation token| GraphQL
```

## Endpoints:

Endpoints to test the flow:
* `/graphql/client-credential` : call microsoft graphql with client credential flow
* `/graphql/on-behalf-of`: call microsoft graphql with on-behalf-of flow

Endpoints to mediate and test the authorization code flow:

* `/oauth/validate`: validate the oauth jwt access token
* `/oauth/authorize`: oauth authorize endpoint mediator
* `/oauth/token`: oauth token endpont mediator 