<template>
  <div
    class="mt-4 border-4 border-solid border-green-500 border-r-white w-8 h-8 rounded-full animate-spin-fast mx-auto"
  />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  methods: {
    scrollToTop(): void {
      if (typeof window !== 'undefined') {
        window.scrollTo({
          top: 0,
        })
      }
    },
  },
})
</script>
