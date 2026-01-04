```markdown
# ChipiAI Chat — Realtime AI‑Powered Chat Platform

ChipiAI Chat is a modular, production‑ready chat system built with Go, gRPC, WebSockets, MongoDB, Kafka, and Nuxt. It supports real‑time messaging, AI responses via callback gRPC, persistent chat history, and multi‑room architecture.

This project demonstrates a complete modern backend pipeline:
WebSocket → gRPC → AI Service → Callback → MongoDB → WebSocket.

---

## 🚀 Features

- **Go/Echo**
- **Realtime WebSocket chat**
- **AI message processing** via gRPC (ContextEnhancementService)
- **Callback server** for AI responses
- **MongoDB message storage**
- **Room creation and message history**
- **JWT authentication**
- **Kafka consumer integration**
- Clean, modular Go architecture

---

## 📦 Project Structure

---

## 🔧 Technologies Used

| Component | Technology |
|-----------|------------|
| Backend | Go 1.22 |
| Realtime | WebSocket (gorilla/websocket) |
| AI Processing | gRPC |
| Storage | MongoDB |
| Authentication | JWT |
| Frontend | Nuxt 3 |
| Messaging | Kafka |
| Deployment | Docker Compose |

---

## ⚙️ Setup & Installation

### 1. Clone the repository

```bash
git clone https://github.com/pasaz/ChipiAiChat.git
cd ChipiAiChat
```

### 2. Start infrastructure (Mongo, Kafka, Postgres)

```bash
docker-compose up -d
```

### 3. Run the backend

```bash
go run main.go
```

Backend starts on:

- **HTTP**: `http://localhost:8080, http://localhost:7070`
- **gRPC**: `:50051, :50052, :50053`

---

## 🔐 Authentication

The chat endpoint requires a valid JWT:

```
Authorization: Bearer <token>
```

Token is validated using:

```
JWT_KEY_CHAT
```

from `.env` file.

---

## 💬 WebSocket API

### Connect

```
ws://localhost:7070/chat?roomId=<roomId>
```

### Send message

Client sends plain text:

```json
"Hello AI"
```

### Receive AI response

Callback server pushes AI messages back through the same WebSocket.

---

## 🧠 gRPC Services

### MessageService (client → AI)

```proto
message NewMessageContent {
  string username = 1;
  string content  = 2;
  string roomId   = 3;
}

service MessageService {
  rpc Message(NewMessageContent) returns (Empty);
}
```

### AiCallback (AI → backend)

```proto
message AiMessage {
  string username = 1;
  string content  = 2;
  string roomId   = 3;
}

service AiCallback {
  rpc SendAiMessage(AiMessage) returns (google.protobuf.Empty);
}
```

---

## 🗄️ MongoDB Models

### Message

```go
type Message struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    RoomID    string             `bson:"roomId"`
    Username  string             `bson:"username"`
    Content   string             `bson:"content"`
    CreatedAt time.Time          `bson:"createdAt"`
}
```

### Room

```go
type Room struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
    CreatedAt time.Time          `bson:"createdAt"`
}
```

---

## 📡 REST API

### Create room

```
POST /create/room
```

Request body:

```json
{
  "name": "General"
}
```

Response:

```json
{
  "roomId": "6778c3f2e4b0a3c1d2f9a123"
}
```

### Get room messages

```
GET /get/room/:id
```

Response:

```json
[
  {
    "roomId": "6778c3f2e4b0a3c1d2f9a123",
    "username": "Sava",
    "content": "Hello",
    "createdAt": "2026-01-04T14:30:00Z"
  }
]
```

---

## 🧪 Kafka Consumer

The consumer listens for user messages and processes them asynchronously.

```go
go consumer.GetMessageUser(&wg)
```

---

## 🧩 Callback Logic

AI sends a gRPC callback:

```go
ws := chat.Connections[req.Username]
ws.WriteMessage(1, []byte(req.Content))

database.SaveMessage(req.RoomId, req.Username, req.Content)
```

