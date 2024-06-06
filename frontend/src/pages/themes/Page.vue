<script setup lang="ts">
//@ts-ignore
import store from "../../store.ts";
import Tabs from "../../components/Tabs.vue";
</script>

<script lang="ts">
import APIService from "../../services/APIService"

export default {
  components: {},
  data() {
    return {
      model: null,
    }
  },
  async mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'themes.index'}},
      {label: 'Темы', route: {name: 'themes.index'}},
      {label: this.$route.params.id, route: {name: 'themes.id', params: {id: this.$route.params.id}}},
    ]

    this.model = await APIService.themesGet(this.$route.params.id)
  },
  methods: {
    save(id: string, payload: object) {
      APIService.themesPut(id, payload)
    }
  }
}
</script>

<template>
  <h1>Настройки темы [{{ $route.params.id }}]</h1>

  <Tabs v-if="model?.payload" v-model="model.payload"/>
    <div class="field gap-1 mt-5 mb-3">
        <Button label="Save" severity="success" @click="save($route.params.id, this.model)"/>
    </div>
</template>

