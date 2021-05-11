<template>
  <div id="reject-rate">
    <bar-chart :data="rejectRate4ChartDash" :style="chartProps.containerStyle" :options="chartProps.options" @selectSeries="onSelectSeries" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { barChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';

export default {
  name: 'RejectRate',
  components: {
    'bar-chart': barChart
  },
  computed: mapState(
    ['rejectRate4ChartDash']
  ),
  data () {
    return {
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '200px',
        },
        options: {
          chart: {  
            title: {
              text: '抛料率(%)',
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
            selectable: true,
            stack: false,
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
      //   categories: ['站料1', '站料2', '站料3', '站料4', '站料5'],
      //   series: [
      //     {
      //       name: '模糊',
      //       data: [0.3, 0.4, 0.2, 0.1, 0.5],
      //     },
      //   ],
      // }
    }
  },
  methods: {
    onSelectSeries(ev) {
    },

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
#reject-rate {
  width: 100%;
  // height: 100%;
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
