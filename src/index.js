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
    Alpine.data("showing", () => {
        return {
            showing: false,
            testData: "Hello",
            testAgain: "Data-2",
            show() {
                this.showing = true
            },
            hide() {
                this.showing = false
            }
        }
    })
})




Alpine.start()
