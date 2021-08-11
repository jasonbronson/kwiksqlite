<template>
  <el-form ref="form" label-width="120px">
    <el-form-item label="Drop Table">
      <el-col :span="11"
        ><b>{{ selectedTable }}</b></el-col
      >
    </el-form-item>
    <el-form-item>
      <el-col :span="11">
        <el-button type="primary" @click="dropTable">Drop Table</el-button>
      </el-col>
    </el-form-item>
  </el-form>
</template>
<script>
import Vue from "vue";
import api from "@/api";

export default {
  name: "Drop",
  data() {
    return {};
  },
  methods: {
    dropTable() {
      api.table.dropTable(this.selectedTable).then((response) => {
        let success = false;
        if (response.status === 200) {
          this.$store.commit("setTables", response.data.tables);
          this.$store.commit("setTableCount", response.data.table_count);
          success = true;

          this.$notify({
            title: "Success",
            message: "Dropping table succeeded",
          });
        }
        if (!success) {
          this.$notify.error({
            title: "Error",
            message: "Dropping table failed",
          });
        }
      });
    },
  },
  computed: {
    selectedTable() {
      return this.$store.state.selectedTable;
    },
  },
};
</script>
