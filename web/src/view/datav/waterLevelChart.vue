<template>
  <div id="finish-rate">
    <gauge-chart :data="chartData" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { gaugeChart } from '@toast-ui/vue-chart';

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'FinishRate',
  mixins: [infoList],
  components: {
    'gauge-chart': gaugeChart
  },
  data () {
    return {
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
          circularAxis: { title: '9000/10000', scale: { min: 0, max: 100 } },
          series: {
            angleRange: {
              start: 270,
              end: 90,
            },
            dataLabels: { visible: true, offsetY: -120, formatter: (value) => `${value}%` },
          },
          plot: {
            bands: [
              { range: [0, 20], color: '#e4a0c2' },
              { range: [20, 50], color: '#dc4d94' },
              { range: [50, 100], color: '#a90757' },
            ],
          },
          theme: {
            chart: {
              fontFamily: 'Verdana',
              backgroundColor: 'rgba(9, 206, 115, 0.1)',
            },
            circularAxis: {
              title: { fontWeight: 500, fontSize: 20, color: '#650434' },
              label: { color: '#650434', fontSize: 15 },
              tick: { strokeStyle: '#650434' },
              strokeStyle: '#650434',
            },
            series: {
              clockHand: {
                color: '#650434',
                baseLine: 10,
              },
              pin: {
                radius: 10,
                color: '#650434',
                borderWidth: 5,
                borderColor: 'rgba(101, 4, 52, 0.3)',
              },
              dataLabels: {
                fontSize: 20,
                color: '#fff',
                textBubble: {
                  visible: true,
                  backgroundColor: '#650434',
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
            data: [90],
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

    setInterval(createData, 3000)
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
