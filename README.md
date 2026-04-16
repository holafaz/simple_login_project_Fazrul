# 🚀 Simple Full-Stac Project

### Simple Login & Register System

Hi! This is my submission for the IT Candidate Challenge for job recrutment process.
This project is a full-stack web application that allows users to **register and log in securely**.

---

## 📦 Tech Stack

* **Frontend:** Vue.js
* **Backend:** Golang (Go)
* **Database:** MySQL (Dockerized)

---

## ✨ Features

* 🔐 **Secure Authentication**

  * Passwords are hashed using **bcrypt** (no plain text storage)

* 🐳 **Dockerized Database**

  * MySQL runs inside Docker — no manual installation needed

* 👋 **User Experience**

  * Displays a personalized welcome message after login

---

## 🛠️ How to Run the Project

### 1️⃣ Start the Database

Make sure **Docker Desktop** is running, then execute:

```bash
docker-compose up -d
```

This will automatically start the MySQL database.

---

### 2️⃣ Run the Backend

Navigate to the backend folder:

```bash
cd backend
go run main.go
```

This starts the API server that handles authentication logic.

---

### 3️⃣ Run the Frontend

Navigate to the frontend folder:

```bash
cd frontend
npm install
npm run dev
```

Open the link shown in your terminal (usually):
👉 http://localhost:5173

---

## 📁 Project Structure

```
project-root/
│── backend/        # Go backend (API)
│── frontend/       # Vue.js frontend
│── docker-compose.yml
```

---

## 👤 Author

**Name:** Fazrul Azim   

**Role:** IT Candidate   

**Status:** Graduating February 2026

---

## 💡 Notes

This project demonstrates:

* Full-stack development skills
* Basic authentication system design
* Docker usage for environment setup
* Clean and simple UI integration with backend

---

## 📬 Contact

If you have any questions or feedback, feel free to reach out!
📧 [fazrulazim37@gmail.com](mailto:fazrulazim37@gmail.com)

---
