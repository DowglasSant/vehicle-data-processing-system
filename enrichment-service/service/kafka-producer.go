package service

import "enrichment-service/model"

type MessageProducer interface {
	ProduceMessage(vehicle *model.Vehicle, topicType string) error
}
