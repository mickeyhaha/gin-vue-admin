<template>
  <div id="date-runtime">
    <column-chart :data="dateRunTime4ChartDash" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { columnChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';

export default {
  name: 'DateRunTime',
  components: {
    'column-chart': columnChart
  },
  computed: mapState(
    ['dateRunTime4ChartDash']
  ),
  data () {
    return {
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '300px',
        },
        options: {
          chart: {  
            title: {
              text: '停机时长(秒)',
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
          // xAxis: { pointOnColumn: false, title: { text: '日期' } },
          // yAxis: { title: 'AOI不良率' },
          theme: {
            title: {
              fontWeight: 'bold',
              fontSize: 15,
              color: '#ffffff',
            },
            xAxis: {
              title: {
                color: '#ffffff'
              },
              label: {
                fontSize: 11,
                fontWeight: 700,
                color: '#ffffff'
              },
              // width: 2,
              color: '#ffffff'
            },
            yAxis: [
              {
                title: {
                  fontSize: 17,
                  fontWeight: 400,
                  color: '#ffffff'
                },
                label: {
                  fontSize: 11,
                  fontWeight: 700,
                  color: '#ffffff'
                },
                // width: 3,
                color: '#ffffff'
              },
              {
                title: {
                  fontSize: 13,
                  fontWeight: 600,
                  color: '#ffffff'
                },
                label: {
                  fontSize: 11,
                  fontWeight: 700,
                  color: '#ffffff'
                },
                width: 3,
                color: '#ffffff'
              }
            ],
            legend: {
              label: {
                // fontSize: 15,
                // fontWeight: 700,
                color: '#ffffff'
              }
            },
            chart: {
              fontFamily: 'Verdana',
              backgroundColor: '#12396a',//rgba(9, 206, 115, 0.1)
            },
            series: {
              dataLabels: {
                color: '#ffffff',
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
      // chartData: {
      //   categories: ['4/18', '4/19', '4/20', '4/18', '4/19', '4/20'],
      //   series: [
      //     {
      //       name: '实际产量',
      //       data: [40, 40, 20, 10, 50],
      //     },
      //     {
      //       name: '标准产量',
      //       data: [32, 42, 21, 11, 52],
      //     },
      //   ],
      // }
    }
  },
  methods: {
    createData () {
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
#date-runtime {
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
