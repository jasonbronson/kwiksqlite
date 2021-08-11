<template>
  <el-form ref="form" :model="form">
    <el-form-item label="Connect to Database">
      <el-input
        v-model="dblocation"
        class="dblocation"
        placeholder="/home/user/mydatabase.sqlite"
        :disabled="isDBConnected"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="connectDB" :disabled="isDBConnected"
        >Connect</el-button
      >
      <el-button @click="disconnectDB" :disabled="isDBDisconnected"
        >Disconnect</el-button
      >
    </el-form-item>
  </el-form>
</template>

<script>
import api from "@/api";
import { mapMutations } from "vuex";
export default {
  name: "DBConnect",
  components: {},
  data() {
    return {
      dblocation: "",
      status: "not connected",
    };
  },
  computed: {
    isDBDisconnected() {
      return this.$store.getters.getTableCount == 0;
    },
    isDBConnected() {
      return this.$store.getters.getTableCount > 0;
    },
  },
  methods: {
    ...mapMutations(["setDatabase", "setTables", "setTableCount"]),
    disconnectDB() {
      this.setDatabase(null);
      this.setTables(null);
      this.setTableCount(0);
      this.$notify({
        title: "Success",
        message: "Disconnecting the database succeeded",
      });
    },
    connectDB() {
      api.database.connectDatabase(this.dblocation).then((response) => {
        let success = false;
        if (response.status === 200) {
          this.setDatabase(this.dblocation);
          this.setTables(response.data.tables);
          this.setTableCount(response.data.table_count);
          success = true;
          this.$notify({
            title: "Success",
            message: "Connecting to the database succeeded",
          });
        }
        if (!success) {
          this.setDatabase(null);
          this.$notify.error({
            title: "Error",
            message: "Connecting to the database failed",
          });
        }
      });
    },
    // checkDatabase() {
    //   this.axios.get("/databaseinfo").then((response) => {
    //     if (response.status === 200) {
    //       console.log("db connection info ", response.data);
    //       if (response.data.db_name) {
    //         this.dblocation = response.data.db_name;
    //         this.status = " Connected";
    //         this.tableCount = response.data.table_count;
    //       }
    //     }
    //   });
    // },
  },
  mounted() {
    this.dblocation = this.$store.getters.getDatabase;
  },
};
</script>
<style scoped>
.dblocation {
}
</style>
