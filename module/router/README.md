# Router

# Schema Design Approach
## Consumer Driven
For when consumer schemas are harder to change.
For example, when interacting with external webhooks, 
or when the consumer has specific field requirements.

## Producer Driven
For when producer schemas are harder to change-
For example: when events are direct copies of a log tail/outbox or an external source, 
and don't have any mapping done before sending the event, 
or if we want to centralize where that mapping happens.

# Patterns
## Event to Command
Producer --> Router --> Consumer

## Fan In
Producers --> Router --> Consumer

## Fan Out
Producer --> Router --> Consumers

## Multicast
Producers --> Router --> Consumers

## Saga Manager 
Producer --> Router --> Consumers
             Router <-- Producers
             Router --> Consumers
Consumer <-- Router <-- Producers

## Event to Events to Command
Producer --> Consumer