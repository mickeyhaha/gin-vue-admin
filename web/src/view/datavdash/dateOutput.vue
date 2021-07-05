<template>
  <div id="date-output">
    <!-- <div class="title">产量</div> -->
    <bar-chart :data="dateOutput4ChartDash" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { barChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';

export default {
  name: 'DateOutput',
  components: {
    'bar-chart': barChart
  },
  computed: mapState(
    ['dateOutput4ChartDash']
  ),
  data () {
    return {
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '350px',
        },
        options: {
          chart: {  
            title: {
              text: '产量(件)',
              align: 'center',
            }, 
          },
          grid: {
              top:"50px",
              left:"50px",
              right:"15px",                
              bottom:"50px",
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
                  fontWeight: 'bold',
                  color: '#ffffff'
                },
                // width: 3,
                color: '#ffffff'
              },
              {
                title: {
                  // fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                label: {
                  // fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                width: 3,
                color: '#ffffff'
              }
            ],
            // plot: {
            //   vertical: {
            //     lineColor: 'white',
            //     lineWidth: 5,
            //     dashSegments: [5, 20],
            //   },
            //   horizontal: {
            //     lineColor: 'rgba(0, 0, 0, 0)',
            //   },
            //   backgroundColor: 'rgba(60, 80, 180, 0.1)'
            // },
            legend: {
              label: {
                // fontSize: 15,
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
#date-output {
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
