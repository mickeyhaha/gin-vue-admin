<template>
  <div id="order-info">
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          工单:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          P201223067-1
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          计划产能(H):
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
          10000
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          产品:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          P23067-1_001
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          产品名称:
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
          abcdefghik
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="col_span.col1" class="block-top-bottom-content">
          制令单:
        </el-col>
        <el-col :span="col_span.col2"  class="block-top-bottom-content">
          321207938
        </el-col>
        <el-col :span="col_span.col3"  class="block-top-bottom-content">
          面别:
        </el-col>
        <el-col :span="col_span.col4"  class="block-top-bottom-content">
          AB
        </el-col>
      </el-row>
  </div>
</template>

<script>
import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'OrderInfo',
  mixins: [infoList],
  data () {
    return {
      col_span: {
        col1: 4,
        col2: 8,
        col3: 6,
        col4: 6,
      },
      listApi: getLackWarnings,
    }
  },
  methods: {
    createData () {
      this.getTableData().then(() => {
        let datas = []
        for (let index = 0; index < this.tableData.length && index < 9; index++) {
          const element = this.tableData[index]
          datas[index] = {
            '机器': element.machineCode+index,  //MatrCode
            '时间': element.leftTime+index
          }
        }
        // this.chartData = {
        //   columns: ['机器', '时间'],
        //   rows: datas
        // }
      });
      
    },
  },
  mounted () {
    const { createData } = this

    createData()

    setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#order-info {
  width: 100%;
  height: 20%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 0px;
  margin-bottom: 10px;

  .dv-conical-column-chart {
    flex: 1;
  }
}
</style>
