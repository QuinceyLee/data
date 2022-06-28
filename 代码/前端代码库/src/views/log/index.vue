<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card" style="width: 100%; margin-top: 10px">
        <div slot="header" class="clearfix">
          <span> 近期日志</span>
        </div>
        <el-button type="success" @click="refreshList" icon="el-icon-refresh"
          >刷 新</el-button
        >
      </el-card>

      <el-table
        :data="logs"
        border
        title="近期日志"
        style="width: 100%; margin-top: 10px"
        height="700"
      >
        <el-table-column prop="id" label="ID" width="70px"> </el-table-column>
        <el-table-column prop="content" label="近期日志"> </el-table-column>
      </el-table> </el-main
  ></el-container>
</template>

<script>
import { Message } from "element-ui";
import { logs } from "@/api/system";

export default {
  name: "Log",
  created() {
    this.getLogs();
  },
  methods: {
    getLogs() {
      logs().then((response) => {
        this.logs = response.data;
      });
    },
    //
    //
    refreshList() {
      this.getLogs();
      Message({
        message: "刷新成功",
        type: "success",
        duration: 5 * 1000,
      });
    },
  },
  data() {
    return {
      logs: [],
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
