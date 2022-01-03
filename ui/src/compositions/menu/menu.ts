import { reactive, toRefs } from 'vue'
import { useRoute } from 'vue-router'
import { Workloads } from '@/apis/interfaces'
import axiosService from '@/apis/axiosService'

export default function () {
  const state = reactive({
    isOpen: false,
    menu: {
      isOpen: {
        dashboards: false,
        platform: false,
      },
    },
    workloads: [] as Workloads,
    workload: [] as Array<string>,
  })

  const route = useRoute()

  axiosService
    .genericGet(
      {
        'Content-Type': 'application/json',
      },
      '/api/v1/versions/read/distinct/workloads',
      {},
    )
    .then((response: any) => {
      switch (response.status) {
        case 200:
          state.workloads = response.data
          state.workloads.forEach((obj) => {
            if (!state.workload.includes(obj.workload)) {
              state.workload.push(obj.workload)
            }
          })
          break
        default:
          state.workloads = []
          break
      }
    })
    .catch((error: any) => {
      throw error
    })

  function getSelectedMenu(data: string): string {
    let classes = ''
    if (route.meta.menu === data) {
      switch (route.meta.menu) {
        case 'dashboards':
          state.menu.isOpen.dashboards = true
          classes = 'bg-green-700'
          break
        case 'platform':
          state.menu.isOpen.platform = true
          classes = 'bg-green-700'
          break
      }
    }
    return classes
  }

  function getActiveLink(data: string): string | undefined {
    if (route.meta.activeLink === data) {
      return 'border-l-4 border-green-700'
    }
  }

  function filteredPlatform(item: string) {
    return state.workloads.filter((items) => item.includes(items.platform))
  }

  function filteredPlatformsByWorkload(item: string): string[] {
    const platforms: string[] = []
    state.workloads.filter((items) => {
      if (item === items.workload) {
        platforms.includes(items.platform) ? '' : platforms.push(items.platform)
      }
    })
    return platforms
  }

  function filteredEnvironmentsByPlatform(workload: string, platform: string) {
    const environment: string[] = []
    state.workloads.filter((items) => {
      if (items.platform == platform && items.workload === workload) {
        environment.includes(items.environment)
          ? ''
          : environment.push(items.environment)
      }
    })
    return environment
  }

  function setWorkloadActiveLink(data: string) {
    if (route.path.includes(data)) {
      return 'text-green-600'
    } else {
      return 'text-white'
    }
  }

  function setPlatformActiveLink(data: string) {
    if (
      (route.path.includes(data) || data.includes(route.path)) &&
      route.path !== '/'
    ) {
      return 'text-green-600 font-bold'
    } else {
      return 'text-white'
    }
  }

  function setActiveLink(data: string) {
    if (route.path === data) {
      return 'text-green-600 font-semibold'
    } else {
      return 'text-white'
    }
  }

  return {
    getSelectedMenu,
    getActiveLink,
    filteredPlatform,
    filteredPlatformsByWorkload,
    filteredEnvironmentsByPlatform,
    setWorkloadActiveLink,
    setPlatformActiveLink,
    setActiveLink,
    ...toRefs(state),
  }
}
