<template>
  <div>
    <template v-if="responseStatus === 200">
      {{ deployments.raw }}
    </template>
    <template v-if="responseStatus === 404">
      <h2 class="text-center text-muted">This version does not exist</h2>
    </template>
    <template v-if="responseStatus === 500">
      <h2 class="text-center text-danger">
        An error occured while retrieving last deployments. Please try later.
      </h2>
    </template>
  </div>
</template>

<script>
export default {
  name: 'Raw',
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
        formData.append('environment', this.$route.params.environment);
        formData.append('version', this.$route.params.version);
        this.$axios
          .get(this.$config.BASE_URL + this.url.api.endpoint, {
            params: formData,
          })
          .then((response) => {
            switch (response.status) {
              case 200:
                if (this.isJson(response.data)) {
                  this.deployments = JSON.stringify(response.data, null, 2);
                } else {
                  this.deployments = response.data;
                }
                this.responseStatus = 200;
                break;
              default:
                this.responseStatus = 500;
            }
          })
          .catch((error) => {
            if (error.response && error.response.status) {
              switch (error.response.status) {
                case 404:
                  this.responseStatus = 404;
                  break;
                default:
                  this.responseStatus = 500;
              }
            } else {
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
      let title =
        'Platform ' +
        this.$route.params.platform +
        ' - environment ' +
        this.$route.params.environment;
      ' - version ' + this.$route.params.version;
      this.title = title;
      return title;
    },
    isJson(str) {
      try {
        JSON.parse(str);
      } catch (e) {
        return false;
      }
      return true;
    },
  },
};
</script>
