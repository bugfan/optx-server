
<template>
  <div id="app">
    <v-header></v-header>
    <div class="detail-wrap">
      <div class="detail-left">
        <div class="product-board">
          <!--<img src="@/assets/img/1.png"/>-->
          <h2>功能</h2>
          <ul>
            <router-link :to="{ name: 'options' }" @click.native="add(1)" tag="li" active-class="active">
              题 库
            </router-link>
            <router-link :to="{ name: 'insert' }" tag="li" @click.native="add(2)" active-class="active">
              批量插入
            </router-link>
          </ul>
          <div class="hr"></div>
        </div>
      </div>
      <div class="detail-right">
        <keep-alive>
          <router-view></router-view>
        </keep-alive>
      </div>
    </div>
    <!--<div class="app-foot">
      <p>© 2018 made by zyy</p>
    </div>-->
    <!--<this-dialog :is-show="isShowAbout" @on-close="hideDialog('isShowAbout')">
      <p>本报告在调研数据的基础上，采用定性与定量相结合的方式深入分析了专车市场发展的驱动因素与阻碍因素、专车市场背后的产业格局、专车企业的竞争格局、用户对专车市场的依赖程度、专车对其他交通工具运力的补充效应等，通过这五个章节的研究反映专车市场的发展态势和面临的问题。报告力求客观、深入、准确地反映中国专车市场发展情况，为政府、企事业单位和社会各界提供决策依据。 </p>
    </this-dialog>
    <this-dialog :is-show="isShowLogin" @on-close="hideDialog('isShowLogin')">
      <login-form @on-success="" @on-error=""></login-form>
    </this-dialog>-->
  </div>
</template>
<script>
import header from "./components/header";

export default {
  name: "App",
  components: {
    "v-header": header
  },
  methods: {
    add (v) {
      // alert(v)
    }
  }
};
</script>

<style>
/* http://meyerweb.com/eric/tools/css/reset/ 
   v2.0 | 20110126
   License: none (public domain)
*/

html,
body,
div,
span,
applet,
object,
iframe,
h1,
h2,
h3,
h4,
h5,
h6,
p,
blockquote,
pre,
a,
abbr,
acronym,
address,
big,
cite,
code,
del,
dfn,
em,
img,
ins,
kbd,
q,
s,
samp,
small,
strike,
strong,
sub,
sup,
tt,
var,
b,
u,
i,
center,
dl,
dt,
dd,
ol,
ul,
li,
fieldset,
form,
label,
legend,
table,
caption,
tbody,
tfoot,
thead,
tr,
th,
td,
article,
aside,
canvas,
details,
embed,
figure,
figcaption,
footer,
header,
hgroup,
menu,
nav,
output,
ruby,
section,
summary,
time,
mark,
audio,
video {
  margin: 0;
  padding: 0;
  border: 0;
  font-size: 100%;
  font: inherit;
  vertical-align: baseline;
}
body {
  line-height: 1;
}
ol,
ul {
  list-style: none;
}
blockquote,
q {
  quotes: none;
}
blockquote:before,
blockquote:after,
q:before,
q:after {
  content: "";
  content: none;
}
table {
  border-collapse: collapse;
  border-spacing: 0;
}
a {
  color: inherit;
  text-decoration: none;
}
body {
  background: #f0f2f5;
  font-family: "Helvetica Neue", Helvetica, Arial, "Hiragino Sans GB",
    "Hiragino Sans GB W3", "Microsoft YaHei UI", "Microsoft YaHei",
    "WenQuanYi Micro Hei", sans-serif;
  font-size: 14px;
  color: #444;
}
.container {
  width: 1200px;
  margin: 0 auto;
}
.detail-wrap {
  width: 1200px;
  margin: 0 auto;
  overflow: hidden;
  padding-top: 20px;
}
.detail-left {
  float: left;
  width: 200px;
}
.detail-right {
  float: left;
  width: 980px;
  margin-left: 20px;
}
.product-board {
  background: #fff;
}
.product-board h2 {
  background: rgb(162, 223, 172);
  color: #fff;
  padding: 10px 15px;
  margin-bottom: 20px;
}
.product-board ul {
  margin-top: 20px;
}
.product-board li {
  text-align: left;
  padding: 10px 15px;
  cursor: pointer;
}
.product-board li.active,
.product-board li:hover {
  cursor: pointer;
}
.product-board li a {
  display: block;
}
.hr {
  height: 1px;
  width: 100%;
  background: #ddd;
}
#app {
  position: relative;
  height: 100%;
}
.app-foot {
  position: absolute;
  bottom: 10px;
  text-align: center;
  height: 50px;
  width: 100%;
  line-height: 50px;
  background: #e3e4e8;
  margin-top: 30px;
}
</style>
