<template>
  <div id="dept-filter">
    <el-row>
      <el-form ref="elForm" :model="formData" :rules="rules" size="small" label-width="100px">
        <el-col :span="4">
          <el-form-item label="线体" prop="LineName" class="item">
            <el-select v-model="formData.LineName" placeholder="请选择线体" @change="lineChanged"
             clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in lines" :key="index" :label="item.lineName"
                :value="item.lineName" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="日期" prop="date" class="item">
            <el-date-picker type="daterange" v-model="formData.date" format="yyyy-MM-dd"
              value-format="yyyy-MM-dd" :style="{width: '100%'}" start-placeholder="开始日期"
              end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="4">
          <el-form-item label="班别" prop="Shift" class="item">
            <el-select v-model="formData.Shift" placeholder="请选择班别" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in shiftOption" :key="index" :label="item.label"
                :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <!-- <el-col :span="4">
          <el-form-item label="工单" prop="order" class="item">
            <el-select v-model="formData.WorkOrderNo" placeholder="请选择工单" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in orderOptions" :key="index" :label="item.WorkOrderNo"
                :value="item.WorkOrderNo" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col> -->
        <el-col :span="10">
          <el-form-item size="small">
            <el-button type="primary" @click="submitForm">查询</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
  </div>
</template>
<script>
import { mapState } from 'vuex';
import { store } from '@/store/index'
import { getWorkOrderListByLine } from "@/api/PVS_Base_Line";

export default {
  name: 'DeptFilter',
  components: {},
  props: [],
  data() {
    return {
      formData: {
        date: null,
        LineName: undefined,
        Shift: undefined,
        WorkOrderNo: undefined,
      },
      rules: {
        date: [{
          required: true,
          message: '日期不能为空',
          trigger: 'change'
        }],
        LineName: [{
          required: true,
          message: '请选择线体',
          trigger: 'change'
        }],
        Shift: [{
          required: false,
          message: '请选择班别',
          trigger: 'change'
        }],
        WorkOrderNo: [{
          required: false,
          message: '请选择工单',
          trigger: 'change'
        }],
      },
      shiftOption: [{
        "label": "白班",
        "value": 1
      }, {
        "label": "晚班",
        "value": 2
      }],
      orderOptions: [],
    }
  },
  computed: mapState(
    ['lines']
  ),
  async created() {
    await store.dispatch("getLines")
  },
  watch: {},
  mounted() {},
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        this.$store.dispatch('submitDeptFilter', this.formData);
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    async lineChanged(val) {
      const res = await getWorkOrderListByLine({ page: 1, pageSize: 100, LineName: val })
      if (res.code == 0) {
        this.orderOptions = res.data.list
      }
    },
  },
}

</script>

<style lang="less">
#dept-filter {
  width: 100%;
  height: 30px;
  // background-color: #030409;
  display: flex;
  // padding: 0px 10px 10px 10px;    // top, right, bottom, left; top/bottom, right/left
  margin-bottom: 10px;

  .item .el-form-item__label {
    color: #fff;
  }
}
</style>
