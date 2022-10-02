function main() {
    highlightActiveNavItem();
    initializeTooltips();
}

function highlightActiveNavItem() {
    var navLinks = document.querySelectorAll("nav a");
    for (var i = 0; i < navLinks.length; i++) {
        var link = navLinks[i]
        if (link.getAttribute('href') == window.location.pathname) {
            link.classList.add("active");
            break;
        }
    }
}

function initializeTooltips() {
    const tooltipTriggerList = document.querySelectorAll('[data-bs-toggle="tooltip"]')
    return [...tooltipTriggerList].map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl))
}

main();