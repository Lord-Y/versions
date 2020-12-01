<template>
  <main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">
    <div
      class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"
    >
      <h1 class="h1">Last deployments</h1>
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
  name: 'Home',
  metaInfo() {
    return {
      title: 'Versions - HOME',
    };
  },
  props: {
    url: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      deployments: [],
      responseStatus: '',
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      try {
        this.$axios
          .get(this.$config.BASE_URL + this.url.api.home)
          .then((response) => {
            switch (response.status) {
              case 200:
                this.deployments = response.data;
                this.responseStatus = 200;
                break;
              case 204:
                this.deployments = response.data;
                this.responseStatus = 204;
                break;
              default:
                this.responseStatus = 500;
            }
          })
          .catch((error) => {
            this.responseStatus = 500;
            console.error(error);
          });
      } catch (error) {
        this.responseStatus = 500;
        throw error;
      }
    },
  },
};
</script>
