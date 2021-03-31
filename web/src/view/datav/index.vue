<template>
  <div id="data-view">
    <dv-full-screen-container>
      <el-row>
        <el-col :span="24"><top-header /></el-col>
      </el-row>
      <el-row>
        <el-col :span="8" class="block-top-bottom-content">
            <reject-rate />
            <lack-warning />
            <scroll-board />
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
              <water-level-chart /> 
              <water-level-chart /> 
              <water-level-chart /> 
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
              <spi-rate /> 
              <spi-rate /> 
              <spi-rate /> 
        </el-col>
      </el-row>

      <!-- <div class="main-content">
         <digital-flop />  

        <div class="block-left-right-content">

          <div class="block-top-bottom-content">
            <reject-rate />
            <lack-warning />
            <scroll-board />
          </div>

          <div class="block-top-bottom-content"> 
              <water-level-chart /> 
              <water-level-chart /> 
              <water-level-chart /> 
          </div>

          <div class="block-top-bottom-content"> 
              <spi-rate /> 
              <spi-rate /> 
              <spi-rate /> 
          </div>
        </div>
      </div> -->
    </dv-full-screen-container>
  </div>
</template>

<script>
import topHeader from './topHeader'
import digitalFlop from './digitalFlop'
import rankingBoard from './rankingBoard'
import roseChart from './roseChart'
import waterLevelChart from './waterLevelChart'
import scrollBoard from './scrollBoard'
import cards from './cards'
import lackWarning from './lackWarning'
import rejectRate from './rejectRate'
import spiRate from './spiRate'

export default {
  name: 'DataView',
  components: {
    topHeader,
    digitalFlop,
    rankingBoard,
    roseChart,
    waterLevelChart,
    scrollBoard,
    cards,
    lackWarning,
    rejectRate,
    spiRate
  },
  data () {
    return {}
  },
  methods: {}
}
</script>

<style lang="less">
  .title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
  }

#data-view {
  width: 100%;
  height: 100%;
  background-color: #030409;
  color: #fff;

  #dv-full-screen-container {
    background-image: url('./img/bg.png');
    background-size: 100% 100%;
    box-shadow: 0 0 3px blue;
    display: flex;
    flex-direction: column;
  }

  .main-content {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  .block-left-right-content {
    flex: 1;
    display: flex;
    margin-top: 20px;
  }

  .block-top-bottom-content {
    // flex: 1;
    // display: flex;
    // width: 33%;
    height: 100%;
    // flex-direction: column;
    // box-sizing: border-box;
    padding-left: 20px;
  }

  .block-top-content {
    height: 55%;
    display: flex;
    flex-grow: 0;
    box-sizing: border-box;
    padding-bottom: 20px;
  }
}
</style>
