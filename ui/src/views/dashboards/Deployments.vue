<template>
  <div class="lg:flex min-h-screen w-full bg-gray-50">
    <Menu />
    <div class="w-full p-4">
      <div class="block">
        <TitleVue :title="meta.title" />
        <SpinnerCommon v-if="loading.loading.active" />
        <AlertMessage :message="alert.message" :classes="alert.class" />
        <div class="mx-auto px-3 w-full" v-if="dashboards">
          <GChart type="BarChart" :data="dashboards.data
            " :options="dashboards.options
    " />
          <!-- <pre>dashboards {{ dashboards }}</pre> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { toRefs } from 'vue'
import Menu from '@/views/menu/Menu.vue'
import TitleVue from '@/components/commons/Title.vue'
import SpinnerCommon from '@/components/commons/SpinnerCommon.vue'
import dashboardsDeployments from '@/compositions/dashboards/deployments'
import AlertMessage from '@/components/commons/AlertMessage.vue'
import { GChart } from 'vue-google-charts'

const props = defineProps({
  url: {
    type: Object,
    required: true,
  },
})

const { url } = toRefs(props)
const { meta, loading, alert, dashboards } = dashboardsDeployments(url)
</script>
