🚀 Nusa IT Candidate Project

Simple Login & Register System

Hi! This is my submission for the IT Candidate Challenge. I built a full-stack application that allows a user to Register and Login.

📦 What is inside?

Website (Frontend): Built with Vue.js.

Server (Backend): Built with Golang.

Database: MySQL running inside Docker.

🛠️ How to run this on your computer

1. The Database
   
Make sure Docker Desktop is open. In your terminal, type:

Bash
docker-compose up -d

This starts the MySQL database automatically.

2. The Backend (The Server)
   
Go into the backend folder and run:

Bash
go run main.go

This starts the API that handles the login logic.

3. The Frontend (The Website)
   
Go into the frontend folder and run:

Bash
npm install
npm run dev

Click the link shown in your terminal (usually http://localhost:5173) to see the website.

✨ Features
Secure Passwords: I use bcrypt so passwords are encrypted (not plain text).

Docker: You don't need to install MySQL manually; Docker handles it.

Welcome Message: Once you login, the website will greet you by name!

👤 My Details
Name: Fazrul

Role: IT Candidate

Status: Graduating February 2026
