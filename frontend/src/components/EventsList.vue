<script>
import {defineComponent} from 'vue'
import APIService from "../services/APIService.ts";

export default defineComponent({
  name: 'EventsList',
  emits: ['update:modelValue'],
  data() {
    return {
      events: null,
      _model: {},
    }
  },
  props: {
      modelValue: {
          required: true,
      },
  },
  methods: {
    modelUpdate() {
      let value = [];

      Object.keys(this._model).forEach(k => {
        if (this._model[k]) value.push(k)
      })

      this.$emit('update:modelValue', value)
      console.log(value)
    }

  },
  async mounted() {
    this.events = await APIService.apiGet('http://127.0.0.1/api/events')
    this.events.payload.forEach(event => {
      this._model[event] = this.modelValue.includes(event)
    })
  }
})
</script>

<template>
  <h3>Events list</h3>

  <DataTable v-if="events" :value="events.payload" stripedRows showGridlines>
    <Column header="Событие">
      <template #body="v">
        {{v.data}}
      </template>
    </Column>
    <Column header="Описание">
      <template #body="v">
        <small><i>Пока не реализовано</i></small>
      </template>
    </Column>
    <Column header="Подписка">
      <template #body="v">
        <InputSwitch v-model="_model[v.data]" @change="modelUpdate" />
      </template>
    </Column>
  </DataTable>
</template>
