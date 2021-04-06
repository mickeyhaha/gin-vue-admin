<template>
  <div id="scroll-board">
    <div class="title">停机原因</div>
    <dv-scroll-board :config="config" />
  </div>
</template>

<script>
export default {
  name: 'ScrollBoard',
  data () {
    return {
      config: {

        header: ['机器', '停机类别', '位置', '停机原因', '停机端口'],
        data: [
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
          [ '机器1', '卡纸',  '1-1',  '太厚',  '1'],
        ],
        index: false,
        align: ['center'],
        rowNum: 6,
        headerBGC: '#5ca8f2',
        headerHeight: 25,
        oddRowBGC: 'rgba(0, 44, 81, 0.8)',
        evenRowBGC: 'rgba(10, 29, 50, 0.8)'
      }
    }
  }
}
</script>

<style lang="less">
#scroll-board {
  width: 100%;
  height: 30%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, .5);
  box-sizing: border-box;
  padding: 0px 10px;
  margin-bottom: 10px;
}
</style>
