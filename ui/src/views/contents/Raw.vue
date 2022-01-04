<template>
  <div>
    <template v-if="responseStatus === 200">
      <pre>{{ deployments }}</pre>
    </template>
    <template v-if="responseStatus === 204">
      <h2 class="text-center text-gray-400 text-xl font-semibold">
        {{ $t('alert.http.noDeployment') }}
      </h2>
    </template>
    <template v-if="responseStatus === 404">
      <h2 class="text-center text-gray-400 text-xl font-semibold">
        {{ $t('alert.http.pageNotFound') }}
      </h2>
    </template>
    <template v-if="responseStatus === 500">
      <h2 class="text-center text-red-800 text-xl font-semibold">
        {{ $t('alert.http.internalServerError') }}
      </h2>
    </template>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import contents from '@/compositions/contents/contents'

const props = defineProps({
  url: {
    type: Object,
    required: true,
  },
})

const { url } = toRefs(props)
const { deployments, responseStatus } = contents(url)
</script>
