# News Aggregator API

Backend system for collecting news from external APIs, publishing raw data to Kafka, processing it, and exposing cleaned news via REST APIs.

---

## Tech

Language: Go 1.25  
Framework: Gin  
Message Broker: Kafka  
External API: NewsAPI  
Architecture: Producer / Consumer  
Data Flow: API → Kafka → Service → REST API  

---

## Modules

API-getnews  
Service to call external News API and publish raw news into Kafka topic.

clawer-news  
Kafka consumer to read raw news, process it, and expose clean news via REST API.

---

## Features

- Fetch news from external APIs
- Push raw news into Kafka topics
- Consume and process Kafka messages
- Expose aggregated news via REST API
- Decoupled producer and consumer
- Scalable data pipeline

---

## Flow

1. API-getnews calls external News API  
2. Raw news is sent to Kafka topic (news.raw)  
3. clawer-news consumes messages from Kafka  
4. Data is processed and normalized  
5. Clean news is exposed via REST API  

---

## API

Trigger fetching news from external API and push to Kafka  

GET http://localhost:8080/test  

Get processed news  

GET http://localhost:8080/news  

Response  
{
  "data": [
    {
      "title": "Some news title",
      "url": "https://example.com",
      "published_at": "2026-01-12T10:00:00Z",
      "source": "BBC"
    }
  ]
}

---

## Kafka

Broker  
localhost:9092  

Topics  
news.raw  

API-getnews acts as Producer  
clawer-news acts as Consumer  

---

## Project Structure

NewsAggregatorAPI  
API-getnews/  
clawer-news/  
models/  
utils/  
docker/  

---

## Run

Start Kafka and Zookeeper  

Start producer  
cd API-getnews  
go run main.go  

Start consumer  
cd clawer-news  
go run main.go  

---

## Use

- News aggregation system  
- Real-time news pipeline  
- Data collection service  
- Kafka practice project  
- Backend for news apps  
