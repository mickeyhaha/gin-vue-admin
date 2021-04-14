<template>
  <div id="finish-rate">
    <gauge-chart :data="chartData" :style="chartProps.containerStyle" :options="chartProps.options" />
    <!-- <input class="uni-input" name="totalPrice" type="number" v-model="getRate" /> -->
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { gaugeChart } from '@toast-ui/vue-chart';

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";
import { store } from '@/store/index';

export default {
  name: 'FinishRate',
  mixins: [infoList],
  components: {
    'gauge-chart': gaugeChart
  },
  computed: {
    getQty() {
      return this.$store.state.selectedLine.Qty;
    },

    getQtyCompleted() {
      return this.$store.state.selectedLine.QtyCompleted;
    },

    getRate() {
      return this.$store.state.selectedLine.QtyCompleted * 100 / this.$store.state.selectedLine.Qty;
    },
  },
  data () {
    return {
      selectItem : store.state.selectedLine,
      listApi: getLackWarnings,
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '280px',
        },
        options: {
          chart: {
            title: {
              text: '完成度%',
              align: 'center',
            }, width: '100%'},
          circularAxis: { title: this.$store.state.selectedLine.QtyCompleted + '/' + this.$store.state.selectedLine.Qty, scale: { min: 0, max: 100 } },
          series: {
            angleRange: {
              start: 270,
              end: 90,
            },
            dataLabels: { visible: true, offsetY: -120, formatter: (value) => `${value}%` },
          },
          plot: {
            bands: [
              { range: [0, 20], color: '#45a8f8' },
              { range: [20, 50], color: '#45a8f8' },
              { range: [50, 100], color: '#45a8f8' },
            ],
          },
          theme: {
            chart: {
              fontFamily: 'Verdana',
              backgroundColor: 'rgba(9, 206, 115, 0.1)',
            },
            circularAxis: {
              title: { fontWeight: 500, fontSize: 20},
              label: { fontSize: 15 },
            },
            series: {
              clockHand: {
                // color: '#650434',
                baseLine: 10,
              },
              pin: {
                radius: 10,
                // color: '#650434',
                borderWidth: 5,
                // borderColor: 'rgba(101, 4, 52, 0.3)',
              },
              dataLabels: {
                fontSize: 20,
                // color: '#fff',
                textBubble: {
                  visible: true,
                  // backgroundColor: '#650434',
                  paddingX: 5,
                  paddingY: 5,
                },
              },
            },
            plot: { bands: { barWidth: 30 } },
          },
        }
      },
      chartData: {
        series: [
          {
            name: '完成度',
            data: [this.$store.state.selectedLine.QtyCompleted * 100 / this.$store.state.selectedLine.Qty],
          },
        ],
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

    console.log(this.selectItem)
    //setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#finish-rate {
  width: 100%;
  // height: 33%;
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
