<template>
  <div id="aoi-rate-spc">
    <!-- <div class="title">AOI不良率</div> -->
    <line-chart :data="TS_AOI_Spc4Chart" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { lineChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';

export default {
  name: 'AoiRateSpc',
  components: {
    'line-chart': lineChart
  },
  computed: mapState(
    ['TS_AOI_Spc4Chart']
  ),
  data () {
    return {
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '100%',
        },
        options: {
          chart: {  
            title: {
              text: 'AOI不良(每30片)',
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
            stack: false,
            dataLabels: { visible: true },
          },
          // xAxis: { pointOnColumn: false, title: { text: '日期' } },
          // yAxis: { title: 'AOI不良率' },
          theme: {
            title: {
              fontWeight: 'bold',
              // fontSize: 20,
              color: '#ffffff',
            },
            xAxis: {
              title: {
                color: '#ffffff'
              },
              label: {
                fontSize: 15,
                // fontWeight: 'bold',
                color: '#ffffff'
              },
              // width: 2,
              color: '#ffffff'
            },
            yAxis: [
              {
                title: {
                  fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                label: {
                  fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                // width: 3,
                color: '#ffffff'
              },
              {
                title: {
                  fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                label: {
                  fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                width: 3,
                color: '#ffffff'
              }
            ],
            legend: {
              label: {
                fontSize: 15,
                // fontWeight: 'bold',
                color: '#ffffff'
              }
            },
            chart: {
              // fontFamily: 'Verdana',
              backgroundColor: '#12396a',//rgba(9, 206, 115, 0.1)
            },
            series: {
              dataLabels: {
                fontSize: 25,
                color: '#ffffff',
                visible: true,
                stack: true,
                // fontFamily: 'Verdana',
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
      // chartData: {
      //   categories: ['Line1', 'Line2', 'Line3', 'Line4', 'Line5'],
      //   series: [
      //     {
      //       name: '模糊',
      //       data: [0.3, 0.4, 0.2, 0.1, 0.5],
      //       stackGroup: 'AOI',
      //     },
      //     {
      //       name: '少料',
      //       data: [0.2, 0.3, 0.3, 0.2, 0.5],
      //       stackGroup: 'AOI',
      //     },
      //     {
      //       name: '错点',
      //       data: [0.4, 0.2, 0.1, 0.3, 0.1],
      //       stackGroup: 'AOI',
      //     },
      //   ],
      // }
    }
  },
  methods: {
    createData () {
      // this.getTableData().then(() => {
      //   this.chartData = this.tableData[0];
      // });
      
    },
  },
  mounted () {
    const { createData } = this

    createData()

    // setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#aoi-rate-spc {
  width: 100%;
  height: 100%;
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
