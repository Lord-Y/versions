<template>
  <div class="flex h-screen">
    <div class="m-auto text-center">
      <p class="text-gray-600 font-black text-9xl">404</p>
      <p class="text-gray-600 font-black text-9xl">
        {{ $t('alert.http.pageNotFound') }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import notFound from '@/compositions/pageStatus/notFound'

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const { meta } = notFound()
</script>
