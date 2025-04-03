<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import BottomBar from './components/BottomBar.vue';
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <Navbar />
      </nav>
    </div>
  </header>

  <body>
  <RouterView />
  </body>
  <footer>
    <div id="footer-blur"></div>
    <div id="footer-blur-2"></div>
    <div id="footer-blur-3"></div>
    <BottomBar />
  </footer>
</template>

<style scoped>

body {
  height: auto;
}

footer {
  position: sticky;
  bottom: 0;
  width: 110%;
  transform: translateX(-5%);
  height: auto;
  min-height: 4rem;
  background-color: #444444;

}

footer > #footer-blur {
  min-height: 1px;
  transform: translateY(3px);
  box-shadow: 0 -3px 0 #aaaaaa;
  filter: blur(7px);
}

footer > #footer-blur-2 {
  display: flex;
  min-height: 1px;
  width: 100%;
  box-shadow: 0 -1px 0 #aaaaaa;
  filter: blur(5px);
}

footer > #footer-blur-3{
  display: flex;
  min-height: 1px;
  margin-left: 5%;
  width: 90%;
  box-shadow: 0 -1px 0 #aaaaaa;
  filter: blur(4px);
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

</style>
