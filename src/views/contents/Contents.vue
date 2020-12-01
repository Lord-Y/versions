<template>
  <Versions
    :url="url"
    :pagination="pagination"
    :title="title"
    :response-status="responseStatus"
    :deployments="deployments"
    :is-open="isOpen"
  />
</template>

<script>
export default {
  name: 'Contents',
  metaInfo() {
    return {
      title: this.title,
    };
  },
  props: {
    isOpen: {
      type: Boolean,
      default: false,
    },
    url: {
      type: Object,
      required: true,
    },
  },
  components: {
    Versions: () => import('@components/Versions.vue'),
  },
  data() {
    return {
      title: this.buildTitle(),
      deployments: [],
      responseStatus: '',
      pagination: {
        enabled: false,
        data: {},
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      try {
        let target = '';
        let page = '';
        let total;
        const formData = new URLSearchParams();
        switch (this.$route.meta.root) {
          case 'home':
            target = this.$config.BASE_URL + this.url.api.endpoint;
            break;
          case 'platform':
            target = this.$config.BASE_URL + this.url.api.endpoint;
            page = this.$route.params.page ? this.$route.params.page : 1;
            formData.append('workload', this.$route.params.workload);
            formData.append('platform', this.$route.params.platform);
            formData.append('page', page);
            break;
          case 'environment':
            target = this.$config.BASE_URL + this.url.api.endpoint;
            page = this.$route.params.page ? this.$route.params.page : 1;
            formData.append('workload', this.$route.params.workload);
            formData.append('platform', this.$route.params.platform);
            formData.append('environment', this.$route.params.environment);
            formData.append('page', page);
            break;
          default:
            title = 'Versions - HOME';
        }
        this.$axios
          .get(target, {
            params: formData,
          })
          .then((response) => {
            switch (response.status) {
              case 200:
                this.deployments = response.data;
                this.responseStatus = 200;
                switch (this.$route.meta.root) {
                  case 'platform':
                    total = parseInt(response.data[0].count);
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
                  case 'environment':
                    total = parseInt(response.data[0].count);
                    if (total > parseInt(this.$config.RANGE_LIMIT)) {
                      this.pagination.data = {
                        url:
                          '/workload/' +
                          this.$route.params.workload +
                          '/platform/' +
                          this.$route.params.platform +
                          '/environment/' +
                          this.$route.params.environment +
                          '/',
                        actualPage: page,
                        total: total,
                      };
                      this.pagination.enabled = true;
                    }
                    break;
                }
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
        console.error(error);
      }
    },
    buildTitle() {
      let title = '';
      switch (this.$route.meta.root) {
        case 'platform':
          title = 'Platform ' + this.$route.params.platform;
          break;
        case 'environment':
          title =
            'Platform ' +
            this.$route.params.platform +
            ' - environment ' +
            this.$route.params.environment;
          break;
        default:
          title = 'Last deployments';
      }
      return title;
    },
  },
};
</script>
