<template>
  <div :id="chartId" :style="{width: '300px', height: '300px',display:'inline-block'}"></div>
</template>

<script>
// 引入基本模板
let echarts = require("echarts");

export default {
  name: "lineChart",
  props: ["chartId", "chartData"],
  data() {
    return {};
  },
  mounted() {
    this.drawLine();
  },
  methods: {
    drawLine() {
      // 基于准备好的dom，初始化echarts实例
      let myChart = echarts.init(document.getElementById(this.chartId));
      var option = {
        tooltip: {
          formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
          feature: {
            // restore: {},
            // saveAsImage: {}
          }
        },
        series: [
          {
            name: "业务指标",
            type: "gauge",
            detail: { fontSize: "20", formatter: "{value}%" }
            //data: [{ value: 10, name: "完成率" }]
          }
        ]
      };

      //setInterval(function() {
      //option.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
      myChart.setOption(option, true);
      // }, 2000);
    }
  }
};
</script>
