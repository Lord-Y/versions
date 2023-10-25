import { reactive, toRefs } from 'vue'
import type { Ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useHead } from '@vueuse/head'
import type { StatsLatest, Options } from '@/apis/interfaces'
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
    dashboards: {
      data: [] as (string | number | object)[][],
      options: {} as Options,
      color: '#002877'
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
      switch (response.status) {
        case 200:
          state.statsLatest = response.data
          state.dashboards.data.push([t('deployments.workload'), t('deployments.details'), { role: 'style' }])
          state.statsLatest.forEach((k) => {
            const datan = [] as (string | number | object)[]
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
            datan.push(label)
            datan.push(k.total)
            datan.push(generateRandomColor())
            state.dashboards.data.push(datan)
          })
          state.dashboards.options.height = 700
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
