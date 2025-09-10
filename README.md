# 🌍 WanderWallet

WanderWallet — сервис для контроля и аналитики личных расходов во время путешествий.  

---

## 🚀 Основные возможности
- ✈️ Управление поездками, категориями расходов и тратами  
- 📊 Аналитика расходов  
- 🔐 JWT-аутентификация пользователей  
- 🗄️ Поддержка PostgreSQL  
- 🌐 REST API + Swagger-документация  

---

## ⚡ Быстрый старт (локальный запуск)

### 1. Клонируйте репозиторий
```bash
git clone https://github.com/<your-username>/wanderwallet.git
cd wanderwallet
```

### 2. Установите зависимости
```bash
go mod download
```

### 3. Настройте переменные окружения
Создайте файл `.env` и задайте параметры подключения, например:
```bash
  RUN_ADDRESS="localhost:3000"
  DATABASE_URI="postgres://postgres:postgres@localhost:5432/Wander_Wallet?sslmode=disable"
  SECRET_KEY=very-secret-key
```

### 4. Стартуйте приложение
Запуск приложения:
```bash
   go run cmd/wanderwallet/main.go
```

### 5. Откройте Swagger
Документация доступна по адресу:

[Swagger](http://localhost:3000/swagger/index.html#/)

(Замените порт/адрес согласно значению RUN_ADDRESS)
