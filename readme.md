# Vehicle Data Processing System 🚗

## Overview

This Vehicle Data Processing System is a comprehensive solution designed for efficient handling, enrichment, and storage of vehicle-related data. Leveraging a microservices architecture, it integrates Kafka for message queuing, MongoDB and PostgreSQL for data storage, among other state-of-the-art technologies. This system is a prime example of a scalable, distributed approach to real-time data processing and management.

## Architecture 🏗️

The system is built on three core microservices:

- **Receiver Service**: Receives incoming vehicle data and publishes it to a Kafka topic.
- **Enrichment Service**: Consumes vehicle data from Kafka, enriches it with additional information (e.g., manufacturer details from PostgreSQL), and forwards it to another Kafka topic.
- **Insertion Service**: Retrieves the enriched vehicle data from Kafka and stores it in MongoDB.

This architecture ensures services are loosely coupled and highly scalable, orchestrated by Docker for smooth deployment and Kafka for efficient message distribution.

## Technologies Used 🔧

- [Kafka](https://kafka.apache.org/) for message queuing and inter-service communication.
- [MongoDB](https://www.mongodb.com/) as a NoSQL database for storing enriched vehicle data.
- [PostgreSQL](https://www.postgresql.org/) for storing additional data for enrichment.
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/) for containerization and simplified deployment.
- [Golang](https://golang.org/) for developing efficient and reliable microservices.
- [Confluent Kafka Docker Images](https://hub.docker.com/u/confluentinc) for Kafka and Zookeeper container management.
- [Kafka UI](https://github.com/provectus/kafka-ui) for a user-friendly interface to visualize Kafka topics and message flows.

## Getting Started 🚀

### Prerequisites

Ensure Docker and Docker Compose are installed on your system.

### Setup

1. **Clone the repository**
   
   ```bash
   git clone https://github.com/DowglasSant/vehicle-data-processing-system.git
   cd vehicle-data-processing-system
