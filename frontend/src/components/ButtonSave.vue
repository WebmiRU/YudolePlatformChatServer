<script>
import { defineComponent } from 'vue'
import APIService from "../services/APIService.ts";

export default defineComponent({
    name: 'ButtonSave',
    props: {
        modelValue: {
            required: true,
        },
    },
  methods: {
      save() {
        APIService.apiPut('http://127.0.0.1/api' + this.$route.fullPath, this.modelValue)
      }
  }
})
</script>

<template>
  <div class="field gap-1 mt-5 mb-3">
    <Button label="Сохранить" severity="success" @click="save($route.params.id, this.modelValue)"/>
  </div>
</template>
