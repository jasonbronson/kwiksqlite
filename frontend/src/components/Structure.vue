<template>
  <el-descriptions title="SQL">
    <el-descriptions-item label="">
      <span class="sql-descrip">{{ sql }}</span></el-descriptions-item
    >
  </el-descriptions>
</template>
<script>
import Vue from "vue";

export default {
  name: "Structure",
  props: {
    sql: String,
  },
  data() {
    return {};
  },
};
</script>
<style scoped>
.sql-descrip {
  padding: 10px;
  background-color: #f5f5f5;
  border: 1px solid gray;
}
</style>
