import { createApp } from 'vue'
import {createRouter, createWebHistory} from 'vue-router'
import PrimeVue from 'primevue/config';
import App from './App.vue'

// PrimeVue components import
import ButtonSave from "./components/ButtonSave.vue"
import Tabs from "./components/Tabs.vue"
import Button from "primevue/button"
import Badge from "primevue/badge"
import Breadcrumb from "primevue/breadcrumb"
import Menubar from "primevue/menubar"
import Column from 'primevue/column'
import DataTable from 'primevue/datatable'
import TabView from "primevue/tabview"
import TabPanel from "primevue/tabpanel"
import InputSwitch from 'primevue/inputswitch'
import InputText from "primevue/inputtext"
import InputNumber from "primevue/inputnumber"
import Dropdown from 'primevue/dropdown'
import InputGroup from 'primevue/inputgroup'



// import './style.css'
// import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/themes/aura-dark-green/theme.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css'
// import 'primeflex/themes/primeone-light.css'
// import 'primeflex/themes/primeone-dark.css'
import './sass/style.sass'


import Index from './Index.vue'
import Modules from './pages/modules/Index.vue'
import ModulesId from './pages/modules/Page.vue'
import ThemeIndex from './pages/themes/Index.vue'
import ThemePage from './pages/themes/Page.vue'
import ChannelIndex from './pages/channels/Index.vue'
import ChannelPage from './pages/channels/Page.vue'

const routes = [
    { path: '/', name: 'index', component: Index },
    { path: '/modules', name: 'modules.index', component: Modules },
    { path: '/modules/:id', name: 'modules.id', component: ModulesId },
    { path: '/themes', name: 'themes.index', component: ThemeIndex },
    { path: '/themes/:id', name: 'themes.page', component: ThemePage },
    { path: '/channels', name: 'channels.index', component: ChannelIndex },
    { path: '/channels/:id', name: 'channels.page', component: ChannelPage },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App)


app.use(router)
app.use(PrimeVue)


app.component('Tabs', Tabs)
app.component('ButtonSave', ButtonSave)
app.component('Badge', Badge)
app.component('Button', Button)
app.component('Breadcrumb', Breadcrumb)
app.component('Column', Column)
app.component('DataTable', DataTable)
app.component('TabView', TabView)
app.component('TabPanel', TabPanel)
app.component('Menubar', Menubar)
app.component('InputSwitch', InputSwitch)
app.component('InputText', InputText)
app.component('InputNumber', InputNumber)
app.component('Dropdown', Dropdown)
app.component('InputGroup', InputGroup)


app.mount('#app')
