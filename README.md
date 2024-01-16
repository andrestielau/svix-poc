# SVIX POC
The goal of this POC is to allow:
- admins to configure notifications-types for web-hooks
- customers to read schemas for web-hook notification-types
- admins to configure schemas for web-hook notification-types
- customers to configure the endpoints they want to receive web-hooks at
- customers to configure notification-types they want for each endpoint
- services to send type-safe messages to pub-sub to be relayed to web-hooks
- admins/tests to validate payloads against json schemas 

This first part is command and consumer driven, 
meaning that the message represents an action to be executed,
and that the schema is defined by the consumer (in svix).
This requires that either:
- the consumer sends messages that are expected to fail.
  (if the customer doesn't have an app created in svix)
- the producer checks the app existence before sending the event.
- the event is filtered after being published but before being sent to svix. 
  
The extended scope of this POC is to allow:
- admins to manage schemas for their event types
- events to be mapped to notifications dynamically
- customers to manage their subscribe to notification-types
- admins to manage transformations and links between events and notifications
- customers to manage the emails they want to receve notifications at

# App

## WebHooks
### API
- Read WebHookTypes
- Manage Endpoints
- Read Sent Messages
- Read Message Attempts
### GRPC
- Manage Apps
- Manage WebHook Schemas
### Subscriber
- Request WebHooks

## Router
### API
- Manage Notification Subscriptions
### GRPC
- Manage Event Schemas
### Subscriber
- Transform payloads between topics
- Publish Notifications for Events into common topic
- Publish Commands for Notifications into provider topics

## Email
### API
- Manage Email
### GRPC
- Manage Template Schemas
### Subscriber
- Render Template
- Add Attachments
- Send Emails

# Package
## app
## files
## schema
## topic
## utils

# Renderer

# Demo