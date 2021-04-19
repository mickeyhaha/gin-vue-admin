<template>
  <div id="data-view-dept">
    <dv-full-screen-container>
      <el-row>
        <el-col :span="24"><top-header /></el-col>
      </el-row>
      <el-row>
        <el-col :span="24"><DeptFilter /></el-col>
      </el-row>
      <el-row>
        <el-col :span="24"><DeptLineSummary /></el-col>
      </el-row>
      <el-row>
        <el-col :span="8" class="block-top-bottom-content">
            <reject-rate />
            <lack-warning />
            <scroll-board />
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <order-info /> 
            <water-level-chart /> 
            <cards />
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
              <aoi-rate /> 
              <spi-rate /> 
              <utiliazation /> 
        </el-col>
      </el-row>
    </dv-full-screen-container>
  </div>
</template>

<script>
import topHeader from './topHeader'
// import digitalFlop from './digitalFlop'
// import rankingBoard from './rankingBoard'
// import roseChart from './roseChart'
import waterLevelChart from './waterLevelChart'
import scrollBoard from './scrollBoard'
import cards from './cards'
import lackWarning from './lackWarning'
import rejectRate from './rejectRate'
import spiRate from './spiRate'
import aoiRate from './aoiRate'
import orderInfo from './orderInfo'
import utiliazation from './utilization'
import DeptLineSummary from './deptLineSummary'
import DeptFilter from './deptFilter'
// import person from './person'

export default {
  name: 'DataViewDept',
  components: {
    topHeader,
    // digitalFlop,
    // rankingBoard,
    // roseChart,
    waterLevelChart,
    scrollBoard,
    cards,
    lackWarning,
    rejectRate,
    spiRate,
    aoiRate,
    orderInfo,
    utiliazation,
    DeptLineSummary,
    DeptFilter,
    // person,
  },
  data () {
    return {  
    }
  },
  methods: {}
}
</script>

<style lang="less">
  .title {
    font-weight: bold;
    height: 30px;
    display: flex;
    align-items: center;
    font-size: 15px;
  }

#data-view-dept {
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