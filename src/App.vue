<template>
  <div id="app">
    <template v-if="$route.meta.root === 'raw'">
      <transition name="slide">
        <router-view :key="$route.fullPath" />
      </transition>
    </template>
    <template
      v-else-if="
        $route.meta.root === 'notFound' ||
        $route.meta.root === 'internalServerError'
      "
    >
      <transition name="slide">
        <router-view :key="$route.fullPath" />
      </transition>
    </template>
    <template v-else>
      <div class="font-sans">
        <ToggleMenu :is-open="isOpen" @update:navIsOpen="navIsOpen = $event" />
        <div class="flex sm:min-h-screen">
          <Menu :is-open="navIsOpen" />
          <transition name="slide">
            <router-view :is-open="navIsOpen" :key="$route.fullPath" />
          </transition>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import ToggleMenu from '@views/menu/ToggleMenu.vue';
import Menu from '@views/menu/Menu.vue';

export default {
  metaInfo: {
    title: 'Versions',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        vmid: 'description',
        name: 'description',
        content: 'Versiions application',
      },
    ],
  },
  components: {
    ToggleMenu,
    Menu,
  },
  data() {
    return {
      isOpen: false,
      navIsOpen: false,
    };
  },
};
</script>
