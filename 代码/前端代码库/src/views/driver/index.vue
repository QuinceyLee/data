<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card" style="width: 100%; margin-top: 10px">
        <div slot="header" class="clearfix">
          <span> 驱动管理</span>
        </div>
        <el-button type="success" @click="refreshList" icon="el-icon-refresh"
          >刷 新</el-button
        >
      </el-card>

      <el-table
        :data="drivers"
        border
        title="插件管理"
        style="width: 100%; margin-top: 10px"
        height="700"
      >
        <el-table-column prop="name" label="驱动名称" width="340px"> </el-table-column>
        <el-table-column prop="type" label="驱动类型"> </el-table-column>
        <el-table-column prop="description" label="驱动信息"> </el-table-column>
      </el-table> </el-main
  ></el-container>
</template>

<script>
import { Message } from "element-ui";
import { list } from "@/api/driver";

export default {
  name: "Driver",
  created() {
    this.getList();
  },
  methods: {
    //
    //
    refreshList() {
      this.getList();
      Message({
        message: "刷新成功",
        type: "success",
        duration: 5 * 1000,
      });
    },
    getList() {
      list().then((response) => {
        this.drivers = response.data;
      });
    },
  },
  data() {
    return {
      drivers: [],
    };
  },
};
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
