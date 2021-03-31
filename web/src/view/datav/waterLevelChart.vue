<template>
  <div id="water-level-chart">
    <div class="chart-container">
      <dv-water-level-pond :config="config" />
    </div>
    <div class="water-level-chart-title">
      工单生产进度
    </div>
  </div>
</template>

<script>
export default {
  name: 'WaterLevelChart',
  data () {
    return {
      config: {
        data: [45],
        shape: 'round',
        waveHeight: 25,
        waveNum: 2
      }
    }
  }
}
</script>

<style lang="less">

#water-level-chart {
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

  .water-level-chart-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
    justify-content: center;
  }

  .water-level-chart-details {
    height: 15%;
    display: flex;
    justify-content: center;
    font-size: 17px;
    align-items: flex-end;

    span {
      font-size: 35px;
      font-weight: bold;
      color: #58a1ff;
      margin: 0 5px;
      margin-bottom: -5px;
    }
  }

  .chart-container {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .dv-water-pond-level {
    width: 100px;
    height: 100px;
    border: 10px solid #19c3eb;
    border-radius: 50%;

    ellipse {
      stroke: transparent !important;
    }

    text {
      font-size: 40px;
    }
  }
}
</style>
