import 'bulma/css/bulma.css'

import { createApp } from 'vue'
import Index from './Index.vue'
import GoodGame from './pages/GoodGame.vue'
import Trovo from './pages/Trovo.vue'
import Twitch from './pages/Twitch.vue'
import YouTube from './pages/YouTube.vue'
// import './sass/template.sass'
import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'

const Home = { template: '<div>Home</div>' }

const routes = [
    { path: '/', component: Home },
    {path: '/chats/goodgame', name: 'chats.goodgame', component: GoodGame},
    {path: '/chats/trovo', name: 'chats.trovo', component: Trovo},
    {path: '/chats/twitch', name: 'chats.twitch', component: Twitch},
    {path: '/chats/youtube', name: 'chats.youtube', component: YouTube},
]


const router = createRouter({
    history: createWebHistory(),
    routes,
})


const app = createApp(Index)
// Make sure to _use_ the router instance to make the
// whole app router-aware.
app.use(router)


// createApp(App1).mount('#app')
app.mount('#app')
