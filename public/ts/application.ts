// Vue.js
import { initVue } from './vue.ts';


// TailwindCSS Styles
import "./tailwind.css";

// Landing page 
// import { setupLanding } from './landing.ts';
import { env } from './config.ts';
// import { setupDaVinci } from './davinci.ts';
// import { setupNavBar } from './navbar.ts';

declare global {
    interface Window {
        vm: any, // Vue.js instance
        observer: any
    }
}

document.addEventListener('DOMContentLoaded', () => {
    // if (env == "development"){
    //     setupDaVinci();
    // }

    // // if we are on the homepage run vue
    // if (window.location.pathname === '/') {
    //     setupLanding()
    // }

    initVue();


});  
