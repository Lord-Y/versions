<template>
  <div class="my-10">
    <div class="flex justify-center">
      <template v-if="actualPage === 1">
        <template v-if="ceil === 2">
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              :title="$t('pagination.begin')"
            >
              {{ $t('pagination.begin') }}
            </router-link>
          </div>
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              :title="$t('pagination.previous')"
            >
              {{ $t('pagination.previous') }}
            </router-link>
          </div>
          <div class="bg-green-700">
            <router-link
              class="py-2 px-3 border border-gray-300 text-white bg-green-700"
              :to="paginationUrl + '1'"
              title="1"
            >
              1
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '2'"
              title="2"
            >
              2
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '2'"
              :title="$t('pagination.next')"
            >
              {{ $t('pagination.next') }}
            </router-link>
          </div>
          <div>
            <router-link
              class="rounded rounded-l-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + ceil"
              :title="$t('pagination.end')"
            >
              {{ $t('pagination.end') }}
            </router-link>
          </div>
        </template>
        <template v-else>
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              :title="$t('pagination.begin')"
            >
              {{ $t('pagination.begin') }}
            </router-link>
          </div>
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              :title="$t('pagination.previous')"
            >
              {{ $t('pagination.previous') }}
            </router-link>
          </div>
          <div class="bg-green-700">
            <router-link
              class="py-2 px-3 border border-gray-300 text-white bg-green-700"
              :to="paginationUrl + '1'"
              title="1"
            >
              1
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '2'"
              title="2"
            >
              2
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '3'"
              title="3"
            >
              3
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '2'"
              :title="$t('pagination.next')"
            >
              {{ $t('pagination.next') }}
            </router-link>
          </div>
          <div>
            <router-link
              class="rounded py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + ceil"
              :title="$t('pagination.end')"
            >
              {{ $t('pagination.end') }}
            </router-link>
          </div>
        </template>
      </template>
      <template v-else>
        <template v-if="actualPage === ceil">
          <div>
            <router-link
              class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + '1'"
              :title="$t('pagination.begin')"
            >
              {{ $t('pagination.begin') }}
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + previous"
              :title="$t('pagination.previous')"
            >
              {{ $t('pagination.previous') }}
            </router-link>
          </div>
          <div v-if="ceil > 2">
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + (previous - 1)"
              :title="previous - 1"
            >
              {{ previous - 1 }}
            </router-link>
          </div>
          <div>
            <router-link
              class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
              :to="paginationUrl + previous"
              :title="previous"
            >
              {{ previous }}
            </router-link>
          </div>
          <div class="bg-green-700">
            <router-link
              class="py-2 px-3 border border-gray-300 text-white bg-green-700"
              :to="paginationUrl + pagination.actualPage"
              :title="pagination.actualPage"
            >
              {{ pagination.actualPage }}
            </router-link>
          </div>
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              :to="paginationUrl + next"
              :title="$t('pagination.next')"
            >
              {{ $t('pagination.next') }}
            </router-link>
          </div>
          <div class="disabled">
            <router-link
              class="cursor-default pointer-events-none rounded rounded-l-none py-2 px-3 border border-gray-300 text-gray-600"
              :to="paginationUrl + ceil"
              :title="$t('pagination.end')"
            >
              {{ $t('pagination.end') }}
            </router-link>
          </div>
        </template>
        <template v-else>
          <template v-if="ceil === 2">
            <div>
              <router-link
                class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                :to="paginationUrl + '1'"
                :title="$t('pagination.begin')"
              >
                {{ $t('pagination.begin') }}
              </router-link>
            </div>
            <div>
              <router-link
                class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                :to="paginationUrl + previous"
                :title="$t('pagination.previous')"
              >
                {{ $t('pagination.previous') }}
              </router-link>
            </div>
            <div>
              <router-link
                class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                :to="paginationUrl + previous"
                :title="$t('pagination.previous')"
              >
                {{ previous }}
              </router-link>
            </div>
            <div class="bg-green-700">
              <router-link
                class="py-2 px-3 border border-gray-300 text-white bg-green-700"
                :to="paginationUrl + pagination.actualPage"
                :title="pagination.actualPage"
              >
                {{ pagination.actualPage }}
              </router-link>
            </div>
            <template v-if="actualPage === ceil">
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="$t('pagination.next')"
                >
                  {{ $t('pagination.next') }}
                </router-link>
              </div>
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600 hover:bg-gray-300"
                  :to="paginationUrl + ceil"
                  :title="$t('pagination.end')"
                >
                  {{ $t('pagination.end') }}
                </router-link>
              </div>
            </template>
            <template v-else>
              <div>
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="next"
                >
                  {{ next }}
                </router-link>
              </div>
              <div>
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="$t('pagination.next')"
                >
                  {{ $t('pagination.next') }}
                </router-link>
              </div>
              <div>
                <router-link
                  class="rounded py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + ceil"
                  :title="$t('pagination.end')"
                >
                  {{ $t('pagination.end') }}
                </router-link>
              </div>
            </template>
          </template>
          <template v-else>
            <template v-if="previous < 1">
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
                  :to="paginationUrl + '1'"
                  :title="$t('pagination.begin')"
                >
                  {{ $t('pagination.begin') }}
                </router-link>
              </div>
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + previous"
                  :title="$t('pagination.previous')"
                >
                  {{ $t('pagination.previous') }}
                </router-link>
              </div>
            </template>
            <template v-else>
              <div>
                <router-link
                  class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + '1'"
                  :title="$t('pagination.begin')"
                >
                  {{ $t('pagination.begin') }}
                </router-link>
              </div>
              <div>
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + previous"
                  :title="$t('pagination.previous')"
                >
                  {{ $t('pagination.previous') }}
                </router-link>
              </div>
            </template>
            <div v-if="actualPage === ceil">
              <router-link
                class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                :to="paginationUrl + (previous - 1)"
                :title="previous - 1"
              >
                {{ previous - 1 }}
              </router-link>
            </div>
            <div v-if="previous > 0">
              <router-link
                class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                :to="paginationUrl + previous"
                :title="previous"
              >
                {{ previous }}
              </router-link>
            </div>
            <div class="bg-green-700">
              <router-link
                class="py-2 px-3 border border-gray-300 text-white bg-green-700"
                :to="paginationUrl + pagination.actualPage"
                :title="pagination.actualPage"
              >
                {{ pagination.actualPage }}
              </router-link>
            </div>
            <template v-if="actualPage === ceil">
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="$t('pagination.next')"
                >
                  {{ $t('pagination.next') }}
                </router-link>
              </div>
              <div class="disabled">
                <router-link
                  class="cursor-default pointer-events-none rounded rounded-l-none py-2 px-3 border border-gray-300 text-gray-600"
                  :to="paginationUrl + ceil"
                  :title="$t('pagination.end')"
                >
                  {{ $t('pagination.end') }}
                </router-link>
              </div>
            </template>
            <template v-else>
              <div>
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="next"
                >
                  {{ next }}
                </router-link>
              </div>
              <div v-if="actualPage === 1">
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + (next + 1)"
                  :title="next + 1"
                >
                  {{ next + 1 }}
                </router-link>
              </div>
              <div>
                <router-link
                  class="py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + next"
                  :title="$t('pagination.next')"
                >
                  {{ $t('pagination.next') }}
                </router-link>
              </div>
              <div>
                <router-link
                  class="rounded rounded-l-none py-2 px-3 border border-gray-300 text-green-600 hover:bg-gray-300"
                  :to="paginationUrl + ceil"
                  :title="$t('pagination.end')"
                >
                  {{ $t('pagination.end') }}
                </router-link>
              </div>
            </template>
          </template>
        </template>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, PropType } from 'vue'
import { Pagination } from '@/apis/interfaces'
import { config } from '@/apis/configs'

const props = defineProps({
  pagination: {
    type: Object as PropType<Pagination>,
    required: true,
  },
})

const actualPage = computed(() => {
  return props.pagination.actualPage
})

const ceil = computed(() => {
  return Math.ceil(props.pagination.total / config.RANGE_LIMIT)
})

const previous = computed(() => {
  return props.pagination.actualPage - 1
})

const next = computed(() => {
  return props.pagination.actualPage + 1
})

const paginationUrl = computed(() => {
  let url = props.pagination.url
  if (props.pagination.mode !== undefined) {
    url = props.pagination.url
  } else {
    const regex = /\/$/g
    const found = props.pagination.url.match(regex)
    if (found === null) {
      url = props.pagination.url + '/'
    } else {
      url = props.pagination.url
    }
  }
  return url
})
</script>
