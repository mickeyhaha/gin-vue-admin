<template>
  <div id="ranking-board">
    <div class="ranking-board-title">缺料预警</div>
    <dv-scroll-ranking-board :config="config" />
  </div>
</template>

<script>
import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'RankingBoard',
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
      this.getTableData().then(() => {
        let datas = []
        for (let index = 0; index < this.tableData.length && index < 9; index++) {
          const element = this.tableData[index]
          datas[index] = {
            name: element.machineCode,  //MatrCode
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

    //setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#ranking-board {
  width: 100%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 30px;

  .ranking-board-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
  }

  .dv-scroll-ranking-board {
    flex: 1;
  }
}
</style>
