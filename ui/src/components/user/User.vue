<script setup>

import {onMounted, ref} from "vue";
import axiosInstance from "@/axios.js";
import Button from "primevue/button";
import {useRouter} from "vue-router";

const userData = ref(null)

const router = useRouter()

onMounted(async () => {
  getPrivateData()
});
const getPrivateData = async () => {
  const accessToken = localStorage.getItem('access_token')
  try {
    const response = await axiosInstance.get('/api/v1/protected/user',
        {
          headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + accessToken
          }
        }
    )
    userData.value = response.data
    console.log(response.data)
  } catch (error) {
    console.error('Error sending data:', error);
  }
};

const logout = () => {
  localStorage.removeItem('access_token')
  router.push('/').then(() => {
    router.go(0)
  })
}

</script>

<template>
  <div>
    <div v-if="userData" class="content">
      <h2>{{ userData.id }}</h2>
      <p>{{ userData.login }}</p>
      <p>{{ userData.email }}</p>
    </div>
    <Button class="input-form-btn" label="Logout" @click="logout"/>
  </div>
</template>

<style scoped>

</style>