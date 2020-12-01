<template>
  <div class="flex h-screen">
    <div class="m-auto">
      <p class="text-gray-600 font-black text-9xl text-center">404</p>
      <p class="text-gray-600 font-black text-9xl">page not found</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
  metaInfo() {
    return {
      title: '404 | page not found',
      meta: [
        {
          name: 'Page not found',
          content: 'Page not found',
        },
      ],
    };
  },
};
</script>
