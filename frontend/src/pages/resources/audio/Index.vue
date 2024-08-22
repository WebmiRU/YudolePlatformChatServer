<script lang="ts" setup>
// @ts-ignore
import store from "../../../store"
</script>

<script lang="ts">
import APIService from "../../../services/APIService"

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

    this.model = await APIService.modulesIndexGet()
  },
  unmounted() {

  },
  methods: {

  }
}
</script>

<template>
  <h1>Ресурсы / Аудио</h1>

  <form method="post" enctype="multipart/form-data" action="/api/resources/audio/upload">
    <input type="file" name="file" accept="audio/*" />
    <br/>
    <br/>
    <button>Загрузить</button>
  </form>

  <br/>
  <br/>

  <DataTable :value="model.payload" tableStyle="min-width: 50rem">
    <Column field="name" header="Name"></Column>

    <Column header="Autostart">
      <template #body="row">
        <InputSwitch v-model="model.payload[row.index].autostart"
                     @change="moduleStateChange(row.index, model.payload[row.index].autostart)"/>
      </template>
    </Column>

    <Column header="State">
      <template #body="row">
        <Badge v-if="row.data.proc_state == 'run'" severity="success">Run</Badge>
        <Badge v-else-if="['stopped', 'failed'].includes(row.data.proc_state)" severity="danger">Stopped</Badge>
      </template>
    </Column>

    <Column header="Start/Stop">
      <template #body="row">
        <Button v-if="row.data.proc_state == 'run'" @click="moduleStop(row.index)" severity="danger">Stop</Button>
        <Button v-if="['stopped', 'failed'].includes(row.data.proc_state)" @click="moduleStart(row.index)"
                severity="success">Start
        </Button>
      </template>
    </Column>

    <Column header="Config">
      <template #body="row">
        <RouterLink :to="{name: 'modules.id', params: {id: row.index}}">
          <Button label="Config" severity="secondary"/>
        </RouterLink>
      </template>
    </Column>
  </DataTable>

</template>

