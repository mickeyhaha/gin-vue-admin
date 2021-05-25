<template>
  <baidu-map class="map" :center="center" :zoom="zoom" :scroll-wheel-zoom="true"
   @ready="handler">
       <bm-city-list anchor="BMAP_ANCHOR_TOP_LEFT"></bm-city-list>
       <bm-scale anchor="BMAP_ANCHOR_TOP_RIGHT"></bm-scale>
       <bm-overview-map anchor="BMAP_ANCHOR_BOTTOM_RIGHT" :isOpen="true"></bm-overview-map>
        <bm-geolocation anchor="BMAP_ANCHOR_BOTTOM_RIGHT" :showAddressBar="true" :autoLocation="true"></bm-geolocation>
        <bm-copyright
          anchor="BMAP_ANCHOR_TOP_RIGHT"
          :copyright="[{id: 1, content: 'Copyright Message', bounds: {ne: {lng: 110, lat: 40}, sw:{lng: 0, lat: 0}}}, {id: 2, content: '<a>我是版权信息</a>'}]">
        </bm-copyright>
        <bm-control>
          <button @click="addZoom(19)">缩放至最大</button>
          <button @click="addZoom(10)">还原</button>
          <button @click="addZoom(3)">缩放至最小</button>
        </bm-control>
    
     <bm-marker :position="{lng: 116.404, lat: 39.915}" :dragging="true" @click="infoWindowOpen">
      <bm-info-window :show="show" @close="infoWindowClose" @open="infoWindowOpen">我爱北京天安门</bm-info-window>
    </bm-marker>

        <bm-point-collection :points="points" shape="BMAP_POINT_SHAPE_STAR" color="red" size="BMAP_POINT_SIZE_SMALL" @click="clickHandler"></bm-point-collection>

  <!-- <bml-marker-clusterer :averageCenter="true">
    <bm-marker v-for="marker of markers" :position="{lng: marker.lng, lat: marker.lat}"></bm-marker>
  </bml-marker-clusterer> -->
  </baidu-map>
</template>

<script>
//https://github.com/Dafrok/vue-baidu-map
//  import {BmlMarkerClusterer} from 'vue-baidu-map'

export default {
  name: 'MyMap',
  data () {
    return {
      center: {lng: 0, lat: 0},
      zoom: 3,
      show: false,
      points: [],
    }
  },
  // components: {
  //   BmlMarkerClusterer
  // },
  methods: {
    handler ({BMap, map}) {
      console.log(BMap, map)
      this.center.lng = 116.404
      this.center.lat = 39.915
      this.zoom = 15
    },
    addZoom (level) {
      this.zoom = level
    },
    infoWindowClose () {
      this.show = false
    },
    infoWindowOpen() {
      this.show = true
    },
    clickHandler(e) {
      alert(`单击点的坐标为：${e.point.lng}, ${e.point.lat}`);
    },
    addPoints() {
      const points = [];
      for (var i = 0; i < 1000; i++) {
        const position = {lng: Math.random() * 40 + 85, lat: Math.random() * 30 + 21}
        points.push(position)
      }
      this.points = points
    },

    mounted () {
      addPoints()
    },
  }
}
</script>

<style>
.map {
  width: 100%;
  height: 300px;
}
</style>
