import {createRouter, createWebHistory} from 'vue-router'

import Home from './components/Home.vue'
import Login from './components/login/Login.vue'
import SignUp from './components/login/SignUp.vue';
import User from './components/user/User.vue';
import GoogleCallback from "@/components/login/GoogleCallback.vue";


const routes = [
    {path: '/', name: "Home", component: Home},
    {path: '/login', name: "Login", component: Login},
    {path: '/signup', name: "SignUp", component: SignUp},
    {path: '/user', name: "User Page", component: User, meta: {requiredAuth: true}},

    {path: '/api/v1/oauth2/callback', name: "Google Callback", component: GoogleCallback},
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const isAuthorized = localStorage.getItem('access_token')
    if (to.meta.requiredAuth && !isAuthorized) {
        next('/login')
    } else next()
})

export default router