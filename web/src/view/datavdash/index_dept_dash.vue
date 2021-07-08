<template>
  <div id="data-view-dept">
    <dv-full-screen-container>
      <div class="top">
          <div class="left" style="text-align: left;">
              <img class="img" src="../datav/img/logo.png">
          </div>
          <div class="middle">
              <p class="center-title">车间看板</p>
          </div>
          <div class="right" style="text-align: right; padding-right:10px">
            <span>{{date}}</span>
          </div>
      </div>
      <!-- <el-row>
        <el-col :span="24"><top-header /></el-col>
      </el-row> 
      <el-row>
        <el-col :span="24"><DeptFilter /></el-col>
      </el-row>-->
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
            <DateOutputEff />
              <!-- <spi-rate />  -->
              <!-- <utiliazation />  -->
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <DateOutputAll /> 
            <!-- <AoiRateSpc />  -->
            <!-- <RejectRate /> -->
            <!-- <lack-warning /> -->
            <!-- <scroll-board /> -->
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8" class="block-top-bottom-content">
              <DateRunTime /> 
        </el-col>
        <el-col :span="8" class="block-top-bottom-content">
            <Oee />
            <!-- <StopReason /> -->
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
import RejectRate from './rejectRate'
// import aoiRate from './aoiRate'
import DeptLineSummary from './deptLineSummary'
import DateOutput from './dateOutput'
import DateOutputEff from './dateOutputEff'
import DateMachineEvent from "./dateMachineEvent";
import DateRunTime from "./dateRunTime";
import StopReason from "./stopReason";
import AoiRateSpc from "./aoiRateSpc";
import DateOutputAll from "./dateOutputAll"
import Oee from './oee'

export default {
  name: 'DataViewDept',
  components: {
    // aoiRate,
    RejectRate,
    DeptLineSummary,
    DateOutput,
    DateOutputEff,
    DateMachineEvent,
    DateRunTime,
    StopReason,
    AoiRateSpc,
    DateOutputAll,
    Oee,
  },
  data () {
    return {  
      date: '',
      formData: {
        date: null,
        LineName: undefined,
        Shift: undefined,
        WorkOrderNo: undefined,
      },
    }
  },
  created() {
            setInterval(() => {
                let date = new Date();
                this.date = formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
            }, 1000)
  },
  methods: {
    createData () {
      this.$store.dispatch('refreshDeptDashboard', this.formData);
    },
  },
  mounted () {
    const { createData } = this

    createData()

    setInterval(createData, 30000)
  }
}
</script>

<style lang="less">
  .title {
    font-weight: bold;
    height: 30px;
    display: flex;
    align-items: center;
    font-size: 30px;
  }

#data-view-dept {
  width: 100%;
  height: 100%;
  background-color: #030409;
  color: #fff;

  #dv-full-screen-container {
    background-image: url('../datav/img/bg.png');
    background-size: 100% 100%;
    box-shadow: 0 0 3px blue;
    display: flex;
    flex-direction: column;
  }
  
  // top header
.top {
    background-image: url('../datav/img/bg_top.png');
    background-size: 100% 100%;
    width: 100%;
    background-size:cover;
    height: 8%;
}
.left img{
    height:70%;
    width:auto;
    text-align: left;
}
.right{
    float: right !important;
    font-size:20;
    font-weight: 'bold';

}
.middle{
    margin-left:23%;
    font-weight: 'bold';
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
    font-size: 30px;
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
    margin-top: 30px;
  }

  .block-top-bottom-content {
    // flex: 1;
    // display: flex;
    // width: 33%;
    height: 100%;
    // flex-direction: column;
    // box-sizing: border-box;
    padding: 0px 30px;
  }

  .block-top-content {
    height: 55%;
    display: flex;
    flex-grow: 0;
    box-sizing: border-box;
    padding-bottom: 30px;
  }
}
</style>