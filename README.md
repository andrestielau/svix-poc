# SVIX POC
This repository contains example code to access svix


# High Level
The concept to prove is the decoupled management of message delivery and filtering of sent messages.

```mermaid
graph TD
    subgraph Mock
        subgraph Org1
            Endpoint1
        end
        subgraph Org2
            Endpoint2
            Endpoint3
        end
    end
    subgraph Svix
        subgraph Apps
            App1
            App2
        end
        subgraph EventTypes
            direction LR
            Withdrawal
            Deposit
            KYCStarted
            KYCApproved
        end
    end
    subgraph Subscriptions
      SA[API]
      SS[Subscriber]
      SR[Repository]
      SvixClient
    end
    App1 -- "KYCApproved,\nDeposit" --> Endpoint1
    App2 -- "Withdrawal,\nDeposit" --> Endpoint2
    App2 -- "KYCStarted,\nKYCApproved" --> Endpoint3
    SvixClient -- "KYCApproved,\nDeposit" --> App1
    SvixClient -- "Withdrawal,\nDeposit,\nKYCStarted,\nKYCApproved" --> App2
    SvixClient --> EventTypes
    Demo --> SA
    SA --> SR
    PubSub --> SS
    SS --> SR
    SS --> SvixClient
    SA --> SvixClient
    Demo --> PubSub

```

# WebHooks Config
This configuration is used by mocks to get the port to emulate, and by the demo to know which ports to register in svix.
 
```yaml
# .wh.yml
events:
  - name: Withdrawal
    schemas:
      1:  {
        opId: 'string',
        status: 'string'
      }
  - name: Deposit
    schemas: 
      1:  {
        opId: 'string',
        status: 'string'
      }
  - name: KYCStarted
    schemas: 
      1: {
        affiliateId: 'string'
      }
  - name: KYCApproved
    schemas: 
      1: {
        affiliateId: 'string'
      }
apps:
  - def:
      name: orgId1
    endpoints:
      - url: http://127.0.0.1:3428
        secret: 
          value: ...
          isSet: true
        filterTypes: 
        - KYCApproved
        - Deposit
  - def: 
      name: orgId2
    endpoints:
      - url: http://127.0.0.1:2574
        secret: 
          value: ...
          isSet: true
        filterTypes: 
        - Withdrawal
        - Deposit
      - url: http://127.0.0.1:3526
        secret: 
          value: ...
          isSet: true
        filterTypes: 
        - KYCStarted
        - KYCApproved
```

# Test
- Start Docker: `docker-compose up -d`
- Check token service logs to get auth token
- Run demo `go run . call demo -t {auth token}`
- Check logs for mock server

## Message Config
```yaml
# .msg.yml
orgId1:
  - eventType: Withdrawal
    payload: 
      opId: foo
      status: bar
  - eventType: Deposit
    payload: 
      opId: foo
      status: bar
  - eventType: KYCStarted
    payload: 
      affiliateId: asd
  - eventType: KYCApproved
    payload: 
      affiliateId: asd
orgId2:
  - eventType: Withdrawal
    payload: 
      opId: foo
      status: bar
  - eventType: Deposit
    payload: 
      opId: foo
      status: bar
  - eventType: KYCStarted
    payload: 
      affiliateId: asd
  - eventType: KYCApproved
    payload: 
      affiliateId: asd
```

# Phase Dependencies
## 1 - Prepare
- Subscribe DB
- Subscribe API
## 2 - Config
- WebHook API
## 3 - Subscribe
- WebHook GRPC
## 4 - Publish
- Subscribe PubSub
## 5 - Sync
- Subscribe GRPC

# Adding New Event Types
## Create Event Mapper/Stream Processor (Go)
## Create Event Types (Svix)
## Create EventNotification (DB)

# Missing
## Sync between svix event types and notification targets
## Validate if user has provider enabled
## Filtering based on content