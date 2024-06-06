<script>
import { defineComponent, reactive } from 'vue'

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
                    placeholder="Select a City"
                    class="w-full md:w-14rem"
                />


                <small v-if="field.description.length">{{ field.description }}</small>
            </div>
        </TabPanel>
    </TabView>
</template>
