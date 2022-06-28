<template>
  <el-container style="height: 800px; border: 1px solid #eee">
    <el-main>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>规则管理</span>
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
      </el-card>

      <!-- table -->
      <el-table
        :data="rulesTableData"
        border
        style="width: 100%; margin-top: 10px"
      >
        <el-table-column
          prop="uuid"
          label="UUID"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          prop="status"
          label="状态"
          width="150px"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag v-if="scope.row.status === 1" type="success">运行中</el-tag>
            <el-tag v-else type="danger">已停止</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="name"
          label="名称"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          prop="description"
          label="备注信息"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column fixed="right" label="操作" width="300" align="center">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="primary"
              @click="openDetailsDialog(scope.row)"
            >
              详情
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="removeRule(scope.row.uuid)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- detail-dialog -->
      <el-dialog
        title="规则详情"
        :visible.sync="status.detailDialogVisible"
        width="1000px"
        top="2%"
      >
        <el-form
          ref="ruleDetailFormRef"
          :model="ruleDetailForm"
          status-icon
          label-width="auto"
        >
          <el-form-item label="规则名称" prop="name">
            <el-input
              v-model="ruleDetailForm.name"
              placeholder="规则名称"
              :disabled="true"
            ></el-input>
          </el-form-item>
          <el-form-item label="输入资源" prop="from">
            <el-select
              v-model="ruleDetailForm.from"
              multiple
              disabled
              placeholder="请选择数据源"
              style="width: 100%"
            >
              <el-option
                v-for="item in inEnds"
                :key="item.uuid"
                :label="item.name"
                :value="item.uuid"
              >
                <span style="float: left">{{ item.uuid }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">
                  {{ item.name }}
                </span>
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="规则回调" prop="config">
            <codemirror
              v-model="ruleDetailForm.actions"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="成功回调" prop="config">
            <codemirror
              v-model="ruleDetailForm.success"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="失败回调" prop="config">
            <codemirror
              v-model="ruleDetailForm.failed"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="备注信息" prop="description">
            <el-input
              v-model="ruleDetailForm.description"
              autocomplete="off"
              :disabled="true"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="status.detailDialogVisible = false">
            确 定
          </el-button>
        </div>
      </el-dialog>

      <!-- 新建规则dialog -->
      <el-dialog
        title="新建规则"
        :visible.sync="status.createDialogVisible"
        width="1000px"
        top="2%"
        @closed="createDialogClosed"
      >
        <el-form
          ref="createFormRef"
          :model="createForm"
          :rules="createFormRules"
          status-icon
          label-width="auto"
        >
          <el-form-item label="规则名称" prop="name">
            <el-input
              v-model="createForm.name"
              placeholder="规则名称"
            ></el-input>
          </el-form-item>
          <el-form-item label="输入资源" prop="from">
            <el-select
              v-model="createForm.from"
              multiple
              placeholder="请选择数据源"
              style="width: 100%"
            >
              <el-option
                v-for="item in inEnds"
                :key="item.uuid"
                :label="item.name"
                :value="item.uuid"
              >
                <span style="float: left">{{ item.uuid }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">
                  {{ item.name }}
                </span>
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="规则回调" prop="config">
            <codemirror
              v-model="createForm.actions"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="成功回调" prop="config">
            <codemirror
              v-model="createForm.success"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="失败回调" prop="config">
            <codemirror
              v-model="createForm.failed"
              :options="options"
              style="line-height: 1"
            />
          </el-form-item>
          <el-form-item label="备注信息" prop="description">
            <el-input
              v-model="createForm.description"
              autocomplete="off"
            ></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="status.createDialogVisible = false">
            取 消
          </el-button>
          <el-button type="warning" @click="validateRule">语法验证</el-button>
          <el-button
            type="primary"
            :loading="status.addRuleLoading"
            @click="createRule"
          >
            提 交
          </el-button>
        </div>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import Vue from "vue";
import { codemirror } from "vue-codemirror";
import "codemirror/lib/codemirror.css";
import "codemirror/theme/darcula.css";
import "codemirror/mode/lua/lua.js";
import VJsonEditor from "v-jsoneditor/src/index";
Vue.use(VJsonEditor);
import { Message } from "element-ui";
import { list as inEndList } from "@/api/inend";
import { list } from "@/api/rule";
import { remove } from "@/api/rule";
import { create } from "@/api/rule";
import { validateRule } from "@/api/rule";

export default {
  name: "Rule",
  data() {
    return {
      options: {
        tabSize: 4,
        mode: "text/x-lua",
        theme: "darcula",
        lineNumbers: true,
        line: true,
        lint: true,
        styleActiveLine: true,
        lineNumbers: true,
        foldGutter: true,
        styleSelectedText: true,
        matchBrackets: true,
        showCursorWhenSelecting: true,
        gutters: ["CodeMirror-lint-markers"],
        hintOptions: {
          completeSingle: true,
        },
        extraKeys: {
          Tab: function (cm) {
            var spaces = Array(cm.getOption("indentUnit") + 1).join(" ");
            cm.replaceSelection(spaces);
          },
        },
        showHint: true,
      },

      detailOptions: {
        modes: ["code"],
        mode: "code",
        ace: ace,
      },

      inEnds: [],
      ruleDetailForm: {},
      rulesTableData: [],

      createForm: {
        name: "",
        description: "",
        from: [],
        actions:
          "Actions = {\n\tfunction(data)\n\t\t\n\t\treturn true, data\n\tend\n}",
        failed:
          "function Failed(error)\n\t\nend\n",
        success:
          'function Success()\n\t\nend\n',
      },
      createFormRules: {
        name: [
          { required: true, message: "输入名称", trigger: ["blur", "change"] },
        ],
        description: [
          { required: true, message: "输入描述", trigger: ["blur", "change"] },
        ],
        from: [
          {
            required: true,
            message: "输入资源入口",
            trigger: ["blur", "change"],
          },
        ],
        actions: [
          {
            required: true,
            message: "输入规则回调",
            trigger: ["blur", "change"],
          },
        ],
        failed: [
          {
            required: true,
            message: "输入成功回调",
            trigger: ["blur", "change"],
          },
        ],
        success: [
          {
            required: true,
            message: "输入失败回调",
            trigger: ["blur", "change"],
          },
        ],
      },

      status: {
        addRuleLoading: false,
        detailDialogVisible: false,
        createDialogVisible: false,
      },
    };
  },

  components: {
    codemirror,
    VJsonEditor,
  },

  created() {
    this.getRuleList();
    this.getInEndList();
  },

  methods: {
    // 删除规则
    removeRule(_uuid) {
      this.$msgbox({
        title: "提示",
        message: "确定要删除该规则吗",
        showCancelButton: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        beforeClose: (action, instance, done) => {
          if (action === "confirm") {
            instance.confirmButtonLoading = true;
            remove(_uuid)
              .then((response) => {
                this.$message.success(response.msg);
                this.getRuleList();
                done();
                setTimeout(() => {
                  instance.confirmButtonLoading = false;
                }, 300);
              })
              .catch(() => {
                instance.confirmButtonLoading = false;
              });
          } else {
            done();
          }
        },
      });
    },

    // 打开详情弹框
    openDetailsDialog(row) {
      this.ruleDetailForm = row;
      this.status.detailDialogVisible = true;
    },

    // 刷新
    refreshList() {
      this.getRuleList();
      this.getInEndList();
    },

    // 新增规则
    createRule() {
      this.$refs.createFormRef.validate((valid) => {
        if (valid) {
          const params = { ...this.createForm };
          this.status.addRuleLoading = true;
          create(params)
            .then(() => {
              this.$message.success("创建成功");
              this.status.createDialogVisible = false;
              this.getRuleList();
            })
            .catch(() => {
              this.status.addRuleLoading = false;
            });
        }
      });
    },

    getRuleList() {
      list().then((response) => {
        this.rulesTableData = response.data;
      });
    },

    getInEndList() {
      inEndList().then((response) => {
        this.inEnds = response.data;
      });
    },

    validateRule() {
      validateRule(this.createForm).then((_response) => {
        this.$message.success("语法验证成功");
      });
    },

    // 新建规则弹框关闭
    createDialogClosed() {
      this.status.addRuleLoading = false;
      this.$refs.createFormRef.resetFields();
      this.createForm = this.$options.data().createForm;
    },
  },
};
</script>
