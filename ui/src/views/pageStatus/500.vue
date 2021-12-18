<template>
  <div class="flex h-screen">
    <div class="m-auto">
      <p class="text-red-600 font-black text-9xl text-center">500</p>
      <p class="text-red-600 font-black text-9xl">Internal server error</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InternalServerError',
  metaInfo() {
    return {
      title: '500 | Internal server error',
      meta: [
        {
          name: 'Internal server error',
          content: 'Internal server error',
        },
      ],
    };
  },
};
</script>
