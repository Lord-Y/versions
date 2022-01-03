<template>
  <div>
    <h1 :class="getClasses">
      {{ title }}
    </h1>
    <canvas class="h-5 border-b-2 border-gray-900 w-full" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  props: {
    title: {
      type: String,
      default: '',
    },
    classes: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      local: {
        classes: 'font-normal text-black text-4xl',
      },
    }
  },
  computed: {
    getClasses(): string {
      let z
      if (this.classes === '') {
        z = this.local.classes
      } else {
        z = this.classes
      }
      return z
    },
  },
})
</script>
