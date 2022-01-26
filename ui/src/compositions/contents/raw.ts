import { reactive, toRefs, Ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useHead } from '@vueuse/head'
import { DeploymentRaw } from '@/apis/interfaces'
import axiosService from '@/apis/axiosService'

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
    deployments: {} as DeploymentRaw,
    classes: {
      aLinks: 'hover:text-green-600 hover:font-extrabold',
    },
    responseStatus: 0 as number,
  })

  const route = useRoute()
  const { t } = useI18n({
    useScope: 'global',
  })

  state.meta.title =
    t('deployments.platform') +
    ' ' +
    route.params.platform +
    ' - ' +
    t('deployments.environment').toLowerCase() +
    ' ' +
    route.params.environment +
    ' - ' +
    t('deployments.version').toLowerCase() +
    ' ' +
    route.params.version
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
      {
        workload: route.params.workload,
        platform: route.params.platform,
        environment: route.params.environment,
        version: route.params.version,
      },
    )
    .then((response: any) => {
      switch (response.status) {
        case 200:
          state.deployments = response.data
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

  function isJSON(str: any): any {
    const zz: string = str.replace(/\\"/g, '"')
    try {
      return JSON.stringify(JSON.parse(zz), null, 2)
    } catch (e) {
      return str
    }
  }

  return {
    isJSON,
    ...toRefs(state),
  }
}
