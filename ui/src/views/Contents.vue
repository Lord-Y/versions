<template>
  <div class="lg:flex min-h-screen w-full bg-gray-50">
    <Menu />
    <div class="w-full p-4">
      <div class="block">
        <Title :title="meta.title" />
        <SpinnerCommon v-if="loading.loading.active" />
        <div class="mx-auto px-3 mt-20 w-full" v-if="!loading.loading.active">
          <Versions
            :url="url"
            :pagination="pagination"
            :response-status="responseStatus"
            :deployments="deployments"
            :is-open="isOpen"
            :alert="alert"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import Menu from '@/views/menu/Menu.vue'
import Title from '@/components/commons/Title.vue'
import SpinnerCommon from '@/components/commons/SpinnerCommon.vue'
import contents from '@/compositions/contents/contents'
import Versions from '@/components/commons/Versions.vue'

const props = defineProps({
  url: {
    type: Object,
    required: true,
  },
})

const { url } = toRefs(props)
const {
  meta,
  loading,
  alert,
  isOpen,
  deployments,
  pagination,
  responseStatus,
} = contents(url)
</script>
