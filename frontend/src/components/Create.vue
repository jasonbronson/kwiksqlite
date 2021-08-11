<template>
  <el-tabs type="card" @tab-click="handleClick">
    <el-tab-pane label="Create">
      <el-form ref="form" label-width="120px">
        <el-form-item label="">
          <el-form-item label="Table name">
            <el-input v-model="tableName"></el-input>
          </el-form-item>
        </el-form-item>
        <el-form-item>
          <el-col :span="11">
            <el-button type="primary" @click="createTable"
              >Create Table</el-button
            >
          </el-col>
        </el-form-item>
      </el-form>
    </el-tab-pane>
    <el-tab-pane label="Import">Not Completed Yet</el-tab-pane>
  </el-tabs>
</template>
<script>
import Vue from "vue";
import api from "@/api";

export default {
  name: "Drop",
  data() {
    return {
      tableName: "",
    };
  },
  methods: {
    handleClick() {},
    createTable() {
      api.table.createTable(this.tableName).then((response) => {
        let success = false;
        if (response.status === 200) {
          this.$store.commit("setTables", response.data.tables);
          this.$store.commit("setTableCount", response.data.table_count);
          success = true;
          this.$notify({
            title: "Success",
            message: "Create table succeeded",
          });
        }
        if (!success) {
          this.$notify.error({
            title: "Error",
            message: "Create table failed",
          });
        }
      });
    },
  },
  computed: {},
};
</script>
