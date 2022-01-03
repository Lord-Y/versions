<template>
  <div class="mt-4 text-center">
    <h4 :class="getClass()">{{ message }}</h4>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  props: {
    message: {
      type: String,
      default: '',
    },
    classes: {
      type: String,
      default: '',
    },
  },
  methods: {
    getClass(): string {
      let c
      switch (this.classes) {
        case 'red':
          c = 'text-red-500 text-xl font-normal'
          break
        case 'green':
          c = 'text-green-500 text-xl font-normal'
          break
        case 'mute':
          c = 'text-gray-500 text-xl font-normal'
          break
        default:
          c = ''
          break
      }
      return c
    },
  },
})
</script>
