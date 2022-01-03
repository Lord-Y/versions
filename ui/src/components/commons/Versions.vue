<template>
  <div>
    <template v-if="responseStatus === 200">
      <div class="overflow-auto">
        <table
          class="table-auto w-full text-left border-collapse divide-y border-t-2"
        >
          <thead>
            <tr>
              <th class="px-2 py-2">{{ $t('deployments.workload') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.platform') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.environment') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.version') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.changelog') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.raw') }}</th>
              <th class="px-2 py-2">{{ $t('deployments.status') }}</th>
              <th class="px-2 py-2">{{ getTimeZone() }}</th>
            </tr>
          </thead>
          <tbody class="divide-y">
            <tr
              class="bg-gray-200 hover:bg-gray-300"
              v-for="(deployment, index) in deployments"
              :key="index"
            >
              <td class="px-2 pt-1 pb-4 uppercase">
                {{ deployment.workload }}
              </td>
              <td class="px-2 pt-1 pb-4">{{ deployment.platform }}</td>
              <td class="px-2 pt-1 pb-4">{{ deployment.environment }}</td>
              <td class="px-2 pt-1 pb-4">{{ deployment.version }}</td>
              <td class="px-2 pt-1 pb-4" v-if="deployment.changelog_url === ''">
                {{ $t('deployments.notApplicable') }}
              </td>
              <td class="px-2 pt-1 pb-4" v-else>
                <a
                  class="hover:text-blue-600"
                  :href="deployment.changelog_url"
                  :title="$t('click.changelog')"
                >
                  {{ $t('click.changelog') }}
                </a>
              </td>
              <td class="px-2 pt-1 pb-4" v-if="deployment.raw === ''">
                {{ $t('deployments.notApplicable') }}
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
                    deployment.version +
                    '/'
                  "
                  :title="$t('click.moreDetails')"
                >
                  {{ $t('click.moreDetails') }}
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
      <h2 class="text-center text-gray-500 text-xl font-semibold">
        {{ alert.message }}
      </h2>
    </template>
    <template v-if="responseStatus === 404">
      <h2 class="text-center text-gray-500 text-xl font-semibold">
        {{ alert.message }}
      </h2>
    </template>
    <template v-if="responseStatus >= 500">
      <h2 class="text-center text-red-500 text-xl font-semibold">
        {{ alert.message }}
      </h2>
    </template>
  </div>
</template>

<script setup lang="ts">
import { PropType } from 'vue'
import moment from 'moment'
import Pagination from './Pagination.vue'
import { Deployments } from '@/apis/interfaces'

defineProps({
  url: {
    type: Object,
    required: true,
  },
  isOpen: {
    type: Boolean,
    default: false,
  },
  deployments: {
    type: Array as PropType<Deployments>,
    default: () => [],
  },
  responseStatus: {
    type: Number,
    required: true,
  },
  pagination: {
    type: Object,
    required: true,
  },
  alert: {
    type: Object,
    required: true,
  },
})

function getTimeZone(): string {
  return (
    'Date (Timezone ' + Intl.DateTimeFormat().resolvedOptions().timeZone + ')'
  )
}

function convertDate(date: Date): string {
  return moment(date).format('LLLL')
}
</script>
