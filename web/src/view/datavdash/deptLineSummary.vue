<template>
  <div id="dept-line-summary">
    <el-table
      :data="deptLineSummary"
      @selection-change="handleSelectionChange"
      @row-click="clickData"
      ref="multipleTable"
      tooltip-effect="dark"
    >
    <!-- <el-table-column label="LineID" prop="lineID" ></el-table-column>  -->
    <el-table-column label="线体名称" prop="LineName" ></el-table-column> 
    <el-table-column label="工单" prop="WorkOrderNo" ></el-table-column> 
    <el-table-column label="制令单" prop="MOrderNo" ></el-table-column> 
    <el-table-column label="计划产量" prop="Qty" ></el-table-column> 
    <el-table-column label="已完成" prop="QtyCompleted" ></el-table-column> 
    <el-table-column label="AOI不良数" prop="errCount" ></el-table-column> 
    <el-table-column label="AOI总数" prop="count" ></el-table-column> 
    <el-table-column label="AOI不良PPM" prop="aoiErrCount" ></el-table-column> 
    <el-table-column label="平衡率%" prop="balance" ></el-table-column> 

    <!-- <el-table-column label="Customer" prop="Customer" ></el-table-column> 
    <el-table-column label="CustOrderNo" prop="CustOrderNo" ></el-table-column> 
    <el-table-column label="MachineType" prop="MachineType" ></el-table-column> 
    <el-table-column label="BOM编号" prop="BOMNo" ></el-table-column> 
    <el-table-column label="BOM版本" prop="BOMVersion" ></el-table-column> 
    <el-table-column label="产品编号" prop="Product" ></el-table-column> 
    <el-table-column label="周期" prop="CycleTime" ></el-table-column> 
    <el-table-column label="版面" prop="PasteSide" ></el-table-column> 
    <el-table-column label="拼版数" prop="PanelCount" ></el-table-column>  -->
    <el-table-column label="开始时间">
        <template slot-scope="scope">{{ scope.row.BeginTime|formatDate }}</template>
    </el-table-column> 
    <el-table-column label="结束时间">
        <template slot-scope="scope">{{ scope.row.EndTime|formatDate }}</template>
    </el-table-column> 
    <!-- <el-table-column label="备注" prop="Remark" ></el-table-column> 
    <el-table-column label="是否启动JIT拉动" prop="JITRunning" ></el-table-column> 
    <el-table-column label="CreateTimeorder" prop="CreateTime" ></el-table-column> 
    <el-table-column label="最后时间" prop="LMTime" ></el-table-column>  -->
    </el-table>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { store } from '@/store/index'
import { formatTimeToStr } from "@/utils/date";

export default {
  name: "DeptLineSummary",
  computed: mapState(
    ['deptLineSummary']
  ),
  data() {
    return {
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
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

<style lang="less">
#dept-line-summary {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  padding-left: 20px;
  padding-right: 20px;

  .el-table__header tr,
  .el-table__header th {
    //此处设置表头样式
    padding: 0;
    height: 60px;
    line-height: 40px;
    background-color: #519ef0;
    color: #fff;
  }
  
  .el-table__body tr,
  .el-table__body td {
    //单元格样式
    padding: 0;
    height: 50px;
    line-height: 40px;
    color: #fff;
    background-color: #12396a;
  }
}
</style>
