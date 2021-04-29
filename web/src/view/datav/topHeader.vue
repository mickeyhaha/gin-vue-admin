<template>
  <div id="top-header">
    <div class="top">
        <div class="left" style="text-align: left;">
            <img class="img" src="./img/logo.png">
        </div>
        <div class="middle">
            <p class="title">车间看板</p>
        </div>
        <div class="right">
            <span id="data"></span> <span id="week"></span>
        </div>
    </div>
    <!-- <dv-decoration-8 class="header-left-decoration" /> -->
    <!-- <dv-decoration-5 class="header-center-decoration" /> -->
    <!-- <dv-decoration-8 class="header-right-decoration" :reverse="true" /> -->
    <!-- <div class="center-title">车间看板</div> -->
  </div>
</template>

<script>
export default {
  name: 'TopHeader'
}
</script>

<style lang="less">
#top-header {
  position: relative;
  width: 100%;
  height: 30px;
  display: flex;
  justify-content: space-between;
  flex-shrink: 0;

  .header-center-decoration {
    width: 40%;
    height: 40px;
    margin-top: 30px;
  }

  .header-left-decoration, .header-right-decoration {
    width: 25%;
    height: 40px;
  }

  .center-title {
    position: absolute;
    font-size: 25px;
    font-weight: bold;
    left: 50%;
    top: 15px;
    transform: translateX(-50%);
  }

.top {
    width: 100%;
    // background: url(../image/2_01.png) no-repeat center center;
    background-size:cover;
    height: 10%;
}
.left img{
    height:55%;
    width:auto;
    text-align: left;
}
.right{
    float: right !important;
    font-size:0.6em;
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

}
</style>
