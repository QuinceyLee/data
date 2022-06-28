<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>输入资源管理</span>
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
          width="1100px"
          title="新建资源"
          top="10vh"
          :close-on-click-modal="false"
          :visible.sync="status.createDialogVisible"
          @closed="createDialogClosed"
        >
          <el-form
            ref="createFormRef"
            :model="createForm"
            :rules="createFormRules"
            label-width="auto"
          >
            <el-form-item label="资源名称" prop="name">
              <el-input
                v-model="createForm.name"
                placeholder="资源名称"
                maxlength="20"
                clearable
              />
            </el-form-item>
            <el-form-item label="资源类型" prop="type">
              <el-select
                v-model="createForm.type"
                placeholder="资源类型"
                filterable
                clearable
                style="width: 100%"
                @change="onSelectChange"
              >
                <el-option
                  v-for="(item, index) in protocolOptions"
                  :key="index"
                  :label="item.text"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item v-if="status.configFormVisible" label="资源配置">
              <dynamic-form
                ref="dynamicFormRef"
                :options="jsonConfig"
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
            <el-button @click="status.createDialogVisible = false"> 取 消 </el-button>
            <el-button
              type="primary"
              :loading="status.createLoading"
              @click="createInEnd"
            >
              提 交
            </el-button>
          </div>
        </el-dialog>
      </el-card>

      <el-table :data="tableData" border style="width: 100%; margin-top: 10px">
        <el-table-column prop="uuid" label="UUID" width="340px" />
        <el-table-column
          prop="type"
          label="类型"
          width="200px"
          :filters="protocolOptions"
          :filter-method="filterHandler"
          filter-placement="bottom-end"
          :column-key="'type'"
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
            <el-button size="mini" type="danger" @click="removeInEnd(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog
        title="输出目标详情"
        :visible.sync="status.detailDialogVisible"
        width="1000px"
      >
        <el-descriptions border :column="1" :label-style="LS">
          <el-descriptions-item label="type">
            <el-tag size="small">{{ inEndDetail.type }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="name" :span="2">
            {{ inEndDetail.name }}
          </el-descriptions-item>
          <el-descriptions-item
            v-for="(item, index) in inEndDetail.config"
            :key="index"
            :label="index"
          >
            {{ item }}
          </el-descriptions-item>
          <el-descriptions-item label="description">
            {{ inEndDetail.description }}
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
import { list } from '@/api/inend'
import { remove } from '@/api/inend'
import { create } from '@/api/inend'
import { Message } from 'element-ui'
import in_types from './in_type'
import DynamicForm from '@/components/DynamicForm/index.vue'
export default {
  name: 'InEnd',
  components: {
    DynamicForm
  },
  data() {
    return {
      LS: {
        width: '200px'
      },

      // 输入资源管理 新增、编辑
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

      protocolOptions: [
        { text: 'MQTT 协议接入', value: 'MQTT' },
        { text: 'HTTP 协议接入', value: 'HTTP' },
        { text: 'UDP 协议接入', value: 'RULEX_UDP' },
        { text: 'COAP 协议接入', value: 'COAP' },
        { text: 'GRPC 协议接入', value: 'GRPC' },
        { text: '通用串口接入', value: 'UART_MODULE' },
        { text: 'MODBUS MASTER 模式', value: 'MODBUS_MASTER' },
        { text: 'SNMP 协议接入', value: 'SNMP_SERVER' },
        { text: '西门子S7系列接入', value: 'SIEMENS_S7' },
        { text: 'NATS 总线接入', value: 'NATS_SERVER' },
      ],

      // 配置文件
      jsonConfig: [],
      tableData: [],
      inEndDetail: {},
      protocol: '',
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
        this.tableData = response.data
      })
    },

    // 新建资源 》资源类型改变
    onSelectChange(v) {
      if (this.isEmpty(v)) {
        this.protocol = ''
        this.jsonConfig = []
        this.status.configFormVisible = false
      } else {
        this.protocol = v
        this.jsonConfig = in_types[v]
        this.status.configFormVisible = true
      }
    },

    // 新建资源 对话框关闭
    createDialogClosed() {
      this.status.configFormVisible = false
      this.status.createLoading = false
      this.$refs.createFormRef.resetFields()
      this.createForm = this.$options.data().createForm
    },

    dynamicFormChanged(val) {
      this.createForm.config = JSON.parse(JSON.stringify(val))
    },

    createInEnd() {
      this.$refs.createFormRef.validate((valid) => {
        if (valid) {
          // 配置表单也验证通过才提交
          const flag = this.$refs.dynamicFormRef.validateForm()
          if (flag) {
            var params = JSON.parse(JSON.stringify(this.createForm))
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

    filterHandler(value, row) {
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
      this.inEndDetail = row
      this.status.detailDialogVisible = true
    },

    removeInEnd(row) {
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
<style lang="less" scoped></style>
