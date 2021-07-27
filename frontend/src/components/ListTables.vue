<template>
  <modal-box
    :is-modal-active="isModalActive"
    title="Please confirm action"
    @confirm="deleteTable($event)"
    button-label="Delete"
    @cancel="isModalActive = false"
    :item-selected="itemSelected"
  >
    <p>Are you sure you want to {{ modaltext }}?</p>
  </modal-box>

  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in tabledata" :key="item.Name">
        <td data-label="Name">{{ item.Name }}</td>
        <td class="actions-cell">
          <div class="buttons right nowrap">
            <button class="button medium blue" type="button" @click="">
              View Table Data
            </button>
          </div>
        </td>
        <td class="actions-cell">
          <div class="buttons right nowrap">
            <button
              class="button medium red"
              type="button"
              @click="storeItemSelected(item.Name)"
            >
              Delete Table
            </button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
  <div class="table-pagination" v-if="pagination">
    <div class="level">
      <div class="level-left">
        <div class="level-item">
          <div class="buttons">
            <button
              v-for="page in pagesList"
              @click="currentPage = page"
              type="button"
              class="button"
              :class="{ active: page === currentPage }"
              :key="page"
            >
              {{ page + 1 }}
            </button>
          </div>
        </div>
      </div>
      <div class="level-right">
        <div class="level-item">
          <small>Page {{ currentPageHuman }} of {{ numPages }}</small>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
//import { computed, ref } from "vue";
//import { mdiEye, mdiTrashCan } from "@mdi/js";
// import axios from "axios";
// import slice from "lodash/slice";
// import remove from "lodash/remove";
import ModalBox from "@/components/ModalBox";
//import Icon from "@/components/Icon";
//import CheckboxCell from "@/components/CheckboxCell";

export default {
  name: "DBTable",
  components: {
    ModalBox,
  },
  props: {
    // checkable: {
    //   type: Boolean,
    //   default: false,
    // },
    pagination: {
      type: Boolean,
      default: false,
    },
    tabledata: {
      type: Array,
    },
  },
  data() {
    return {
      modaltext: "proceed",
      isModalActive: false,
      itemSelected: "",
    };
  },
  methods: {
    deleteTable(event) {
      console.log("this is emitted from child: ", event);
      this.axios.delete("/table/drop?table=" + event).then((response) => {
        if (response.status === 200) {
          console.log("dropped table ", event);
          this.isModalActive = false;
          this.$emit("reload");
        }
      });
    },
    storeItemSelected(item) {
      this.itemSelected = item;
      this.isModalActive = true;
    },
  },
  mounted() {},
};
</script>
