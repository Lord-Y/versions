<template>
  <div class="w-full lg:w-96 bg-gray-900">
    <header class="block md:px-4 md:py-1">
      <div
        class="flex justify-between lg:justify-start items-center px-4 py-1 md:p-0"
      >
        <div>
          <img
            width="32"
            height="32"
            class="h-8 w-8 items-center"
            src="/logo.png"
            alt="Logo"
          />
        </div>
        <div class="px-2 pt-2 pb-4">
          <router-link
            class="block text-green-600 font-extrabold hover:text-white px-2 py-1 uppercase"
            to="/"
            :title="$t('brand.name')"
            >{{ $t('brand.name') }}</router-link
          >
        </div>
        <div class="lg:hidden">
          <button
            type="button"
            @click="isOpen = !isOpen"
            class="block text-green-600 focus:text-white focus:outline-none"
          >
            <svg class="h-6 w-6 fill-current" viewBox="0 0 24 24">
              <path
                v-if="isOpen"
                fill-rule="evenodd"
                d="M18.278 16.864a1 1 0 0 1-1.414 1.414l-4.829-4.828-4.828 4.828a1 1 0 0 1-1.414-1.414l4.828-4.829-4.828-4.828a1 1 0 0 1 1.414-1.414l4.829 4.828 4.828-4.828a1 1 0 1 1 1.414 1.414l-4.828 4.829 4.828 4.828z"
              />
              <path
                v-if="!isOpen"
                fill-rule="evenodd"
                d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
              />
            </svg>
          </button>
        </div>
      </div>
      <div :class="isOpen ? 'block' : 'hidden'" class="pt-2 pb-4 lg:flex">
        <nav class="w-full">
          <div>
            <button
              type="button"
              class="block text-white w-full text-left p-3 border-t-2 border-b-2 border-black"
              :class="getSelectedMenu('dashboards')"
              @click="menu.isOpen.dashboards = !menu.isOpen.dashboards"
            >
              {{ $t('dashboards.dashboards') }}
            </button>
            <div class="bg-black" v-if="menu.isOpen.dashboards">
              <router-link
                class="block text-white hover:text-green-600 hover:font-extrabold p-2"
                :class="getActiveLink('/dashboards/deployments')"
                to="/dashboards/deployments"
                :title="$t('dashboards.deployments')"
                >{{ $t('dashboards.deployments') }}</router-link
              >
            </div>
          </div>
          <div v-if="workload.length > 0">
            <button
              type="button"
              class="block text-white w-full text-left p-3 border-b-2 border-black"
              :class="getSelectedMenu('platform')"
              @click="menu.isOpen.platform = !menu.isOpen.platform"
            >
              {{ $t('deployments.workload', 2) }}
            </button>
            <template v-for="(w, index) in workload" :key="index">
              <div class="bg-black" v-if="menu.isOpen.platform">
                <div
                  class="block font-extrabold p-2 uppercase"
                  :class="setWorkloadActiveLink('/workload/' + w)"
                >
                  {{ w }}
                </div>
                <template
                  v-for="(platform, indexWS) in filteredPlatformsByWorkload(w)"
                  :key="indexWS + platform + w"
                >
                  <div class="bg-black ml-4">
                    <router-link
                      class="block hover:text-green-600 hover:font-extrabold p-2"
                      :class="
                        setPlatformActiveLink(
                          '/workload/' + w + '/platform/' + platform,
                        )
                      "
                      :to="'/workload/' + w + '/platform/' + platform"
                      :title="platform"
                      >{{ platform }}</router-link
                    >
                  </div>
                  <div
                    class="bg-black ml-8"
                    v-for="(
                      environment, indexEnv
                    ) in filteredEnvironmentsByPlatform(w, platform)"
                    :key="indexEnv + platform + environment"
                  >
                    <router-link
                      class="block hover:text-green-600 hover:font-extrabold p-2"
                      :class="
                        setActiveLink(
                          '/workload/' +
                            w +
                            '/platform/' +
                            platform +
                            '/environment/' +
                            environment,
                        )
                      "
                      :to="
                        '/workload/' +
                        w +
                        '/platform/' +
                        platform +
                        '/environment/' +
                        environment
                      "
                      :title="environment"
                      >{{ environment }}</router-link
                    >
                  </div>
                </template>
              </div>
            </template>
          </div>
        </nav>
      </div>
    </header>
  </div>
</template>

<script setup lang="ts">
import menus from '@/compositions/menu/menu'

const {
  isOpen,
  menu,
  workload,
  getSelectedMenu,
  getActiveLink,
  setWorkloadActiveLink,
  setPlatformActiveLink,
  setActiveLink,
  filteredPlatformsByWorkload,
  filteredEnvironmentsByPlatform,
} = menus()
</script>
