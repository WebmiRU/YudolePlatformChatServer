<script lang="ts" setup>
//@ts-ignore
import store from "../../store"
</script>

<script lang="ts">
import APIService from "../../services/APIService"

export default {
  computed: {},
  components: {},
  data() {
    return {
      updateInterval: null,
      model: {payload: {}},
    }
  },
  async mounted() {
    store.breadcrumbs = [
      {icon: 'pi pi-home', route: {name: 'index'}},
      {label: 'Каналы', route: {name: 'channels.index'}}
    ]

    this.model = await APIService.channelsIndex()
  },
  unmounted() {},
  methods: {},
}
</script>

<template>
  <h1>Доступные каналы</h1>

  <br/>
  <br/>

  <DataTable :value="model.payload" tableStyle="min-width: 50rem">
    <Column field="title" header="Название"></Column>

    <Column header="Настройки">
      <template #body="row">
        <RouterLink :to="{name: 'channels.page', params: {id: row.index}}">
          <Button label="Настройка" severity="secondary"/>
        </RouterLink>
      </template>
    </Column>
  </DataTable>

</template>

