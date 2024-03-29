import { reactive, toRefs } from 'vue'
import type { Ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useHead } from '@vueuse/head'
import type { Deployments, GenericObject } from '@/apis/interfaces'
import axiosService from '@/apis/axiosService'
import { config } from '@/apis/configs'

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
    isOpen: false,
    deployments: [] as Deployments,
    pagination: {
      enabled: false,
      data: {
        url: '',
        actualPage: 1,
        total: 0,
      },
    },
    classes: {
      aLinks: 'hover:text-green-600 hover:font-extrabold',
    },
    responseStatus: 0 as number,
  })

  const route = useRoute()
  const { t } = useI18n({
    useScope: 'global',
  })
  const dataRoute: GenericObject = route.meta as GenericObject

  let page: number

  if (!route.params.page) {
    page = 1
  } else {
    page = Number(route.params.page)
  }

  let getData: Record<string, unknown> = {}
  switch (dataRoute.root.toString()) {
    case 'platform':
      state.meta.title = t('deployments.platform') + ' ' + route.params.platform
      state.meta.description = state.meta.title
      getData = {
        workload: route.params.workload,
        platform: route.params.platform,
        page: page,
      }
      break
    case 'environment':
      state.meta.title =
        t('deployments.platform') +
        ' ' +
        route.params.platform +
        ' - ' +
        t('deployments.environment').toLowerCase() +
        ' ' +
        route.params.environment
      state.meta.description = state.meta.title
      getData = {
        workload: route.params.workload,
        platform: route.params.platform,
        environment: route.params.environment,
        page: page,
      }
      break
    default:
      state.meta.title = t('deployments.last')
      state.meta.description = t('deployments.last')
      break
  }

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
      getData,
    )
    .then((response: any) => {
      let total: number
      switch (response.status) {
        case 200:
          state.deployments = response.data
          switch (dataRoute.root.toString()) {
            case 'platform':
              total = state.deployments[0].total
              if (total > config.RANGE_LIMIT) {
                state.pagination.enabled = true
                  ; (state.pagination.data.url = `/workload/${route.params.workload}/platform/${route.params.platform}`),
                    (state.pagination.data.actualPage = page)
                state.pagination.data.total = total
                state.pagination.enabled = true
              }
              break
            case 'environment':
              total = state.deployments[0].total
              if (total > config.RANGE_LIMIT) {
                state.pagination.enabled = true
                  ; (state.pagination.data.url = `/workload/${route.params.workload}/platform/${route.params.platform}/environment/${route.params.environment}/`),
                    (state.pagination.data.actualPage = page)
                state.pagination.data.total = total
                state.pagination.enabled = true
              }
              break
          }
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
      state.responseStatus = response.status
      state.loading.loading.active = false
    })
    .catch((error: any) => {
      state.alert.class = 'red'
      state.alert.message = t('alert.http.errorOccured')
      if (error.response && error.response.data) {
        switch (error.response.status) {
          case 404:
          case 500:
            state.responseStatus = error.response.status
            break
          default:
            state.responseStatus = error.response.status
            break
        }
      } else {
        state.responseStatus = error.response.status
      }
      state.loading.loading.active = false
      throw error
    })

  return {
    ...toRefs(state),
  }
}
