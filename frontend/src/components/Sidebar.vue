<template>
  <el-aside class="sidebar">
    <p>KwikSQLite</p>
    <el-menu
      :default-active="activeMenuItem"
      class="el-menu"
      @open="handleOpen"
      @close="handleClose"
    >
      <router-link to="/">
        <el-menu-item index="1">
          <i class="el-icon-menu"></i>
          <span>Database</span>
        </el-menu-item>
      </router-link>
      <router-link to="/query">
        <el-menu-item index="2" :disabled="activateMenu">
          <i class="el-icon-menu"></i>
          <span>Query</span>
        </el-menu-item>
      </router-link>
      <router-link to="/create">
        <el-menu-item index="3" :disabled="activateMenu">
          <i class="el-icon-menu"></i>
          <span>Create</span>
        </el-menu-item>
      </router-link>
      <el-submenu index="4" :disabled="activateMenu">
        <template #title><i class="el-icon-menu"></i>Tables</template>
        <el-menu-item
          :index="index + 1"
          v-for="(table, index) in tables"
          @click="showTable(table.Name)"
        >
          <i class="el-icon-caret-right"></i>
          <span>{{ table.Name }}</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: "HomeView",
  props: {
    menu: {
      required: true,
    },
  },
  data() {
    return {
      activeMenuItem: "1",
    };
  },
  computed: {
    activateMenu() {
      console.log(this.$store.state.dbName, "***");
      if (
        this.$store.state.dbName &&
        this.$store.state.dbName != null &&
        this.$store.state.dbName !== "" &&
        this.$store.state.dbName.length > 0
      ) {
        return false;
      } else {
        return true;
      }
    },
    tables() {
      let tables = this.$store.state.tables;
      if (tables) {
        tables.forEach((table, index) => {
          if (table.Name == this.$store.state.selectedTable) {
            this.activeMenuItem = index + 1;
          }
        });
      }
      return tables;
    },
  },
  methods: {
    showTable(tableName) {
      console.log("Select Table:", tableName);
      this.$store.commit("setTable", tableName);
      this.$router.push(`/table/${tableName}`);
    },
    tablesClick() {
      console.log("Tables Click");
      this.$router.push(`/tables`);
    },
  },
};
</script>
<style scoped>
.sidebar {
}
</style>
