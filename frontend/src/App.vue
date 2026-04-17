<template>
  <div class="container">
    <h1 class="top-headline">
      {{ loggedIn ? 'Welcome, ' + currentUser : 'Simple Log In Website' }}
    </h1>

    <div class="card">
      <div v-if="!loggedIn">
        <h2>{{ isLogin ? 'Login' : 'Create Account' }}</h2>
        
        <div class="form-group">
          <input v-model="form.email" type="email" placeholder="Email Address" />
          <input v-model="form.password" type="password" placeholder="Password" />
        </div>

        <button @click="handleAction">{{ isLogin ? 'Login' : 'Register' }}</button>

        <p @click="isLogin = !isLogin">
          {{ isLogin ? "No account? Register here" : "Have an account? Login here" }}
        </p>
      </div>

      <div v-else>
        <p style="color: white;">You are currently logged in.</p>
        <button @click="logout" style="background: #dc3545;">Logout</button>
      </div>

      <div v-if="statusMsg" :class="statusType">{{ statusMsg }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import axios from 'axios'

const isLogin = ref(true)
const statusMsg = ref('')
const statusType = ref('')
const form = reactive({ username: '', password: '' })

// NEW: State for tracking login status
const loggedIn = ref(false)
const currentUser = ref('')

const handleAction = async () => {
// Simple Email Validation (Check for @)
  const emailRegex = /\S+@\S+\.\S+/
  
  if (!form.email || !form.password) {
    statusMsg.value = "Please fill all fields"
    statusType.value = "error"
    return
  }

  if (!emailRegex.test(form.email)) {
    statusMsg.value = "Please enter a valid email address (missing @ or .)"
    statusType.value = "error"
    return
  }


  const url = isLogin.value ? 'http://localhost:8080/login' : 'http://localhost:8080/register'
  
  try {
    const res = await axios.post(url, form)
    statusMsg.value = res.data.message
    statusType.value = "success"

    // If login is successful, update the headline
    if (isLogin.value) {
      loggedIn.value = true
      currentUser.value = res.data.user // This matches the "user" key from your Go backend
    }
  } catch (err) {
    statusMsg.value = err.response?.data?.error || "Something went wrong"
    statusType.value = "error"
  }
}

// Optional: Logout function to reset the view
const logout = () => {
  loggedIn.value = false
  currentUser.value = ''
  statusMsg.value = ''
  form.username = ''
  form.password = ''
}
</script>

<style scoped>
/* Keeping your custom Green theme */
.container { 
  display: flex; 
  flex-direction: column; /* Stacks headline above card */
  justify-content: center; 
  align-items: center; 
  height: 100vh; 
  background: #034f26; 
  font-family: sans-serif; 
}

.top-headline {
  color: white;
  margin-bottom: 20px;
  font-size: 2.5rem;
}

.card { 
  background: rgb(5, 97, 6); 
  padding: 2rem; 
  border-radius: 8px; 
  box-shadow: 0 4px 6px rgba(0,0,0,0.1); 
  width: 300px; 
  text-align: center; 
}

.card h2 { color: #DFFF00; }

.form-group { 
  display: flex; 
  flex-direction: column; 
  gap: 10px; 
  margin-bottom: 20px; 
}

input { 
  padding: 10px; 
  border: 1px solid #ddd; 
  border-radius: 4px; 
}

button { 
  width: 100%; 
  padding: 10px; 
  background: #007bff; 
  color: white; 
  border: none; 
  border-radius: 4px; 
  cursor: pointer; 
}

button:hover { background: #0056b3; }

p { 
  margin-top: 15px;
  font-size: 0.9rem; 
  color: #DFFF00; /* Changed to match your h2 for better visibility on green */
  cursor: pointer; 
}

.success { color: #DFFF00; margin-top: 10px; }
.error { color: #ff4d4d; margin-top: 10px; }
</style>
