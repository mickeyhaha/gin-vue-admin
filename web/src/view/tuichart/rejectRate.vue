<template>
  <div id="reject-rate">
    <div class="reject-rate-title">抛料率</div>
    <bar-chart :data="config" />
    <!-- <dv-conical-column-chart :config="config" style="width:400px;height:200px;"/> -->
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { barChart, lineChart } from '@toast-ui/vue-chart';

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'RejectRate',
  mixins: [infoList],
  components: {
    'bar-chart': barChart,
  },
  data () {
    return {
      listApi: getLackWarnings,
      config: {
        categories: ['July', 'Aug', 'Sep', 'Oct', 'Nov'],
        series: [
          {
            name: 'Budget',
            data: [3000, 5000, 7000, 6000, 4000],
          },
          {
            name: 'Income',
            data: [1000, 7000, 2000, 5000, 3000],
          },
        ],
      }
    }
  },
  methods: {
    createData () {
      this.getTableData().then(res => {
        let datas = []
        for (let index = 0; index < this.tableData.length && index < 9; index++) {
          const element = this.tableData[index]
          datas[index] = {
            name: element.machineCode+index,  //MatrCode
            value: element.leftTime+index
          }
        }
        // this.config = {
        //   data: datas,
        //   rowNum: 9
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
#reject-rate {
  width: 100%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 30px;

  .reject-rate-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
  }

  .dv-conical-column-chart {
    flex: 1;
  }
}
</style>
