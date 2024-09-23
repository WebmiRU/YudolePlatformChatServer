<script setup lang="ts">
//@ts-ignore
import store from "../../store.ts"
import EventsList from "../../components/EventsList.vue";
</script>

<script lang="ts">
import APIService from "../../services/APIService"

export default {
  components: {},
  data() {
    return {
      model: null,
      themes: null,
      resources: {
        audio: {},
        images: {},
      },
      modules: null,
      // themesList: [],
    }
  },
  async mounted() {
    console.log('MDLS START')
    console.log('MDLS', await APIService.apiGet('http://127.0.0.1/api/modules'))


    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'channels.index'}},
      {label: 'Каналы', route: {name: 'channels.index'}},
      {label: this.$route.params.id, route: {name: 'channels.page', params: {id: this.$route.params.id}}},
    ]

    this.model = await APIService.channelsGet(this.$route.params.id)
    this.themes = await APIService.apiGet('http://127.0.0.1/api/themes')
    // this.themesList = Object.keys(this.themes.payload)
    this.resources.audio = await APIService.apiGet('http://127.0.0.1/api/resources/audio')
    this.resources.images = await APIService.apiGet('http://127.0.0.1/api/resources/images')
    this.modules = await APIService.apiGet('http://127.0.0.1/api/modules')


    Object.keys(this.themes.payload).forEach(v => {
      if (!this.model.payload.theme.config[v]) {
        this.model.payload.theme.config[v] = this.themes.payload[v]
      }
    })
  },
  methods: {},
  computed: {
    themesList() {
      if (!this.themes) return null

      const result = []

      Object.keys(this.themes.payload).forEach(k => {
        result.push({
          key: k,
          title: this.themes.payload[k].title
        })
      })

      return result
    },
    /**
     * Формирует список доступных в данный момент модулей-клиентов,
     * для возможности выбрать иконку из списка доступных,
     * для каждого из возможных клиентов5
     */
    clientsList() {
      if (!this.modules || !this.model) return null

      const services = []

      Object.keys(this.modules.payload).forEach(key => {
        if (this.modules.payload[key].type == 'client') {
          services.push(this.modules.payload[key].service)
        }
      })

      return Object.keys(this.model.payload.service_icons).filter(v => services.includes(v))
    },
    serviceIcons() {
      if (!this.resources?.service_icon || !this.model) return null
    }
  }
}
</script>

<template>
  <h1>Настройки канала [{{ $route.params.id }}]</h1>

<!--  {{resources.images.payload}}-->

  <TabView>
    <TabPanel header="Настройки канала">
      <div class="field flex flex-column gap-1 mb-5 mt-3">
        <label>Выберите тему</label>
        <Dropdown
          v-if="model?.payload"
          v-model="model.payload.theme.name"
          :options="themesList"
          option-label="title"
          option-value="key"
          placeholder="Выберите тему"
          class="w-full"
        />
      </div>
      <div class="field flex flex-column gap-1 mb-5 mt-3">
        <label>Иконки сервисов</label>
        <table border="1" style="border-collapse: collapse">
          <thead>
          <tr>
            <th>Сервис</th>
            <th>Иконка</th>
          </tr>
          </thead>
          <tr v-for="v in clientsList">
            <td>{{ v }}</td>
            <td class="w-full">
              <Dropdown
                v-if="resources.images.payload"
                v-model="model.payload.service_icons[v]"
                :options="resources.images.payload"
                option-label=""
                option-value=""
                placeholder="Выберите тему"
                class="w-full"
              >
                <template #value="v">
                  <img :src="'/api/resources/images/'+v.value.sha256" alt="#" style="max-width: 100px; height: 16px" />
                </template>
                <template #option="v">
<!--                  {{v.option}}-->
                  <img :src="'/api/resources/images/'+v.option.sha256" alt="#" style="max-width: 200px; max-height: 48px" />
                </template>
              </Dropdown>
            </td>
          </tr>
        </table>
      </div>
    </TabPanel>

    <TabPanel
      header="Настройки темы"
      :disabled="!(Object.keys(themes?.payload ?? {}).find(v => v == model?.payload.theme.name))"
    >
      <Tabs v-if="model?.payload" v-model="model.payload.theme.config[model.payload.theme.name]"/>
    </TabPanel>

    <TabPanel header="События">
      <EventsList v-if="model" v-model="model.payload.events" />
    </TabPanel>
  </TabView>

  <ButtonSave v-if="model" v-model="model"/>
</template>

