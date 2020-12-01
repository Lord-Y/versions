<template>
  <main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">
    <div
      class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"
    >
      <h1 class="h1">{{ title }}</h1>
    </div>
    <canvas class="my-4 w-100" id="myChart" width="900" height="25"></canvas>
    <template v-if="responseStatus === 200">
      <div class="table-responsive">
        <table class="table table-hover table-striped table-sm">
          <thead>
            <tr>
              <th>Workload</th>
              <th>Platform</th>
              <th>Environment</th>
              <th>Version</th>
              <th>Changelog</th>
              <th>Raw</th>
              <th>Date</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(deployment, index) in deployments" :key="index">
              <td>{{ deployment.workload | toUpperCase }}</td>
              <td>{{ deployment.platform }}</td>
              <td>{{ deployment.environment }}</td>
              <td>{{ deployment.version }}</td>
              <td
                v-if="
                  deployment.changelog_url === 'NULL' ||
                  deployment.changelog_url === ''
                "
              >
                N/A
              </td>
              <td v-else>
                <a
                  class="nav-link"
                  :href="deployment.changelog_url"
                  title="See changelog"
                >
                  {{ deployment.version }}
                </a>
              </td>
              <td v-if="deployment.raw === 'NULL' || deployment.raw === ''">
                N/A
              </td>
              <td v-else>
                <a
                  class="nav-link"
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
                  title="See version content"
                >
                  {{ deployment.version }}
                </a>
              </td>
              <td>{{ deployment.date }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <Pagination v-if="pagination.enabled" :pagination="pagination.data" />
    </template>
    <template v-if="responseStatus === 204">
      <h2 class="text-center text-muted">
        So far no deployment has been registered
      </h2>
    </template>
    <template v-if="responseStatus === 500">
      <h2 class="text-center text-danger">
        An error occured while retrieving last deployments. Please try later.
      </h2>
    </template>
  </main>
</template>

<script>
export default {
  name: 'Environment',
  metaInfo() {
    return {
      title: this.buildTitle(),
    };
  },
  props: {
    url: {
      type: Object,
      required: true,
    },
  },
  components: {
    Pagination: () => import('@components/Pagination.vue'),
  },
  data() {
    return {
      deployments: [],
      responseStatus: '',
      title: '',
      pagination: {
        data: {},
        enabled: false,
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      try {
        let page = this.$route.params.page ? this.$route.params.page : 1;
        const formData = new URLSearchParams();
        formData.append('workload', this.$route.params.workload);
        formData.append('platform', this.$route.params.platform);
        formData.append('page', page);
        this.$axios
          .get(this.$config.BASE_URL + this.url.api.platform, {
            params: formData,
          })
          .then((response) => {
            switch (response.status) {
              case 200:
                this.deployments = response.data;
                this.responseStatus = 200;
                let total = parseInt(response.data[0].count);
                if (total > parseInt(this.$config.RANGE_LIMIT)) {
                  this.pagination.data = {
                    url:
                      '/workload/' +
                      this.$route.params.workload +
                      '/platform/' +
                      this.$route.params.platform,
                    actualPage: page,
                    total: total,
                  };
                  this.pagination.enabled = true;
                }
                break;
              default:
                this.responseStatus = 500;
            }
          })
          .catch((error) => {
            switch (error.response.status) {
              case 404:
                this.responseStatus = 404;
                break;
              default:
                this.responseStatus = 500;
            }
            console.error(error);
          });
      } catch (error) {
        this.responseStatus = 500;
        throw error;
      }
    },
    buildTitle() {
      let title = 'Platform ' + this.$route.params.platform;
      this.title = title;
      return title;
    },
  },
};
</script>
