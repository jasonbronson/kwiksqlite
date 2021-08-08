<template>
  <el-aside class="sidebar">
    <p>KwikSQLite</p>
    <el-menu
      default-active="1"
      class="el-menu"
      @open="handleOpen"
      @close="handleClose"
    >
      <el-menu-item index="1">
        <i class="el-icon-menu"></i>
        <span>Database</span>
      </el-menu-item>
      <el-menu-item index="2" disabled>
        <i class="el-icon-menu"></i>
        <span>Query</span>
      </el-menu-item>
      <el-menu-item index="3" disabled>
        <i class="el-icon-menu"></i>
        <span>Create</span>
      </el-menu-item>
      <el-submenu index="4" disabled>
        <template #title><i class="el-icon-menu"></i>Tables</template>
        <el-menu-item-group>
          <template #title>Group 1</template>
          <el-menu-item index="1-1">Option 1</el-menu-item>
          <el-menu-item index="1-2">Option 2</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="Group 2">
          <el-menu-item index="1-3">Option 3</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="1-4">
          <template #title>Option4</template>
          <el-menu-item index="1-4-1">Option 4-1</el-menu-item>
        </el-submenu>
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
    return {};
  },
};
</script>
<style scoped>
.sidebar {
}
</style>
