<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span> 输出目标管理</span>
        </div>
        <el-button
          type="primary"
          icon="el-icon-help"
          @click="status.createDialogVisible = true"
        >
          新 建
        </el-button>
        <el-button type="success" icon="el-icon-refresh" @click="refreshList">
          刷 新
        </el-button>

        <el-dialog
          top="10vh"
          width="1100"
          title="新建目标"
          :show-close="false"
          :close-on-click-modal="false"
          :visible.sync="status.createDialogVisible"
          @closed="createDialogClosed"
        >
          <el-form
            ref="createFormRef"
            :model="createForm"
            :rules="createFormRules"
            label-width="78px"
          >
            <el-form-item label="目标名称" prop="name">
              <el-input
                v-model="createForm.name"
                placeholder="输出目标名称"
                maxlength="20"
                clearable
              />
            </el-form-item>
            <el-form-item label="目标类型" prop="type">
              <el-select
                v-model="createForm.type"
                placeholder="输出目标类型"
                filterable
                clearable
                style="width: 100%"
                @change="onSelectChange"
              >
                <el-option label="MQTT 桥接" value="MQTT" />
                <el-option label="MQTT 遥测" value="MQTT_TELEMETRY_TARGET" />
                <el-option label="MONGO 单节点" value="MONGO_SINGLE" />
                <el-option label="TdEngine" value="TDENGINE" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="status.configFormVisible" label="目标配置">
              <dynamic-form
                ref="dynamicFormRef"
                :options="jsonConfig"
                :protocol="protocol"
                @onChanged="(val) => dynamicFormChanged(val)"
              />
            </el-form-item>
            <el-form-item label="备注信息" prop="description">
              <el-input
                v-model="createForm.description"
                autocomplete="off"
                placeholder="备注信息"
                maxlength="50"
                clearable
              />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="status.createDialogVisible = false">
              取 消
            </el-button>
            <el-button
              type="primary"
              :loading="status.createLoading"
              @click="createOutEnd"
            >
              提 交
            </el-button>
          </div>
        </el-dialog>
      </el-card>

      <el-table
        :data="outendsTableData"
        border
        style="width: 100%; margin-top: 10px"
      >
        <el-table-column prop="uuid" label="UUID" width="360px" />
        <el-table-column
          prop="type"
          label="类型"
          width="240px"
          :filters="[
            { text: 'MQTT 协议桥接', value: 'MQTT' },
            { text: 'MQTT 遥测', value: 'MQTT_TELEMETRY_TARGET' },
            { text: 'MongoDb 存储', value: 'MONGO_SINGLE' },
            { text: 'TdEngine', value: 'TDENGINE' },
          ]"
          :filter-method="filterType"
        />
        <el-table-column prop="name" label="名称" />
        <el-table-column prop="description" label="信息" />
        <el-table-column
          prop="state"
          label="状态"
          :filters="[
            { text: '正常', value: 1 },
            { text: '异常', value: 0 },
          ]"
          :filter-method="filterState"
          width="80px"
        >
          <template slot-scope="scope">
            <el-tag v-if="scope.row.state === 1" type="success">正常</el-tag>
            <el-tag v-else type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="160">
          <template slot-scope="scope">
            <el-button size="mini" type="primary" @click="details(scope.row)">
              详情
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="removeOutEnd(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog
        title="输出目标详情"
        :visible.sync="status.detailDialogVisible"
        width="800px"
      >
        <el-descriptions border :column="1" :label-style="LS">
          <el-descriptions-item label="type">
            <el-tag size="small">{{ outEndDetail.type }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="name">
            {{ outEndDetail.name }}
          </el-descriptions-item>
          <el-descriptions-item
            v-for="(item, index) in outEndDetail.config"
            :key="index"
            :label="index"
          >
            {{ item }}
          </el-descriptions-item>
          <el-descriptions-item label="description">
            {{ outEndDetail.description }}
          </el-descriptions-item>
        </el-descriptions>
        <span slot="footer">
          <el-button type="primary" @click="status.detailDialogVisible = false">
            确 定
          </el-button>
        </span>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import { list } from '@/api/outend'
import { remove } from '@/api/outend'
import { create } from '@/api/outend'
import { Message } from 'element-ui'
import out_types from './out_type'
import DynamicForm from '@/components/DynamicForm/index.vue'

export default {
  name: 'OutEnd',
  components: {
    DynamicForm
  },
  data() {
    return {
      LS: {
        width: '200px'
      },

      createForm: {
        type: '',
        name: '',
        description: '',
        config: {}
      },
      createFormRules: {
        type: [
          { required: true, message: '输入类型', trigger: ['blur', 'change'] }
        ],
        name: [
          { required: true, message: '输入名称', trigger: ['blur', 'change'] }
        ],
        description: [
          { required: true, message: '输入描述', trigger: ['blur', 'change'] }
        ]
      },

      protocol: '',
      outEndDetail: {},
      outendsTableData: [],

      status: {
        createLoading: false,
        createDialogVisible: false,
        detailDialogVisible: false,
        configFormVisible: false
      }
    }
  },

  created() {
    this.getList()
  },

  methods: {
    getList() {
      list().then((response) => {
        this.outendsTableData = response.data
      })
    },

    onSelectChange(v) {
      this.protocol = v
      this.jsonConfig = out_types[v]
      this.status.configFormVisible = true
    },

    // 新建输出目标 对话框关闭
    createDialogClosed() {
      this.status.configFormVisible = false
      this.status.createLoading = false
      this.$refs.createFormRef.resetFields()
      this.createForm = this.$options.data().createForm
    },

    dynamicFormChanged(val) {
      this.createForm.config = JSON.parse(JSON.stringify(val))
    },

    createOutEnd() {
      this.$refs.createFormRef.validate((valid) => {
        if (valid) {
          // 配置表单也验证通过才提交
          const flag = this.$refs.dynamicFormRef.validateForm()
          if (flag) {
            var params = JSON.parse(JSON.stringify(this.createForm))
            delete params.config.moreRegisterParams
            this.status.createLoading = true
            create(params)
              .then((_response) => {
                Message({
                  message: '创建成功',
                  type: 'success',
                  duration: 5 * 1000
                })
                this.getList()
                this.status.createDialogVisible = false
              })
              .catch(() => {
                this.status.createLoading = false
              })
            return true
          } else {
            return false
          }
        } else {
          return false
        }
      })
    },

    filterType(value, row) {
      return row.type === value
    },

    filterState(value, row) {
      return row.state === value
    },

    refreshList() {
      this.getList()
      Message({
        message: '刷新成功',
        type: 'success',
        duration: 5 * 1000
      })
    },

    details(row) {
      this.outEndDetail = row
      this.status.detailDialogVisible = true
    },

    removeOutEnd(row) {
      remove(row.uuid).then((response) => {
        Message({
          message: response.msg,
          type: 'success',
          duration: 5 * 1000
        })
        this.getList()
      })
    },

    isEmpty(obj) {
      if (typeof obj === 'number' && !isNaN(obj)) {
        return false
      }
      if (!obj) {
        return true
      }
      return Object.keys(obj).length < 1
    }
  }
}
</script>
