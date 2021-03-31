<template>
  <div id="reject-rate">
    <div class="title">抛料率</div>
    <ve-histogram :data="chartData" width="200px" height="200px"></ve-histogram>
    <!-- <bar-chart :data="chartData" :style="chartData.containerStyle" /> -->
    <!-- <dv-conical-column-chart :config="chartData" style="width:400px;height:200px;"/> -->
  </div>
</template>

<script>
// import 'tui-chart/dist/tui-chart.css'
// import { barChart, lineChart } from '@toast-ui/vue-chart'
import VeHistogram from 'v-charts/lib/histogram.common'

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'RejectRate',
  mixins: [infoList],
  components: {
    VeHistogram,
    // 'bar-chart': barChart
  },
  data () {
    return {
      listApi: getLackWarnings,
      chartData: {
      },
      // chartData: {
      //   categories: ['July', 'Aug', 'Sep', 'Oct', 'Nov'],
      //   series: [
      //     {
      //       name: 'Budget',
      //       data: [3000, 5000, 7000, 6000, 4000],
      //     },
      //     {
      //       name: 'Income',
      //       data: [1000, 7000, 2000, 5000, 3000],
      //     }
      //   ],
      //   containerStyle: {
      //     width: '600px',
      //     height: '700px',
      //   }
      //}
    }
  },
  methods: {
    createData () {
      this.getTableData().then(res => {
        let datas = []
        for (let index = 0; index < this.tableData.length && index < 9; index++) {
          const element = this.tableData[index]
          datas[index] = {
            '机器': element.machineCode+index,  //MatrCode
            '时间': element.leftTime+index
          }
        }
        this.chartData = {
          columns: ['机器', '时间'],
          rows: datas
        }
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
#reject-rate {
  width: 100%;
  height: 22%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 30px;
  margin-bottom: 10px;

  .dv-conical-column-chart {
    flex: 1;
  }
}
</style>
