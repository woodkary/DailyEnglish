<template>
  <view class="tab-bar">
    <router-link class="tab-item" to="/pages/Home/Home">Home</router-link>
    <router-link class="tab-item" to="/pages/About/About">About</router-link>
    <text class="tab-item">Helloworld</text>
  </view>
</template>

<style>
.tab-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 50px;
  background-color: #fff;
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.tab-item {
  color: #333;
  text-decoration: none;
}
</style>
