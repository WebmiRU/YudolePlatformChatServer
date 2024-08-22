<script lang="ts" setup>
// @ts-ignore
import store from "../../../store"
</script>

<script lang="ts">
import APIService from "../../../services/APIService.ts"

export default {
  inject: ['sse'],
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
      {label: 'Ресурсы'},
      {label: 'Аудио', route: {name: 'resources.audio'}},
    ]

    console.log(APIService)
    this.model = await APIService.resourcesAudioGet()
    console.log('RAA')
  },
  unmounted() {

  },
  methods: {}
}
</script>

<template>
  <h1>Ресурсы / Аудио</h1>

  <form method="post" enctype="multipart/form-data" action="/api/resources/audio/upload">
    <input type="file" name="file" accept="audio/*"/>
    <br/>
    <br/>
    <button>Загрузить</button>
  </form>

  <br/>
  <br/>

  <DataTable :value="model.payload" tableStyle="min-width: 50rem">
    <Column field="name" header="Имя"/>
    <Column field="source" header="Источник"/>

    <Column header="Размер">
      <template #body="row">
        {{ row.data.size }}
      </template>
    </Column>

    <Column header="Проигрыватель">
      <template #body="row">
        <audio controls>
          <source :src="'/api/resources/audio/' + row.data.sha256" type="audio/mpeg"/>
        </audio>
      </template>
    </Column>
  </DataTable>

</template>

