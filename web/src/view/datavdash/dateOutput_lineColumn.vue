<template>
  <div id="date-output">
    <!-- <div class="title">产量</div> -->
    <columnLine-chart :data="dateOutput4ChartDash" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { columnLineChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';

export default {
  name: 'DateOutput',
  components: {
    'columnLine-chart': columnLineChart
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
          legend: {
            visible: true,
            align: 'bottom'
          },
          series: {
            column: {
              dataLabels: {
                visible: true,
              }
            },
            line: {
              showDot: true,
              spline: true,
              dataLabels: {
                visible: true,
                formatter: function(data){
                  if (data <= 80) {
                    return  "!! " + data.toFixed(2) + " !!"
                  } else {
                    return data.toFixed(2) + ""
                  }
                }
              }
            }
          },
          theme: {
            title: {
              fontWeight: 'bold',
              // fontSize: 20,
              color: '#ffffff',
            },
            chart: {
              // fontFamily: 'Verdana',
              backgroundColor: '#12396a',//rgba(9, 206, 115, 0.1)
            },
            xAxis: {
              title: {
                // fontSize: 25,
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
                  // fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                label: {
                  fontSize: 15,
                  fontWeight: 'bold',
                  color: '#ffffff'
                },
                width: 2,
                color: '#ffffff'
              },
              {
                title: {
                  // fontSize: 15,
                  // fontWeight: 'bold',
                  color: '#ffffff'
                },
                label: {
                  fontSize: 15,
                  fontWeight: 'bold',
                  color: '#ffffff'
                },
                width: 2,
                color: '#ffffff'
              },
            ],
            legend: {
              label: {
                fontSize: 15,
                // fontWeight: 'bold',
                color: '#ffffff'
              }
            },
            series: {
              column: {
                dataLabels: {
                  fontSize: 25,
                  color: '#ffffff',
                }
              },
              line: {
                dataLabels: {
                  fontSize: 25,
                  color: '#ffffff',
                  // visible: false,
                  // stack: false,
                  // fontFamily: 'Verdana',
                //   lineWidth: 2,
                //   textStrokeColor: '#ffffff',
                //   shadowColor: '#ffffff',
                //   shadowBlur: 4,
                  // stackTotal: {
                    // fontFamily: 'Verdana',
                    // fontSize: 25,
                    // fontWeight: 14,
                    // color: '#ffffff',
                    // textBubble: {
                    //   visible: true,
                    //   paddingY: 6,
                    //   borderWidth: 3,
                    //   borderColor: '#00bcd4',
                    //   borderRadius: 7,
                    //   backgroundColor: '#041367',
                    //   shadowOffsetX: 0,
                    //   shadowOffsetY: 0,
                    //   shadowBlur: 0,
                    //   shadowColor: 'rgba(0, 0, 0, 0)'
                    // }
                  // }
                }
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
