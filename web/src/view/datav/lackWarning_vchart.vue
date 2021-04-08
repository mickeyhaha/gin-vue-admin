<template>
  <div id="reject-rate">
    <div class="reject-rate-title">抛料预警</div>
    <ve-bar :data="chartData" width="200px" height="200px"></ve-bar>
  </div>
</template>

<script>
import VeBar from 'v-charts/lib/bar.common'

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'RejectRate',
  mixins: [infoList],
  components: {
    VeBar,
  },
  data () {
    return {
      listApi: getLackWarnings,
      chartData: {
      },
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

    //setInterval(createData, 3000)
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
