<template>
  <div id="data-view-dept">
    <dv-full-screen-container>
      <div class="top">
          <div class="left" style="text-align: left;">
              <img class="img" src="./img/logo.png">
          </div>
          <div class="middle">
              <p class="center-title">报表查询</p>
          </div>
          <div class="right" style="text-align: right; padding-right:10px">
            <span>{{date}}</span>
          </div>
      </div>
      <!-- <el-row>
        <el-col :span="24"><top-header /></el-col>
      </el-row> -->
      <el-row>
        <el-col :span="24"><DeptFilter /></el-col>
      </el-row>
      <el-row>
        <el-col :span="24"><DeptLineSummary /></el-col>
      </el-row>
      <!-- <el-row>
        <el-col :span="24"><TimeOutput /></el-col>
      </el-row> -->
      <!-- <el-row>
        <el-col :span="24"><DateOutput /></el-col>
      </el-row> -->
      <el-row>
        <el-col :span="8" class="block-top-bottom-content">
            <DateOutput />
            <!-- <water-level-chart />  -->
            <!-- <cards /> -->
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
              <AoiRateSpc /> 
              <!-- <spi-rate />  -->
              <!-- <utiliazation />  -->
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <RejectRate />
            <!-- <lack-warning /> -->
            <!-- <scroll-board /> -->
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8" class="block-top-bottom-content">
              <DateRunTime /> 
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <StopReason />
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <DateMachineEvent />
        </el-col>
      </el-row>
    </dv-full-screen-container>
  </div>
</template>

<script>
import { formatTimeToStr } from "@/utils/date";
// import topHeader from './topHeader'
// import digitalFlop from './digitalFlop'
// import rankingBoard from './rankingBoard'
// import roseChart from './roseChart'
// import waterLevelChart from './waterLevelChart'
// import scrollBoard from './scrollBoard'
// import cards from './cards'
// import lackWarning from './lackWarning'
import RejectRate from './rejectRate'
// import spiRate from './spiRate'
// import aoiRate from './aoiRate'
// import orderInfo from './orderInfo'
// import utiliazation from './utilization'
import DeptLineSummary from './deptLineSummary'
import DeptFilter from './deptFilter'
// import TimeOutput from './timeOutput'
import DateOutput from './dateOutput'
import DateMachineEvent from "./dateMachineEvent";
import DateRunTime from "./dateRunTime";
import StopReason from "./stopReason"
import AoiRateSpc from "./aoiRateSpc"

export default {
  name: 'DataViewDept',
  components: {
    // topHeader,
    // digitalFlop,
    // rankingBoard,
    // roseChart,
    // waterLevelChart,
    // scrollBoard,
    // cards,
    // lackWarning,
    RejectRate,
    // spiRate,
    // aoiRate,
    // orderInfo,
    // utiliazation,
    DeptLineSummary,
    DeptFilter,
    // TimeOutput,
    DateOutput,
    DateMachineEvent,
    DateRunTime,
    StopReason,
    // person,
    AoiRateSpc,
  },
  data () {
    return {  
      date: '',
    }
  },
  created() {
            // let target = new Date("2019-6-6").getTime(); //获取目标时间戳
            setInterval(() => {
                let date = new Date();
                this.date = formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
            }, 1000)
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
    width: 100%;
    height: 2048px;
    background-image: url('./img/bg.png');
    background-size: 100% 100%;
    box-shadow: 0 0 3px blue;
    display: flex;
    flex-direction: column;
  }
  
  // top header
.top {
    background-image: url('./img/bg_top.png');
    background-size: 100% 100%;
    width: 100%;
    background-size:cover;
    height: 6%;
}
.left img{
    height:70%;
    width:auto;
    text-align: left;
}
.right{
    float: right !important;
    font-size:15;
}
.middle{
    margin-left:23%;
    font-weight: 700;
}
.top div *{
    position: relative;
    top: 10%;
}
.top > div{
    float:left;
    text-align: center;
    height:100%;
    width:18%;
    color: #fff;
    line-height: 40px;
}

  .center-title {
    position: absolute;
    font-size: 25px;
    font-weight: bold;
    left: 50%;
    top: 15px;
    transform: translateX(-50%);
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
    padding: 0px 20px;
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