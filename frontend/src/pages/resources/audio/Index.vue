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
  methods: {
    uploadClick() {
      this.$refs.file.click()
    },

    async fileChange() {
      const fd = new FormData()
      fd.append('file', this.$refs.file.files[0])

      let response = await fetch('/api/resources/audio', {
        method: 'POST',
        body: fd
      });

      this.$refs.file.value = null

      this.model = await response.json();

    }
  }
}
</script>

<template>
  <input @change="fileChange" ref="file" type="file" accept="audio/*" style="display: none"/>

  <button @click="uploadClick">Загрузить</button>
  <h1>Ресурсы / Аудио</h1>

  <br/>
  <br/>

  <DataTable :value="model.payload">
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
          <source :src="'/api/resources/audio/' + row.data.sha256" :type="row.data.mime_type"/>
        </audio>
      </template>
    </Column>
  </DataTable>
</template>

