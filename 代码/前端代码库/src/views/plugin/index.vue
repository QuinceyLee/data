<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card" style="width: 100%; margin-top: 10px">
        <div slot="header" class="clearfix">
          <span> 插件管理</span>
        </div>

        <el-button type="primary" icon="el-icon-help" @click="importPlugin"
          >导入插件</el-button
        >
        <el-button type="success" @click="refreshList" icon="el-icon-refresh"
          >刷 新</el-button
        >
      </el-card>

      <el-table
        :data="plugins"
        border
        title="插件管理"
        style="width: 100%; margin-top: 10px"
        height="700"
      >
        <el-table-column prop="name" label="插件名称" width="340px"> </el-table-column>
        <el-table-column prop="version" label="版本"> </el-table-column>
        <el-table-column prop="homepage" label="主页"> </el-table-column>
        <el-table-column prop="helpLink" label="文档地址"> </el-table-column>
        <el-table-column prop="author" label="作者"> </el-table-column>
        <el-table-column prop="email" label="邮箱"> </el-table-column>
        <el-table-column prop="license" label="协议"> </el-table-column>
      </el-table> </el-main
  ></el-container>
</template>

<script>
import { Message } from "element-ui";
import { plugins } from "@/api/plugin";

export default {
  name: "Plugin",
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
      plugins().then((response) => {
        this.plugins = response.data;
      });
    },
    //
    //
    importPlugin() {
      Message({
        message: "当前版本暂不支持导入插件, 请关注后续版本",
        type: "warning",
        duration: 5 * 1000,
      });
    },
  },
  data() {
    return {
      plugins: [],
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
