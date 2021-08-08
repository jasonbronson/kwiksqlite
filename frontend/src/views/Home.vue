<template>
  <el-main>
    <el-table :data="tableData">
      <el-table-column prop="date" label="Date" width="140"> </el-table-column>
      <el-table-column prop="name" label="Name" width="120"> </el-table-column>
      <el-table-column prop="address" label="Address"> </el-table-column>
    </el-table>
  </el-main>
</template>

<script>
// @ is an alias to /src
import { ref, onMounted } from "vue";

export default {
  name: "Home",
  components: {},
  data() {
    const item = {
      date: "2016-05-02",
      name: "Tom",
      address: "No. 189, Grove St, Los Angeles",
    };
    return {
      tableData: Array(20).fill(item),
    };
  },
  methods: {
    test() {
      console.log("test");
      this.axios.get("/").then((response) => (this.data = response.data));
    },
  },
};
</script>
