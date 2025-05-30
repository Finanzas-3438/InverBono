export function fancyTransition() {
  let overlay = document.createElement("div");
  overlay.classList.add(
    "bg-black",
    "fixed",
    "top-0",
    "left-0",
    "w-full",
    "h-full",
    "z-50",
    "transition-transform",
    "duration-300",
    "ease-in-out"
  );
  overlay.style.transform = "translateX(-100%)";
  overlay.style.opacity = "1";
  document.body.appendChild(overlay);

  document.querySelectorAll("a").forEach((el) => {
    let href = el.getAttribute("href");
    if (!href) return;

    el.classList.add("cursor-pointer");

    el.onclick = (e) => {
      if (href.startsWith("#")) {
        return; 
      }
      if (href === window.location.href || new URL(href, window.location.origin).href === window.location.href) {

         const isExternal = new URL(href, window.location.origin).origin !== window.location.origin;
         if (!isExternal && new URL(href, window.location.origin).pathname === window.location.pathname && new URL(href, window.location.origin).search === window.location.search) {
            return;
         }
      }

      e.preventDefault(); 

      overlay.style.transform = "translateX(0%)";

      setTimeout(() => {
        window.location.href = href;
      }, 300);
    };
  });

  window.addEventListener('pageshow', function(event) {
    overlay.style.transition = 'none';
    overlay.style.transform = 'translateX(-100%)';
    setTimeout(() => {
        overlay.style.transition = 'transform 0.3s ease-in-out';
    }, 0);
  });
}