<template>
  <div id="dateoutput-all">
    <gauge-chart :data="dateOutputAll4ChartDash" :style="chartProps.containerStyle" :options="chartProps.options" />
  </div>
</template>

<script>
import '@toast-ui/chart/dist/toastui-chart.min.css';
import { gaugeChart } from '@toast-ui/vue-chart';
import { mapState } from 'vuex';
import { store } from '@/store/index';

export default {
  name: 'DateOutputAll',
  components: {
    'gauge-chart': gaugeChart
  },
  computed: mapState(
    ['dateOutputAll4ChartDash']
  ),
  data () {
    return {
      selectItem : store.state.selectedLine,
      chartProps: {
        containerStyle: {
          width: '100%',
          height: '350px',
        },
        options: {
          chart: {  
            title: {
              text: '车间效率',
              align: 'center',
            }, 
          },
          circularAxis: { scale: { min: 0, max: 100 } },
          series: {
            angleRange: {
              start: 270,
              end: 90,
            },
            dataLabels: { visible: true, offsetY: -120, formatter: function(data){
              return data.toFixed(2) + "%"
              }
            //  (value) => `${value}%` 
             },
          },
          plot: {
            bands: [
              { range: [0, 20], color: 'red' },
              { range: [20, 50], color: 'yellow' },
              { range: [50, 80], color: 'green' },
              { range: [80, 100], color: '#45a8f8' },
            ],
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
            circularAxis: {
              title: { fontWeight: 500, fontSize: 20},
              label: { fontSize: 15, color: '#ffffff' },
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
                fontSize: 25,
                // color: '#fff',
                textBubble: {
                  fontSize: 25,
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
    },
  },
  mounted () {
  }
}
</script>

<style lang="less">
#dateoutput-all {
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
