<template>
  <div class="tabs-container">
    <div class="tabs">
      <div 
        class="tab" 
        :class="{ active: selectedTab === 'tab1' }" 
        @click="selectTab('tab1')">
        {{ firstTab }}
      </div>
      <div 
        class="tab" 
        :class="{ active: selectedTab === 'tab2' }" 
        @click="selectTab('tab2')">
        {{ secondTab }}
      </div>
      <div class="underline" :style="underlineStyle"></div>
    </div>
    <div class="content">
      <div v-if="selectedTab === 'tab1'">
        <slot name="tab1-content"></slot>
      </div>
      <div v-if="selectedTab === 'tab2'">
        <slot name="tab2-content"></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props:{
    firstTab: String,
    secondTab: String
  },
  data() {
    return {
      selectedTab: 'tab1'
    };
  },
  computed: {
    underlineStyle() {
      return {
        transform: this.selectedTab === 'tab1' ? 'translateX(0%)' : 'translateX(100%)'
      };
    }
  },
  methods: {
    selectTab(tab) {
      this.selectedTab = tab;
    }
  }
};
</script>

<style scoped>
.tabs-container {
  width: 90%;
  margin: 0 auto;
  font-family: Arial, sans-serif;
}

.tabs {
  display: flex;
  position: relative;
  border-bottom: 2px solid #ccc;
}

.tab {
  flex: 1;
  font-size: 20px;
  font-weight: 550;
  text-align: center;
  padding: 5px 0;
  cursor: pointer;
}

.tab.active {
  color: #42b983;
}

.underline {
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 50%;
  height: 2px;
  background-color: #42b983;
  transition: transform 0.3s ease;
}

.content {
  /* padding: 20px 0; */
  /* text-align: center; */
}
</style>
