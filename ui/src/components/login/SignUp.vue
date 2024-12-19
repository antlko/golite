<script setup>

import Button from "primevue/button";
import FloatLabel from "primevue/floatlabel";
import InputText from "primevue/inputtext";
import Toast from "primevue/toast";
import {ref} from "vue";
import {useRouter} from "vue-router";
import {useToast} from "primevue/usetoast";
import axiosInstance from "@/axios.js";

const toast = useToast();
const router = useRouter()

const signUpForm = ref({
  login: "",
  email: "",
  password: ""
});

const routeLogin = () => {
  router.push('/login')
}

const sendData = async () => {
  try {
    const loginFormValue = signUpForm.value
    const response = await axiosInstance.post('/api/v1/auth/signup',
        loginFormValue,
        {
          headers: {"Content-Type": "application/json"}
        }
    )
    console.log('Server response:', response.data);
    showInfo('You successfully registered! Try to login.')
    signUpForm.value = {
      login: "",
      email: "",
      password: ""
    }

  } catch (error) {
    console.error('Error sending data:', error);
    showError('Failed to create a user!')
  }
};

const googleLogin = async () => {
  try {
    // Make a request to your backend to get the Google OAuth URL
    const response = await axiosInstance.post('/api/v1/oauth2/google/signin');
    // Redirect the user to the Google OAuth2 authorization page
    window.location.href = response.data.url;

  } catch (error) {
    if (error.response) {
      console.error('Error response:', error.response.data);
    } else if (error.request) {
      console.error('No response received:', error.request);
    } else {
      console.error('Error in setting up request:', error.message);
    }
    showError('Failed to initiate Google sign-in!');
  }
};

const showInfo = (message) => {
  toast.add({severity: 'info', summary: 'Info', detail: message, life: 3000});
};

const showError = (message) => {
  toast.add({severity: 'error', summary: 'Error', detail: message, life: 3000});
};
</script>

<template>
  <div>
    <Toast/>
  </div>
  <div class="login-form-container">
    <h2>Sign Up Form</h2>
    <div class="login-form">
      <FloatLabel class="input-form">
        <label for="username">Login</label>
        <InputText id="username" v-model="signUpForm.login"/>
      </FloatLabel>
      <FloatLabel class="input-form">
        <label for="username">Email</label>
        <InputText id="email" type="email" v-model="signUpForm.email"/>
      </FloatLabel>
      <FloatLabel class="input-form">
        <label for="username">Password</label>
        <InputText id="password" type="password" v-model="signUpForm.password"/>
      </FloatLabel>
    </div>
    <div class="buttons-box">
      <div class="standard-login">
        <Button class="input-form-btn" label="Register" @click="sendData"/>
        <Button class="input-form-btn" label="Login" severity="contrast" @click="routeLogin"/>
      </div>
      <div class="google-login">
        <Button class="input-form-btn" label="Login With Google" severity="info" @click="googleLogin"/>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-form-container {
  min-height: 80vh;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.input-form {
  margin-top: 2em;
}

.input-form-btn {
  margin-left: 1em;
  margin-right: 1em;
}

.buttons-box {
  margin-top: 2em;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.google-login {
  display: flex;
  width: 100%;
  margin-top: 1em;
  justify-content: center;
}
</style>