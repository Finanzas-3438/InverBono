import { createApp } from 'vue/dist/vue.esm-bundler.js';
import { env } from './config.ts';
import Sidebar from './components/sidebar-component.vue';
function NewApp() {
  let app = createApp({})
  app.component("sidebar-component", Sidebar);
  // Common
  return app;
}

let lastApp: any = null;

export function initVue(callbackFn?: Function){
  if (window.vm) {
    lastApp?.unmount()
    lastApp = null;
  }
  let mountPoint = document.querySelector("[data-behavior='vue']");
  if (!mountPoint) {
    if (env == "dev"){
      console.warn("vue mount not found")
    }
    return;
  }

  // Save a copy of the original HTML
  const originalHTML = mountPoint.innerHTML;


  let app = NewApp();
  try {
    window.vm = app.mount("[data-behavior='vue']")
    if(callbackFn) callbackFn()
  } catch(e){
    alert("Vue mount error:" + e);
    console.error("Vue mount error", e);
    // Restore the original HTML
    mountPoint.innerHTML = originalHTML;
  }
  lastApp = app;
}