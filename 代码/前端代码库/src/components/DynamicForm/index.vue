<template>
  <div class="dynamic-form">
    <el-form ref="configFormRef" :model="configForm" label-width="auto">
      <div v-for="(item, index) in options_" :key="index">
        <el-form-item
          v-if="item.type !== 'el-inline' && item.type !== 'flag'"
          :label="item.label"
          :prop="item.name"
          :rules="[
            {
              required: item.required,
              message: `请输入${item.label}`,
              trigger: 'blur',
            },
          ]"
        >
          <el-input
            v-if="item.type == 'el-input-text'"
            v-model.trim="configForm[item.name]"
            type="text"
            size="small"
            :placeholder="item.placeholder"
            clearable
            maxlength="100"
            @change="changed($event, item.name, '')"
          />
          <el-input-number
            v-if="item.type == 'el-input-number'"
            v-model.trim="configForm[item.name]"
            size="small"
            style="width: 100%"
            controls-position="right"
            :placeholder="item.placeholder"
            @change="changed($event, item.name, '')"
          />
          <el-select
            v-if="item.type === 'el-select'"
            v-model="configForm[item.name]"
            style="width: 100%"
            clearable
            size="small"
            :placeholder="item.placeholder"
            @change="changed($event, item.name, '')"
          >
            <el-option
              v-for="itemChild in item.selectOptions"
              :key="itemChild.name"
              :label="itemChild.label"
              :value="itemChild.value"
            />
          </el-select>
        </el-form-item>

        <el-row v-if="item.type === 'el-inline'" :key="item.name" :gutter="10">
          <el-col
            v-for="childrenItem in item.children"
            :key="childrenItem.name"
            :span="calcSpan(item.children)"
          >
            <el-form-item
              :label="childrenItem.label"
              :prop="item.name + '.' + childrenItem.name"
              :rules="[
                {
                  required: childrenItem.required,
                  message: `请输入${childrenItem.label}`,
                  trigger: 'blur',
                },
              ]"
            >
              <el-input
                v-if="childrenItem.type == 'el-input-text'"
                v-model="configForm[item.name][childrenItem.name]"
                size="small"
                clearable
                :placeholder="childrenItem.placeholder"
                @change="changed($event, item.name, childrenItem.name)"
              />
              <el-input-number
                v-if="childrenItem.type == 'el-input-number'"
                v-model="configForm[item.name][childrenItem.name]"
                size="small"
                style="width: 100%"
                controls-position="right"
                clearable
                :placeholder="childrenItem.placeholder"
                @change="changed($event, item.name, childrenItem.name)"
              />
              <el-select
                v-if="childrenItem.type == 'el-select'"
                v-model="configForm[item.name][childrenItem.name]"
                style="width: 100%"
                clearable
                size="small"
                :placeholder="childrenItem.placeholder"
                @change="changed($event, item.name, childrenItem.name)"
              >
                <el-option
                  v-for="itemSon in childrenItem.selectOptions"
                  :key="itemSon.value"
                  :label="itemSon.label"
                  :value="itemSon.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <dynamic-params
          v-if="item.type == 'dynamicParams'"
          ref="dynamicParamsRef"
          v-model="configForm[item.name]"
          :dynparams="item"
          :name="item.name"
          @childChanged="(key, val) => dynamicParamsChanged(key, val)"
        />
      </div>
    </el-form>
  </div>
</template>

<script>
import in_types from './../../views/inend/in_type'
import dynamicParams from '@/components/DynamicParams/index.vue'
import { isEmpty, deepClone } from '@/utils/index'
export default {
  name: 'DynamicForm',
  components: {
    dynamicParams
  },
  props: {
    options: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      configForm: {},
      options_: [],
      tcpConfig: [],
      rtuConfig: []
    }
  },

  watch: {
    options() {
      this.options_ = this.options
    },
    options_() {
      this.setDefaultValue()
    }
  },

  created() {
    this.options_ = this.options
    this.initModbus()
  },

  methods: {
    // 赋初值
    setDefaultValue() {
      this.configForm = {}
      if (this.options_ && this.options_.length > 0) {
        this.options_.forEach((item1, index1) => {
          const obj1 = item1
          if (item1.children) {
            this.configForm[item1.name] = {}
            item1.children.forEach((item2, index2) => {
              this.configForm[item1.name][item2.name] = item2.value
            })
          } else {
            if (Object.prototype.toString.call(obj1) === '[object Object]') {
              this.configForm[item1.name] = deepClone(item1.value)
            }
          }
        })
        this.configForm = Object.assign({}, this.configForm)
      }
    },

    initModbus() {
      this.tcpConfig = in_types['TCP']
      this.rtuConfig = in_types['RTU']
    },

    // 校验数据
    validateForm() {
      let flag = null
      this.$refs.configFormRef.validate((valid) => {
        if (valid) {
          this.$emit('onChanged', this.configForm)
          flag = true
        } else {
          flag = false
        }
      })
      return flag
    },

    dynamicParamsChanged(key, val) {
      this.configForm[key] = JSON.parse(JSON.stringify(val))
    },

    calcSpan(e) {
      let colSpan = 24
      switch (e.length) {
        case 2:
          colSpan = 12
          break
        case 3:
          colSpan = 8
          break
        case 4:
          colSpan = 6
          break
        default:
          colSpan = 24
      }
      return colSpan
    },

    changed(val, key1, key2) {
      if (val === 'TCP') {
        this.options_ = this.tcpConfig
        this.options_[0].value = 'TCP'
      } else if (val === 'RTU') {
        this.options_ = this.rtuConfig
        this.options_[0].value = 'RTU'
      }
      if (isEmpty(key2)) {
        this.configForm[key1] = val
      } else {
        this.configForm[key1][key2] = val
      }
    }
  }
}
</script>

<style lang="less" scoped>
.dynamic-form {
  .button-item {
    /deep/ .el-form-item__content {
      margin-left: 0 !important;
      text-align: center;
    }
  }
  /deep/ .el-form-item {
    margin-bottom: 15px;
  }
  /deep/ .el-form-item__label {
    text-align: left;
  }
  /deep/ .el-input-number .el-input__inner {
    text-align: left;
  }
  /deep/ .el-form-item__error {
    padding-top: 0;
  }
  /deep/ .el-form--label-top .el-form-item__label {
    padding: 0;
  }
  .delete-icon {
    text-align: center;
  }
}
</style>
