<script>
import {defineComponent, reactive} from 'vue'

export default defineComponent({
  name: 'Tabs',
  emits: ['update:modelValue'],
  props: {
    modelValue: {
      type: Object,
      required: false,
      default: reactive({}),
    },
  },
  computed: {
    value: {
      get() {
        return this.modelValue
      },
      set(value) {
        this.$emit('update:modelValue', value)
      },
    },
  },
})
</script>

<template>
  <TabView>
    <TabPanel v-if="value" v-for="(tab, tabKey) in value.tabs" :key="tab.id" :header="tab.title">
      <div v-for="(field, fieldKey) in value.tabs[tabKey].fields" class="field flex flex-column gap-1 mb-5">
        <label v-if="field.label.length">{{ field.label }}</label>
        <InputText
          v-if="field.type == 'string'"
          v-model="value.tabs[tabKey]['fields'][fieldKey].value"
          :placeholder="value.tabs[tabKey]['fields'][fieldKey].placeholder"
        />
        <InputNumber
          v-else-if="field.type == 'integer'"
          v-model="value.tabs[tabKey]['fields'][fieldKey].value"
          :placeholder="value.tabs[tabKey]['fields'][fieldKey].placeholder"
        />
        <Dropdown
          v-else-if="field.type == 'select'"
          v-model="value.tabs[tabKey]['fields'][fieldKey].value"
          :options="value.tabs[tabKey]['fields'][fieldKey].items"
          optionLabel="title"
          option-value="value"
          placeholder="Выберите значение"
          class="w-full"
        />
        <InputSwitch
          v-if="field.type == 'checkbox'"
          v-model="value.tabs[tabKey]['fields'][fieldKey].value"
        />

        <InputGroup v-if="field.type == 'string_list'" v-for="(v, i) in value.tabs[tabKey]['fields'][fieldKey].value ?? ['']">
          <InputText placeholder="Channel name" v-model="value.tabs[tabKey]['fields'][fieldKey].value[i]" />
          <Button @click="value.tabs[tabKey]['fields'][fieldKey].value.push('')" icon="pi pi-plus" severity="success" />
          <Button v-if="value.tabs[tabKey]['fields'][fieldKey].value?.length > 1" @click="value.tabs[tabKey]['fields'][fieldKey].value.splice(i, 1)" icon="pi pi-times" severity="danger" />
        </InputGroup>

        <small v-if="field.description.length" v-html="field.description"></small>
      </div>
    </TabPanel>
  </TabView>
</template>
