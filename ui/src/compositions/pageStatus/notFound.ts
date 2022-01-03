import { reactive, toRefs } from 'vue'
import { useI18n } from 'vue-i18n'
import { useHead } from '@vueuse/head'

export default function () {
  const state = reactive({
    meta: {
      title: '',
      description: '',
    },
  })

  const { t } = useI18n({
    useScope: 'global',
  })

  state.meta.title = t('alert.http.pageNotFound')
  state.meta.description = t('alert.http.pageNotFound')

  useHead({
    title: state.meta.title,
    meta: [
      {
        name: 'description',
        content: state.meta.description,
      },
      {
        property: 'og:title',
        content: state.meta.title,
      },
      {
        property: 'og:description',
        content: state.meta.description,
      },
    ],
  })

  return {
    ...toRefs(state),
  }
}
