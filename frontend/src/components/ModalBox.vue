<template>
  <div
    v-show="showModal"
    class="flex items-center flex-col justify-center overflow-hidden fixed inset-0 z-40"
  >
    <div
      class="absolute inset-0 bg-gray-900 bg-opacity-80"
      @click="cancel"
    ></div>
    <div
      class="w-full max-h-modal flex flex-col overflow-hidden relative md:mx-auto md:w-2/5"
    >
      <header
        v-if="title"
        class="flex items-center flex-shrink-0 justify-start px-6 py-4 relative bg-gray-100 border-gray-200 border-b rounded-t"
      >
        <p class="modal-card-title">{{ title }}</p>
      </header>
      <section
        class="bg-white flex-grow flex-shrink overflow-auto p-6 space-y-2"
      >
        <slot />
      </section>
      <footer
        class="flex items-center flex-shrink-0 justify-start px-6 py-4 relative bg-gray-100 border-gray-200 border-t rounded-b"
      >
        <button class="button mr-2" @click="cancelClicked">Cancel</button>
        <button class="button" :class="button" @click="confirmClicked">
          {{ buttonLabel }}
        </button>
      </footer>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";

export default {
  name: "ModalBox",
  props: {
    title: String,
    button: {
      type: String,
      default: "blue",
    },
    buttonLabel: {
      type: String,
      default: "Confirm",
    },
    itemSelected: {
      type: String,
    },
    isModalActive: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["update:confirm", "cancel", "confirm"],
  data() {
    return {};
  },
  methods: {
    cancelClicked() {
      this.$emit("cancel");
    },
    confirmClicked() {
      console.log("confirm clicked ", this.itemSelected);
      this.$emit("confirm", this.itemSelected);
    },
  },
  computed: {
    showModal() {
      return this.isModalActive === true ? true : false;
    },
  },
  mounted() {},
};
</script>
