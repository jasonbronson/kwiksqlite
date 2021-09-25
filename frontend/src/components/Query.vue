<template>
  <el-form ref="form" label-width="120px">
    <el-form-item label="Custom Query">
      <el-col :span="11"></el-col>
    </el-form-item>
    <el-form-item label="SQL">
      <el-input
        type="textarea"
        v-model="query"
        placeholder="Select * from tableName"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-col :span="11">
        <el-button type="primary" @click="customQuery">Execute</el-button>
      </el-col>
    </el-form-item>
  </el-form>
</template>
<script>
import Vue from "vue";
import api from "@/api";

export default {
  name: "Query",
  data() {
    return {
      query: "",
    };
  },
  methods: {
    customQuery() {
      api.query.customQuery(this.query).then((response) => {
        let success = false;
        if (response.status === 200) {
          success = true;
          this.$notify.success({
            title: "Query Executed",
            message: "query was successfully run"
          })
        }
        if (!success) {
          console.log(response, "*");
          this.$notify.error({
            title: "Query failed",
            message: "Error:" + response.error,
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
