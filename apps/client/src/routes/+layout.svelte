<script>
    import "@app/app.css";
    import Footer from "@components/Footer/footer.svelte";
    import Navbar from "@components/Navbar/Navbar.svelte";
    import PageTopGradient from "@components/Ui_decorations/page_top_gradient.svelte";
    import ImageDisplayerOverlay from "@components/Images/ImageDisplayerOverlay.svelte";
    import { hasChangedLayout, defineLayout, setVirtualMobile } from "@stores/layout";
    import { setSiteLanguage } from "@stores/site_content";
    import { browser } from "$app/environment";
    import { onMount } from "svelte";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /** @type {import('./$types').LayoutData} */
        export let data;

        /** @type {string} */
        let site_detected_language = data.language ?? "en";

        setSiteLanguage(site_detected_language);
        setVirtualMobile(data.is_virtual_mobile);
    
    /*=====  End of Properties  ======*/
    
    

    onMount(() => {
        defineLayout();
    })

    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Handles the layout resize event
         * @param {UIEvent} event
         */
         const handleLayoutResize = e => {
            if (!browser) return;

            if (hasChangedLayout()) {
                defineLayout();
            }
        }
    
    /*=====  End of Methods  ======*/
    
    
</script>

<svelte:window 
    on:resize|passive={handleLayoutResize}
/>
<div id="app-root" class:adebug={false}>
    <ImageDisplayerOverlay />
    <PageTopGradient />
    <Navbar />
    <slot></slot>
    <Footer />
</div>

<style>
    :global(html) {
        background: var(--body-bg-color);
        color: var(--body-text-color);
    }

    @media (prefers-reduced-motion) {
        :global(html) {
            scroll-behavior: auto;
        }
    }

    #app-root {
        position: relative;
    }
</style>