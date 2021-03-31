<template>
  <div id="lack-warning">
    <div class="lack-warning-title">缺料预警</div>
    <dv-capsule-chart :config="config" />
  </div>
</template>

<script>
import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'LackWarning',
  mixins: [infoList],
  data () {
    return {
      listApi: getLackWarnings,
      config: {
        data: [
        ],
        rowNum: 9
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
            value: element.leftTime
          }
        }
        this.config = {
          data: datas,
          rowNum: 9
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
#lack-warning {
  width: 100%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 30px;

  .lack-warning-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
  }

  .dv-capsule-chart {
    flex: 1;
  }
}
</style>
