import './assets/main.css'

import {createApp} from 'vue'
import App from './App.vue'
import Aura from '@primevue/themes/aura';
import PrimeVue from 'primevue/config';
import router from "@/router.js";
import ToastService from 'primevue/toastservice';

const app = createApp(App)
app.use(PrimeVue, {
    theme: {
        preset: Aura,
        options: {
            darkModeSelector: 'light',
        }
    }
});
app.use(router)
app.use(ToastService);
app.mount('#app')
