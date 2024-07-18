<template>
  <div id="app">
    <h1>Frontend Page</h1>
    <div class="wrapper-div">
      <iframe ref="myIframe" :src="iframeUrl" width="800" height="300"></iframe>
      <button @click="changeFontColor">Modify Div in Iframe</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      iframeUrl: '/py/',
    };
  },
  methods: {
    changeFontColor() {
      const iframe = this.$refs.myIframe;
      const iframeDocument = iframe.contentDocument || iframe.contentWindow.document;
      const targetDiv = iframeDocument.getElementById('targetDiv');
      if (targetDiv) {
        targetDiv.style.color = 'red';
      }
    }
  }
};

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.wrapper-div {
  position: relative;
  height: 400px; /* Set desired height for the wrapper div */
}
button {
   position: absolute;
   left: 50%;
   transform: translateX(-50%);
   bottom: 10px; /* Adjust as needed to set distance from bottom */
}
</style>