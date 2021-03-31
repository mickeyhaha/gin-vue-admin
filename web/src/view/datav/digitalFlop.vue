<template>
  <div id="digital-flop">
    <div
      class="digital-flop-item"
      v-for="item in tableData"
      :key="item.lineName"
    >
      <div class="digital-flop-title">{{ item.lineName }}</div>
      <div class="digital-flop">
        <!-- <dv-digital-flop
          :config="item.lineID"
          style="width:100px;height:50px;"
        /> -->
          <div class="unit">{{ item.lineID }}</div>
      </div>
    </div>

    <dv-decoration-10 />
  </div>
</template>

<script>
import {
  getPVS_Base_LineList
} from "@/api/PVS_Base_Line";
import infoList from "@/mixins/infoList";

export default {
  name: 'DigitalFlop',
  mixins: [infoList],
  data () {
    return {
      listApi: getPVS_Base_LineList,
    }
  },
  methods: {
    createData () {
      this.getTableData();
    },
  },
  mounted () {
    const { createData } = this

    createData()

    setInterval(createData, 3000)
  }
}
</script>

<style lang="less">
#digital-flop {
  position: relative;
  height: 15%;
  flex-shrink: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgba(6, 30, 93, 0.5);

  .dv-decoration-10 {
    position: absolute;
    width: 95%;
    left: 2.5%;
    height: 5px;
    bottom: 0px;
  }

  .digital-flop-item {
    width: 11%;
    height: 80%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    border-left: 3px solid rgb(6, 30, 93);
    border-right: 3px solid rgb(6, 30, 93);
  }

  .digital-flop-title {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .digital-flop {
    display: flex;
  }

  .unit {
    margin-left: 10px;
    display: flex;
    align-items: flex-end;
    box-sizing: border-box;
    padding-bottom: 13px;
  }
}
</style>
