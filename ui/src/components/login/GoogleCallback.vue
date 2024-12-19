<script setup>

import {onMounted} from 'vue'
import {useRouter} from "vue-router";

const router = useRouter()

onMounted(() => {
  let uri = window.location.search.substring(1);
  let params = new URLSearchParams(uri);

  localStorage.setItem("refresh_token", params.get("refresh"))
  localStorage.setItem("access_token", params.get("access"))

  window.location.href = '/user'
})

</script>

<template>
  <h1>Google login...</h1>
</template>

<style scoped>

</style>