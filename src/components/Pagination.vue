<template>
  <div class="mt-3">
    <template v-if="pagination.actualPage == 1">
      <template v-if="ceil === 2">
        <nav aria-label="navigation">
          <ul class="pagination justify-content-center">
            <li class="page-item disabled">
              <a
                class="page-link"
                href="#"
                tabindex="-1"
                :title="text.beginning"
              >
                {{ text.beginning }}
              </a>
            </li>
            <li class="page-item disabled">
              <a
                class="page-link"
                href="#"
                tabindex="-1"
                :title="text.previous"
              >
                {{ text.previous }}
              </a>
            </li>
            <li class="page-item active">
              <RouterLink
                class="page-link"
                :to="pagination.url + '1'"
                exact
                title="1"
              >
                1
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '2'"
                exact
                title="2"
              >
                2
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '2'"
                exact
                :title="text.next"
              >
                {{ text.next }}
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + ceil"
                exact
                :title="text.last"
              >
                {{ text.last }}
              </RouterLink>
            </li>
          </ul>
        </nav>
      </template>
      <template v-else>
        <nav aria-label="navigation">
          <ul class="pagination justify-content-center">
            <li class="page-item disabled">
              <a
                class="page-link"
                href="#"
                tabindex="-1"
                :title="text.beginning"
              >
                {{ text.beginning }}
              </a>
            </li>
            <li class="page-item disabled">
              <a
                class="page-link"
                href="#"
                tabindex="-1"
                :title="text.previous"
              >
                {{ text.previous }}
              </a>
            </li>
            <li class="page-item active">
              <RouterLink
                class="page-link"
                :to="pagination.url + '1'"
                exact
                title="1"
              >
                1
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '2'"
                exact
                title="2"
              >
                2
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '3'"
                exact
                title="3"
              >
                3
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '2'"
                exact
                :title="text.next"
              >
                {{ text.next }}
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + ceil"
                exact
                :title="text.last"
              >
                {{ text.last }}
              </RouterLink>
            </li>
          </ul>
        </nav>
      </template>
    </template>
    <template v-else>
      <template v-if="pagination.actualPage == ceil">
        <nav aria-label="navigation">
          <ul class="pagination justify-content-center">
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + '1'"
                exact
                :title="text.beginning"
              >
                {{ text.beginning }}
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + previous"
                exact
                :title="text.previous"
              >
                {{ text.previous }}
              </RouterLink>
            </li>
            <li class="page-item" v-if="ceil > 2">
              <RouterLink
                class="page-link"
                :to="pagination.url + (previous - 1)"
                exact
                :title="previous - 1"
              >
                {{ previous - 1 }}
              </RouterLink>
            </li>
            <li class="page-item">
              <RouterLink
                class="page-link"
                :to="pagination.url + previous"
                exact
                :title="previous"
              >
                {{ previous }}
              </RouterLink>
            </li>
            <li class="page-item active">
              <RouterLink
                class="page-link"
                :to="pagination.url + pagination.actualPage"
                exact
                :title="pagination.actualPage"
              >
                {{ pagination.actualPage }}
              </RouterLink>
            </li>
            <li class="page-item disabled">
              <RouterLink
                class="page-link"
                :to="pagination.url + next"
                exact
                :title="text.next"
              >
                {{ text.next }}
              </RouterLink>
            </li>
            <li class="page-item disabled">
              <RouterLink
                class="page-link"
                :to="pagination.url + ceil"
                exact
                :title="text.last"
              >
                {{ text.last }}
              </RouterLink>
            </li>
          </ul>
        </nav>
      </template>
      <template v-else>
        <template v-if="ceil === 2">
          <nav aria-label="navigation">
            <ul class="pagination justify-content-center">
              <li class="page-item">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + '1'"
                  exact
                  :title="text.beginning"
                >
                  {{ text.beginning }}
                </RouterLink>
              </li>
              <li class="page-item">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + previous"
                  exact
                  :title="text.previous"
                >
                  {{ text.previous }}
                </RouterLink>
              </li>
              <li class="page-item">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + previous"
                  exact
                  :title="$t('previous')"
                >
                  {{ previous }}
                </RouterLink>
              </li>
              <li class="page-item active">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + pagination.actualPage"
                  exact
                  :title="pagination.actualPage"
                >
                  {{ pagination.actualPage }}
                </RouterLink>
              </li>
              <template v-if="pagination.actualPage == ceil">
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="text.next"
                  >
                    {{ text.next }}
                  </RouterLink>
                </li>
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + ceil"
                    exact
                    :title="text.last"
                  >
                    {{ text.last }}
                  </RouterLink>
                </li>
              </template>
              <template v-else>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="next"
                  >
                    {{ next }}
                  </RouterLink>
                </li>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="text.next"
                  >
                    {{ text.next }}
                  </RouterLink>
                </li>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + ceil"
                    exact
                    :title="text.last"
                  >
                    {{ text.last }}
                  </RouterLink>
                </li>
              </template>
            </ul>
          </nav>
        </template>
        <template v-else>
          <nav aria-label="navigation">
            <ul class="pagination justify-content-center">
              <template v-if="previous < 1">
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + '1'"
                    exact
                    :title="text.beginning"
                  >
                    {{ text.beginning }}
                  </RouterLink>
                </li>
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + previous"
                    exact
                    :title="text.previous"
                  >
                    {{ text.previous }}
                  </RouterLink>
                </li>
              </template>
              <template v-else>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + '1'"
                    exact
                    :title="text.beginning"
                  >
                    {{ text.beginning }}
                  </RouterLink>
                </li>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + previous"
                    exact
                    :title="text.previous"
                  >
                    {{ text.previous }}
                  </RouterLink>
                </li>
              </template>
              <li class="page-item" v-if="pagination.actualPage == ceil">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + (previous - 1)"
                  exact
                  :title="previous - 1"
                >
                  {{ previous - 1 }}
                </RouterLink>
              </li>
              <li class="page-item" v-if="previous > 0">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + previous"
                  exact
                  :title="previous"
                >
                  {{ previous }}
                </RouterLink>
              </li>
              <li class="page-item active">
                <RouterLink
                  class="page-link"
                  :to="pagination.url + pagination.actualPage"
                  exact
                  :title="pagination.actualPage"
                >
                  {{ pagination.actualPage }}
                </RouterLink>
              </li>
              <template v-if="pagination.actualPage == ceil">
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="text.next"
                  >
                    {{ text.next }}
                  </RouterLink>
                </li>
                <li class="page-item disabled">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + ceil"
                    exact
                    :title="text.last"
                  >
                    {{ text.last }}
                  </RouterLink>
                </li>
              </template>
              <template v-else>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="next"
                  >
                    {{ next }}
                  </RouterLink>
                </li>
                <li class="page-item" v-if="pagination.actualPage == 1">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + (next + 1)"
                    exact
                    :title="next + 1"
                  >
                    {{ next + 1 }}
                  </RouterLink>
                </li>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + next"
                    exact
                    :title="text.next"
                  >
                    {{ text.next }}
                  </RouterLink>
                </li>
                <li class="page-item">
                  <RouterLink
                    class="page-link"
                    :to="pagination.url + ceil"
                    exact
                    :title="text.last"
                  >
                    {{ text.last }}
                  </RouterLink>
                </li>
              </template>
            </ul>
          </nav>
        </template>
      </template>
    </template>
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
      return Math.ceil(this.pagination.total / parseInt(2));
      // this.pagination.total / parseInt(this.$config.RANGE_LIMIT),
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
