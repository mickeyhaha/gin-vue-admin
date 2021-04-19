<template>
  <div>
    <el-table
      :data="deptLineSummary"
      @selection-change="handleSelectionChange"
      @row-click="clickData"
      border
      ref="multipleTable"
      stripe
      height="250"
      style="width: 100%;"
      tooltip-effect="dark"
    >
    <el-table-column label="LineID" prop="lineID" width="120"></el-table-column> 
    <el-table-column label="AOI不良数" prop="errCount" width="120"></el-table-column> 
    <el-table-column label="AOI总数" prop="count" width="120"></el-table-column> 
    <el-table-column label="AOI不良率" prop="aoiErrCount" width="120"></el-table-column> 

    <el-table-column label="LineName字段" prop="LineName" width="120"></el-table-column> 
    <el-table-column label="MOrderNo字段" prop="MOrderNo" width="120"></el-table-column> 
    <el-table-column label="WorkOrderNo字段" prop="WorkOrderNo" width="120"></el-table-column> 
    <el-table-column label="Customer字段" prop="Customer" width="120"></el-table-column> 
    <el-table-column label="CustOrderNo字段" prop="CustOrderNo" width="120"></el-table-column> 
    <el-table-column label="MachineType字段" prop="MachineType" width="120"></el-table-column> 
    <el-table-column label="BOM编号" prop="BOMNo" width="120"></el-table-column> 
    <el-table-column label="BOM版本" prop="BOMVersion" width="120"></el-table-column> 
    <el-table-column label="产品编号" prop="Product" width="120"></el-table-column> 
    <el-table-column label="周期" prop="CycleTime" width="120"></el-table-column> 
    <el-table-column label="版面" prop="PasteSide" width="120"></el-table-column> 
    <el-table-column label="拼版数" prop="PanelCount" width="120"></el-table-column> 
    <el-table-column label="计划数量" prop="Qty" width="120"></el-table-column> 
    <el-table-column label="QtyCompleted字段" prop="QtyCompleted" width="120"></el-table-column> 
    <el-table-column label="开始时间" prop="BeginTime" width="120"></el-table-column> 
    <el-table-column label="结束时间" prop="EndTime" width="120"></el-table-column> 
    <el-table-column label="备注" prop="Remark" width="120"></el-table-column> 
    <el-table-column label="是否启动JIT拉动" prop="JITRunning" width="120"></el-table-column> 
    <el-table-column label="CreateTime字段" prop="CreateTime" width="120"></el-table-column> 
    <el-table-column label="最后时间" prop="LMTime" width="120"></el-table-column> 
    </el-table>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { store } from '@/store/index'

export default {
  name: "DeptLineSummary",
  computed: mapState(
    ['deptLineSummary']
  ),
  data() {
    return {
    };
  },
  methods: {
      clickData(val) {        
        this.$store.dispatch('selectLine', val);
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
  },
  async created() {
    await store.dispatch("getDeptLineSummary")
  }
};
</script>

<style>
</style>
