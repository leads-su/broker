# Broker Package for Go Lang
This package provides ability for internal application messages management.

## Creating New Broker
In order to initialize new instance of broker, you just have to add the following line of code
```go
brokerInstance := broker.NewBroker()
```

## Starting Broker
To start broker, you need to call `Start()` method in a separate goroutine, this will only start broker, which will have no channels to subscribe to.
```go
go brokerInstance.Start()
```

## Subscribing To Channel
To subscribe to a new channel, the following piece of code can be used.
```go
channel := brokerInstance.Subscribe()
```

## Working With The Channel
### Publishing messages to the channel
To publish new message to the channel, you can call `Publish()` method which accepts `interface{}` as a parameter for a value.
```go
channel.Publish("this can be string/number/structure/another interface/anything")
```

### Listening to messages in the channel
Before you will start publishing messages to the newly created channel, it is recommented to create logic responsible for retrival of these messages, here is an example
```go
for {
    switch <-channel {
        case "string":
            fmt.Println("Retrieved string `string` from the channel")
        case "number":
            fmt.Println("Retrieved string `number` from the channel")
        case "array":
            fmt.Println("Retrieved string `array` from the channel")
        case "struct":
            fmt.Println("Retrieved string `struct` from the channel")
        case "interface":
            fmt.Println("Retrieved string `interface` from the channel")
    }
}
```

## Complete Example
```go

func main() {
    brokerInstance, channel := initializeBroker()

    go func() {
        switch <-channel {
            case "string":
                fmt.Println("Retrieved string `string` from the channel")
            case "number":
                fmt.Println("Retrieved string `number` from the channel")
            case "array":
                fmt.Println("Retrieved string `array` from the channel")
            case "struct":
                fmt.Println("Retrieved string `struct` from the channel")
            case "interface":
                fmt.Println("Retrieved string `interface` from the channel")
        }
    }()

    channel.Publish("string")
    channel.Publish("array")
    channel.Publish("number")
    channel.Publish("interface")
    channel.Publish("struct")
}

func initializeBroker() (*broker.Broker, chan interface{}) {
    brokerInstance := broker.NewBroker()
    go brokerInstance.Start()
    channel := brokerInstance.Subscribe()
    return brokerInstance, channel
}
```