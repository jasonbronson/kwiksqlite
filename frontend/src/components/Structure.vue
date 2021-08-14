<template>
  <el-descriptions title="SQL">
    <el-descriptions-item label="">
      <span class="sql-descrip">{{ sql }}</span></el-descriptions-item
    >
  </el-descriptions>

  <h3 class="header-text">Columns</h3>
  <el-table :data="columns" stripe style="width: 100%">
    <el-table-column prop="ColumnName" label="Column" width="180">
    </el-table-column>
    <el-table-column prop="ColumnType" label="Data type" width="180">
    </el-table-column>
    <el-table-column prop="allow_null" label="Allow null"> </el-table-column>
    <el-table-column prop="primary_key" label="Primary key"> </el-table-column>
    <el-table-column prop="actions" label="Actions">
      <template #default="scope">
        <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
          >Edit</el-button
        >
        <el-button
          size="mini"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)"
          >Delete</el-button
        >
      </template></el-table-column
    >
  </el-table>
</template>
<script>
export default {
  name: "Structure",
  props: {
    sql: String,
    columns: Array,
  },
  data() {
    return {
      tableData: [],
    };
  },
  computed: {},
  mounted() {},
};
</script>
<style scoped>
.sql-descrip {
  padding: 10px;
  background-color: #f5f5f5;
  border: 1px solid gray;
}
.header-text {
  text-align: left;
  padding-top: 20px;
}
</style>
