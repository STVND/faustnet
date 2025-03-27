<script setup lang="ts">
import { RouterView } from 'vue-router'
import navbar from './components/Navbar.vue'
import BottomBar from './components/BottomBar.vue';
</script>

<template>
  <div id="active-view">
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />
    <div class="wrapper">
      <nav>
        <navbar />
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
</div>
</template>

<style scoped>

header {
  line-height: 1.5;
  max-height: 100vh;
}

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

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
