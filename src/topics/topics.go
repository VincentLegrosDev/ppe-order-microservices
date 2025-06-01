package topics 

type TopicName string

const (
	OrderReceived        		TopicName = "OrderReceived"
	OrderConfirmed       		TopicName = "OrderConfirmed"
	OrderPackedAndPicked 		TopicName = "OrderPackedAndPicked"
	Error 							 		TopicName = "Error"
	Notification         		TopicName = "Notification"
	DeadQueueLetter      		TopicName = "DeadQueueLetter"
	OrderTimeMetric 				TopicName = "OrderTimeMetric"
	OrderCountMetric				TopicName = "OrderCountMetric"
) 

func (topic TopicName) String() string {
	return string(topic)
}
