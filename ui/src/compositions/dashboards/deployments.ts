import { reactive, toRefs, Ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useHead } from '@vueuse/head'
import type { StatsLatest, BarData, dataset } from '@/apis/interfaces'
import axiosService from '@/apis/axiosService'
import moment from 'moment'

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export default function (
  url: Ref<{
    [x: string]: any
  }>,
) {
  const state = reactive({
    meta: {
      title: '',
      description: '',
    },
    loading: {
      loading: {
        active: true,
      },
    },
    alert: {
      class: '',
      message: '',
    },
    statsLatest: [] as StatsLatest,
    classes: {
      aLinks: 'hover:text-green-600 hover:font-extrabold',
    },
    barData: {} as BarData,
    options: {
      indexAxis: 'y',
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
        title: {
          display: false,
          text: 'Chart.js Horizontal Bar Chart',
        },
      },
    },
  })

  const { t } = useI18n({
    useScope: 'global',
  })

  state.meta.title =
    t('statistics.statistics') +
    ' - ' +
    t('statistics.latest', {
      field: 10,
    }).toLowerCase()
  state.meta.description = state.meta.title

  useHead({
    title: state.meta.title,
    meta: [
      {
        name: 'description',
        content: state.meta.description,
      },
      {
        property: 'og:title',
        content: state.meta.title,
      },
      {
        property: 'og:description',
        content: state.meta.description,
      },
    ],
  })

  axiosService
    .genericGet(
      {
        'Content-Type': 'application/json',
      },
      url.value.api.default,
      {},
    )
    .then((response: any) => {
      const data: Array<number> = []
      const backgroundColor: Array<string> = []
      const dataset: dataset = {
        label: '',
        data: [],
        backgroundColor: [],
      }
      switch (response.status) {
        case 200:
          state.statsLatest = response.data
          state.statsLatest.forEach((k, index) => {
            if (index == 0) {
              state.barData.labels = []
              state.barData.datasets = []
            }
            const label: string =
              t('deployments.workload').toLowerCase() +
              ' ' +
              k.workload +
              ' - ' +
              t('deployments.environment').toLowerCase() +
              ' ' +
              k.environment +
              ' - ' +
              t('deployments.status').toLowerCase() +
              ' ' +
              k.status +
              ' - ' +
              t('deployments.date').toLowerCase() +
              ' ' +
              moment(k.date).format('LL')
            state.barData.labels.push(label)
            data.push(k.total)
            dataset.label = t('deployments.details')
            backgroundColor.push(generateRandomColor())
          })
          dataset.data = data
          dataset.backgroundColor = backgroundColor
          state.barData.datasets.push(dataset)
          break
        case 204:
          state.alert.class = 'mute'
          state.alert.message = t('alert.http.noDeployment')
          break
        default:
          state.alert.class = 'red'
          state.alert.message = t('alert.http.errorOccured')
          break
      }
      state.loading.loading.active = false
    })
    .catch((error: any) => {
      state.alert.class = 'red'
      state.alert.message = t('alert.http.errorOccured')
      state.loading.loading.active = false
      throw error
    })

  function generateRandomColor(): string {
    const letters = '0123456789ABCDEF'
    let color = '#'
    for (let i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)]
    }
    return color
  }

  return {
    ...toRefs(state),
  }
}
