<template>
  <h1 :class="classes">
    {{ title }}
  </h1>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: '',
    },
    classes: {
      type: String,
      default: 'text-center font-semibold text-green-600 text-3xl mb-10',
    },
  },
}
</script>
