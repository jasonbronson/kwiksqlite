<template>
  <footer v-show="isFooterBarVisible" class="bg-white py-2 px-6">
    <div class="container-fluid">
      <div class="level">
        <div class="level-left">
          <div class="level-item">
            <div class="footer-text">
              <b></b>Web-based Sqlite Database Browser
            </div>
          </div>
        </div>
        <div class="level-right">
          <div class="level-item">
            <div class="logo">
              <a
                href="https://github.com/jasonbronson/kwiksqlite-admin"
                target="_new"
              >
                Kwiksqlite Admin
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import dayjs from "dayjs";
import { computed } from "vue";
import { useStore } from "vuex";

export default {
  name: "FooterBar",
  setup() {
    const store = useStore();

    const year = computed(() => dayjs().year());

    const isFooterBarVisible = computed(() => !store.state.isFormScreen);

    return {
      year,
      isFooterBarVisible,
    };
  },
};
</script>
