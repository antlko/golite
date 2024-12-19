<script setup>
import {onMounted, ref} from "vue";

const isAuth = ref(false)

onMounted(async () => {
  isAuth.value = !!localStorage.getItem('access_token')
});
</script>

<template>
  <div class="header-box">
    <div class="header-links">
      <RouterLink to="/">
        <div class="header-link-elem">Home</div>
      </RouterLink>

      <RouterLink v-if="!isAuth" to="/login">
        <div class="header-link-elem">Login</div>
      </RouterLink>
      <RouterLink v-if="!isAuth" to="/signup">
        <div class="header-link-elem">SignUp</div>
      </RouterLink>
      <RouterLink v-if="isAuth" to="/user">
        <div class="header-link-elem">Profile</div>
      </RouterLink>
    </div>
  </div>
</template>

<style scoped>
a {
  color: inherit; /* blue colors for links too */
  text-decoration: inherit; /* no underline */
}

.header-box {
  display: flex;
  align-items: center;
}

.header-links {
  display: flex;
  margin-left: auto;
}

.header-link-elem {
  padding: 1em;
  border: solid transparent;
}

.header-link-elem:hover {
  border: solid;
  cursor: pointer;
}
</style>