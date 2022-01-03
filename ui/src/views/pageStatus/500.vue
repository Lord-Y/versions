<template>
  <div class="flex h-screen">
    <div class="m-auto text-center">
      <p class="text-red-600 font-black text-9xl">500</p>
      <p class="text-red-600 font-black text-9xl">
        {{ $t('alert.http.internalServerError') }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import internalError from '@/compositions/pageStatus/internalError'

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const { meta } = internalError()
</script>
