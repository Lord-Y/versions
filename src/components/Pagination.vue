<template>
  <div class="mt-5">
    <div class="flex justify-center">
      <template v-if="parseInt(pagination.actualPage) === 1">
        <template v-if="ceil === 2">
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              exact
              :title="text.beginning"
            >
              {{ text.beginning }}
            </RouterLink>
          </div>
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              exact
              :title="text.previous"
            >
              {{ text.previous }}
            </RouterLink>
          </div>
          <div class="bg-blue-500">
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-white bg-blue-500"
              :to="pagination.url + '1'"
              exact
              title="1"
            >
              1
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '2'"
              exact
              title="2"
            >
              2
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '2'"
              exact
              :title="text.next"
            >
              {{ text.next }}
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="rounded rounded-l-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + ceil"
              exact
              :title="text.last"
            >
              {{ text.last }}
            </RouterLink>
          </div>
        </template>
        <template v-else>
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              exact
              :title="text.beginning"
            >
              {{ text.beginning }}
            </RouterLink>
          </div>
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              to="#"
              exact
              :title="text.previous"
            >
              {{ text.previous }}
            </RouterLink>
          </div>
          <div class="bg-blue-500">
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-white bg-blue-500"
              :to="pagination.url + '1'"
              exact
              title="1"
            >
              1
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '2'"
              exact
              title="2"
            >
              2
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '3'"
              exact
              title="3"
            >
              3
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '2'"
              exact
              :title="text.next"
            >
              {{ text.next }}
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="rounded py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + ceil"
              exact
              :title="text.last"
            >
              {{ text.last }}
            </RouterLink>
          </div>
        </template>
      </template>
      <template v-else>
        <template v-if="parseInt(pagination.actualPage) === ceil">
          <div>
            <RouterLink
              class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + '1'"
              exact
              :title="text.beginning"
            >
              {{ text.beginning }}
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + previous"
              exact
              :title="text.previous"
            >
              {{ text.previous }}
            </RouterLink>
          </div>
          <div v-if="ceil > 2">
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + (previous - 1)"
              exact
              :title="previous - 1"
            >
              {{ previous - 1 }}
            </RouterLink>
          </div>
          <div>
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
              :to="pagination.url + previous"
              exact
              :title="previous"
            >
              {{ previous }}
            </RouterLink>
          </div>
          <div class="bg-blue-500">
            <RouterLink
              class="py-2 px-3 border border-gray-300 text-white bg-blue-500"
              :to="pagination.url + pagination.actualPage"
              exact
              :title="pagination.actualPage"
            >
              {{ pagination.actualPage }}
            </RouterLink>
          </div>
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600"
              :to="pagination.url + next"
              exact
              :title="text.next"
            >
              {{ text.next }}
            </RouterLink>
          </div>
          <div class="disabled">
            <RouterLink
              class="cursor-default pointer-events-none rounded rounded-l-none py-2 px-3 border border-gray-300 text-gray-600"
              :to="pagination.url + ceil"
              exact
              :title="text.last"
            >
              {{ text.last }}
            </RouterLink>
          </div>
        </template>
        <template v-else>
          <template v-if="ceil === 2">
            <div>
              <RouterLink
                class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                :to="pagination.url + '1'"
                exact
                :title="text.beginning"
              >
                {{ text.beginning }}
              </RouterLink>
            </div>
            <div>
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                :to="pagination.url + previous"
                exact
                :title="text.previous"
              >
                {{ text.previous }}
              </RouterLink>
            </div>
            <div>
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                :to="pagination.url + previous"
                exact
                :title="text.previous"
              >
                {{ previous }}
              </RouterLink>
            </div>
            <div class="bg-blue-500">
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-white bg-blue-500"
                :to="pagination.url + pagination.actualPage"
                exact
                :title="pagination.actualPage"
              >
                {{ pagination.actualPage }}
              </RouterLink>
            </div>
            <template v-if="parseInt(pagination.actualPage) === ceil">
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="text.next"
                >
                  {{ text.next }}
                </RouterLink>
              </div>
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-gray-600 hover:bg-gray-300"
                  :to="pagination.url + ceil"
                  exact
                  :title="text.last"
                >
                  {{ text.last }}
                </RouterLink>
              </div>
            </template>
            <template v-else>
              <div>
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="next"
                >
                  {{ next }}
                </RouterLink>
              </div>
              <div>
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="text.next"
                >
                  {{ text.next }}
                </RouterLink>
              </div>
              <div>
                <RouterLink
                  class="rounded py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + ceil"
                  exact
                  :title="text.last"
                >
                  {{ text.last }}
                </RouterLink>
              </div>
            </template>
          </template>
          <template v-else>
            <template v-if="previous < 1">
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none rounded rounded-r-none py-2 px-3 border border-gray-300 text-gray-600"
                  :to="pagination.url + '1'"
                  exact
                  :title="text.beginning"
                >
                  {{ text.beginning }}
                </RouterLink>
              </div>
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + previous"
                  exact
                  :title="text.previous"
                >
                  {{ text.previous }}
                </RouterLink>
              </div>
            </template>
            <template v-else>
              <div>
                <RouterLink
                  class="rounded rounded-r-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + '1'"
                  exact
                  :title="text.beginning"
                >
                  {{ text.beginning }}
                </RouterLink>
              </div>
              <div>
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + previous"
                  exact
                  :title="text.previous"
                >
                  {{ text.previous }}
                </RouterLink>
              </div>
            </template>
            <div v-if="parseInt(pagination.actualPage) === ceil">
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                :to="pagination.url + (previous - 1)"
                exact
                :title="previous - 1"
              >
                {{ previous - 1 }}
              </RouterLink>
            </div>
            <div v-if="previous > 0">
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                :to="pagination.url + previous"
                exact
                :title="previous"
              >
                {{ previous }}
              </RouterLink>
            </div>
            <div class="bg-blue-500">
              <RouterLink
                class="py-2 px-3 border border-gray-300 text-white bg-blue-500"
                :to="pagination.url + pagination.actualPage"
                exact
                :title="pagination.actualPage"
              >
                {{ pagination.actualPage }}
              </RouterLink>
            </div>
            <template v-if="parseInt(pagination.actualPage) === ceil">
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="text.next"
                >
                  {{ text.next }}
                </RouterLink>
              </div>
              <div class="disabled">
                <RouterLink
                  class="cursor-default pointer-events-none rounded rounded-l-none py-2 px-3 border border-gray-300 text-gray-600"
                  :to="pagination.url + ceil"
                  exact
                  :title="text.last"
                >
                  {{ text.last }}
                </RouterLink>
              </div>
            </template>
            <template v-else>
              <div>
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="next"
                >
                  {{ next }}
                </RouterLink>
              </div>
              <div v-if="parseInt(pagination.actualPage) === 1">
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + (next + 1)"
                  exact
                  :title="next + 1"
                >
                  {{ next + 1 }}
                </RouterLink>
              </div>
              <div>
                <RouterLink
                  class="py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + next"
                  exact
                  :title="text.next"
                >
                  {{ text.next }}
                </RouterLink>
              </div>
              <div>
                <RouterLink
                  class="rounded rounded-l-none py-2 px-3 border border-gray-300 text-blue-600 hover:bg-gray-300"
                  :to="pagination.url + ceil"
                  exact
                  :title="text.last"
                >
                  {{ text.last }}
                </RouterLink>
              </div>
            </template>
          </template>
        </template>
      </template>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Pagination',
  props: {
    pagination: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      page: {
        actualPage: this.pagination.actualPage,
      },
      text: {
        beginning: 'Beginning',
        last: 'Last',
        next: 'Next',
        previous: 'Previous',
      },
    };
  },
  computed: {
    ceil: function () {
      return Math.ceil(
        this.pagination.total / parseInt(this.$config.RANGE_LIMIT),
      );
    },
    previous: function () {
      return parseInt(this.pagination.actualPage) - 1;
    },
    next: function () {
      return parseInt(this.pagination.actualPage) + 1;
    },
    scrollToTop: function () {
      return window.scrollTo(0, 0);
    },
  },
};
</script>
