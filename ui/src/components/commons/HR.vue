<template>
  <hr :class="classes" />
</template>

<script>
export default {
  props: {
    classes: {
      type: String,
      default: 'my-5 border-2 border-gray-300',
    },
  },
}
</script>
