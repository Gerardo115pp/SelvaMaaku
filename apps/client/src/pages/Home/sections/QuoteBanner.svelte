<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import viewport from "@components/viewport_actions/useViewportActions";
    import { writable } from "svelte/store";
    import { fade } from "svelte/transition";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /** 
         * Whether the component is visible or not 
         * @type {import("svelte/store").Writable<boolean>} 
         */
        let component_visible = writable(false);
    
        
        /*----------  Entry animation  ----------*/
        
            /** 
             * Duration of the entry animation in milliseconds 
             * @type {number} 
             */
            const entry_animation_duration = 1500;
    
    /*=====  End of Properties  ======*/
    
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        const setVisibility = is_visible => {
            if ($component_visible === is_visible) return;
    
            component_visible.set(is_visible);
        }
    
    /*=====  End of Methods  ======*/
    
    
    
</script>

<section id="smk-quote-banner-section">
    <div id="smk-qbs-content-wrapper" class="design-content-width">
        <div id="smk-qbs-accent-icon-wrapper" class="wixarica-icons">
            <WixaricaIcon with_accent/>
        </div>
        <blockquote id="smk-qbs-quote-information"
            on:viewportEnter={() => setVisibility(true)}
            on:viewportLeave={() => setVisibility(false)}
            use:viewport={{height_offset: 0.2}}
        >
            {#if $component_visible}
                 <h2 id="smk-qbs-qi-quote" transition:fade={{duration: entry_animation_duration, delay: entry_animation_duration * 0.4}}>
                     Is there a reward greater than life?
                 </h2>
                 <footer transition:fade={{duration: entry_animation_duration, delay: entry_animation_duration * 1.3}}>
                     <cite id="smk-qbs-quote-source">
                         From the book Life of Pi by Yann Martel
                     </cite>
                 </footer>
            {/if}
        </blockquote>
    </div>
</section>

<style>
    #smk-quote-banner-section {
        display: flex;
        justify-content: center;
        padding: var(--spacing-5) 0;
    }

    #smk-qbs-content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }

    section#smk-quote-banner-section .wixarica-icons {
        width: calc(var(--wixarica-icons-size) * 1.25);
    }

    blockquote#smk-qbs-quote-information {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-3);
        align-items: center;
        text-align: center;
    }
    
    blockquote#smk-qbs-quote-information h2 {
        font-size: var(--font-size-h3);

        &::before {
            content: "“";
        }

        &::after {
            content: "”";
        }
    }

    blockquote#smk-qbs-quote-information footer cite::before {
        content: "— ";
    }
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            blockquote#smk-qbs-quote-information {
                gap: var(--spacing-2);  
            }

            blockquote#smk-qbs-quote-information h2 {
                font-size: var(--font-size-h2);
            }

            blockquote#smk-qbs-quote-information footer cite {
                font-size: var(--font-size-fineprint);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>