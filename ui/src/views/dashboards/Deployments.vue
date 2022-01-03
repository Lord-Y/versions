<template>
  <div class="lg:flex min-h-screen w-full bg-gray-50">
    <Menu />
    <div class="w-full p-4">
      <div class="block">
        <TitleVue :title="meta.title" />
        <SpinnerCommon v-if="loading.loading.active" />
        <AlertMessage :message="alert.message" :classes="alert.class" />
        <div class="mx-auto px-3 mt-20 w-full" v-if="barData">
          <BarChart :chart-data="barData" :options="options" />
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
// https://www.chartjs.org/docs/master/getting-started/integration.html#bundlers-webpack-rollup-etc
import { BarChart } from 'vue-chart-3'
import {
  Chart,
  BarController,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
} from 'chart.js'
Chart.register(
  BarController,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
)

const props = defineProps({
  url: {
    type: Object,
    required: true,
  },
})

const { url } = toRefs(props)
const { meta, loading, alert, barData, options } = dashboardsDeployments(url)
</script>
