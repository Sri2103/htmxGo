import Alpine from 'alpinejs'
import 'htmx.org'
window.Alpine = Alpine
window.htmx = require('htmx.org')
document.addEventListener("alpine:init", () => {
    Alpine.data('dropdown', () => ({
        open: false,
        toggle() {
            this.open = !this.open
        }
    }))
})




Alpine.start()