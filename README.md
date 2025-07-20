# Distributed Logging System with Go, Kafka, PostgreSQL, gRPC, Prometheus & Grafana

## 📝 Project Overview

This project is a **Distributed Logging System** designed to collect logs from multiple services. It stores logs in **Kafka**, processes them, and saves them into **PostgreSQL**. The system includes a **gRPC log collector**, a **log processor**, and a **RESTful API** for viewing logs. Real-time monitoring is enabled through **Prometheus** and **Grafana**.

## 🧩 Services

* **log\_collector**: A gRPC service that receives logs and pushes them to Kafka.
* **log\_processor**: Consumes logs from Kafka and stores them in PostgreSQL.
* **log\_dashboard**: A REST API (Gin + PostgreSQL) for viewing logs.
* **prometheus**: Collects performance and metrics data from services.
* **grafana**: Visualizes real-time metrics via dashboards.

## 🚀 Running the Project

Ensure you have Docker and Docker Compose installed. Then run:

```bash
docker-compose up --build
```

This will build and start all necessary services.

## 🌐 Access Points

| Service        | URL / Port                                               |
| -------------- | -------------------------------------------------------- |
| REST API       | [http://localhost:8080/logs](http://localhost:8080/logs) |
| gRPC Collector | Internal port: 50051                                     |
| PostgreSQL     | Internal port: 5432                                      |
| Kafka          | Internal port: 9092                                      |
| Prometheus     | [http://localhost:9090](http://localhost:9090)           |
| Grafana        | [http://localhost:3000](http://localhost:3000)           |

**Grafana Login Credentials**:

* **Username**: `admin`
* **Password**: `admin` *(or as set in `docker-compose.yml`)*

## 📊 Monitoring

Prometheus gathers metrics such as:

* Number of processed logs
* Error rates
* Kafka lag and throughput
* gRPC request durations

Grafana visualizes these metrics in real-time, allowing easy tracking of system performance and health.

---

Feel free to contribute or raise issues. This system can be extended with authentication, alerting, and advanced analytics features.
