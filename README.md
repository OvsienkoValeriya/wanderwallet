# 🌍 WanderWallet

WanderWallet is a backend service for tracking and analyzing personal travel expenses.

The application is built with Go, deployed on AWS EC2, and uses Amazon SNS and Amazon SQS for asynchronous processing of expense-related events.

## 🚀 Features

- ✈️ Manage trips, expense categories, and individual expenses
- 📊 View travel expense analytics
- 🔐 JWT-based user authentication and authorization
- 🗄️ PostgreSQL data storage
- 🌐 REST API with Swagger/OpenAPI documentation
- 📨 Event-driven processing with Amazon SNS and Amazon SQS
- ⚙️ Separate background worker for processing SQS messages
- ☁️ Deployment on Amazon EC2
- 🔑 Secure AWS access through IAM roles and policies
- 🐳 Dockerized API and worker services

## 🏗️ Architecture

WanderWallet consists of two main components:

- **API service** — handles authentication, trips, expenses, categories, and analytics
- **Expense worker** — consumes and processes expense-related events from Amazon SQS

Expense events are published to an Amazon SNS topic and delivered to an Amazon SQS queue for asynchronous processing.

The application is hosted on Amazon EC2. IAM roles and policies are used to grant the EC2 instance access to AWS services without storing AWS credentials directly in the application.

## ⚡ Local Setup

### 1. Clone the repository

```bash
git clone https://github.com/knitleria/wanderwallet.git
cd wanderwallet
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment variables
Create a .env file and specify the required configuration values:
```bash
  RUN_ADDRESS="localhost:3000"
  DATABASE_URI="postgres://postgres:postgres@localhost:5432/Wander_Wallet?sslmode=disable"
  SECRET_KEY=very-secret-key
```
AWS-related environment variables may also be required when running the event-processing functionality locally.

### 4. Start the application
```bash
   go run cmd/wanderwallet/main.go
```

### 5. Open Swagger documentation
Once the application is running, Swagger documentation is available at:

[Swagger](http://localhost:3000/swagger/index.html#/)
Replace the host and port according to the value of RUN_ADDRESS.

🛠️ Technologies
- Go
- Gin
- PostgreSQL
- GORM
- JWT
- REST API
- Swagger 
- Amazon EC2
- Amazon SNS
- Amazon SQS
- AWS IAM
- Docker
- GoMock
- Testify


