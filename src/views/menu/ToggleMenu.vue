<template>
  <header
    class="bg-gray-800 sm:flex sm:justify-between sm:items-center sm:px-4 sm:py-2"
  >
    <div class="flex items-center justify-between px-4 py-3 sm:p-0">
      <div class="text-white text-xl">
        <a class="hover:text-blue-600" href="/" title="Versions">Versions</a>
      </div>
      <div class="sm:hidden">
        <button
          type="button"
          class="block text-gray-500 hover:text-white focus:text-white focus:outline-none"
          @click="
            navIsOpen = !navIsOpen;
            $emit('update:navIsOpen', navIsOpen);
          "
        >
          <span v-if="navIsOpen"> X </span>
          <span v-if="!navIsOpen">
            <svg viewBox="0 0 24 24" class="h-6 w-6 fill-current">
              <path
                fill-rule="evenodd"
                d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
              ></path>
            </svg>
          </span>
        </button>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'ToggleMenu',
  props: {
    isOpen: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      navIsOpen: this.isOpen,
    };
  },
};
</script>
