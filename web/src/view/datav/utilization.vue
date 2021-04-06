<template>
  <div id="utilization">
    <!-- <div class="title">稼动率</div> -->
    <line-chart :data="chartData" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { lineChart } from '@toast-ui/vue-chart';

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'Utilization',
  mixins: [infoList],
  components: {
    'line-chart': lineChart
  },
  data () {
    return {
      listApi: getLackWarnings,
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '190px',
        },
        options: {
          chart: {  
            title: {
              text: '稼动率%',
              align: 'center',
            }, 
          },
          tooltip: {
            visible: false,
          },
          legend: {
            visible: true,
            align: 'bottom'
          },
          series: {
            stack: true,
            dataLabels: { visible: true },
          },
          xAxis: { pointOnColumn: false, title: { text: '时间' } },
          // yAxis: { title: 'AOI不良率' },
          theme: {
            // title: {
            //   fontFamily: 'Comic Sans MS',
            //   fontSize: 45,
            //   fontWeight: 100,
            //   color: '#ff416d'
            // },
            chart: {
              fontFamily: 'Verdana',
              backgroundColor: 'rgba(9, 206, 115, 0.1)',
            },
            series: {
              showDot: true,
              dataLabels: {
                visible: true,
                stack: true,
                fontFamily: 'Verdana',
              //   lineWidth: 2,
              //   textStrokeColor: '#ffffff',
              //   shadowColor: '#ffffff',
              //   shadowBlur: 4,
                // stackTotal: {
                //   fontFamily: 'Verdana',
                //   fontWeight: 14,
                //   color: '#ffffff',
                //   textBubble: {
                //     visible: true,
                //     paddingY: 6,
                //     borderWidth: 3,
                //     borderColor: '#00bcd4',
                //     borderRadius: 7,
                //     backgroundColor: '#041367',
                //     shadowOffsetX: 0,
                //     shadowOffsetY: 0,
                //     shadowBlur: 0,
                //     shadowColor: 'rgba(0, 0, 0, 0)'
                //   }
                // }
              }
            }
          }
        }
      },
      chartData: {
        categories: ['13:00', '14:00', '15:00', '16:00', '17:00'],
        series: [
          {
            name: 'Line1',
            data: [0.3, 0.4, 0.2, 0.1, 0.5],
          },
          {
            name: 'Line2',
            data: [0.2, 0.3, 0.3, 0.2, 0.5],
          },
          {
            name: 'Line3',
            data: [0.4, 0.2, 0.1, 0.3, 0.1],
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
#utilization {
  width: 100%;
  height: 30%;
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
