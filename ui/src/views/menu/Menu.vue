<template>
  <div
    class="bg-gray-800 h-screen w-full sm:items-center sm:w-80 sm:block flex"
    :class="isOpen ? 'block' : 'hidden'"
  >
    <div class="w-full" v-if="workload.length > 0">
      <template v-for="(w, index) in workload">
        <div :key="index" :class="setWorkloadActiveLink('/workload/' + w)">
          <div class="px-4 py-1 text-white" :key="index">
            {{ w | toUpperCase }}
          </div>
          <div :key="index + w">
            <template
              v-for="(platform, indexWS) in filteredPlatformsByWorkload(w)"
            >
              <div
                class="ml-8 hover:text-blue-600"
                :class="
                  setPlatformActiveLink(
                    '/workload/' + w + '/platform/' + platform,
                  )
                "
                :key="indexWS + platform + w"
              >
                <a
                  :key="indexWS + platform"
                  :href="'/workload/' + w + '/platform/' + platform"
                  class="text-white hover:text-blue-600"
                  :class="
                    setPlatformActiveLink(
                      '/workload/' + w + '/platform/' + platform,
                    )
                  "
                  :title="platform"
                >
                  {{ platform }}
                </a>
              </div>
              <div class="ml-14" :key="indexWS + platform">
                <template
                  v-for="(
                    environment, indexEnv
                  ) in filteredEnvironmentsByPlatform(w, platform)"
                >
                  <a
                    :key="indexEnv + platform + environment"
                    class="block mt-1 p-1 text-white hover:text-blue-600"
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
                    :href="
                      '/workload/' +
                      w +
                      '/platform/' +
                      platform +
                      '/environment/' +
                      environment
                    "
                    :title="environment"
                  >
                    {{ environment }}
                  </a>
                </template>
              </div>
            </template>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Menu',
  props: {
    isOpen: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      workloads: [],
      workload: [],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      try {
        this.$axios
          .get(
            this.$config.BASE_URL + '/api/v1/versions/read/distinct/workloads',
          )
          .then((response) => {
            if (response.status === 200) {
              this.workloads = response.data;
              response.data.forEach((obj) => {
                if (!this.workload.includes(obj.workload)) {
                  this.workload.push(obj.workload);
                }
              });
            }
          })
          .catch((error) => {
            console.error(error);
          });
      } catch (error) {
        throw error;
      }
    },
    filteredPlatform(item) {
      return this.workloads.filter((items) => item.includes(items.platform));
    },
    filteredPlatformsByWorkload(item) {
      let platforms = [];
      this.workloads.filter((items) => {
        if (item === items.workload) {
          platforms.includes(items.platform)
            ? ''
            : platforms.push(items.platform);
        }
      });
      return platforms;
    },
    filteredEnvironmentsByPlatform(workload, item) {
      let environment = [];
      this.workloads.filter((items) => {
        if (items.platform == item && items.workload === workload) {
          environment.includes(items.environment)
            ? ''
            : environment.push(items.environment);
        }
      });
      return environment;
    },
    setWorkloadActiveLink(data) {
      if (this.$route.path.includes(data)) {
        return 'bg-gray-500 text-gray-400';
      }
    },
    setPlatformActiveLink(data) {
      if (this.$route.path.includes(data)) {
        return 'text-gray-900 font-bold';
      }
    },
    setActiveLink(data) {
      if (this.$route.path === data) {
        return 'text-gray-900 font-semibold';
      }
    },
  },
};
</script>
