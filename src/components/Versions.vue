<template>
  <div class="bg-gray-50 w-full p-4" :class="isOpen ? 'hidden' : 'block'">
    <div>
      <h1 class="text-4xl font-normal">{{ title }}</h1>
    </div>
    <canvas class="my-4 h-12"></canvas>
    <template v-if="responseStatus === 200">
      <div class="overflow-auto">
        <table
          class="table-auto w-full text-left border-collapse divide-y border-t-2"
        >
          <thead>
            <tr>
              <th class="px-2 py-2">Workload</th>
              <th class="px-2 py-2">Platform</th>
              <th class="px-2 py-2">Environment</th>
              <th class="px-2 py-2">Version</th>
              <th class="px-2 py-2">Changelog</th>
              <th class="px-2 py-2">Raw</th>
              <th class="px-2 py-2">Status</th>
              <th class="px-2 py-2" v-html="getTimeZone()"></th>
            </tr>
          </thead>
          <tbody class="divide-y">
            <tr
              class="bg-gray-200 hover:bg-gray-300"
              v-for="(deployment, index) in deployments"
              :key="index"
            >
              <td class="px-2 pt-1 pb-4">
                {{ deployment.workload | toUpperCase }}
              </td>
              <td class="px-2 pt-1 pb-4">{{ deployment.platform }}</td>
              <td class="px-2 pt-1 pb-4">{{ deployment.environment }}</td>
              <td class="px-2 pt-1 pb-4">{{ deployment.version }}</td>
              <td
                class="px-2 pt-1 pb-4"
                v-if="
                  deployment.changelog_url === 'NULL' ||
                  deployment.changelog_url === ''
                "
              >
                N/A
              </td>
              <td class="px-2 pt-1 pb-4" v-else>
                <a
                  class="hover:text-blue-600"
                  :href="deployment.changelog_url"
                  title="Click here to see changelog"
                >
                  Click here to see changelog
                </a>
              </td>
              <td
                class="px-2 pt-1 pb-4"
                v-if="deployment.raw === 'NULL' || deployment.raw === ''"
              >
                N/A
              </td>
              <td class="px-2 pt-1 pb-4" v-else>
                <a
                  class="hover:text-blue-600"
                  :href="
                    '/workload/' +
                    deployment.workload +
                    '/platform/' +
                    deployment.platform +
                    '/environment/' +
                    deployment.environment +
                    '/raw/' +
                    deployment.version
                  "
                  title="Click here for more details"
                >
                  Click here for more details
                </a>
              </td>
              <td class="px-2 pt-1 pb-4">{{ deployment.status }}</td>
              <td
                class="px-2 pt-1 pb-4"
                v-html="convertDate(deployment.date)"
              ></td>
            </tr>
          </tbody>
        </table>
      </div>
      <Pagination v-if="pagination.enabled" :pagination="pagination.data" />
    </template>
    <template v-if="responseStatus === 204">
      <h2 class="text-center text-gray-400 text-xl font-semibold">
        So far no deployment has been registered
      </h2>
    </template>
    <template v-if="responseStatus === 404">
      <h2 class="text-center text-gray-400 text-xl font-semibold">
        This environment does not exist
      </h2>
    </template>
    <template v-if="responseStatus === 500">
      <h2 class="text-center text-red-800 text-xl font-semibold">
        An error occured while retrieving last deployments. Please try later.
      </h2>
    </template>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  name: 'Versions',
  props: {
    isOpen: {
      type: Boolean,
      default: false,
    },
    url: {
      type: Object,
      required: true,
    },
    pagination: {
      type: Object,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    responseStatus: {
      type: [String, Number],
      required: true,
    },
    deployments: {
      type: Array,
      required: true,
    },
  },
  components: {
    Pagination: () => import('@components/Pagination.vue'),
  },
  metaInfo() {
    return {
      title: this.title,
    };
  },
  methods: {
    getTimeZone() {
      return (
        'Date (Timezone ' +
        Intl.DateTimeFormat().resolvedOptions().timeZone +
        ')'
      );
    },
    convertDate(date) {
      return moment(date).format('LLLL');
    },
  },
};
</script>
