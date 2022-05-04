package broker

// Broker represents structure of broker
type Broker struct {
	stopChannel        chan struct{}
	publishChannel     chan interface{}
	subscribeChannel   chan chan interface{}
	unsubscribeChannel chan chan interface{}
}

// NewBroker creates new instance of broker
func NewBroker() *Broker {
	return &Broker{
		stopChannel:        make(chan struct{}),
		publishChannel:     make(chan interface{}, 1),
		subscribeChannel:   make(chan chan interface{}, 1),
		unsubscribeChannel: make(chan chan interface{}, 1),
	}
}

// Start starts broker
func (broker *Broker) Start() {
	subscribers := map[chan interface{}]struct{}{}
	for {
		select {
		case <-broker.stopChannel:
			return
		case messageChannel := <-broker.subscribeChannel:
			subscribers[messageChannel] = struct{}{}
		case messageChannel := <-broker.unsubscribeChannel:
			delete(subscribers, messageChannel)
		case message := <-broker.publishChannel:
			for messageChannel := range subscribers {
				select {
				case messageChannel <- message:
				default:
				}
			}
		}
	}
}

// Stop stops broker
func (broker *Broker) Stop() {
	close(broker.stopChannel)
}

// Subscribe creates new channel subscription
func (broker *Broker) Subscribe() chan interface{} {
	messageChannel := make(chan interface{}, 5)
	broker.subscribeChannel <- messageChannel
	return messageChannel
}

// Unsubscribe removes channel subscription
func (broker *Broker) Unsubscribe(messageChannel chan interface{}) {
	broker.unsubscribeChannel <- messageChannel
}

// Publish publishes new message
func (broker *Broker) Publish(message interface{}) {
	broker.publishChannel <- message
}
