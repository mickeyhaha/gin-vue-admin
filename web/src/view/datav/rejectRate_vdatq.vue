<template>
  <div id="reject-rate">
    <div class="title">抛料率</div>
    <column-chart :data="chartData" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { columnChart } from '@toast-ui/vue-chart';

import {
  getLackWarnings
} from "@/api/MoniWholeView";
import infoList from "@/mixins/infoList";

export default {
  name: 'RejectRate',
  mixins: [infoList],
  components: {
    'column-chart': columnChart
  },
  data () {
    return {
      listApi: getLackWarnings,
      // chartData: {
      // },
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '100%',
        },
        options: {
          title: {
            text: '抛料率%',
            align: 'center'
          },
          legend: {
            visible: false,
            //align: 'bottom'
          },
          series: {
            stack: true,
            dataLabels: { visible: true }
          },
          theme: {
            title: {
              fontFamily: 'Comic Sans MS',
              fontSize: 45,
              fontWeight: 100,
              color: '#ff416d'
            },
            chart: {
              fontFamily: 'Verdana',
              backgroundColor: 'rgba(9, 206, 115, 0.2)',
            },
            series: {
              dataLabels: {
                visible: true,
                stack: true,
                fontFamily: 'monaco',
                lineWidth: 2,
                textStrokeColor: '#ffffff',
                shadowColor: '#ffffff',
                shadowBlur: 4,
                stackTotal: {
                  fontFamily: 'monaco',
                  fontWeight: 14,
                  color: '#ffffff',
                  textBubble: {
                    visible: true,
                    paddingY: 6,
                    borderWidth: 3,
                    borderColor: '#00bcd4',
                    borderRadius: 7,
                    backgroundColor: '#041367',
                    shadowOffsetX: 0,
                    shadowOffsetY: 0,
                    shadowBlur: 0,
                    shadowColor: 'rgba(0, 0, 0, 0)'
                  }
                }
              }
            }
          }
        }
      },
      chartData: {
        categories: ['July', 'Aug', 'Sep', 'Oct', 'Nov'],
        series: [
          {
            name: '抛料率%',
            data: [0.3, 0.4, 0.2, 0.1, 0.5],
          }
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

  .dv-conical-column-chart {
    flex: 1;
  }
}
</style>
