# CRM App

A simple Customer Relationship Management (CRM) application built with:

- 🧠 **Backend**: Golang + Fiber  
- 🗄️ **Database**: PostgreSQL  
- 💻 **Frontend**: Next.js (React)

## Features

- Customer management (create, update, delete, view)
- Secure API using Fiber
- PostgreSQL for data storage
- Responsive frontend built with Next.js

## Tech Stack

- **Backend**: Go, Fiber, GORM
- **Frontend**: Next.js, React, Tailwind CSS (optional)
- **Database**: PostgreSQL

## Getting Started

### Backend (Go + Fiber)

1 setup postgresql create crm_system database and load database.sql

2. Navigate to the backend directory:
```bash
cd CRM-APP
go mod tidy
```
   
3 run the server
```bash
go run cmd/server/main.go
```
4 Frontend (Next.js)
Navigate to the frontend directory:
```bash
cd frontend
npm install
npm run dev
```

and good luck
