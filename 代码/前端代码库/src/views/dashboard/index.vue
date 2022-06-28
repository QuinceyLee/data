<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-row :gutter="24" style="width: 100%; margin-top: 10px">
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>当前版本</span>
              </div>
              <h2>{{ systemInfo.version }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>操作系统</span>
              </div>
              <h2>{{ systemInfo.osArch }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>CPU使用</span>
              </div>
              <el-progress
                :percentage="systemInfo.system > 100 ? 100 : systemInfo.system"
                :color="customColors"
                :stroke-width="26"
                style="margin-top: 16px"
              />
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>磁盘使用</span>
              </div>
              <div style="text-align: center">
                <el-progress
                  :width="64"
                  :stroke-width="6"
                  type="circle"
                  :percentage="systemInfo.diskInfo"
                  :color="customColors"
                />
              </div>
            </el-card>
          </div>
        </el-col>
      </el-row>
      <!-- ---------------资源统计------------------ -->
      <el-row :gutter="24" style="width: 100%; margin-top: 10px">
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>入口总数</span>
              </div>
              <h2>{{ resourceCount.inends }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>规则总数</span>
              </div>
              <h2>{{ resourceCount.rules }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>出口总数</span>
              </div>
              <h2>{{ resourceCount.outends }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>插件总数</span>
              </div>
              <h2>{{ resourceCount.plugins }}</h2>
            </el-card>
          </div>
        </el-col>
      </el-row>
      <!-- ------------操作统计---------------- -->

      <el-row :gutter="24" style="width: 100%; margin-top: 10px">
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>输入成功</span>
              </div>
              <h2>{{ statistics.inSuccess }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>输入失败</span>
              </div>
              <h2>{{ statistics.inFailed }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>输出成功</span>
              </div>
              <h2>{{ statistics.outSuccess }}</h2>
            </el-card>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-card shadow="always" class="box-card" style="height: 160px">
              <div slot="header" class="clearfix">
                <span>输出失败</span>
              </div>
              <h2>{{ statistics.outFailed }}</h2>
            </el-card>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="24" style="width: 100%; margin-top: 10px">
        <el-card class="box-card" style="width: 100%; margin-top: 10px">
          <div slot="header" class="clearfix">
            <span>近期日志</span>
            <el-button style="float: right; padding: 3px 0" type="text">
              清空日志
            </el-button>
          </div>
          <el-table
            :data="logs"
            border
            style="width: 100%; margin-top: 10px"
            height="500"
          >
            <el-table-column prop="content" />
          </el-table>
        </el-card>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import { system } from '@/api/system'
import { statistics } from '@/api/system'
import { resourceCount } from '@/api/system'
import { logs } from '@/api/system'

export default {
  name: 'Dashboard',
  data() {
    return {
      resourceCount: {
        inends: 0,
        outends: 0,
        rules: 0,
        plugins: 0
      },
      statistics: {
        inFailed: 0,
        inSuccess: 0,
        outFailed: 0,
        outSuccess: 0
      },
      systemInfo: {
        alloc: 0,
        cpuPercent: [],
        diskInfo: 0,
        osArch: '',
        system: 0,
        total: 0,
        version: ''
      },
      logs: [],
      // 进度条颜色
      customColors: [
        { color: '#58D68D', percentage: 20 },
        { color: '#27AE60', percentage: 40 },
        { color: '#D4AC0D', percentage: 60 },
        { color: '#E59866', percentage: 80 },
        { color: '#D35400', percentage: 100 }
      ]
    }
  },
  created() {
    const thiz = this
    thiz.getSystem()
    thiz.getStatistics()
    thiz.getResourceCount()
    thiz.getLogs()
    setInterval(() => {
      thiz.getSystem()
      thiz.getStatistics()
      thiz.getResourceCount()
      thiz.getLogs()
    }, 3000)
  },
  destroyed() {
    for (let i = 0; i <= 3; i++) {
      clearInterval(i)
    }
  },
  methods: {
    getSystem() {
      system().then((response) => {
        this.systemInfo = response.data
      })
    },
    getStatistics() {
      statistics().then((response) => {
        this.statistics = response.data
      })
    },
    getResourceCount() {
      resourceCount().then((response) => {
        this.resourceCount = response.data
      })
    },
    getLogs() {
      logs().then((response) => {
        this.logs = response.data
      })
    }
  }
}
</script>
