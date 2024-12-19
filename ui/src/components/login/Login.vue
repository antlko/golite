<script setup>
import FloatLabel from "primevue/floatlabel";
import InputText from "primevue/inputtext";
import Button from "primevue/button";
import Toast from "primevue/toast";
import {ref} from "vue";
import {useRouter} from "vue-router";
import axiosInstance from "@/axios.js";
import {useToast} from "primevue/usetoast";

const router = useRouter()
const toast = useToast();

const loginForm = ref({
  login: "",
  password: ""
});

const routeRegister = () => {
  router.push('/signup')
}

const routeUser = () => {
  router.push('/user').then(() => {
    router.go(0)
  })
}

const sendData = async () => {
  try {
    const loginFormValue = loginForm.value
    const response = await axiosInstance.post('/api/v1/auth/signin',
        loginFormValue,
        {
          headers: {"Content-Type": "application/json"}
        }
    )

    localStorage.setItem("access_token", response.data.access_token)
    localStorage.setItem("refresh_token", response.data.refresh_token)
    routeUser()
  } catch (error) {
    toast.add({severity: 'error', summary: 'Error', detail: error.response.data.message, life: 3000});
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

</script>
<template>
  <Toast/>
  <div class="login-form-container">
    <h2>Login Form</h2>
    <div class="login-form">
      <FloatLabel class="input-form">
        <label for="username">Login</label>
        <InputText id="username" v-model="loginForm.login"/>
      </FloatLabel>
      <FloatLabel class="input-form">
        <label for="username">Password</label>
        <InputText id="password" type="password" v-model="loginForm.password"/>
      </FloatLabel>
    </div>
    <div class="buttons-box">
      <div class="standard-login">
        <Button class="input-form-btn" label="Login" @click="sendData"/>
        <Button class="input-form-btn" label="Register" @click="routeRegister" severity="contrast"/>
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